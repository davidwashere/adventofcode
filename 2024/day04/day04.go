package day04

import (
	"aoc/util"
	"strings"
)

var (
	grid *util.InfGrid[string]
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	grid = util.NewInfGrid[string]()

	for y, line := range data {
		for x := 0; x < len(line); x++ {
			c := string(line[x])
			grid.Set(c, x, y)
		}
	}
}

func part1(inputFile string) int {
	load(inputFile)

	count := 0
	grid.VisitAll(func(val string, x, y int, dims ...int) {
		if val != "X" {
			return
		}

		fs := []func(x, y int, count int, dims ...int) []string{
			grid.GetNMany,
			grid.GetEMany,
			grid.GetSMany,
			grid.GetWMany,
			grid.GetNEMany,
			grid.GetSEMany,
			grid.GetSWMany,
			grid.GetNWMany,
		}

		for _, f := range fs {
			if strings.Join(f(x, y, 3), "") == "MAS" {
				count++
			}
		}
	})

	return count
}

func part2(inputFile string) int {
	load(inputFile)

	totCount := 0
	grid.VisitAll(func(val string, x, y int, dims ...int) {
		if val != "A" {
			return
		}

		count := 0
		if grid.Get(x-1, y-1) == "M" && grid.Get(x+1, y+1) == "S" {
			count++
		}

		if grid.Get(x-1, y-1) == "S" && grid.Get(x+1, y+1) == "M" {
			count++
		}

		if grid.Get(x-1, y+1) == "M" && grid.Get(x+1, y-1) == "S" {
			count++
		}

		if grid.Get(x-1, y+1) == "S" && grid.Get(x+1, y-1) == "M" {
			count++
		}

		if count == 2 {
			totCount++
		}

	})

	return totCount
}
