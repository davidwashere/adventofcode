package day14

import (
	"aoc/util"
	"fmt"
)

func loadGrid(inputfile string) *util.InfGrid[string] {
	grid := util.NewInfGrid[string]()
	grid.WithDefaultValue(".")

	lines, _ := util.ReadFileToStringSlice(inputfile)

	for _, line := range lines {
		ints := util.ParseInts(line)

		for i := 2; i < len(ints); i = i + 2 {
			px := ints[i-2]
			py := ints[i-1]
			x := ints[i]
			y := ints[i+1]

			if px == x {
				s := util.Min(py, y)
				e := util.Max(py, y)

				for j := s; j <= e; j++ {
					grid.Set("#", x, j)
				}

			} else if py == y {
				s := util.Min(px, x)
				e := util.Max(px, x)

				for j := s; j <= e; j++ {
					grid.Set("#", j, y)
				}
			}
		}
	}

	// for i, line := range data {
	// 	tokens := util.ParseTokens(line)

	// 	fmt.Println(tokens)
	// }

	// grid.FlipH()
	grid.SetExtents(grid.GetMinX(), 0, grid.GetMaxX(), grid.GetMaxY())
	return grid
}

func part1(inputfile string) int {
	grid := loadGrid(inputfile)

	// grid.Dump()

	floor := grid.GetMaxY()
	fmt.Println("floor at:", floor)

	downV := util.VectorNorth
	leftDiagV := util.VectorNorthWest
	rightDiagV := util.VectorNorthEast

	sandCount := 0
	for {

		curPt := util.Point{X: 500, Y: 0}
		for {
			if curPt.Y >= floor {
				break
			}

			nextPt := curPt.Apply(downV)
			next := grid.Get(nextPt.X, nextPt.Y)

			if next == "#" {
				nextPt = curPt.Apply(leftDiagV)
				next = grid.Get(nextPt.X, nextPt.Y)
				if next == "#" {
					nextPt = curPt.Apply(rightDiagV)
					next = grid.Get(nextPt.X, nextPt.Y)
					if next == "#" {
						grid.Set("#", curPt.X, curPt.Y)
						break
					} else {
						curPt = nextPt
						continue
					}
				} else {
					curPt = nextPt
					continue
				}
			} else {
				curPt = nextPt
				continue
			}
		}
		if curPt.Y >= floor {
			break
		}
		sandCount++
	}

	return sandCount
}

func part2(inputfile string) int {
	grid := loadGrid(inputfile)

	// grid.Dump()

	floor := grid.GetMaxY() + 2
	fmt.Println("floor at:", floor)

	downV := util.VectorNorth
	leftDiagV := util.VectorNorthWest
	rightDiagV := util.VectorNorthEast

	sandCount := 0
	for {

		curPt := util.Point{X: 500, Y: 0}
		for {
			nextPt := curPt.Apply(downV)
			next := grid.Get(nextPt.X, nextPt.Y)

			if nextPt.Y == floor {
				grid.Set("#", curPt.X, curPt.Y)
				break
			}

			if next == "#" {
				nextPt = curPt.Apply(leftDiagV)
				next = grid.Get(nextPt.X, nextPt.Y)
				if next == "#" {
					nextPt = curPt.Apply(rightDiagV)
					next = grid.Get(nextPt.X, nextPt.Y)
					if next == "#" {
						grid.Set("#", curPt.X, curPt.Y)
						break
					} else {
						curPt = nextPt
						continue
					}
				} else {
					curPt = nextPt
					continue
				}
			} else {
				curPt = nextPt
				continue
			}
		}
		sandCount++
		if curPt.Y == 0 && curPt.X == 500 {
			break
		}
	}

	return sandCount
}
