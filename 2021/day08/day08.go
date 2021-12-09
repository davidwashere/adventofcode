package day08

import (
	"aoc/util"
	"sort"
	"strconv"
	"strings"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	result := 0
	for _, line := range data {
		allwords := util.ParseWords(line)

		words := allwords[len(allwords)-4:]

		for _, word := range words {
			if len(word) == 2 || len(word) == 3 || len(word) == 4 || len(word) == 7 {
				result++
			}
		}
	}

	return result
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	result := 0
	digitMap := map[int]string{}
	for _, line := range data {
		allwords := util.ParseWords(line)

		// for i, word := range allwords {
		// 	word := SortString(word)
		// 	allwords[i] = word
		// 	fmt.Println(word)
		// }

		pattern := allwords[:10]
		signal := allwords[10:]

		// length to #
		// 2 = #1
		// 3 = #7
		// 4 = #4
		// 7 = #8

		// solve for known 1, 7, 4, 8
		rem := []string{}
		for _, word := range pattern {
			switch len(word) {
			case 2:
				digitMap[1] = word
			case 3:
				digitMap[7] = word
			case 4:
				digitMap[4] = word
			case 7:
				digitMap[8] = word
			default:
				rem = append(rem, word)
			}
		}

		// solve for 3 and 6
		// a 3 is the only 5 digit word with the two letters from #1
		// a 6 is the only 6 digit word without the two letters from #1
		rem2 := []string{}
		for _, word := range rem {
			switch len(word) {
			case 5:
				if AllLettersIn(digitMap[1], word) {
					// must be 3
					digitMap[3] = word
				} else {
					rem2 = append(rem2, word)
				}

			case 6:
				if !AllLettersIn(digitMap[1], word) {
					// must be 6
					digitMap[6] = word
				} else {
					rem2 = append(rem2, word)
				}
			default:
				rem2 = append(rem2, word)
			}
		}

		letsToFind := RemoveLetters(digitMap[1], digitMap[4])
		// a 9 is only 6 digits word left with letsToFind in it
		// the remaining 6 digit word is a 0
		rem = []string{}
		for _, word := range rem2 {
			switch len(word) {
			case 6:
				if AllLettersIn(letsToFind, word) {
					// must be 9
					digitMap[9] = word
				} else {
					// must be 0
					digitMap[0] = word
				}
			default:
				rem = append(rem, word)
			}
		}

		// a 5 is the remaining 5 digit word containing only chars also in 9
		// a 2 has a letter that is not found in 9
		if AllLettersIn(rem[0], digitMap[9]) {
			digitMap[5] = rem[0]
			digitMap[2] = rem[1]
		} else {
			digitMap[5] = rem[1]
			digitMap[2] = rem[0]
		}

		// reverse map and make value a string
		rMap := map[string]string{}

		for k, v := range digitMap {
			v := SortString(v)
			rMap[v] = strconv.Itoa(k)
		}

		final := ""
		for _, word := range signal {
			word := SortString(word)
			final += rMap[word]
		}

		temp, _ := strconv.Atoi(final)
		result += temp
	}

	return result
}

// AllLettersIn will try to find the characters from 'search' in 'sample'
func AllLettersIn(search string, sample string) bool {
	searchM := map[rune]bool{}
	for _, char := range search {
		searchM[char] = false
	}

	for _, char := range sample {
		delete(searchM, char)
	}

	return len(searchM) == 0
}

// RemoveLetters will remove the chars in 'letters' from 'from'
func RemoveLetters(letters, from string) string {
	searchM := map[rune]bool{}
	for _, char := range letters {
		searchM[char] = false
	}

	result := ""
	for _, char := range from {
		if _, ok := searchM[char]; !ok {
			result += string(char)
		}
	}

	return result
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
