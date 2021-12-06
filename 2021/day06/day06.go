package day06

import (
	"aoc/util"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	ages := util.ParseInts(data[0])

	for days := 80; days > 0; days-- {
		initialLen := len(ages)
		for i := 0; i < initialLen; i++ {
			age := ages[i]

			if age == 0 {
				ages = append(ages, 8)
				ages[i] = 6
			} else {
				ages[i] = age - 1
			}
		}

	}

	return len(ages)
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	ages := util.ParseInts(data[0])

	// 0, 1, 2, 3, 4, 5, 6, 7, 8
	numAtAge := make([]int, 9)
	for _, age := range ages {
		numAtAge[age]++
	}

	for days := 256; days > 0; days-- {
		// zeros = # of new 6s and 8s
		zeros := numAtAge[0]

		numAtAge = numAtAge[1:]
		numAtAge[6] += zeros
		numAtAge = append(numAtAge, zeros)
	}

	sum := 0
	for _, num := range numAtAge {
		sum += num
	}

	return sum
}
