package day06

import (
	"aoc2020/util"
	"fmt"
	"strings"
)

func part1(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	grps := []map[string]int{}
	grp := map[string]int{}
	for _, line := range data {
		if strings.TrimSpace(line) == "" {
			grps = append(grps, grp)
			grp = map[string]int{}
		}

		for _, rune := range line {
			char := string(rune)
			val, ok := grp[char]
			if !ok {
				grp[char] = 1
			} else {
				grp[char] = val + 1
			}
		}
	}
	grps = append(grps, grp)
	fmt.Println(grps)

	result := 0
	for _, val := range grps {
		result += len(val)
	}

	return result
}

func part2(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	// grps := []map[string]int{}
	grp := map[string]int{}
	numfolks := 0
	result := 0
	for _, line := range data {
		if strings.TrimSpace(line) == "" {
			// grps = append(grps, grp)
			for _, val := range grp {
				if val == numfolks {
					result++
				}
			}

			grp = map[string]int{}
			numfolks = 0
			continue
		}

		numfolks++
		for _, rune := range line {
			char := string(rune)
			val, ok := grp[char]
			if !ok {
				grp[char] = 1
			} else {
				grp[char] = val + 1
			}
		}
	}

	for _, val := range grp {
		if val == numfolks {
			result++
		}
	}

	return result
}

func part1Optimize(inputfile string) int {
	data, err := util.ReadFileToStringSliceWithDelim(inputfile, "\n\n")
	util.Check(err)

	result := 0
	for _, grp := range data {
		result += util.CountUniqueChars(grp)
	}

	return result
}

func part2Optimize(inputfile string) int {
	data, err := util.ReadFileToStringSliceWithDelim(inputfile, "\n\n")
	util.Check(err)

	result := 0
	for _, grp := range data {
		chars := util.MapUniqueChars(grp)
		folks := len(strings.Split(grp, "\n"))

		for _, count := range chars {
			if count == folks {
				result++
			}
		}
	}

	return result
}
