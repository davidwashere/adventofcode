package day07

import (
	"aoc/util"
	"fmt"
)

func gimmieCrabsMinAndMax(inputfile string) (crabs map[int]int, min int, max int) {
	data, _ := util.ReadFileToStringSlice(inputfile)
	ints := util.ParseInts(data[0])

	min = util.MaxInt
	max = util.MinInt

	// crab pos -> num crabs at pos
	crabs = map[int]int{}

	for _, i := range ints {
		min = util.Min(i, min)
		max = util.Max(i, max)

		if _, ok := crabs[i]; !ok {
			crabs[i] = 0
		}

		crabs[i] += 1
	}

	return
}

func part1(inputfile string) int {
	crabs, min, max := gimmieCrabsMinAndMax(inputfile)

	minSumPos := -1
	minSum := util.MaxInt
	for i := min; i <= max; i++ {
		sum := 0
		for pos, numCrabs := range crabs {
			left := util.Min(i, pos)
			right := util.Max(i, pos)
			sum += (right - left) * numCrabs
		}

		if sum < minSum {
			minSum = sum
			minSumPos = i
		}
	}

	fmt.Printf("Pos: %v with sum %v\n", minSumPos, minSum)

	return minSum
}

func part2(inputfile string) int {
	crabs, min, max := gimmieCrabsMinAndMax(inputfile)

	minSumPos := -1
	minSum := util.MaxInt
	for i := min; i <= max; i++ {
		sum := 0
		for pos, numCrabs := range crabs {
			left := util.Min(i, pos)
			right := util.Max(i, pos)

			dist := right - left
			fuel := (dist * (dist + 1)) / 2
			sum += fuel * numCrabs
		}

		if sum < minSum {
			minSum = sum
			minSumPos = i
		}
	}

	fmt.Printf("Pos: %v with sum %v\n", minSumPos, minSum)

	return minSum
}
