package day19

import (
	"fmt"
	"regexp"
	"testing"
)

var vf = func(t *testing.T, got, want interface{}) {
	if got != want {
		t.Helper()
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestRe(t *testing.T) {
	var re *regexp.Regexp
	re = regexp.MustCompile("^a((c)|(d))b$")
	vf(t, match(re, "acdab"), false)
	vf(t, match(re, "acdb"), false)
	vf(t, match(re, "acb"), true)
	vf(t, match(re, "adb"), true)

	re = regexp.MustCompile("^(a((aa|bb)(ab|ba)|(ab|ba)(aa|bb))b)$")
	vf(t, match(re, "ababbb"), true)
	vf(t, match(re, "bababa"), false)
	vf(t, match(re, "abbbab"), true)
	vf(t, match(re, "aaabbb"), false)
	vf(t, match(re, "aaaabbb"), false)

	// Recursive regex not supported by Go Lang
	re = regexp.MustCompile("^a((cd)|(ccdd)|(cccddd))b$")
	vf(t, match(re, "ab"), false)
	vf(t, match(re, "acdb"), true)
	vf(t, match(re, "accdb"), false)
	vf(t, match(re, "acddb"), false)
	vf(t, match(re, "accddb"), true)
	vf(t, match(re, "accdddb"), false)
	vf(t, match(re, "acccddb"), false)
	vf(t, match(re, "acccdddb"), true)
}

func TestBuildRe(t *testing.T) {
	rules, _ := parse("sample.txt")
	reString := buildReString("0", rules)

	fmt.Println(reString)
}

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 2
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 109
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sampleP2.txt")
	// got := part2("sampleP3.txt")
	fmt.Printf("Got: %v\n", got)
	want := 12
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	// got := part2("input.txt")
	got := part2("inputP2.txt")
	fmt.Printf("Got: %v\n", got)
	want := 301
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
