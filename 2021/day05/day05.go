package day05

import (
	"aoc/util"
)

type pt struct {
	x int
	y int
}

type IntInfinityGrid struct {
	data map[int]map[int]int
}

func NewIntInfinityGrid() *IntInfinityGrid {
	grid := new(IntInfinityGrid)
	grid.data = map[int]map[int]int{}
	return grid
}

func (g *IntInfinityGrid) Set(val, x, y int) {
	data := g.data
	if _, ok := data[x]; !ok {
		data[x] = map[int]int{}
	}

	data[x][y] = val
}

func (g *IntInfinityGrid) Get(x, y int) int {
	data := g.data

	if _, ok := data[x]; !ok {
		return 0
	}

	if _, ok := data[x][y]; !ok {
		return 0
	}

	return data[x][y]
}

func (g *IntInfinityGrid) VisitAllPopulatedVals(visitFunc func(val, x, y int)) {
	for x, ymap := range g.data {
		for y, val := range ymap {
			visitFunc(val, x, y)
		}
	}
}

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	grid := NewIntInfinityGrid()

	for _, line := range data {
		ints := util.ParseInts(line)
		p1 := pt{x: ints[0], y: ints[1]}
		p2 := pt{x: ints[2], y: ints[3]}

		if p1.x != p2.x && p1.y != p2.y {
			// Skip diagnal
			continue
		}

		pts := []pt{}
		for x := util.Min(p1.x, p2.x); x <= util.Max(p1.x, p2.x); x++ {
			for y := util.Min(p1.y, p2.y); y <= util.Max(p1.y, p2.y); y++ {
				pts = append(pts, pt{x: x, y: y})
			}
		}

		for _, p := range pts {
			val := grid.Get(p.x, p.y)
			grid.Set(val+1, p.x, p.y)
		}
	}

	result := 0
	grid.VisitAllPopulatedVals(func(val, x, y int) {
		if val >= 2 {
			result++
		}
	})

	return result
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)
	grid := NewIntInfinityGrid()

	for _, line := range data {
		ints := util.ParseInts(line)
		p1 := pt{x: ints[0], y: ints[1]}
		p2 := pt{x: ints[2], y: ints[3]}

		dirX := 0
		if p2.x < p1.x {
			dirX = -1
		} else if p2.x > p1.x {
			dirX = 1
		}
		dirY := 0
		if p2.y < p1.y {
			dirY = -1
		} else if p2.y > p1.y {
			dirY = 1
		}

		pts := []pt{}
		for x, y := p1.x, p1.y; x != p2.x || y != p2.y; x, y = x+dirX, y+dirY {
			pts = append(pts, pt{x: x, y: y})
		}
		pts = append(pts, pt{x: p2.x, y: p2.y})

		for _, p := range pts {
			val := grid.Get(p.x, p.y)
			grid.Set(val+1, p.x, p.y)
		}
	}

	result := 0
	grid.VisitAllPopulatedVals(func(val, x, y int) {
		if val >= 2 {
			result++
		}
	})

	return result
}
