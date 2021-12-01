package util

import (
	"testing"
)

func TestParseTokens(t *testing.T) {
	vf := func(got, want interface{}) {
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	}
	str := "departure location: 32-842 or 854-967"
	tokens := ParseTokens(str)

	vf(len(tokens.Words), 3)
	vf(tokens.Words[0], "departure")
	vf(tokens.Words[1], "location")
	vf(tokens.Words[2], "or")

	vf(len(tokens.Ints), 4)
	vf(tokens.Ints[0], 32)
	vf(tokens.Ints[1], 842)
	vf(tokens.Ints[2], 854)
	vf(tokens.Ints[3], 967)

	str = "-`~=!@#$%^&*_+',.?;:/|\\\"<>()[]{}"
	tokens = ParseTokens(str)
	vf(len(tokens.Symbols), 32)

	str = "mask = 0010011010X1000100X101011X10010X1010"
	tokens = ParseTokens(str)
	vf(len(tokens.Strs), 2)
	vf(tokens.Strs[0], "mask")
	vf(tokens.Strs[1], "0010011010X1000100X101011X10010X1010")

	str = "acc -17"
	tokens = ParseTokens(str)
	vf(tokens.Words[0], "acc")
	vf(tokens.Ints[0], 17)
	vf(tokens.Symbols[0], "-")

	str = "R90"
	tokens = ParseTokens(str)
	vf(tokens.Words[0], "R")
	vf(tokens.Ints[0], 90)

	str = "pid:827837505 byr:1976"
	tokens = ParseTokens(str)
	vf(tokens.Words[0], "pid")
	vf(tokens.Words[1], "byr")
	vf(tokens.Ints[0], 827837505)
	vf(tokens.Ints[1], 1976)
	vf(tokens.Strs[0], "pid")
	vf(tokens.Strs[1], "827837505")
	vf(tokens.Strs[2], "byr")
	vf(tokens.Strs[3], "1976")
}
