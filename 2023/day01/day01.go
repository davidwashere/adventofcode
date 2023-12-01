package day01

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

var digits = []string{
	"0",
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
}

var words = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func extents(line string, lists ...[]string) (int, int) {
	lowI := util.MaxInt
	lowVal := 0

	highI := util.MinInt
	highVal := 0

	for _, list := range lists {
		for val, item := range list {
			curI := strings.Index(line, item)
			if curI == -1 {
				continue
			}

			if curI < lowI {
				lowVal = val
				lowI = curI
			}

			curI = strings.LastIndex(line, item)
			if curI > highI {
				highVal = val
				highI = curI
			}
		}
	}

	return lowVal, highVal
}

func part1(inputFile string) int {
	data, _ := util.ReadFileToStringSlice(inputFile)

	sum := 0
	for _, line := range data {
		l, h := extents(line, digits)

		val, _ := strconv.Atoi(fmt.Sprintf("%v%v", l, h))

		sum += val
	}
	return sum
}

func part2(inputFile string) int {
	data, _ := util.ReadFileToStringSlice(inputFile)

	sum := 0
	for _, line := range data {
		l, h := extents(line, digits, words)

		val, _ := strconv.Atoi(fmt.Sprintf("%v%v", l, h))

		sum += val
	}
	return sum
}
