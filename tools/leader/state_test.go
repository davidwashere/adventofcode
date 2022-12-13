package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLeaderStateDiff(t *testing.T) {
	dataB, _ := ioutil.ReadFile("testdata/leader_state.json")

	var lOld Leaderboard
	var lNew Leaderboard

	json.Unmarshal(dataB, &lOld)

	dataB, _ = ioutil.ReadFile("testdata/leader_state_update.json")

	json.Unmarshal(dataB, &lNew)

	state := LeaderState{}
	state.OldBoard = &lOld
	diff := state.GenerateDiff(lNew)

	var want = 3
	var got = len(diff.Members)

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	if m, ok := diff.Members["264332"]; !ok {
		t.Errorf("davidwashere's entry is missing")
	} else {
		got := len(m.CompletionDayLevel)
		want := 1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

		got = len(m.CompletionDayLevel["11"])
		want = 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	if m, ok := diff.Members["676994"]; !ok {
		t.Errorf("Kris Heyden's entry is missing")
	} else {
		got := len(m.CompletionDayLevel)
		want := 1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

		got = len(m.CompletionDayLevel["11"])
		want = 1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	if m, ok := diff.Members["759608"]; !ok {
		t.Errorf("Steven Goff's entry is missing")
	} else {
		got := len(m.CompletionDayLevel)
		want := 1

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}

		got = len(m.CompletionDayLevel["4"])
		want = 2

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	result := buildResultsContent(diff)

	// new davidwashere day 11 star 2
	// new kris day 11 star 1
	// new steven day 4 star 1 & 2

	fmt.Println(result.String())

}
