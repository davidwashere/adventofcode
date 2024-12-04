package day03

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 161
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 175700056
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample2.txt")
	fmt.Printf("Got: %v\n", got)
	want := 48
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 71668682
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
