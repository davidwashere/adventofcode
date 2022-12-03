package day03

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 0
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestPriority(t *testing.T) {
	tt := []struct {
		in   byte
		want int
	}{
		{'p', 16},
		{'L', 38},
		{'P', 42},
		{'v', 22},
		{'t', 20},
		{'s', 19},
	}

	for _, test := range tt {
		got := priority(test.in)

		if got != test.want {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 0
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 70
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 0
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
