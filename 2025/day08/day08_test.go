package day08

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt", 10)
	fmt.Printf("Got: %v\n", got)
	want := 40
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt", 1000)
	fmt.Printf("Got: %v\n", got)
	want := 0
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 25272
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 0
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
