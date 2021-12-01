package day01

import (
	"aoc/util"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToIntSlice(inputfile)

	result := 0
	last := data[0]

	for i := 1; i < len(data); i++ {
		cur := data[i]

		if cur > last {
			result++
		}

		last = cur
	}

	return result
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToIntSlice(inputfile)
	result := 0

	last := data[0] + data[1] + data[2]

	for i := 3; i < len(data); i++ {
		cur := data[i] + data[i-1] + data[i-2]

		if cur > last {
			result++
		}

		last = cur
	}

	return result
}
