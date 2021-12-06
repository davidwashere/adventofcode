package day06

import (
	"aoc/util"
)

func runSim(inputfile string, days int) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	ages := util.ParseInts(data[0])

	// num of fish at particular age, length never exceeds 9 (0-8)
	numAtAge := make([]int, 9)
	for _, age := range ages {
		numAtAge[age]++
	}

	for day := days; day > 0; day-- {
		zeros := numAtAge[0]

		numAtAge = numAtAge[1:]            // shift array dropping the fish at age 0
		numAtAge[6] += zeros               // the zeros now become 6's
		numAtAge = append(numAtAge, zeros) // is also new fish at age 8 for each zero dropped
	}

	sum := 0
	for _, num := range numAtAge {
		sum += num
	}

	return sum
}

func part1(inputfile string) int {
	return runSim(inputfile, 80)
}

func part2(inputfile string) int {
	return runSim(inputfile, 256)
}
