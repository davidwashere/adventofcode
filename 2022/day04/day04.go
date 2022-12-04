package day04

import (
	"aoc/util"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	result := 0
	for _, line := range data {
		ints := util.ParseInts(line)

		r1 := []int{ints[0], ints[1]}
		r2 := []int{ints[2], ints[3]}

		if util.RangeFullyContains(r1, r2) {
			result++
		}
	}

	return result
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	result := 0
	for _, line := range data {
		ints := util.ParseInts(line)

		r1 := []int{ints[0], ints[1]}
		r2 := []int{ints[2], ints[3]}

		if util.RangeOverlaps(r1, r2) {
			result++
		}
	}

	return result
}
