package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
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
	StartHourEnvKey    = EnvPrefix + "START_HOUR"
	EndHourEnvKey      = EnvPrefix + "END_HOUR"
	DebugEnvKey        = EnvPrefix + "DEBUG"
)

var (
	GHToken      = ""
	GHOwner      = ""
	GHRepo       = ""
	GHActionFile = ""
	GHRef        = ""
	Interval     = 30
	StartHour    = 9
	EndHour      = 23
	Debug        = false

	loc, _ = time.LoadLocation("America/Chicago")
)

func main() {
	log.Printf("loading env")
	dumpEnv()
	loadAndValidateEnv()

	// log.Printf("doing thing - initial")
	// err := doImportantThing()
	// if err != nil {
	// 	panic(err)
	// }

	log.Printf("running forever every %v sec(s)", Interval)
	runForever()

	log.Printf("done")
}

func dumpEnv() {
	log.Printf("%v = %v", GHOwnerEnvKey, os.Getenv(GHOwnerEnvKey))
	log.Printf("%v = %v", GHRepoEnvKey, os.Getenv(GHRepoEnvKey))
	log.Printf("%v = %v", GHActionFileEnvKey, os.Getenv(GHActionFileEnvKey))
	log.Printf("%v = %v", GHRefEnvKey, os.Getenv(GHRefEnvKey))
	log.Printf("%v = %v", IntervalEnvKey, os.Getenv(IntervalEnvKey))
	log.Printf("%v = %v", StartHourEnvKey, os.Getenv(StartHourEnvKey))
	log.Printf("%v = %v", EndHourEnvKey, os.Getenv(EndHourEnvKey))
	log.Printf("%v = %v", DebugEnvKey, os.Getenv(DebugEnvKey))
}

// untilNextEvent will return the 'duration' until the next event should trigger
func untilNextEvent() time.Duration {
	now := time.Now().In(loc).Round(0)

	nowPlusInterval := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()+Interval, now.Nanosecond(), loc)
	startToday := time.Date(now.Year(), now.Month(), now.Day(), StartHour, 0, 0, 0, loc)
	endToday := time.Date(now.Year(), now.Month(), now.Day(), EndHour, 0, 0, 0, loc)
	startTomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, StartHour, 0, 0, 0, loc)

	if now.After(startToday) && nowPlusInterval.Before(endToday) {
		// we're within start/end window for today, return a duration representing interval
		return time.Duration(Interval) * time.Second
	}

	if now.After(endToday) {
		// next event would be tomorrow because 'now' represents 'today'
		return time.Until(startTomorrow)
	}

	// now is before startHour for today
	return time.Until(startToday)
}

func loadAndValidateEnv() {
	var err error

	GHToken = panicOnEmpty(GHTokenEnvKey)
	GHOwner = panicOnEmpty(GHOwnerEnvKey)
	GHRepo = panicOnEmpty(GHRepoEnvKey)
	GHActionFile = panicOnEmpty(GHActionFileEnvKey)
	GHRef = panicOnEmpty(GHRefEnvKey)

	intervalStr := os.Getenv(IntervalEnvKey)
	if len(intervalStr) > 0 {
		Interval, err = strconv.Atoi(intervalStr)
		if err != nil {
			panic(fmt.Sprintf("%v is not a valid integer: %v", IntervalEnvKey, intervalStr))
		}

		if Interval < 5 {
			panic(fmt.Sprintf("%v must be 5 greater", IntervalEnvKey))
		}
	}

	startHourStr := os.Getenv(StartHourEnvKey)
	if len(startHourStr) > 0 {
		StartHour, err = strconv.Atoi(startHourStr)
		if err != nil {
			panic(fmt.Sprintf("%v is not a valid integer: %v", StartHourEnvKey, startHourStr))
		}

		if StartHour < 0 || StartHour > 23 {
			panic(fmt.Sprintf("%v must be between 0 and 23", StartHourEnvKey))
		}
	}

	endHourStr := os.Getenv(EndHourEnvKey)
	if len(endHourStr) > 0 {
		EndHour, err = strconv.Atoi(endHourStr)
		if err != nil {
			panic(fmt.Sprintf("%v is not a valid integer: %v", EndHourEnvKey, endHourStr))
		}

		if EndHour < 1 || EndHour > 24 {
			panic(fmt.Sprintf("%v must be between 0 and 23", EndHourEnvKey))
		}
	}

	if StartHour >= EndHour {
		panic(fmt.Sprintf("%v [%v] must be before %v [%v] be between 0 and 23", StartHourEnvKey, StartHour, EndHourEnvKey, EndHour))
	}

	debugStr := os.Getenv(DebugEnvKey)
	if strings.EqualFold(debugStr, "true") {
		Debug = true
	}
}

func runForever() {
	done := make(chan bool)
	dur := untilNextEvent()
	log.Printf("next trigger in: %v", dur)
	timer := time.NewTimer(dur)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-timer.C:
				err := doImportantThing()
				if err != nil {
					log.Print(err)
				}
				dur := untilNextEvent()
				log.Printf("next trigger in: %v", dur)
				timer.Reset(dur)
			}
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan
	done <- true
	timer.Stop()
}

// panicOnEmpty will trigger panic if the env variable is missing
func panicOnEmpty(envKey string) string {
	v := os.Getenv(envKey)
	if len(v) == 0 {
		panic(fmt.Sprintf("set %v in env", envKey))
	}

	return v
}

func doImportantThing() error {
	if Debug {
		log.Printf("triggered - in debug mode - doing nothing")
		return nil
	} else {

	}
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

	body, err := io.ReadAll(resp.Body)
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

// func fTime(t time.Time) string {
// 	return t.In(loc).Format("2006-01-02 03:04:05 PM")
// }
