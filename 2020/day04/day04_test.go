package day04

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	part1("sample.txt")
}

func TestPart1SampleInput(t *testing.T) {
	got := part1("sample.txt")
	want := 2
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
	got := part2("sample.valid.txt")
	t.Logf("Got: %v", got)

	// if want != got {
	// 	t.Errorf("Expected %v but got %v", want, got)
	// }
}

func TestPart2(t *testing.T) {
	got := part2("input.txt")
	t.Logf("Got: %v", got)
}

func TestPoint(t *testing.T) {

	type Passport struct {
		fields map[string]string
		valid  bool
	}

	p := Passport{}
	p.fields = map[string]string{}
	p.fields["hi"] = "world"

	passports := []Passport{}
	passports = append(passports, p)

	for _, item := range passports {
		item.fields["good"] = "bye" // Works because map is pointer
		item.valid = true           // Doesn't work bool is not pointer
	}

	fmt.Printf("%+v\n", passports[0])

	passports[0].valid = true // works because direct ref
	fmt.Printf("%+v\n", passports[0])

	passports2 := []*Passport{}
	p = Passport{}
	p.fields = map[string]string{}
	p.fields["sup"] = "yo"

	passports2 = append(passports2, &p)

	for _, item := range passports2 {
		item.fields["still"] = "works" // Works because map is pointer
		item.valid = true              // Works
	}

	fmt.Printf("%+v\n", passports2[0])
}
