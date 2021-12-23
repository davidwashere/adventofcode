package day15

import (
	"aoc/util"
	"strconv"
)

type pt struct {
	x int
	y int
}

// smallestDistPt returns pt with the smallest distance in map
func smallestDistPt(grid *util.InfGrid, q map[pt]bool) pt {
	var lp pt
	min := util.MaxInt
	for k := range q {
		v := grid.Get(k.x, k.y).(int)

		if v < min {
			lp = k
			min = v
		}
	}

	return lp
}

func loadFile(inputfile string) *util.InfGrid {
	data, _ := util.ReadFileToStringSlice(inputfile)

	weightGrid := util.NewInfGrid()

	for y, line := range data {
		for x, char := range line {
			val, _ := strconv.Atoi(string(char))
			weightGrid.Set(val, x, y)
		}
	}

	return weightGrid
}

func calc(weightGrid *util.InfGrid) int {
	// pending := []pt{}
	pending := map[pt]bool{}
	processed := map[pt]bool{}
	distGrid := util.NewInfGrid()
	distGrid.WithDefaultValue(util.MaxInt)
	distGrid.SetExtents(0, 0, weightGrid.Width()-1, weightGrid.Height()-1)
	distGrid.LockBounds()
	distGrid.Set(0, 0, 0) // set 0,0 as dist 0

	pending[pt{x: 0, y: 0}] = true

	for len(pending) > 0 {
		cur := smallestDistPt(distGrid, pending)
		delete(pending, cur)
		processed[cur] = true

		baseDist := distGrid.Get(cur.x, cur.y).(int)

		distGrid.VisitOrtho(cur.x, cur.y, func(val interface{}, x, y int) {
			p := pt{x: x, y: y}
			if processed[p] {
				return
			}

			if _, ok := pending[p]; !ok {
				pending[p] = true
			}

			dist := val.(int)
			weight := weightGrid.Get(x, y).(int)

			dist = util.Min(dist, baseDist+weight)

			distGrid.Set(dist, x, y)
		})
	}

	return distGrid.Get(distGrid.Width()-1, distGrid.Height()-1).(int)
}

func part1(inputfile string) int {
	weightGrid := loadFile(inputfile)
	return calc(weightGrid)
}

func part2(inputfile string) int {
	bumps := [][]int{
		{0, 1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
		{3, 4, 5, 6, 7},
		{4, 5, 6, 7, 8},
	}

	weightGrid := loadFile(inputfile)

	width := weightGrid.Width()
	height := weightGrid.Height()

	// Duplicate grid
	for dx := 0; dx < 5; dx++ {
		for dy := 0; dy < 5; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			for x := 0; x < width; x++ {
				for y := 0; y < height; y++ {
					nx := x + (dx * width)
					ny := y + (dy * height)

					inc := bumps[dx][dy]

					val := weightGrid.Get(x, y).(int)
					val += inc
					val = (val-1)%9 + 1

					weightGrid.Set(val, nx, ny)
				}
			}
		}
	}

	// repeat p1
	return calc(weightGrid)
}
