package main

import (
	"aoc/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

var (
	sessionTokenEnvKey   = "SESSION_COOKIE"
	leaderboardURLEnvKey = "LEADER_URL"
	discordURLEnvKey     = "DISCORD_URL"
	numStarsEnvKey       = "NUM_STARS_TO_PRINT"
	leaderboardCacheFile = ".cache.leader"
	timeLocation, _      = time.LoadLocation("America/Chicago")
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

var force bool

func main() {
	l := pullAndCacheLeaderboard()

	diff := loadAndSaveDiff(l)

	results := buildResultsContent(diff)

	fmt.Println()
	fmt.Print(results.String())

	if len(results.String()) > 0 {
		publishToDiscord(results)
	}
}

func loadAndSaveDiff(l Leaderboard) Leaderboard {
	state := LeaderState{}
	if err := state.LoadState(); err != nil {
		log.Fatal(err)
	}

	diff := state.GenerateDiff(l)
	if err := state.SaveState(l); err != nil {
		log.Fatal(err)
	}

	log.Printf("state updated: %v", state.GetStateFile())
	return diff
}

func buildResultsContent(l Leaderboard) *strings.Builder {
	results := new(strings.Builder)

	numStars := util.MaxInt
	numStarsStr := os.Getenv(numStarsEnvKey)
	if len(numStarsStr) > 0 {
		t, err := strconv.Atoi(numStarsStr)
		if err == nil {
			numStars = t
		}
	}

	memberKeys := sortMemberKeysByName(l.Members)

	// Print the users, days, stars, and timestamps
	for _, mKey := range memberKeys {
		m := l.Members[mKey]

		dayKeys := util.SortMapKeysInt(m.CompletionDayLevel)
		if len(dayKeys) == 0 {
			continue
		}

		fmt.Fprintf(results, "%+v\n", m.Name)

		daysToUse := dayKeys
		if len(daysToUse) > numStars {
			daysToUse = daysToUse[len(daysToUse)-numStars:]
		}

		for _, day := range daysToUse {
			stars := m.CompletionDayLevel[day]

			sortedStarsKeys := util.SortMapKeys(stars)
			for i, k := range sortedStarsKeys {
				tms := stars[k]
				ti := time.Unix(tms.StarTms, 0)
				t := fTime(ti)
				if i == 0 {
					fmt.Fprintf(results, "  Day %2s - %v\n", day, t)
				} else {
					// grab first star details again
					tms1 := stars[sortedStarsKeys[0]]
					ti1 := time.Unix(tms1.StarTms, 0)
					dur := ti.Sub(ti1)

					// second star
					fmt.Fprintf(results, "           %v - %v\n", t, dur)

				}
			}
		}
		fmt.Fprintln(results)
	}

	return results
}

// fTime formats time
func fTime(t time.Time) string {
	return t.In(timeLocation).Format("2006-01-02 03:04 PM")
}

// func dayOfMonth() int {
// 	return time.Now().In(timeLocation).Day()
// }

func sortMemberKeysByName(m map[string]Member) []string {
	keys := []string{}

	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		a := strings.ToLower(m[keys[i]].Name)
		b := strings.ToLower(m[keys[j]].Name)

		return a < b
	})

	return keys
}

func publishToDiscord(results *strings.Builder) {
	discordURL := os.Getenv(discordURLEnvKey)
	log.Printf("%v length: %v", discordURLEnvKey, len(discordURL))
	if len(discordURL) > 0 {
		// GITHUB_EVENT_NAME=schedule for cron workflows
		data := struct {
			Content string `json:"content"`
		}{
			"```\nNew Stars:\n\n" + results.String() + "\n```",
		}

		dataB, err := json.Marshal(data)
		util.Check(err)

		reader := bytes.NewReader(dataB)
		// reader := strings.NewReader(results.String())

		resp, err := http.Post(discordURL, "application/json", reader)
		util.Check(err)
		defer resp.Body.Close()

		dataB, err = io.ReadAll(resp.Body)
		util.Check(err)

		log.Printf("Discord POST StatusCode [%v], Body: %v", resp.StatusCode, string(dataB))
	}
}

func pullAndCacheLeaderboard() Leaderboard {
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
			log.Printf("Missing env key: %v\n", sessionTokenEnvKey)
			os.Exit(1)
		}

		if os.Getenv(leaderboardURLEnvKey) == "" {
			log.Printf("Missing env key: %v\n", leaderboardURLEnvKey)
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
			os.WriteFile(leaderboardCacheFile, raw, 0644)
		} else {
			os.WriteFile(leaderboardCacheFile, prettyJSON.Bytes(), 0644)
		}
	}

	return l
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

	data, err := os.ReadFile(leaderboardCacheFile)
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
		log.Printf("Error, status code != 200: [%v]:%v\n", resp.StatusCode, resp.Status)
		os.Exit(1)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	util.Check(err)

	return data

}
