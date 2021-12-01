package day17

import (
	"aoc2020/util"
)

// .
const (
	ACTIVE   = "#"
	INACTIVE = "."
)

type coord struct {
	x int
	y int
	z int
	w int
}

type chg struct {
	c   coord
	val string
}

func NewCoord(x, y int, more ...int) coord {
	z := 0
	if len(more) > 0 {
		z = more[0]
	}

	w := 0
	if len(more) > 1 {
		w = more[1]
	}

	return coord{
		x: x,
		y: y,
		z: z,
		w: w,
	}
}

// If a cube is active and exactly 2 or 3 of its neighbors are also active
// the cube remains active. Otherwise, the cube becomes inactive.

// If a cube is inactive but exactly 3 of its neighbors are active
// the cube becomes active. Otherwise, the cube remains inactive.
func part1(inputfile string) int {
	// Generate coordinates for adjacent cubes
	adjs := []coord{}
	util.PermsIntOfLen([]int{-1, 0, 1}, 3, true, func(perm []int) {
		x, y, z := perm[0], perm[1], perm[2]

		if !(x == 0 && y == 0 && z == 0) {
			adjs = append(adjs, NewCoord(x, y, z))
		}
	})

	// Create the infinity grid
	iGrid := util.NewInfinityGridFromFile(inputfile, INACTIVE)
	iGrid.AddDimension() // make it 3d

	for cycle := 0; cycle < 6; cycle++ {
		iGrid.Grow(1)

		changes := []chg{}
		iGrid.VisitAll(func(val string, x, y int, dims ...int) {
			z := dims[0]

			c := NewCoord(x, y, z)
			nActive := activeNeightbors(adjs, iGrid, c)

			if val == ACTIVE && (nActive < 2 || nActive > 3) {
				changes = append(changes, chg{c, INACTIVE})
			} else if val == INACTIVE && nActive == 3 {
				changes = append(changes, chg{c, ACTIVE})
			}
		})

		for _, change := range changes {
			c := change.c
			iGrid.Set(change.val, c.x, c.y, c.z)
		}
	}

	result := 0
	iGrid.VisitAll(func(val string, x, y int, dims ...int) {
		if val == ACTIVE {
			result++
		}
	})

	return result
}

func part2(inputfile string) int {
	// Generate coordinates for adjacent cubes
	adjs := []coord{}
	util.PermsIntOfLen([]int{-1, 0, 1}, 4, true, func(perm []int) {
		x, y, z, w := perm[0], perm[1], perm[2], perm[3]

		if !(x == 0 && y == 0 && z == 0 && w == 0) {
			adjs = append(adjs, NewCoord(x, y, z, w))
		}
	})

	iGrid := util.NewInfinityGridFromFile(inputfile, INACTIVE)
	iGrid.AddDimension() // make it 3d
	iGrid.AddDimension() // make it 4d

	for cycle := 0; cycle < 6; cycle++ {
		iGrid.Grow(1)

		changes := []chg{}
		iGrid.VisitAll(func(val string, x, y int, dims ...int) {
			z, w := dims[0], dims[1]

			c := coord{x, y, z, w}
			nActive := activeNeightbors(adjs, iGrid, c)

			if val == ACTIVE && (nActive < 2 || nActive > 3) {
				changes = append(changes, chg{c, INACTIVE})
			} else if val == INACTIVE && nActive == 3 {
				changes = append(changes, chg{c, ACTIVE})
			}
		})

		for _, change := range changes {
			c := change.c
			iGrid.Set(change.val, c.x, c.y, c.z, c.w)
		}
	}

	result := 0
	iGrid.VisitAll(func(val string, x, y int, dims ...int) {
		if val == ACTIVE {
			result++
		}
	})

	return result
}

func activeNeightbors(adjs []coord, iGrid *util.InfinityGrid, o coord) int {
	count := 0
	for _, c := range adjs {
		x, y, z, w := c.x, c.y, c.z, c.w

		c := coord{x: o.x + x, y: o.y + y, z: o.z + z, w: o.w + w}

		if iGrid.Get(c.x, c.y, c.z, c.w) == ACTIVE {
			count++
		}
	}

	return count
}
