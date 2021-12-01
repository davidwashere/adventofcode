package day05

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	got := part1("sample.txt")
	//FBFBBFFRLR
	fmt.Println(got)

	// F = 6
	// B = 5
	// F = 4
	// B = 3
	// B = 2
	// F = 1
	// F = 0

	num := 0
	num += (1 << 5)
	num += (1 << 3)
	num += (1 << 2)
	fmt.Println("Row: ", num)

	num = 0
	num += (1 << 2)
	num += (1 << 0)
	fmt.Println("Col: ", num)
}

func TestPart1SampleInput(t *testing.T) {
	got := part1("sample.txt")
	want := 357
	t.Logf("Got: %v", got)

	if want != got {
		t.Errorf("Expected %v but got %v", want, got)
	}
}

func TestPart1(t *testing.T) {
	got := part1("input.txt")
	t.Logf("Got: %v", got)
}

func TestPart2SampleInput(t *testing.T) {
	got := part2("sample.txt")
	t.Logf("Got: %v", got)

	// if want != got {
	// 	t.Errorf("Expected %v but got %v", want, got)
	// }
}

func TestPart2(t *testing.T) {
	got := part2("input.txt")
	t.Logf("Got: %v", got)
}
