package day18

import (
	"fmt"
	"testing"
)

var vf = func(t *testing.T, got, want interface{}) {
	if got != want {
		t.Helper()
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestP1(t *testing.T) {
	// got := part1("sample.txt")
	// fmt.Printf("Got: %v\n", got)
	var got int

	got = calcLine("(9 + (7 * 5 + 9 + 2) * 7 * 7 * (7 * 9)) * (7 + 4 + (2 + 9 + 7 + 5 * 9) * 7) + 5 + 6 + 6 + 5")
	vf(t, got, 259091932)
	got = calcLine("1 + 2 * 3 + 4 * 5 + 6")
	vf(t, got, 71)
	got = calcLine("1 + (2 * 3) + (4 * (5 + 6))")
	vf(t, got, 51)
	got = calcLine("2 * 3 + (4 * 5)")
	vf(t, got, 26)
	got = calcLine("5 + (8 * 3 + 9 + 3 * 4 * 3)")
	vf(t, got, 437)
	got = calcLine("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")
	vf(t, got, 12240)
	got = calcLine("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")
	vf(t, got, 13632)
	got = calcLine("2 * (8 + 3 * 3 + (4 + 5)) + 7 * 2")
	vf(t, got, 182)
	got = calcLine("2 * (5 * 7 + (2 * 2 + 7) + (5 + 2 * 8 + 4 * 4) + 9) * (9 * (7 * 2) * (7 + 3 * 3 * 2 + 4) * 9 + (2 * 4 * 9) + 8) + 6 * (8 + 2 * 8)")
	vf(t, got, 3429363680)
	got = calcLine("6 * (6 + 6 * (2 + 2 + 6 * 9 + 4 * 3) + 4 * 4 * (3 * 2 * 8 * 2)) * 7 * 2")
	vf(t, got, 109283328)
	got = calcLine("(3 + 9 * 8 * 7 * 9) + 4 + 4 * 4 * (9 * 6 + (2 * 7 + 3 + 6 * 5 * 5) + 8 + 7) * (2 * (2 * 6) + 9 + (8 * 6 + 6 + 8))")
	vf(t, got, 1482024320)
	got = calcLine("1 + ((2 + 3) * 6 + 7) + 5 + 6 * 6")
	vf(t, got, 294)
	got = calcLine("((3 + 7 + 6 + 5) * (7 * 5 * 6 + 9) * (8 * 5 * 4 * 5) * 9 * 3) + ((9 * 7 * 5 + 8) * 6 + 7 + 6) + 5 + 6 * 6")
	vf(t, got, 596042172)
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 650217205854
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	var got int
	got = calcLineP2("1 + 2 * 3 + 4 * 5 + 6")
	vf(t, got, 231)
	got = calcLineP2("1 + (2 * 3) + (4 * (5 + 6))")
	vf(t, got, 51)
	got = calcLineP2("2 * 3 + (4 * 5)")
	vf(t, got, 46)
	got = calcLineP2("5 + (8 * 3 + 9 + 3 * 4 * 3)")
	vf(t, got, 1445)
	got = calcLineP2("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))")
	vf(t, got, 669060)
	got = calcLineP2("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2")
	vf(t, got, 23340)
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 20394514442037
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
