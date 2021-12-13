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

type Coord struct {
	x int
	y int
}

var EMPTY = Coord{x: -1, y: -1}

func NewCoord(x, y int) Coord {
	return Coord{
		x: x,
		y: y,
	}
}

func popItemFromMap(m map[Coord]bool) (Coord, error) {
	for k := range m {
		c := k
		delete(m, k)
		return c, nil
	}

	return EMPTY, fmt.Errorf("Map Empty")
}

func part1(inputfile string) int {
	grid := loadGrid(inputfile)

	numFlashes := 0
	steps := 100
	for steps > 0 {
		steps--
		pending := map[Coord]bool{}

		// increase all by 1
		grid.VisitAll2D(func(val interface{}, x, y int) {
			n := val.(int)

			n++
			if n > 9 {
				pending[NewCoord(x, y)] = true
			}

			grid.Set(n, x, y)
		})

		// doo flashes
		flashed := map[Coord]bool{}
		for len(pending) > 0 {
			c, _ := popItemFromMap(pending)

			grid.VisitOrthoAndDiag(c.x, c.y, func(val interface{}, x, y int) {
				// if val == nil, that means is outside bounds of grid
				if val == nil {
					return
				}

				curC := NewCoord(x, y)

				_, f := flashed[curC]
				_, p := pending[curC]

				// if this coord is pending flash or has flashed
				// ignore it cause doesn't matter if new value inceased (will be reset to 0)
				if f || p {
					return
				}

				n := val.(int)
				n++
				if n > 9 {
					pending[curC] = true
				}

				grid.Set(n, x, y)
			})

			flashed[c] = true
			numFlashes++
		}

		// set all flashed coords to 0
		for k, _ := range flashed {
			grid.Set(0, k.x, k.y)
		}

	}

	return numFlashes
}

func part2(inputfile string) int {
	grid := loadGrid(inputfile)

	target := grid.Width() * grid.Height()

	steps := 0
	flashed := map[Coord]bool{}
	for len(flashed) != target {
		steps++
		pending := map[Coord]bool{}

		// increase all by 1
		grid.VisitAll2D(func(val interface{}, x, y int) {
			n := val.(int)

			n++
			if n > 9 {
				pending[NewCoord(x, y)] = true
			}

			grid.Set(n, x, y)
		})

		flashed = map[Coord]bool{}
		// doo flashes
		for len(pending) > 0 {
			c, _ := popItemFromMap(pending)

			grid.VisitOrthoAndDiag(c.x, c.y, func(val interface{}, x, y int) {
				// if val == nil, that means is outside bounds of grid
				if val == nil {
					return
				}

				curC := NewCoord(x, y)

				_, f := flashed[curC]
				_, p := pending[curC]

				// if this coord is pending flash or has flashed
				// ignore it cause doesn't matter if new value inceased (will be reset to 0)
				if f || p {
					return
				}

				n := val.(int)
				n++
				if n > 9 {
					pending[curC] = true
				}

				grid.Set(n, x, y)
			})

			flashed[c] = true
		}

		// set all flashed coords to 0
		for k, _ := range flashed {
			grid.Set(0, k.x, k.y)
		}
	}

	return steps
}
