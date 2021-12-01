package day23

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("389125467")
	fmt.Printf("Got: %v\n", got)
	want := "67384529"
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("215694783")
	fmt.Printf("Got: %v\n", got)
	want := "38925764"
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("389125467")
	fmt.Printf("Got: %v\n", got)
	want := 149245887792
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("215694783")
	fmt.Printf("Got: %v\n", got)
	// want := 0
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}
