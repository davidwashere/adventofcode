package day22

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

const (
	Void = iota
	Wall
	Open
)

var (
	dirPoints = map[util.Vector]int{
		util.VectorEast:  0,
		util.VectorNorth: 1,
		util.VectorWest:  2,
		util.VectorSouth: 3,
	}
)

func load(inputFile string) ([]string, *util.InfGrid[int]) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	steps := []string{}
	re := regexp.MustCompile("[0-9]+|[LR]+")

	gridDone := false
	gridData := []string{}
	for _, line := range data {
		if len(line) == 0 {
			gridDone = true
			continue
		}

		if gridDone {
			steps = re.FindAllString(line, -1)
			continue
		}

		gridData = append(gridData, line)
	}

	grid := util.NewInfGrid[int]()
	grid.WithDefaultValue(Void)

	for y := 0; y < len(gridData); y++ {
		line := gridData[y]
		for x := 0; x < len(line); x++ {
			c := line[x]

			if c == '.' {
				grid.Set(Open, x, y)
			} else if c == '#' {
				grid.Set(Wall, x, y)
			}
		}
	}

	grid.WithDumpFunc(func(val int, freshRow bool) {
		if freshRow {
			fmt.Println()
			return
		}

		if val == Void {
			fmt.Print(" ")
		} else if val == Open {
			fmt.Print(".")
		} else if val == Wall {
			fmt.Print("#")
		}
	})

	return steps, grid
}

type MinMax struct {
	min int
	max int
}

func (m MinMax) String() string {
	return fmt.Sprintf("{%v,%v}", m.min, m.max)
}

func calcAllMinMax(grid *util.InfGrid[int]) ([]*MinMax, []*MinMax) {
	colsMinMax := make([]*MinMax, grid.GetMaxX()-grid.GetMinX()+1)
	rowsMinMax := []*MinMax{}

	for y := grid.GetMinY(); y <= grid.GetMaxY(); y++ {
		rowMin := util.MaxInt
		rowMax := util.MinInt

		for x := grid.GetMinX(); x <= grid.GetMaxX(); x++ {
			if colsMinMax[x] == nil {
				mm := new(MinMax)
				mm.min = util.MaxInt
				mm.max = util.MinInt
				colsMinMax[x] = mm
			}

			if grid.Get(x, y) != Void {
				colsMinMax[x].min = util.Min(colsMinMax[x].min, y)
				colsMinMax[x].max = util.Max(colsMinMax[x].max, y)
				rowMin = util.Min(rowMin, x)
				rowMax = util.Max(rowMax, x)
			}
		}
		rowsMinMax = append(rowsMinMax, &MinMax{rowMin, rowMax})
	}

	return rowsMinMax, colsMinMax
}

func dumpGrid(grid *util.InfGrid[int]) {
	grid.FlipH()
	grid.Dump()
	grid.FlipH()
}

func part1(inputFile string) int {
	steps, grid := load(inputFile)
	// fmt.Println(steps)
	dumpGrid(grid)

	rowsMinMax, colsMinMax := calcAllMinMax(grid)
	// fmt.Println("rows", rowsMinMax)
	// fmt.Println("cols", colsMinMax)

	curDir := util.VectorEast
	curPos := util.NewPoint(rowsMinMax[0].min, 0)

	for _, step := range steps {
		if step == "R" {
			curDir.Rotate(90)
			continue
		}

		if step == "L" {
			curDir.Rotate(-90)
			continue
		}

		moves, _ := strconv.Atoi(step)
		for i := 0; i < moves; i++ {
			p := curPos.Apply(curDir)

			if grid.Get(p.X, p.Y) == Void {
				if curDir.X > 0 {
					// right
					p.X = rowsMinMax[curPos.Y].min
				} else if curDir.X < 0 {
					// left
					p.X = rowsMinMax[curPos.Y].max
				} else if curDir.Y > 0 {
					// south
					p.Y = colsMinMax[curPos.X].min
				} else if curDir.Y < 0 {
					// North
					p.Y = colsMinMax[curPos.X].max
				}
			}

			if grid.Get(p.X, p.Y) == Wall {
				break
			}

			if grid.Get(p.X, p.Y) == Open {
				curPos = p
				continue
			}
		}
	}

	row := curPos.Y + 1
	col := curPos.X + 1

	return (1000 * row) + (4 * col) + dirPoints[curDir]
}

func getNextPointDir(p util.Point, curDir util.Vector) (util.Point, util.Vector) {
	// hardcoding for part 2
	np := util.Point{}

	// left edges
	if p.X == 49 && p.Y < 100 {
		if p.Y < 50 {
			// A1 - OO
			np.X = 0
			np.Y = 149 - p.Y
			dir := util.VectorEast
			return np, dir
		} else {
			// B1 - OO
			// 99y = 49x
			// 50y = 0x
			np.X = p.Y - 50
			np.Y = 100
			dir := util.VectorNorth
			return np, dir
		}
	} else if p.X == -1 {
		if p.Y >= 100 && p.Y <= 149 {
			// A2 - OO
			// 100y = 49y (-51)
			// 149y = 0y  (-149)
			np.X = 50
			np.Y = 149 - p.Y
			dir := util.VectorEast
			return np, dir
		} else {
			// C2 - OO
			// 150y=50x
			// 199y=99x
			np.X = p.Y - 100
			np.Y = 0
			dir := util.VectorNorth
			return np, dir
		}
	}

	// right edges
	if p.X == 150 {
		// F1 - OO
		// 0y = 149y
		// 49y = 100y
		np.X = 99
		np.Y = 149 - p.Y
		dir := util.VectorWest
		return np, dir
	} else if p.X == 100 && p.Y >= 50 {
		if p.Y < 100 {
			// E2 - OO
			// 50y = 100x
			// 99y = 149x
			np.X = p.Y + 50
			np.Y = 49
			dir := util.VectorSouth
			return np, dir
		} else {
			// F2 - OO
			// 100y = 49y
			// 149y = 0y
			np.X = 149
			np.Y = 149 - p.Y
			dir := util.VectorWest
			return np, dir
		}

	} else if p.X == 50 && p.Y >= 150 {
		// G2
		// 150y=50x
		// 199y=99x

		np.X = p.Y - 100
		np.Y = 149
		dir := util.VectorSouth
		return np, dir
	}

	// top edges
	if p.Y < 0 {
		if p.X < 100 {
			// C1 - OO
			// 50x = 150y
			// 99x = 199y
			np.X = 0
			np.Y = p.X + 100
			dir := util.VectorEast
			return np, dir
		} else {
			// D1 - OO
			// 100x = 0x
			// 149x = 49x

			np.X = p.X - 100
			np.Y = 199
			dir := util.VectorSouth
			return np, dir
		}

	} else if p.Y == 99 && p.X < 50 {
		// B2 - OO
		// 0x = 50y
		// 49x = 99y
		np.X = 50
		np.Y = p.X + 50
		dir := util.VectorEast
		return np, dir
	}

	// bottom edges
	if p.Y == 50 && p.X >= 100 {
		// E1 - OO

		// 100x=50y
		// 149x-99y

		np.X = 99
		np.Y = p.X - 50
		dir := util.VectorWest

		return np, dir
	} else if p.Y == 150 && p.X >= 50 && p.X < 100 {
		// G1 - OO

		// 50x=150y
		// 99x=199y
		np.X = 49
		np.Y = p.X + 100
		dir := util.VectorWest

		return np, dir
	} else if p.Y >= 200 {
		// D2

		// 0x = 100x
		// 49x = 149x

		np.X = p.X + 100
		np.Y = 0
		dir := util.VectorNorth
		return np, dir
	}

	panic("shouldn't ever get here")
}

func part2(inputFile string) int {
	steps, grid := load(inputFile)
	// fmt.Println(steps)
	// dumpGrid(grid)

	rowsMinMax, _ := calcAllMinMax(grid)
	// fmt.Println("rows", rowsMinMax)
	// fmt.Println("cols", colsMinMax)

	curDir := util.VectorEast
	curPos := util.NewPoint(rowsMinMax[0].min, 0)

	for _, step := range steps {
		if step == "R" {
			curDir.Rotate(90)
			continue
		}

		if step == "L" {
			curDir.Rotate(-90)
			continue
		}

		moves, _ := strconv.Atoi(step)
		for i := 0; i < moves; i++ {
			p := curPos.Apply(curDir)

			if grid.Get(p.X, p.Y) == Void {
				p, curDir = getNextPointDir(p, curDir)
				// if curDir.X > 0 {
				// 	// right
				// 	p.X = rowsMinMax[curPos.Y].min
				// } else if curDir.X < 0 {
				// 	// left
				// 	p.X = rowsMinMax[curPos.Y].max
				// } else if curDir.Y > 0 {
				// 	// south
				// 	p.Y = colsMinMax[curPos.X].min
				// } else if curDir.Y < 0 {
				// 	// North
				// 	p.Y = colsMinMax[curPos.X].max
				// }
			}

			if grid.Get(p.X, p.Y) == Wall {
				break
			}

			if grid.Get(p.X, p.Y) == Open {
				curPos = p
				continue
			}
		}
	}

	row := curPos.Y + 1
	col := curPos.X + 1

	return (1000 * row) + (4 * col) + dirPoints[curDir]
}
