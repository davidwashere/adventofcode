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

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	folds := []fold{}
	grid := util.NewInfGrid()
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
	// grid.Dump()

	fold := folds[0]
	nGrid := util.NewInfGrid()
	nGrid.WithDefaultValue(".")
	grid.VisitAll2D(func(val interface{}, x, y int) {
		if val.(string) == "." {
			return
		}

		if fold.horizontal {
			if y < fold.pos {
				nGrid.Set(val, x, y)
			} else if y > fold.pos {
				nGrid.Set(val, x, grid.Height()-1-y)
			}
		} else {
			if x < fold.pos {
				nGrid.Set(val, x, y)
			} else if x > fold.pos {
				nGrid.Set(val, grid.Width()-1-x, y)
			}
		}
	})

	// nGrid.FlipH()
	// fmt.Println("")
	// nGrid.Dump()

	result := 0
	nGrid.VisitAll2D(func(val interface{}, x, y int) {
		if val.(string) == "#" {
			result++
		}
	})

	return result
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	folds := []fold{}
	grid := util.NewInfGrid()
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

	for _, fold := range folds {
		nGrid := util.NewInfGrid()
		nGrid.WithDefaultValue(".")

		if fold.horizontal {
			nGrid.SetExtents(0, 0, grid.Width()-1, fold.pos-1)
		} else {
			nGrid.SetExtents(0, 0, fold.pos-1, grid.Height()-1)
		}

		grid.VisitAll2D(func(val interface{}, x, y int) {
			if val.(string) == "." {
				return
			}

			if fold.horizontal {
				if y < fold.pos {
					nGrid.Set(val, x, y)
				} else if y > fold.pos {
					nGrid.Set(val, x, grid.Height()-1-y)
				}
			} else {
				if x < fold.pos {
					nGrid.Set(val, x, y)
				} else if x > fold.pos {
					nGrid.Set(val, grid.Width()-1-x, y)
				}
			}
		})

		grid = nGrid
	}

	fmt.Println("")
	grid.FlipH()
	grid.WithDefaultValue(" ")
	grid.Dump()

	return 0
}
