package day09

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func loadGrid(inputfile string) *util.InfGrid[int] {
	data, _ := util.ReadFileToStringSlice(inputfile)

	grid := util.NewInfGrid[int]()
	grid.WithDefaultValue(util.MinInt)

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

	return grid
}

func part1(inputfile string) int {
	grid := loadGrid(inputfile)

	lowPoints := []int{}
	grid.VisitAll2D(func(val int, x, y int) {
		if val == util.MinInt {
			return
		}

		neighborsRaw := grid.GetOrtho(x, y)
		neighbors := []int{}

		for _, neighbor := range neighborsRaw {
			if neighbor == util.MinInt {
				continue
			}

			neighbors = append(neighbors, neighbor)
		}

		v := val
		for _, n := range neighbors {
			if n <= v {
				return
			}
		}

		lowPoints = append(lowPoints, v)
	})

	result := 0
	for _, p := range lowPoints {
		result += p + 1
	}

	return result
}

func addToMap(m map[string]int, x, y, val int) {
	key := fmt.Sprintf("%v,%v", x, y)
	m[key] = val
}

func getItemFromMap(m map[string]int) (x int, y int, val int) {
	for k, v := range m {
		sp := strings.Split(k, ",")

		x, _ := strconv.Atoi(sp[0])
		y, _ := strconv.Atoi(sp[1])

		return x, y, v
	}

	return 0, 0, 0
}

func isIn(m map[string]int, x, y int) bool {
	key := fmt.Sprintf("%v,%v", x, y)
	if _, ok := m[key]; !ok {
		return false
	}

	return true
}

func removeFromMap(m map[string]int, x, y int) (bool, int) {
	key := fmt.Sprintf("%v,%v", x, y)
	if _, ok := m[key]; !ok {
		return false, 0
	}

	val := m[key]
	delete(m, key)

	return true, val
}

func part2(inputfile string) int {
	// like path finding, with 9's as the walls
	// create a map of all 'avail coords' (in this omitting 9's)
	// start with the first coord, and 'walk' until walls hit, every spot touched is removed from map and added to 'curent basin'
	// then pick next item in map until all gone

	grid := loadGrid(inputfile)

	// visited := util.NewInfGrid()
	// unvisited := util.NewInfGrid()

	// key = "x,y"
	unvisited := map[string]int{}
	visited := map[string]int{}

	basins := [][]int{}
	// basins = append(basins, []int{})

	grid.VisitAll2D(func(val int, x, y int) {
		v := val
		if v != 9 {
			addToMap(unvisited, x, y, v)
		} else {
			addToMap(visited, x, y, v)
		}
	})

	for len(unvisited) > 0 {
		x, y, _ := getItemFromMap(unvisited)
		// basin := basins[len(basins)-1]

		basin := spread(grid, unvisited, visited, []int{}, x, y)
		basins = append(basins, basin)
	}

	sort.SliceStable(basins, func(i, j int) bool {
		return len(basins[i]) > len(basins[j])
	})

	result := 1
	for i := 0; i < 3; i++ {
		l := len(basins[i])
		result = result * l
	}

	return result
}

func spread(grid *util.InfGrid[int], unvisited, visited map[string]int, basin []int, x, y int) []int {
	if isIn(unvisited, x, y) {
		removeFromMap(unvisited, x, y)

		v := grid.Get(x, y)
		addToMap(visited, x, y, v)
		basin = append(basin, v)
		basin = spread(grid, unvisited, visited, basin, x, y+1)
		basin = spread(grid, unvisited, visited, basin, x, y-1)
		basin = spread(grid, unvisited, visited, basin, x+1, y)
		basin = spread(grid, unvisited, visited, basin, x-1, y)
	}

	return basin
}
