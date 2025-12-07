package day07

import (
	"aoc/util"
)

var (
	grid = util.NewInfGrid[string]().WithDefaultValue(".")
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	for row, line := range data {
		for col, c := range line {
			grid.Set(string(c), col, row)
		}
	}

	grid.Dump()
}

func part1(inputFile string) int {
	load(inputFile)

	result := 0
	grid.VisitAll(func(val string, x, y int, dims ...int) {
		above := grid.Get(x, y-1)
		c := grid.Get(x, y)
		if c != "^" && (above == "S" || above == "|") {
			grid.Set("|", x, y)
			return
		}

		if c == "^" && above == "|" {
			grid.Set("|", x-1, y)
			grid.Set("|", x+1, y)
			result++
		}
	})

	grid.Dump()

	return result
}

func part2(inputFile string) int {
	load(inputFile)

	counts := util.NewInfGrid[int]().WithDefaultValue(0)

	grid.VisitAll(func(val string, x, y int, dims ...int) {
		above := grid.Get(x, y-1)
		abovecount := counts.Get(x, y-1)
		c := grid.Get(x, y)
		if c != "^" && (above == "S" || above == "|") {
			switch above {
			case "S":
				grid.Set("|", x, y)
				counts.Set(1, x, y)
			case "|":
				grid.Set("|", x, y)
				counts.Set(counts.Get(x, y)+abovecount, x, y)
			}
			return
		}

		if c == "^" && above == "|" {
			counts.Set(counts.Get(x-1, y)+abovecount, x-1, y)
			counts.Set(counts.Get(x+1, y)+abovecount, x+1, y)
			grid.Set("|", x-1, y)
			grid.Set("|", x+1, y)
		}
	})

	// counts.WithDumpFunc(func(val int, freshRow bool) {
	// 	if freshRow {
	// 		fmt.Println("")
	// 		return
	// 	}
	// 	if val == 0 {
	// 		fmt.Printf(" ")
	// 		return
	// 	}

	// 	fmt.Printf("%d", val)
	// })
	// counts.Dump()

	sum := 0
	for _, v := range counts.GetRow(counts.Height() - 1) {
		sum += v
	}

	return sum
}
