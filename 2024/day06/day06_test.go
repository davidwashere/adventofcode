package day06

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 41
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input2.txt")
	fmt.Printf("Got: %v\n", got)
	want := 5551
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 6
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input2.txt")
	fmt.Printf("Got: %v\n", got)
	want := 0
	// 1523 // TO LOW
	// 2235 // TO HIGH
	// 2239 // TO HIGH
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
