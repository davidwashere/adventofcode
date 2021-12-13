package day11

import (
	"aoc/util"
	"fmt"
	"strconv"
)

func loadGrid(inputfile string) *util.InfGrid {
	data, _ := util.ReadFileToStringSlice(inputfile)

	grid := util.NewInfGrid()

	y := 0
	for _, line := range data {

		x := 0
		for _, c := range line {
			n, _ := strconv.Atoi(string(c))
			grid.Set(n, x, y)
			x++
		}
		y++
	}

	grid.FlipH()

	return grid
}

type node struct {
	x int
	y int
}

func part1(inputfile string) int {
	grid := loadGrid(inputfile)

	grid.Dump()

	// increase all by 1
	grid.VisitAll2D(func(val interface{}, x, y int) {
		n := val.(int)

		n++
		if n > 9 {
			n = 0
		}

		grid.Set(n, x, y)
	})

	fmt.Println()
	grid.Dump()
	result := 0

	return result
}

func part2(inputfile string) int {
	return 0
}
