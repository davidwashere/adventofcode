package day09

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt", 5)
	fmt.Printf("Got: %v\n", got)
	want := 127
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt", 25)
	fmt.Printf("Got: %v\n", got)
	want := 1398413738
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt", 5)
	fmt.Printf("Got: %v\n", got)
	want := 62
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt", 25)
	fmt.Printf("Got: %v\n", got)
	want := 169521051
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
