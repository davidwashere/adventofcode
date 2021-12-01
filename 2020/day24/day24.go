package day24

import (
	"aoc2020/util"
)

func parsefile(inputfile string) [][]string {
	data, _ := util.ReadFileToStringSlice(inputfile)

	tilesToFlip := [][]string{}

	for _, line := range data {
		steps := []string{}
		inS := false
		inN := false
		for _, char := range line {
			switch char {
			case 's':
				inS = true
			case 'n':
				inN = true
			case 'e':
				if inS {
					inS = false
					steps = append(steps, "se")
				} else if inN {
					inN = false
					steps = append(steps, "ne")
				} else {
					steps = append(steps, "e")
				}
			case 'w':
				if inS {
					inS = false
					steps = append(steps, "sw")
				} else if inN {
					inN = false
					steps = append(steps, "nw")
				} else {
					steps = append(steps, "w")
				}
			}
		}
		// fmt.Println(steps)
		tilesToFlip = append(tilesToFlip, steps)
	}

	return tilesToFlip
}

func part1(inputfile string) int {
	tilesToFlip := parsefile(inputfile)

	grid := util.NewInfinityGrid(".")

	oddV := map[string]util.Vector{
		"e":  util.NewVector(1, 0, 1),
		"w":  util.NewVector(-1, 0, 1),
		"se": util.NewVector(1, 1, 1),
		"sw": util.NewVector(0, 1, 1),
		"ne": util.NewVector(1, -1, 1),
		"nw": util.NewVector(0, -1, 1),
	}

	evenV := map[string]util.Vector{
		"e":  util.NewVector(1, 0, 1),
		"w":  util.NewVector(-1, 0, 1),
		"se": util.NewVector(0, 1, 1),
		"sw": util.NewVector(-1, 1, 1),
		"ne": util.NewVector(0, -1, 1),
		"nw": util.NewVector(-1, -1, 1),
	}

	for _, steps := range tilesToFlip {
		x := 0
		y := 0
		for _, step := range steps {
			var v util.Vector
			if y%2 == 0 {
				v = evenV[step]
			} else {
				v = oddV[step]
			}
			x, y = v.Apply(x, y)

		}
		val := grid.Get(x, y)
		if val == "." {
			grid.Set("#", x, y)
		} else {
			grid.Set(".", x, y)
		}
	}

	grid.Dump()

	result := 0
	grid.VisitAll2D(func(val string, x, y int) {
		if val == "#" {
			result++
		}
	})

	return result
}

type chg struct {
	x   int
	y   int
	val string
}

func part2(inputfile string) int {
	tilesToFlip := parsefile(inputfile)

	grid := util.NewInfinityGrid(".")

	oddV := map[string]util.Vector{
		"e":  util.NewVector(1, 0, 1),
		"w":  util.NewVector(-1, 0, 1),
		"se": util.NewVector(1, 1, 1),
		"sw": util.NewVector(0, 1, 1),
		"ne": util.NewVector(1, -1, 1),
		"nw": util.NewVector(0, -1, 1),
	}

	evenV := map[string]util.Vector{
		"e":  util.NewVector(1, 0, 1),
		"w":  util.NewVector(-1, 0, 1),
		"se": util.NewVector(0, 1, 1),
		"sw": util.NewVector(-1, 1, 1),
		"ne": util.NewVector(0, -1, 1),
		"nw": util.NewVector(-1, -1, 1),
	}

	for _, steps := range tilesToFlip {
		x := 0
		y := 0
		for _, step := range steps {
			var v util.Vector
			if y%2 == 0 {
				v = evenV[step]
			} else {
				v = oddV[step]
			}
			x, y = v.Apply(x, y)

		}
		val := grid.Get(x, y)
		if val == "." {
			grid.Set("#", x, y)
		} else {
			grid.Set(".", x, y)
		}
	}

	grid.Dump()

	result := 0
	for i := 1; i <= 100; i++ {
		result = 0
		changes := []chg{}
		grid.Grow(1)
		grid.VisitAll2D(func(val string, x, y int) {
			var vecs map[string]util.Vector
			if y%2 == 0 {
				vecs = evenV
			} else {
				vecs = oddV
			}

			numAdj := adjBlack(grid, x, y, vecs)

			if val == "#" {
				if numAdj == 0 || numAdj > 2 {
					changes = append(changes, chg{x: x, y: y, val: "."})
				}
			} else {
				if numAdj == 2 {
					changes = append(changes, chg{x: x, y: y, val: "#"})
				}
			}
		})

		for _, c := range changes {
			grid.Set(c.val, c.x, c.y)
		}

		grid.VisitAll2D(func(val string, x, y int) {
			if val == "#" {
				result++
			}
		})
		// fmt.Printf("Day %v: %v\n", i, result)
	}

	return result
}

func adjBlack(grid *util.InfinityGrid, x, y int, vecs map[string]util.Vector) int {
	result := 0
	for _, v := range vecs {
		nx, ny := v.Apply(x, y)
		if grid.Get(nx, ny) == "#" {
			result++
		}
	}

	return result
}
