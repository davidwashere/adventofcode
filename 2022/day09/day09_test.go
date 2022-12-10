package day09

import (
	"aoc/util"
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 13
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 6266
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestTowardVec(t *testing.T) {
	p1 := util.NewPoint(-3, 6)
	p2 := util.NewPoint(-2, 8)

	v := p1.TowardVector(p2)

	if v.X != -1 || v.Y != -1 {
		t.Errorf("ya nope %+v", v)
	}

}

func TestP2(t *testing.T) {
	var got int
	var want int
	got = part2("sample.txt")
	want = 1
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}

	got = part2("sample2.txt")
	want = 36
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 2369
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
