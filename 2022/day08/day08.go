package day08

import (
	"aoc/util"
)

func loadGrid(inputfile string) *util.InfGrid[int] {
	data, _ := util.ReadFileToStringSlice(inputfile)

	grid := util.NewInfGrid[int]()

	for y, line := range data {
		for x := 0; x < len(line); x++ {
			c := line[x]
			c = c - '0'
			grid.Set(int(c), x, y)
		}
	}

	return grid
}

func part1(inputfile string) int {
	grid := loadGrid(inputfile)
	grid.WithDefaultValue(util.MaxInt)

	result := 0
	grid.VisitAll2D(func(vR int, x, y int) {
		if x == grid.GetMinX() || x == grid.GetMaxX() || y == grid.GetMinY() || y == grid.GetMaxY() {
			result++
			return
		}

		if allLower(grid, x, y, grid.VisitN2D) ||
			allLower(grid, x, y, grid.VisitS2D) ||
			allLower(grid, x, y, grid.VisitE2D) ||
			allLower(grid, x, y, grid.VisitW2D) {
			result++

		}
	})

	return result
}

func allLower(grid *util.InfGrid[int], x, y int, visitFunc func(x, y int, visitFunc func(val int, x int, y int) bool)) bool {
	base := grid.Get(x, y)

	result := true
	visitFunc(x, y, func(v int, x, y int) bool {
		if v >= base {
			result = false
			return false
		}
		return true
	})

	return result
}

func part2(inputfile string) int {
	grid := loadGrid(inputfile)
	grid.WithDefaultValue(util.MaxInt)

	result := util.MinInt
	grid.VisitAll2D(func(vR int, x, y int) {
		n := calcDist(grid, x, y, grid.VisitN2D)
		s := calcDist(grid, x, y, grid.VisitS2D)
		e := calcDist(grid, x, y, grid.VisitE2D)
		w := calcDist(grid, x, y, grid.VisitW2D)

		score := n * s * e * w
		result = util.Max(result, score)
	})

	return result
}

func calcDist(grid *util.InfGrid[int], x, y int, visitFunc func(x, y int, visitFunc func(val int, x int, y int) bool)) int {
	base := grid.Get(x, y)
	dist := 0
	visitFunc(x, y, func(v int, x, y int) bool {
		dist++
		return v < base
	})

	return dist
}
