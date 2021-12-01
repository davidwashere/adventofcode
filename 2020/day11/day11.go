package day11

import (
	"aoc2020/util"
)

type change struct {
	x   int
	y   int
	val string
}

var vec = util.NewNormalizedVector

var adjVectors = []util.Vector{
	vec(-1, -1),
	vec(0, -1),
	vec(1, -1),
	vec(1, 0),
	vec(1, 1),
	vec(0, 1),
	vec(-1, 1),
	vec(-1, 0),
}

func part1(inputfile string) int {
	grid := util.NewInfinityGridFromFile(inputfile, "X") // X = out of bounds

	for {
		changes := []change{}

		grid.VisitAll(func(val string, x int, y int, dims ...int) {
			adjOccSeats := adjOccupiedSeats(grid, x, y)
			if val == "L" {
				if adjOccSeats == 0 {
					changes = append(changes, change{x, y, "#"})
				}

			} else if val == "#" {
				if adjOccSeats >= 4 {
					changes = append(changes, change{x, y, "L"})
				}

			}
		})

		if len(changes) == 0 {
			break
		}

		for _, c := range changes {
			grid.Set(c.val, c.x, c.y)
		}
	}

	occSeats := 0
	grid.VisitAll(func(val string, x int, y int, dims ...int) {
		if val == "#" {
			occSeats++
		}
	})

	return occSeats
}

// adjOccupiedSeats counts number of occupied seats immediatly surrounding
// the x, y coord
func adjOccupiedSeats(grid *util.InfinityGrid, x, y int) int {
	occSeats := 0
	for _, v := range adjVectors {
		if grid.Get(v.Apply(x, y)) == "#" {
			occSeats++
		}
	}

	return occSeats
}

func part2(inputfile string) int {
	grid := util.NewInfinityGridFromFile(inputfile, "X") // X = out of bounds

	for {
		changes := []change{}
		grid.VisitAll(func(val string, x int, y int, dims ...int) {
			adjOctSeats := adjSeenOccupiedSeats(grid, x, y)
			if val == "L" {
				if adjOctSeats == 0 {
					changes = append(changes, change{x, y, "#"})
				}

			} else if val == "#" {
				if adjOctSeats >= 5 {
					changes = append(changes, change{x, y, "L"})
				}

			}
		})

		if len(changes) == 0 {
			break
		}

		for _, c := range changes {
			grid.Set(c.val, c.x, c.y)
		}
	}

	occSeats := 0
	grid.VisitAll(func(val string, x int, y int, dims ...int) {
		if val == "#" {
			occSeats++
		}
	})

	return occSeats
}

// adjSeenOccupiedSeats counts number of occupied seats 'visible' to
// the x, y coord
func adjSeenOccupiedSeats(grid *util.InfinityGrid, x, y int) int {
	occSeats := 0
	for _, v := range adjVectors {
		curX, curY := v.Apply(x, y)
		for {
			if grid.Get(curX, curY) == "#" {
				occSeats++
				break
			}

			if grid.Get(curX, curY) == "L" || grid.Get(curX, curY) == "X" {
				break
			}

			curX, curY = v.Apply(curX, curY)
		}
	}

	return occSeats
}
