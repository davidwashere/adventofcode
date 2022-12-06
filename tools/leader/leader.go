package main

import (
	"aoc/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"

	"golang.org/x/net/publicsuffix"
)

var (
	sessionTokenEnvKey   = "SESSION_COOKIE"
	leaderboardURLEnvKey = "LEADER_URL"
	leaderboardCacheFile = ".cache.leader"
)

type Leaderboard struct {
	OwnerId int               `json:"owner_id"`
	Event   string            `json:"event"`
	Members map[string]Member `json:"members"`
}

type Member struct {
	ID                 int                                 `json:"id"`
	Name               string                              `json:"name"`
	Stars              int                                 `json:"stars"`
	LocalScore         int                                 `json:"local_score"`
	GlobalScore        int                                 `json:"global_score"`
	CompletionDayLevel map[string]map[string]StarTimestamp `json:"completion_day_level"`
}

type StarTimestamp struct {
	StarTms int64 `json:"get_star_ts"`
}

func pullCachedLeaderboard() ([]byte, error) {
	info, err := os.Stat(leaderboardCacheFile)
	if err != nil {
		return nil, err
	}

	elapsed := time.Since(info.ModTime())

	log.Printf("Cache file last update: %v (%.2f mins ago)", fTime(info.ModTime()), elapsed.Minutes())

	if elapsed.Minutes() > 15 {
		return nil, fmt.Errorf("cache expired")
	}

	data, err := ioutil.ReadFile(leaderboardCacheFile)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func pullLeaderboard() []byte {
	leaderURL := os.Getenv(leaderboardURLEnvKey)
	session := os.Getenv(sessionTokenEnvKey)

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	util.Check(err)

	url, err := url.Parse(leaderURL)
	util.Check(err)

	cookie := http.Cookie{
		Name:  "session",
		Value: session,
	}

	jar.SetCookies(url, []*http.Cookie{&cookie})

	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Get(leaderURL)
	util.Check(err)

	if resp.StatusCode != 200 {
		fmt.Printf("Error, status code != 200: [%v]:%v\n", resp.StatusCode, resp.Status)
		os.Exit(1)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	util.Check(err)

	return data

}

var force bool

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "force" {
			force = true
		}
	}

	l := Leaderboard{}
	var err error
	var raw []byte

	if !force {
		// Load Leaderboard data from cache (aoc asked do not pull more than once / 15 mins)
		raw, err = pullCachedLeaderboard()
		if err == nil {
			log.Printf("%v found", leaderboardCacheFile)
			err = json.Unmarshal(raw, &l)
			if err == nil {
				log.Printf("%v parsed successfully", leaderboardCacheFile)
			}
		}
	}

	// Load leaderboard directly from AoC if didn't load from cache
	if err != nil || len(raw) == 0 {
		err := util.LoadEnv()
		if err != nil {
			log.Printf(".env.local not found or no worky, checking raw env: %v", err)
		}

		if os.Getenv(sessionTokenEnvKey) == "" {
			fmt.Printf("Missing env key: %v\n", sessionTokenEnvKey)
			os.Exit(1)
		}

		if os.Getenv(leaderboardURLEnvKey) == "" {
			fmt.Printf("Missing env key: %v\n", leaderboardURLEnvKey)
			os.Exit(1)
		}

		log.Printf("No cache, cache invalid, or forced - pulling from https://adventofcode.com/")
		raw = pullLeaderboard()
		err = json.Unmarshal(raw, &l)
		util.Check(err)

		// Cache output
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, raw, "", "\t")
		if err != nil {
			ioutil.WriteFile(leaderboardCacheFile, raw, 0644)
		} else {
			ioutil.WriteFile(leaderboardCacheFile, prettyJSON.Bytes(), 0644)
		}

	}

	// Print the users, days, stars, and timestamps
	for _, m := range l.Members {

		dayKeys := util.SortMapKeys(m.CompletionDayLevel)
		if len(dayKeys) == 0 {
			continue
		}

		fmt.Printf("%+v\n", m.Name)
		for _, day := range dayKeys {
			stars := m.CompletionDayLevel[day]

			sortedStarsKeys := util.SortMapKeys(stars)
			for i, k := range sortedStarsKeys {
				tms := stars[k]
				ti := time.Unix(tms.StarTms, 0)
				t := fTime(ti)
				if i == 0 {
					fmt.Printf("  Day %2s - %v\n", day, t)
				} else {
					// grab first star details again
					tms1 := stars[sortedStarsKeys[0]]
					ti1 := time.Unix(tms1.StarTms, 0)
					dur := ti.Sub(ti1)

					// second star
					fmt.Printf("           %v - %v\n", t, dur)

				}
			}
		}
		fmt.Println()
	}
}

// fTime formats time
func fTime(t time.Time) string {
	return t.Format("2006-01-02 03:04 PM")
}
