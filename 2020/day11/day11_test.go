package day11

import (
	"aoc2020/util"
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 37
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 2368
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestOccupiedSeats(t *testing.T) {
	grid := util.NewInfinityGridFromFile("sample2.txt", "X")
	got := adjSeenOccupiedSeats(grid, 3, 4)
	fmt.Printf("Got: %v\n", got)
	want := 8
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	var got int
	var want int
	got = part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want = 26
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 2124
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
