package util

import (
	"regexp"
	"strconv"
	"strings"
)

// Tokens represents the results of parsing a string
// See parse methods for description of what each field holds
type Tokens struct {
	Strs    []string
	Ints    []int
	Words   []string
	Symbols []string
}

// ParseTokens will parse string and int values out from `input`
func ParseTokens(input string) Tokens {
	tokens := Tokens{}

	tokens.Strs = ParseStrs(input)
	tokens.Words = ParseWords(input)
	tokens.Ints = ParseInts(input)
	tokens.Symbols = ParseSymbols(input)

	return tokens
}

func _tokenParseWithRegex(regex, input string) []string {
	re := regexp.MustCompile(regex)
	res := []string{}
	for _, val := range re.FindAllString(input, -1) {
		val = strings.TrimSpace(val)
		if val != "" {
			res = append(res, val)
		}
	}

	return res
}

// ParseWords will extract alpha words (no nums) from the input
func ParseWords(input string) []string {
	return _tokenParseWithRegex("[a-zA-Z]+", input)
}

// ParseStrs will extract alphanumeric strings from the input
func ParseStrs(input string) []string {
	return _tokenParseWithRegex("[a-zA-Z0-9]+", input)
}

// ParseSymbols will extract symbols from the input (such as `~` and `:`)
func ParseSymbols(input string) []string {
	return _tokenParseWithRegex("[-`~=!@#$%^&*_+',.?;:/|\\\\\"<>()[\\]{}]?", input)
}

// ParseInts will extract numbers from the input
func ParseInts(input string) []int {
	intRe := regexp.MustCompile("[0-9]+")
	res := []int{}
	for _, valS := range intRe.FindAllString(input, -1) {
		val, _ := strconv.Atoi(valS)

		res = append(res, val)
	}

	return res
}
