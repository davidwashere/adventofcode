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
	Interval     = 1800
	StartHour    = 9
	EndHour      = 23
	Debug        = false

	loc, _ = time.LoadLocation("America/Chicago")

	timeNow   = time.Now
	timeUntil = func(t time.Time) time.Duration { return t.Sub(timeNow()) }
)

func main() {
	time.Local = loc
	dumpEnv()
	loadAndValidateEnv()

	log.Printf("current time %v", fTime(timeNow()))
	log.Printf("current time %v", timeNow())

	runForever()

	log.Printf("done")
}

func dumpEnv() {
	max := len(GHActionFileEnvKey)
	f := fmt.Sprintf("  %%-%vv = %%v", max)

	log.Printf("Env:")
	log.Printf(f, GHOwnerEnvKey, os.Getenv(GHOwnerEnvKey))
	log.Printf(f, GHRepoEnvKey, os.Getenv(GHRepoEnvKey))
	log.Printf(f, GHActionFileEnvKey, os.Getenv(GHActionFileEnvKey))
	log.Printf(f, GHRefEnvKey, os.Getenv(GHRefEnvKey))
	log.Printf(f, IntervalEnvKey, os.Getenv(IntervalEnvKey))
	log.Printf(f, StartHourEnvKey, os.Getenv(StartHourEnvKey))
	log.Printf(f, EndHourEnvKey, os.Getenv(EndHourEnvKey))
	log.Printf(f, DebugEnvKey, os.Getenv(DebugEnvKey))
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

		if Interval < 1 {
			panic(fmt.Sprintf("%v must be greater than 0", IntervalEnvKey))
		}
	}

	startHourStr := os.Getenv(StartHourEnvKey)
	if len(startHourStr) > 0 {
		StartHour, err = strconv.Atoi(startHourStr)
		if err != nil {
			panic(fmt.Sprintf("%v is not a valid integer: %v", StartHourEnvKey, startHourStr))
		}

		if StartHour < 0 || StartHour > 24 {
			panic(fmt.Sprintf("%v must be between 0 and 23", StartHourEnvKey))
		}
	}

	endHourStr := os.Getenv(EndHourEnvKey)
	if len(endHourStr) > 0 {
		EndHour, err = strconv.Atoi(endHourStr)
		if err != nil {
			panic(fmt.Sprintf("%v is not a valid integer: %v", EndHourEnvKey, endHourStr))
		}

		if EndHour < 0 || EndHour > 24 {
			panic(fmt.Sprintf("%v must be between 0 and 23", EndHourEnvKey))
		}
	}

	debugStr := os.Getenv(DebugEnvKey)
	if strings.EqualFold(debugStr, "true") {
		Debug = true
	}
}

func runForever() {
	done := make(chan bool)
	dur := untilNextEventSafe()
	// durInterval := time.Duration(Interval) * time.Second
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
				dur := untilNextEventSafe()
				if dur < time.Duration(0) {
					panic(fmt.Sprintf("%v is less than zero", dur))

				}
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

func untilNextEventSafe() time.Duration {
	dur := untilNextEvent()
	if dur < 0 {
		now := timeNow()
		log.Printf("error, duration less than zero - now: %v startHour: %v endHour %v", now, StartHour, EndHour)
		panic("duration less than zero")
	}

	return dur
}

// untilNextEvent will calculate the duration until the next event should occur
func untilNextEvent() time.Duration {
	durInterval := time.Duration(Interval) * time.Second
	if StartHour == EndHour {
		// Start and End are the same = running constant for 24 hours
		return durInterval
	}

	now := timeNow().Round(0)
	nowPlusInterval := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()+Interval, now.Nanosecond(), loc)

	// Example: StartHour = 9 (9am), EndHour = 21 (9pm)
	//
	//            11111111112222           11111111112222
	//  012345678901234567890123 012345678901234567890123
	// │         S───────────E  │         S───────────E  │
	//
	if StartHour < EndHour {
		start := time.Date(now.Year(), now.Month(), now.Day(), StartHour, 0, 0, 0, loc)
		end := time.Date(now.Year(), now.Month(), now.Day(), EndHour, 0, 0, 0, loc)

		startTomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, StartHour, 0, 0, 0, loc)

		if now.After(start) && nowPlusInterval.Before(end) {
			return durInterval
		}

		// if now.After(end) {
		if nowPlusInterval.After(end) {
			// next event would be tomorrow because 'now' represents 'today'
			return timeUntil(startTomorrow)
		}

		return timeUntil(start)
	}

	// Example: StartHour = 21 (9pm), EndHour = 9 (9am)
	//
	//            11111111112222           11111111112222
	//  012345678901234567890123 012345678901234567890123
	// ├─────────E           S──┼─────────E           S──┤
	//
	start := time.Date(now.Year(), now.Month(), now.Day(), StartHour, 0, 0, 0, loc)
	end := time.Date(now.Year(), now.Month(), now.Day(), EndHour, 0, 0, 0, loc)

	if nowPlusInterval.After(end) && now.Before(start) {
		return timeUntil(start)
	}

	return durInterval
}

func fTime(t time.Time) string {
	return t.Format("2006-01-02 03:04 PM")
}
