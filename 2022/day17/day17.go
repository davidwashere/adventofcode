package day17

import (
	"aoc/util"
	"fmt"
)

const (
	Solid = "#"
)

var RockPts = [][]util.Point{
	{ // HLine
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 3, Y: 0},
	},
	{ // Plus
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		// {X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 1, Y: 2},
	},
	{ // Corner
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 2, Y: 0},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
	},
	{ // VLine
		{X: 0, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: 2},
		{X: 0, Y: 3},
	},
	{ // Box
		{X: 0, Y: 0},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: 1, Y: 1},
	},
}

type WindDir util.Vector

var (
	WindDirLeft  = util.NewVector(-1, 0, 1)
	WindDirRight = util.NewVector(1, 0, 1)

	Down = util.NewVector(0, -1, 1)
)

type Rock struct {
	// BotLeft
	Pos util.Point
	Pts []util.Point
}

func (r Rock) String() string {
	return fmt.Sprintf("{%v - %v}", r.Pos, r.Pts)
}

func load(inputFile string) []util.Vector {
	data, _ := util.ReadFileToStringSlice(inputFile)

	jetPattern := []util.Vector{}
	for _, line := range data {
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c == '<' {
				jetPattern = append(jetPattern, WindDirLeft)
			} else if c == '>' {
				jetPattern = append(jetPattern, WindDirRight)
			} else {
				panic("unknown char in input")
			}
		}
	}

	return jetPattern
}

// getWindDir given an index will return the nextIndex and the appropriate jetPattern
func getWindDir(jetPattern []util.Vector, index int) (int, util.Vector) {
	i := index
	dir := jetPattern[i]
	i = (i + 1) % len(jetPattern)
	return i, dir
}

// pushRock will attempt to move rock in windDir, and alter rock accordingly
func pushRock(windDir util.Vector, rock *Rock, grid *util.InfGrid[bool], minX, maxX int) {
	bump := false
	newPos := rock.Pos.Apply(windDir)
	for _, pt := range rock.Pts {
		x := pt.X + newPos.X
		y := pt.Y + newPos.Y

		if x < minX || x > maxX {
			bump = true
			break
		}

		if grid.Get(x, y) {
			bump = true
			break
		}
	}

	if !bump {
		rock.Pos = newPos
	}
}

// fallRock will attempt to move a rock one unit down, and return whether or not move was
// successful
func fallRock(rock *Rock, grid *util.InfGrid[bool], floor int) bool {
	newPos := rock.Pos.Apply(Down)
	for _, pt := range rock.Pts {
		x := pt.X + newPos.X
		y := pt.Y + newPos.Y

		if y < floor {
			return false
		}

		if grid.Get(x, y) {
			return false
		}
	}

	rock.Pos = newPos
	return true
}

func part1(inputFile string) int {
	jetPattern := load(inputFile)

	maxRocks := 2022
	minX := 0
	maxX := 6
	minY := 0

	curTop := -1

	grid := util.NewInfGrid[bool]()

	windI := 0
	for rockI := 0; rockI < maxRocks; rockI++ {
		// grid.Dump()
		// fmt.Println()
		rock := Rock{
			Pos: util.NewPoint(2, curTop+4),
			Pts: RockPts[rockI%len(RockPts)],
		}

		// rock falling
		for {
			var windDir util.Vector
			windI, windDir = getWindDir(jetPattern, windI)

			pushRock(windDir, &rock, grid, minX, maxX)

			if !fallRock(&rock, grid, minY) {
				for i := 0; i < len(rock.Pts); i++ {
					pt := rock.Pts[i]
					x := pt.X + rock.Pos.X
					y := pt.Y + rock.Pos.Y
					grid.Set(true, x, y)
					curTop = util.Max(curTop, y)
				}
				break
			}
		}

	}

	return curTop + 1
}

func part2(inputFile string) int {
	jetPattern := load(inputFile)

	// wtf - this time s out
	maxRocks := 1000000000000
	minX := 0
	maxX := 6
	minY := 0

	curTop := -1

	grid := util.NewInfGrid[bool]()

	seen := map[string][]int{}

	rockI := 0
	windI := 0
	for rockNum := 0; rockNum < maxRocks; rockNum++ {
		key := fmt.Sprintf("%v,%v", rockI, windI)
		if v, ok := seen[key]; ok {
			// we've seen this combination of rock index + jet index before
			prevRockNum, prevTop := v[0], v[1]
			rem := (maxRocks - rockNum) % (rockNum - prevRockNum)
			// if between this and last time this rock and jet seen before
			// can be repeated to exactly reach maxRocks, lets use it
			// TODO: may be more efficient on first cycle found to repeat
			// until can't anymore, then just play rocks until reach maxRocks
			if rem == 0 {
				numRepeats := (maxRocks - rockNum) / (rockNum - prevRockNum)
				top := curTop + 1
				heightOfRepeatingSeq := top - prevTop
				return top + (heightOfRepeatingSeq * numRepeats)
			}

		} else {
			seen[key] = []int{rockNum, curTop + 1}
		}

		rock := Rock{
			Pos: util.NewPoint(2, curTop+4),
			Pts: RockPts[rockI],
		}
		rockI = (rockNum + 1) % len(RockPts)

		// rock falling
		for {
			var windDir util.Vector
			windI, windDir = getWindDir(jetPattern, windI)

			pushRock(windDir, &rock, grid, minX, maxX)

			if !fallRock(&rock, grid, minY) {
				for i := 0; i < len(rock.Pts); i++ {
					pt := rock.Pts[i]
					x := pt.X + rock.Pos.X
					y := pt.Y + rock.Pos.Y
					grid.Set(true, x, y)
					curTop = util.Max(curTop, y)
				}
				break
			}
		}
	}

	return curTop + 1
}
