package day13

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 295
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 4808
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	var got int
	var want int
	got = part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want = 1068781
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}

	got = part2("sample2.txt")
	fmt.Printf("Got: %v\n", got)
	want = 3417
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}

	got = part2("sample3.txt")
	fmt.Printf("Got: %v\n", got)
	want = 754018
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}

	got = part2("sample4.txt")
	fmt.Printf("Got: %v\n", got)
	want = 779210
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}

	got = part2("sample5.txt")
	fmt.Printf("Got: %v\n", got)
	want = 1261476
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}

	got = part2("sample6.txt")
	fmt.Printf("Got: %v\n", got)
	want = 1202161486
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 741745043105674
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
