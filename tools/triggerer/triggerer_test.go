package main

import (
	"testing"
	"time"
)

func TestUntilNextEventStartBeforeEnd(t *testing.T) {
	StartHour = 10
	EndHour = 20
	Interval = 1800 // 30 mins
	durInterval := time.Duration(Interval) * time.Second

	forcedTime := time.Date(2022, 12, 11, 19, 50, 0, 0, loc)
	timeNow = func() time.Time { return forcedTime }

	// ensure duration > 0 when now is in window but interval is after EndHour
	dur := untilNextEvent()
	if dur < 0 {
		t.Errorf("dur [%v] less than zero, should not be possible", dur)
	}

	// ensure duration = interval within window and will start before end
	forcedTime = time.Date(2022, 12, 11, 19, 29, 0, 0, loc)
	dur = untilNextEvent()
	if dur != durInterval {
		t.Errorf("dur [%v] should be equal to [%v[] ", dur, durInterval)
	}

	// ensure duration == exactly 14 hours
	forcedTime = time.Date(2022, 12, 11, 20, 0, 0, 0, loc)
	dur = untilNextEvent()
	want := time.Duration(14) * time.Hour
	if dur != want {
		t.Errorf("dur [%v] should be [%v]", dur, want)
	}

	// ensure duration != exactly 14 hours
	forcedTime = time.Date(2022, 12, 11, 8, 0, 0, 0, loc)
	dur = untilNextEvent()
	want = time.Duration(2) * time.Hour
	if dur == want {
		t.Errorf("dur [%v] should be [%v]", dur, want)
	}
}

func TestUntilNextEventEndBeforeStartj(t *testing.T) {
	StartHour = 20
	EndHour = 10
	Interval = 1800 // 30 mins
	durInterval := time.Duration(Interval) * time.Second

	forcedTime := time.Date(2022, 12, 11, 9, 50, 0, 0, loc)
	timeNow = func() time.Time { return forcedTime }

	// ensure duration > 0 when now is in window but interval is after EndHour
	dur := untilNextEvent()
	if dur < 0 {
		t.Errorf("dur [%v] less than zero, should not be possible", dur)
	}

	// ensure duration = interval within window and will start before end
	forcedTime = time.Date(2022, 12, 11, 9, 30, 0, 0, loc)
	dur = untilNextEvent()
	if dur != durInterval {
		t.Errorf("dur [%v] should be equal to [%v[] ", dur, durInterval)
	}

	// ensure duration == exactly 9 hours
	forcedTime = time.Date(2022, 12, 11, 11, 0, 0, 0, loc)
	dur = untilNextEvent()
	want := time.Duration(9) * time.Hour
	if dur != want {
		t.Errorf("dur [%v] should be [%v]", dur, want)
	}
}
