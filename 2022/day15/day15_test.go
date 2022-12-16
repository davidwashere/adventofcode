package day15

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt", 10)
	fmt.Printf("Got: %v\n", got)
	want := 26
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt", 2000000)
	fmt.Printf("Got: %v\n", got)
	want := 4717631
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt", 20)
	fmt.Printf("Got: %v\n", got)
	want := 56000011
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt", 4000000)
	fmt.Printf("Got: %v\n", got)
	want := 13197439355220
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
