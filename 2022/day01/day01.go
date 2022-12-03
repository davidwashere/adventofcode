package day01

import (
	"aoc/util"
	"sort"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	elfs := []int{}
	elfs = append(elfs, 0)

	for _, line := range data {
		ints := util.ParseInts(line)
		if len(ints) == 0 {
			elfs = append(elfs, 0)
			continue
		}

		index := len(elfs) - 1
		elfs[index] += ints[0]
	}

	max := 0
	for _, cals := range elfs {
		max = util.Max(max, cals)
	}

	return max
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	elfs := []int{}
	elfs = append(elfs, 0)

	for _, line := range data {
		ints := util.ParseInts(line)
		if len(ints) == 0 {
			elfs = append(elfs, 0)
			continue
		}

		index := len(elfs) - 1
		elfs[index] += ints[0]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))

	return elfs[0] + elfs[1] + elfs[2]
}
