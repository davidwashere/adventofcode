package day02

import (
	"aoc2020/util"
	"fmt"
	"strconv"
	"strings"
)

func part1(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	found := 0
	for _, line := range data {
		// 1-3 a: abcde
		rulePassSplit := strings.Split(line, ":")
		ruleSplit := rulePassSplit[0]                         // 1-3 a
		pass := strings.TrimSpace(rulePassSplit[1])           // abcde
		boundsCharSplit := strings.Split(ruleSplit, " ")      // 1-3 a
		char := boundsCharSplit[1]                            // a
		boundsSplit := strings.Split(boundsCharSplit[0], "-") // 1-3
		min, _ := strconv.Atoi(boundsSplit[0])
		max, _ := strconv.Atoi(boundsSplit[1])

		count := strings.Count(pass, char)

		valid := false
		if count <= max && count >= min {
			valid = true
			found++
		}

		fmt.Printf("[%v] [%v] [%v] [%v] [%v] = %v\n", min, max, char, pass, count, valid)
	}

	return found
}

func part2(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	found := 0
	for _, line := range data {
		// 1-3 a: abcde
		rulePassSplit := strings.Split(line, ":")
		ruleSplit := rulePassSplit[0]                         // 1-3 a
		pass := strings.TrimSpace(rulePassSplit[1])           // abcde
		boundsCharSplit := strings.Split(ruleSplit, " ")      // 1-3 a
		char := boundsCharSplit[1]                            // a
		boundsSplit := strings.Split(boundsCharSplit[0], "-") // 1-3
		pos1, _ := strconv.Atoi(boundsSplit[0])
		pos2, _ := strconv.Atoi(boundsSplit[1])

		left := false
		right := false
		if string(pass[pos1-1]) == char {
			left = true
		}

		if string(pass[pos2-1]) == char {
			right = true
		}

		if left && right || (!left && !right) {
			continue
		}
		found++

		fmt.Printf("[%v] [%v] [%v] [%v] = valid\n", pos1, pos2, char, pass)
	}

	return found
}
