package day03

import (
	"aoc/util"
	"fmt"
	"strconv"
)

var (
	digits = [][]int{}
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	for _, line := range data {
		row := []int{}
		for _, ds := range line {
			d, err := strconv.Atoi(string(ds))
			if err != nil {
				panic(err)
			}

			row = append(row, d)
		}

		digits = append(digits, row)
	}

	// fmt.Printf("%v\n", digits)
}

func part1(inputFile string) int {
	load(inputFile)

	result := 0
	for _, bank := range digits {
		top := 0
		for i := 0; i < len(bank)-1; i++ {
			l := bank[i] * 10
			for j := i + 1; j < len(bank); j++ {
				r := bank[j]
				top = max(top, l+r)
			}
		}

		result += top
	}

	return result
}

func part2(inputFile string) int {
	load(inputFile)

	numDigitsToSum := 12

	result := 0
	for _, bank := range digits {
		top := 0

		highest := 0
		highestIndex := -1
		startAt := -1
		for k := 0; k < numDigitsToSum; k++ {
			for i := startAt + 1; i <= len(bank)-numDigitsToSum+k; i++ {
				cur := bank[i]
				if cur > highest {
					highest = cur
					highestIndex = i
				}
			}
			startAt = highestIndex
			top = (top * 10) + highest
			highest = -1
			highestIndex = -1
		}

		fmt.Printf("%d\n", top)
		result += top
	}

	return result
}
