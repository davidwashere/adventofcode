package day04

import (
	"aoc/util"
	"fmt"
)

var (
	grid = util.NewInfGrid[string]().WithDefaultValue(".")
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	// data, _ := util.ReadFileToIntSlice(inputFile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputFile, "\n\n")
	// grid := util.NewInfGrid[string]().WithDefaultValue(".")

	for row, line := range data {
		for col, c := range line {
			grid.Set(string(c), col, row)
		}
	}

	grid.FlipH()
	grid.Dump()
}

func part1(inputFile string) int {
	load(inputFile)

	rolls := 0
	grid.VisitAll(func(val string, x, y int, dims ...int) {
		if val != "@" {
			return
		}

		adjacentRolls := 0
		grid.VisitOrthoAndDiag(x, y, func(val string, x, y int) {
			if val == "@" {
				adjacentRolls++
			}
		})
		if adjacentRolls < 4 {
			rolls++
		}
	})

	return rolls
}

func part2(inputFile string) int {
	load(inputFile)

	rolls := 0
	for {
		pointsToRemove := []util.Point{}
		grid.VisitAll(func(val string, x, y int, dims ...int) {
			if val != "@" {
				return
			}

			adjacentRolls := 0
			grid.VisitOrthoAndDiag(x, y, func(val string, x, y int) {
				if val == "@" {
					adjacentRolls++
				}
			})
			if adjacentRolls < 4 {
				rolls++
				pointsToRemove = append(pointsToRemove, util.NewPoint(x, y))
			}
		})
		if len(pointsToRemove) == 0 {
			break
		}

		for _, p := range pointsToRemove {
			grid.Delete(p.X, p.Y)
		}
		fmt.Println("===")
		grid.Dump()
	}

	return rolls
}
