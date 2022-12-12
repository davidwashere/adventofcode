package day13

import (
	"aoc/util"
	"fmt"
	"strings"
)

type fold struct {
	horizontal bool
	pos        int
}

// Grid default value
var Def = "."

func loadData(inputfile string) (*util.InfGrid[string], []fold) {
	data, _ := util.ReadFileToStringSlice(inputfile)

	folds := []fold{}
	grid := util.NewInfGrid[string]()
	grid.WithDefaultValue(".")

	for _, line := range data {
		if len(line) == 0 {
			continue
		}

		tokens := util.ParseTokens(line)
		if strings.Contains(line, "fold") {
			fold := fold{}
			fold.pos = tokens.Ints[0]
			if strings.Contains(line, "y") {
				fold.horizontal = true
			}
			folds = append(folds, fold)
			continue
		}

		grid.Set("#", tokens.Ints[0], tokens.Ints[1])
	}

	grid.FlipH()
	return grid, folds
}

func part1(inputfile string) int {
	grid, folds := loadData(inputfile)

	fold := folds[0]
	nGrid := util.NewInfGrid[string]()
	nGrid.WithDefaultValue(Def)

	grid.VisitAll2D(func(val string, x, y int) {
		if val == Def {
			return
		}

		if fold.horizontal && y > fold.pos {
			y = grid.Height() - 1 - y
		} else if !fold.horizontal && x > fold.pos {
			x = grid.Width() - 1 - x
		}

		nGrid.Set(val, x, y)
	})

	result := 0
	nGrid.VisitAll2D(func(val string, x, y int) {
		if val == "#" {
			result++
		}
	})

	return result
}

func part2(inputfile string) int {
	grid, folds := loadData(inputfile)

	for _, fold := range folds {
		nGrid := util.NewInfGrid[string]()
		nGrid.WithDefaultValue(Def)

		if fold.horizontal {
			nGrid.SetExtents(0, 0, grid.Width()-1, fold.pos-1)
		} else {
			nGrid.SetExtents(0, 0, fold.pos-1, grid.Height()-1)
		}

		grid.VisitAll2D(func(val string, x, y int) {
			if val == Def {
				return
			}

			if fold.horizontal && y > fold.pos {
				y = grid.Height() - 1 - y
			} else if !fold.horizontal && x > fold.pos {
				x = grid.Width() - 1 - x
			}

			nGrid.Set(val, x, y)
		})

		grid = nGrid
	}

	grid.FlipH()
	grid.WithDefaultValue(" ")
	fmt.Println()
	grid.Dump()
	fmt.Println()

	return 0
}
