package day07

import (
	"aoc2020/util"
	"fmt"
	"testing"
)

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 4
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 238
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 32
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 82930
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

// func TestOptimize(t *testing.T) {
// 	filename := "sample.txt"
// 	got := part1Optimize(filename)
// 	fmt.Printf("Got: %v\n", got)

// 	got = part2Optimize(filename)
// 	fmt.Printf("Got: %v\n", got)
// }

// func TestOptimize_Actual(t *testing.T) {
// 	filename := "input.txt"
// 	got := part1Optimize(filename)
// 	fmt.Printf("Got: %v\n", got)

// 	got = part2Optimize(filename)
// 	fmt.Printf("Got: %v\n", got)
// }

func TestSetup(t *testing.T) {
	var err error

	err = util.ConvertFromCRLFtoLF("sample.txt")
	util.Check(err)

	err = util.ConvertFromCRLFtoLF("input.txt")
	util.Check(err)
}
