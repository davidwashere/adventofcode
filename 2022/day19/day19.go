package day19

import (
	"aoc/util"
	"fmt"
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

		fmt.Println(tokens)
	}
}

func part1(inputFile string) int {
	load(inputFile)

	return 0
}

func part2(inputFile string) int {
	load(inputFile)

	return 0
}