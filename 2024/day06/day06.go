package day06

import (
	"aoc/util"
	"fmt"
)

var (
	grid           *util.InfGrid[string]
	startX, startY int
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	grid = util.NewInfGrid[string]()

	// for y := len(data) - 1; y >= 0; y-- {
	for num, line := range data {
		// for y, line := range data {
		y := len(data) - 1 - num
		for x := 0; x < len(line); x++ {
			c := string(line[x])

			if c == "^" {
				startX = x
				startY = y
			}
			grid.Set(c, x, y)
		}
	}

	// grid.Dump()
}

func part1(inputFile string) int {
	load(inputFile)

	v := util.VectorNorth
	x := startX
	y := startY
	grid.Set("X", x, y)

	count := 1
	for {
		tx, ty := v.Apply(x, y)
		c := grid.Get(tx, ty)

		if c == "#" {
			v.RotateInt(-90)
			continue
		}

		if c == "" {
			break
		}

		if c != "X" {
			count++
		}

		x, y = tx, ty
		grid.Set("X", x, y)
	}

	return count
}

func setupGrid() {
	v := util.VectorNorth
	x := startX
	y := startY
	grid.Set("X", x, y)

	for {
		tx, ty := v.Apply(x, y)
		c := grid.Get(tx, ty)

		if c == "#" {
			v.RotateInt(-90)
			continue
		}

		if c == "" {
			break
		}

		x, y = tx, ty
		grid.Set("X", x, y)
	}
}

func part2(inputFile string) int {
	load(inputFile)

	// load the spots that dude will walk
	setupGrid()

	// grid.Dump()

	v := util.VectorNorth
	x := startX
	y := startY

	steps := 0
	count := 0
	for {
		tx, ty := v.Apply(x, y)
		c := grid.Get(tx, ty)

		if c == "#" {
			v.RotateInt(-90)
			continue
		}

		if c == "" {
			break
		}

		tv := util.Vector(v)
		ttx, tty := tv.Apply(tx, ty)
		if !(ttx == startX && tty == startY) {
			if isLoop(tx, ty, util.Vector(v)) {
				count++
			}
		}

		x, y = tx, ty

		// fmt.Println("=====")
		// grid.Dump()
		steps++
		if steps%100 == 0 {
			fmt.Println(steps)
		}
	}

	fmt.Println(steps)
	return count
}

func isLoop(sx, sy int, v util.Vector) bool {
	v.Rotate(-90)

	x, y := sx, sy
	for {
		tx, ty := v.Apply(x, y)

		if tx == sx && ty == sy {
			// TODO: Must ALSO be going in the original direction (before the rotate)
			return true
		}

		c := grid.Get(tx, ty)
		if c == "" {
			// we're off grid
			return false
		}

		if c == "#" {
			// if this is the first step, and we already hit a wall, not possible to loop
			if x == sx && y == sy {
				return false
			}

			v.Rotate(-90)
			continue
		}

		if c != "X" {
			// we're off walking path
			grid.Dump()
			return false
		}

		x, y = tx, ty
	}
}
