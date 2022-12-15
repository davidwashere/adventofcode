package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const (
	EnvPrefix          = "TRIGGERER_"
	GHTokenEnvKey      = EnvPrefix + "GITHUB_TOKEN"
	GHOwnerEnvKey      = EnvPrefix + "GITHUB_OWNER"
	GHRepoEnvKey       = EnvPrefix + "GITHUB_REPO"
	GHActionFileEnvKey = EnvPrefix + "GITHUB_ACTION_FILE"
	GHRefEnvKey        = EnvPrefix + "GITHUB_REF"
	IntervalEnvKey     = EnvPrefix + "INTERVAL"
)

var (
	GHToken         = ""
	GHOwner         = ""
	GHRepo          = ""
	GHActionFile    = ""
	GHRef           = ""
	Interval        = 30
	timeLocation, _ = time.LoadLocation("America/Chicago")
)

func main() {
	log.Printf("loading env")
	loadAndValidateEnv()

	log.Printf("doing thing - initial")
	err := doImportantThing()
	if err != nil {
		panic(err)
	}

	log.Printf("running forever every %v min(s)", Interval)
	runForever()

	log.Printf("done")
}

func loadAndValidateEnv() {
	GHToken = panicOnEmpty(GHTokenEnvKey)
	GHOwner = panicOnEmpty(GHOwnerEnvKey)
	GHRepo = panicOnEmpty(GHRepoEnvKey)
	GHActionFile = panicOnEmpty(GHActionFileEnvKey)
	GHRef = panicOnEmpty(GHRefEnvKey)

	intervalStr := panicOnEmpty(IntervalEnvKey)

	var err error
	Interval, err = strconv.Atoi(intervalStr)
	if err != nil {
		panic(fmt.Sprintf("%v is not a valid integer: %v", IntervalEnvKey, intervalStr))
	}

	if Interval < 5 {
		panic(fmt.Sprintf("%v must be 5 greater", IntervalEnvKey))
	}
}

func runForever() {
	ticker := time.NewTicker(time.Duration(Interval) * time.Minute)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				log.Printf("doing thing at %v", fTime(t))
				err := doImportantThing()
				if err != nil {
					log.Print(err)
				}
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan
	done <- true
}

// panciOnEmpty will trigger panic if the env variable is missing
func panicOnEmpty(envKey string) string {
	v := os.Getenv(envKey)
	if len(v) == 0 {
		panic(fmt.Sprintf("set %v in env", envKey))
	}

	return v
}

func doImportantThing() error {
	req, err := createHttpReq()
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("result: Status[%v] Body [%v]", resp.Status, string(body))
	return nil
}

func createHttpReq() (*http.Request, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/actions/workflows/%v/dispatches", GHOwner, GHRepo, GHActionFile)
	dataB := []byte(fmt.Sprintf("{\"ref\":\"%v\"}", GHRef))
	reader := bytes.NewReader(dataB)
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", GHToken))
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	return req, nil
}

func fTime(t time.Time) string {
	return t.In(timeLocation).Format("2006-01-02 03:04:05 PM")
}
