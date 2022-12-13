package day12

import (
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 31
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 520
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 29
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 508
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestNextLetter(t *testing.T) {

	tt := []struct {
		val  string
		want string
	}{
		{"S", "a"},
		{"a", "b"},
		{"z", "E"},
	}

	for _, test := range tt {
		got := nextLetter(test.val)

		if got != test.want {
			t.Errorf("got %v want %v", got, test.want)
		}
	}

}
