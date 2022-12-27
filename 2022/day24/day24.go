package day24

import (
	"aoc/util"
	"fmt"
)

var (
	minX = 0
	minY = 0

	maxX = 0
	maxY = 0

	dirs = map[byte]util.Vector{
		'<': util.VectorWest,
		'>': util.VectorEast,
		// 'v': util.VectorSouth,
		'^': util.VectorSouth,
		// '^': util.VectorNorth,
		'v': util.VectorNorth,
	}
)

type Blizzard struct {
	Pos util.Point
	Dir util.Vector
}

func (b *Blizzard) Update() {
	newPos := b.Pos.Apply(b.Dir)

	if newPos.X > maxX {
		newPos.X = 0
	}
	if newPos.X < minX {
		newPos.X = maxX
	}
	if newPos.Y > maxY {
		newPos.Y = 0
	}
	if newPos.Y < minY {
		newPos.Y = maxY
	}

	b.Pos = newPos
}

var blizzards []*Blizzard
var blizzGrid *util.InfGrid[bool]

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	blizzards = []*Blizzard{}
	for y := 0; y < len(data)-2; y++ {
		line := data[y+1]
		for x := 0; x < len(line)-2; x++ {
			c := line[x+1]

			if dir, ok := dirs[c]; ok {
				blizzard := &Blizzard{
					Pos: util.NewPoint(x, y),
					Dir: dir,
				}

				blizzards = append(blizzards, blizzard)
			}

			maxY = util.Max(maxY, y)
			maxX = util.Max(maxX, x)
		}
	}
}

func updateBlizzards() {
	for _, blizz := range blizzards {
		blizz.Update()
	}
}

func updateGrid() {
	blizzGrid = util.NewInfGrid[bool]()

	for _, blizz := range blizzards {
		blizzGrid.Set(true, blizz.Pos.X, blizz.Pos.Y)
	}
	blizzGrid.SetExtents(minX, minY, maxX, maxY)
	blizzGrid.LockBounds()
	// blizzGrid.FlipH()
	blizzGrid.WithDumpFunc(gridDumpFunc)
}

func gridDumpFunc(val, freshRow bool) {
	if freshRow {
		fmt.Println()
		return
	}

	if val {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}

func part1(inputFile string) int {
	load(inputFile)
	updateGrid()

	start := util.NewPoint(0, -1)
	end := util.NewPoint(blizzGrid.GetMaxX(), blizzGrid.GetMaxY())

	return calcMinTime(start, end)
}

func getMoves(player util.Point) []util.Point {
	moves := []util.Point{}
	if !blizzGrid.Get(player.X, player.Y) {
		// no blizz here
		moves = append(moves, player)
	}

	blizzGrid.VisitOrtho(player.X, player.Y, func(blizz bool, x, y int) {
		if !blizz {
			moves = append(moves, util.NewPoint(x, y))
		}
	})

	return moves
}

func calcMinTime(start, end util.Point) int {
	players := map[util.Point]bool{}
	players[start] = true

	min := 0
	for {
		min++
		updateBlizzards()
		updateGrid()
		// fmt.Printf("==== [%v] ====\n", min)
		// blizzGrid.FlipH()
		// blizzGrid.Dump()
		// blizzGrid.FlipH()

		nPlayers := map[util.Point]bool{}
		for player := range players {
			moves := getMoves(player)

			for _, move := range moves {
				nPlayers[move] = true
			}
		}

		if len(nPlayers) == 0 {
			panic(fmt.Sprintf("no moves - all players died at min %v", min))
		}

		if _, ok := nPlayers[end]; ok {
			// player is at the position just outside exit
			break
		}

		players = map[util.Point]bool{}
		for move := range nPlayers {
			players[move] = true
		}
	}

	return min + 1
}

func part2(inputFile string) int {
	load(inputFile)
	updateGrid()

	start := util.NewPoint(0, -1)
	end := util.NewPoint(blizzGrid.GetMaxX(), blizzGrid.GetMaxY())

	total := 0
	total += calcMinTime(start, end)

	start = util.NewPoint(blizzGrid.GetMaxX(), blizzGrid.GetMaxY()+1)
	end = util.NewPoint(0, 0)
	total += calcMinTime(start, end) - 1

	start = util.NewPoint(0, -1)
	end = util.NewPoint(blizzGrid.GetMaxX(), blizzGrid.GetMaxY())
	total += calcMinTime(start, end) - 1

	return total
}
