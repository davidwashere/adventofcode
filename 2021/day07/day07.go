package day07

import (
	"aoc/util"
)

func gimmieCrabsMinAndMax(inputfile string) (crabs map[int]int, min int, max int) {
	data, _ := util.ReadFileToStringSlice(inputfile)
	ints := util.ParseInts(data[0])

	min = util.MaxInt
	max = util.MinInt

	// pos -> num crabs
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

func crawlCrabbiesCrawl(crabs map[int]int, min int, max int, calc func(dist, numCrabs int) int) int {
	minSum := util.MaxInt
	for i := min; i <= max; i++ {
		sum := 0
		for pos, numCrabs := range crabs {
			left := util.Min(i, pos)
			right := util.Max(i, pos)

			dist := right - left

			sum += calc(dist, numCrabs)
		}

		minSum = util.Min(sum, minSum)
	}

	return minSum
}

func part1(inputfile string) int {
	crabs, min, max := gimmieCrabsMinAndMax(inputfile)

	minSum := crawlCrabbiesCrawl(crabs, min, max, func(dist, numCrabs int) int {
		return dist * numCrabs
	})

	return minSum
}

func part2(inputfile string) int {
	crabs, min, max := gimmieCrabsMinAndMax(inputfile)

	minSum := crawlCrabbiesCrawl(crabs, min, max, func(dist, numCrabs int) int {
		fuel := util.SumArithmeticProgression(dist, 1, dist)
		return fuel * numCrabs
	})

	return minSum
}
