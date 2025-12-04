package day04

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 13
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 1553
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 43
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 8442
	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}
}
