package day06

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	m := map[byte]int{}
	add(m, 'a')
	add(m, 'a')
	add(m, 'b')

	if len(m) != 2 {
		t.Errorf("Wrong len")
	}

	sub(m, 'a')
	if len(m) != 2 {
		t.Errorf("Wrong len")
	}

	sub(m, 'a')
	if len(m) != 1 {
		t.Errorf("Wrong len")
	}

	sub(m, 'b')
	if len(m) != 0 {
		t.Errorf("Wrong len")
	}

}

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 7
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 1140
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 19
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 3495
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
