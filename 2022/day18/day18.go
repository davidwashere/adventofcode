package day18

import (
	"aoc/util"
	"fmt"
)

type Cube struct {
	X int
	Y int
	Z int
}

func (c Cube) String() string {
	return fmt.Sprintf("{%v, %v, %v}", c.X, c.Y, c.Z)
}

var (
	Cubes []*Cube
	Grid  *util.InfGrid[bool]
	minX  = util.MaxInt
	maxX  = util.MinInt
	minY  = util.MaxInt
	maxY  = util.MinInt
	minZ  = util.MaxInt
	maxZ  = util.MinInt
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	Grid = util.NewInfGrid[bool]()
	for _, line := range data {
		ints := util.ParseInts(line)
		x := ints[0]
		y := ints[1]
		z := ints[2]

		c := &Cube{x, y, z}
		Cubes = append(Cubes, c)
		Grid.Set(true, x, y, z)

		minX = util.Min(minX, x)
		maxX = util.Max(maxX, x)
		minY = util.Min(minY, y)
		maxY = util.Max(maxY, y)
		minZ = util.Min(minZ, z)
		maxZ = util.Max(maxZ, z)
	}
}

func part1(inputFile string) int {
	load(inputFile)

	sides := 0
	for _, c := range Cubes {
		Grid.VisitOrtho3D(c.X, c.Y, c.Z, func(otherCube bool, x, y, z int) {
			if !otherCube {
				sides++
			}
		})
	}

	return sides
}

func part2(inputFile string) int {
	load(inputFile)

	// plan: start outside the extents, and crawl the full space removing any empty space found, stopping
	// when hitting an actual cube, the remaining cubes will be the truely 'empty' space

	// find all cubes
	innerCubes := map[Cube]bool{}
	for z := minZ - 1; z <= maxZ+1; z++ {
		for y := minY - 1; y <= maxY+1; y++ {
			for x := minX - 1; x <= maxX+1; x++ {
				innerCubes[Cube{x, y, z}] = true
			}
		}
	}

	// remove actual cubes from innerCubes
	for _, c := range Cubes {
		if innerCubes[*c] {
			delete(innerCubes, *c)
		}
	}

	// start outside the cube 'structure' and purge all cubes from 'innerCubes' that are actually outerCubes
	processed := map[Cube]bool{}

	q := util.Queue[Cube]{}
	start := Cube{minX - 1, minY - 1, minZ - 1}
	q.Enqueue(start)
	processed[start] = true

	for !q.IsEmpty() {
		cur := q.Dequeue()

		if innerCubes[cur] {
			// not an innerCube
			delete(innerCubes, cur)
		} else {
			continue
		}

		// added other empty cubes to test / purge
		Grid.VisitOrtho3D(cur.X, cur.Y, cur.Z, func(isCube bool, x, y, z int) {
			// if space holds cube, ignore it
			if isCube {
				return
			}

			c := Cube{x, y, z}
			// if we've already crawled this cube, ignore it
			if processed[c] {
				return
			}

			q.Enqueue(c)
			processed[c] = true
		})
	}

	// calc sides like before
	sides := 0
	for _, c := range Cubes {
		Grid.VisitOrtho3D(c.X, c.Y, c.Z, func(otherCube bool, x, y, z int) {
			if !otherCube {
				sides++
			}
		})
	}

	// for the cubes that are 'actually' inner cubes
	// reduce side count
	for c := range innerCubes {
		Grid.VisitOrtho3D(c.X, c.Y, c.Z, func(isCube bool, x, y, z int) {
			if isCube {
				sides--
			}
		})
	}

	return sides
}
