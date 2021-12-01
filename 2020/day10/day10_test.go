package day10

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 35
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}

	got = part1("sample2.txt")
	fmt.Printf("Got: %v\n", got)
	want = 220
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 2664
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	var got int
	var want int
	got = part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want = 8
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}

	got = part2("sample2.txt")
	fmt.Printf("Got: %v\n", got)
	want = 19208
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 148098383347712
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
