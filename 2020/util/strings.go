package util

import "unicode"

// MapUniqueChars will create a map for each character found
// in the given string and the count of each string occurance
// ignoring whitespace
func MapUniqueChars(inputString string) map[rune]int {
	data := map[rune]int{}
	for _, char := range inputString {

		// Ignore line breaks
		if unicode.IsSpace(char) {
			continue
		}

		val, ok := data[char]
		if !ok {
			data[char] = 1
		} else {
			data[char] = val + 1
		}
	}
	return data
}

// CountUniqueChars will return the number of unique characters
// found in the provided string, ignoring whitespace chars
func CountUniqueChars(inputString string) int {
	data := MapUniqueChars(inputString)
	return len(data)
}
