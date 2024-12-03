package day01

import (
	"aoc/util"
	"fmt"
	"sort"
)

var (
	l1 []int
	l2 []int
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	// data, _ := util.ReadFileToIntSlice(inputFile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputFile, "\n\n")
	// grid := util.NewInfinityGridFromFile(inputfile, ".")

	for _, line := range data {
		tokens := util.ParseTokens(line)
		// ints := util.ParseInts(line)
		// strs := util.ParseStrs(line)
		// words := util.ParseWords(line)
		l1 = append(l1, tokens.Ints[0])
		l2 = append(l2, tokens.Ints[1])

		// fmt.Println(tokens)
	}

	sort.Ints(l1)
	sort.Ints(l2)

	if len(l1) != len(l2) {
		panic(fmt.Sprintf("input not same len %v != %v", len(l1), len(l2)))
	}
	// fmt.Println(l1)
	// fmt.Println(l2)
}

func part1(inputFile string) int {
	load(inputFile)

	sum := 0
	for i, l := range l1 {
		r := l2[i]

		hi := max(l, r)
		lo := min(l, r)

		sum += hi - lo
	}

	return sum
}

func part2(inputFile string) int {
	load(inputFile)

	counts := map[int]int{}
	for _, r := range l2 {
		counts[r]++
	}

	sum := 0
	for _, l := range l1 {
		sum += l * counts[l]
	}

	return sum
}
