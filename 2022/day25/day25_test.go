package day25

import (
	"fmt"
	"testing"
)

func TestSnafuToDec(t *testing.T) {
	tt := []struct {
		in   string
		want int
	}{
		{"1=-0-2", 1747},
		{"12111", 906},
		{"2=0=", 198},
		{"21", 11},
		{"2=01", 201},
		{"111", 31},
		{"20012", 1257},
		{"112", 32},
		{"1=-1=", 353},
		{"1-12", 107},
		{"12", 7},
		{"1=", 3},
		{"122", 37},
	}

	for _, tc := range tt {
		got := snafuToDec(tc.in)
		if got != tc.want {
			t.Errorf("got %v want %v", got, tc.want)
		}
	}
}

func TestDecToSnafu(t *testing.T) {
	tt := []struct {
		in   int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{2, "2"},
		{3, "1="},
		{4, "1-"},
		{5, "10"},
		{1747, "1=-0-2"},
		{906, "12111"},
		{198, "2=0="},
		{11, "21"},
		{201, "2=01"},
		{31, "111"},
		{1257, "20012"},
		{32, "112"},
		{353, "1=-1="},
		{107, "1-12"},
		{7, "12"},
		{3, "1="},
		{37, "122"},
	}

	for _, tc := range tt {
		got := decToSnafu(tc.in)
		if got != tc.want {
			t.Errorf("got %v want %v for in %v", got, tc.want, tc.in)
		}
	}
}

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := "2=-1=0"
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := "20-1-0=-2=-2220=0011"
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := ""
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := ""
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
