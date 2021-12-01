package day20

import (
	"aoc2020/util"
	"math"
	"strings"
)

type tilesT struct {
	asMap      map[int]tile
	asSlice    []tile
	edgeCounts map[string]int
}

type tile struct {
	id   int
	data []string
	grid *util.InfinityGrid
}

func parsefile(inputfile string) tilesT {
	data, _ := util.ReadFileToStringSlice(inputfile)

	var tiles tilesT
	tiles.asMap = map[int]tile{}
	var t tile
	for _, line := range data {
		// fmt.Println(line)
		if strings.Index(line, ":") >= 0 {
			ints := util.ParseInts(line)
			t = tile{id: ints[0]}
			continue
		}

		if strings.TrimSpace(line) == "" {
			t.grid = util.NewInfinityGridFromSlice(t.data, "X")
			tiles.asSlice = append(tiles.asSlice, t)
			tiles.asMap[t.id] = t
			continue
		}

		t.data = append(t.data, line)
	}
	t.grid = util.NewInfinityGridFromSlice(t.data, "X")
	tiles.asSlice = append(tiles.asSlice, t)
	tiles.asMap[t.id] = t

	return tiles
}

func populateEdgeCounts(tiles *tilesT) {
	tiles.edgeCounts = map[string]int{}
	for _, tile := range tiles.asSlice {
		for _, edge := range tile.grid.Edges() {
			if _, ok := tiles.edgeCounts[edge]; !ok {
				tiles.edgeCounts[edge] = 1
			} else {
				tiles.edgeCounts[edge] = tiles.edgeCounts[edge] + 1
			}
		}

		for _, edge := range tile.grid.EdgesFlipped() {
			if _, ok := tiles.edgeCounts[edge]; !ok {
				tiles.edgeCounts[edge] = 1
			} else {
				tiles.edgeCounts[edge] = tiles.edgeCounts[edge] + 1
			}
		}
	}
}

func cornerTileIDs(tiles tilesT) []int {
	var tileIDs []int

	for _, tile := range tiles.asSlice {
		uniqueEdges := 0
		for _, edge := range tile.grid.Edges() {
			if tiles.edgeCounts[edge] == 1 {
				uniqueEdges++
			}
		}

		if uniqueEdges == 2 {
			tileIDs = append(tileIDs, tile.id)
		}
	}

	return tileIDs
}

func edgeTileIDs(tiles tilesT) []int {
	var tileIDs []int

	for _, tile := range tiles.asSlice {
		uniqueEdges := 0
		for _, edge := range tile.grid.Edges() {
			if tiles.edgeCounts[edge] == 1 {
				uniqueEdges++
			}
		}

		if uniqueEdges == 1 {
			tileIDs = append(tileIDs, tile.id)
		}
	}

	return tileIDs
}

func middleTileIDs(tiles tilesT, cornerAndEdgeIds []int) []int {
	var tileIDs []int

	for _, tile := range tiles.asSlice {
		if !util.IsIntIn(cornerAndEdgeIds, tile.id) {
			tileIDs = append(tileIDs, tile.id)
		}
	}

	return tileIDs
}

func part1(inputfile string) int {
	tiles := parsefile(inputfile)
	populateEdgeCounts(&tiles)
	cornerIDs := cornerTileIDs(tiles)

	idSum := 1
	for _, cornerID := range cornerIDs {
		idSum *= cornerID
	}

	return idSum
}

func alignGridLeft(right string, grid *util.InfinityGrid) {
	left := grid.LeftEdge()
	rotCount := 0
	flipped := false
	for right != left {

		rotCount++
		if rotCount > 3 && flipped {
			panic("not good!")
		}

		if rotCount > 3 {
			grid.FlipH()
			rotCount = 0
			flipped = true
			left = grid.LeftEdge()
			continue
		}

		grid.Rotate(-90)
		left = grid.LeftEdge()
	}
}

func alignGridTop(bottom string, grid *util.InfinityGrid) {
	top := grid.TopEdge()
	rotCount := 0
	flipped := false
	for bottom != top {

		rotCount++
		if rotCount > 3 && flipped {
			panic("not good!")
		}

		if rotCount > 3 {
			grid.FlipH()
			rotCount = 0
			flipped = true
			top = grid.TopEdge()
			continue
		}

		grid.Rotate(-90)
		top = grid.TopEdge()
	}
}

func alignGridLeftTop(right string, bottom string, grid *util.InfinityGrid) {
	top := grid.TopEdge()
	left := grid.LeftEdge()
	rotCount := 0
	flipped := false
	for bottom != top && right != left {

		rotCount++
		if rotCount > 3 && flipped {
			panic("not good!")
		}

		if rotCount > 3 {
			grid.FlipH()
			rotCount = 0
			flipped = true
			top = grid.TopEdge()
			left = grid.LeftEdge()
			continue
		}

		grid.Rotate(-90)
		top = grid.TopEdge()
		left = grid.LeftEdge()
	}
}

func findTileIDWithMatchingEdge(tiles tilesT, edge string, tileIDs []int) int {
	for _, tileID := range tileIDs {
		tile := tiles.asMap[tileID]
		tileEdges := append(tile.grid.Edges(), tile.grid.EdgesFlipped()...)
		for _, tileEdge := range tileEdges {
			if edge == tileEdge {
				return tileID
			}
		}
	}

	return -1
}

func findTileIDWithMatchingLeftTop(tiles tilesT, right string, bottom string, tileIDs []int) int {
	var tile tile
	for _, tileID := range tileIDs {
		foundCount := 0
		tile = tiles.asMap[tileID]
		tileEdges := append(tile.grid.Edges(), tile.grid.EdgesFlipped()...)
		for _, tileEdge := range tileEdges {
			if right == tileEdge || bottom == tileEdge {
				foundCount++
			}
		}

		if foundCount >= 2 {
			break
		}
	}

	return tile.id
}

func part2(inputfile string) int {
	tiles := parsefile(inputfile)

	populateEdgeCounts(&tiles)
	cornerIDs := cornerTileIDs(tiles)
	edgeIDs := edgeTileIDs(tiles)
	midIDs := middleTileIDs(tiles, append(edgeIDs, cornerIDs...))

	startTile := tiles.asMap[cornerIDs[0]]

	startTile.grid.Rotate(-90)

	b := startTile.grid.BottomEdge()
	r := startTile.grid.RightEdge()

	for !(tiles.edgeCounts[r] == 2 && tiles.edgeCounts[b] == 2) {
		startTile.grid.Rotate(-90)
		b = startTile.grid.BottomEdge()
		r = startTile.grid.RightEdge()
	}

	maxIndex := int(math.Sqrt(float64(len(tiles.asSlice)))) - 1
	var orgGrid = [][]int{}
	// Init Grid
	for y := 0; y <= maxIndex; y++ {
		row := []int{}
		for x := 0; x <= maxIndex; x++ {
			row = append(row, 0)
		}
		orgGrid = append(orgGrid, row)
	}
	orgGrid[0][0] = startTile.id

	cornerIDs = util.RemoveIntFromSlice(cornerIDs, startTile.id)

	// ******************************
	// resolve top row tiles
	// ******************************
	for x := 1; x <= maxIndex; x++ {
		// Find the next tile to the right
		prevTile := tiles.asMap[orgGrid[x-1][0]]
		prevTileR := prevTile.grid.RightEdge()

		var slice *[]int
		if x != maxIndex {
			// Looking for edges
			slice = &edgeIDs
		} else {
			// looking for next corner
			slice = &cornerIDs
		}

		tileID := findTileIDWithMatchingEdge(tiles, prevTileR, *slice)
		tile := tiles.asMap[tileID]
		alignGridLeft(prevTileR, tile.grid)
		*slice = util.RemoveIntFromSlice(*slice, tile.id)
		orgGrid[x][0] = tile.id
	}

	// ******************************
	// resolve middle rows tiles
	// ******************************
	var prevTile tile
	for y := 1; y <= maxIndex-1; y++ {
		for x := 0; x <= maxIndex; x++ {
			if x == 0 {
				prevTile = tiles.asMap[orgGrid[x][y-1]]
				slice := &edgeIDs
				prevTileB := prevTile.grid.BottomEdge()
				tileID := findTileIDWithMatchingEdge(tiles, prevTileB, *slice)
				tile := tiles.asMap[tileID]
				alignGridTop(prevTileB, tile.grid)
				*slice = util.RemoveIntFromSlice(*slice, tile.id)
				orgGrid[x][y] = tile.id
				prevTile = tile
				continue
			}

			var slice *[]int
			if x == maxIndex {
				slice = &edgeIDs
			} else {
				slice = &midIDs
			}

			aboveTileB := tiles.asMap[orgGrid[x][y-1]].grid.BottomEdge()
			prevTileR := prevTile.grid.RightEdge()
			tileID := findTileIDWithMatchingLeftTop(tiles, prevTileR, aboveTileB, *slice)
			tile := tiles.asMap[tileID]
			// findTileIDWithMatchingLeftTop(tiles, prevTileR, aboveTileB, *slice)
			alignGridLeftTop(prevTileR, aboveTileB, tile.grid)
			*slice = util.RemoveIntFromSlice(*slice, tile.id)
			orgGrid[x][y] = tile.id
			prevTile = tile
		}
	}

	// ******************************
	// resolve bottom rows tiles
	// ******************************
	for x := 0; x <= maxIndex; x++ {
		y := maxIndex
		if x == 0 {
			prevTile = tiles.asMap[orgGrid[x][y-1]]
			slice := &cornerIDs
			prevTileB := prevTile.grid.BottomEdge()
			tileID := findTileIDWithMatchingEdge(tiles, prevTileB, *slice)
			tile := tiles.asMap[tileID]
			alignGridTop(prevTileB, tile.grid)
			*slice = util.RemoveIntFromSlice(*slice, tile.id)
			orgGrid[x][y] = tile.id
			prevTile = tile
			continue
		}

		var slice *[]int
		if x == maxIndex {
			slice = &cornerIDs
		} else {
			slice = &edgeIDs
		}

		aboveTileB := tiles.asMap[orgGrid[x][y-1]].grid.BottomEdge()
		prevTileR := prevTile.grid.RightEdge()
		tileID := findTileIDWithMatchingLeftTop(tiles, prevTileR, aboveTileB, *slice)
		tile := tiles.asMap[tileID]
		// findTileIDWithMatchingLeftTop(tiles, prevTileR, aboveTileB, *slice)
		alignGridLeftTop(prevTileR, aboveTileB, tile.grid)
		*slice = util.RemoveIntFromSlice(*slice, tile.id)
		orgGrid[x][y] = tile.id
		prevTile = tile
	}

	for tileY := 0; tileY <= maxIndex; tileY++ {
		for tileX := 0; tileX <= maxIndex; tileX++ {
			tile := tiles.asMap[orgGrid[tileX][tileY]]
			// tile.grid.Dump()
			tile.grid.Shrink(1)
			// fmt.Println()
			// tile.grid.Dump()
			// fmt.Println()
			// fmt.Println()
			// if tile.id == 3079 {
			// 	row := tile.grid.GetRow(tile.grid.GetMaxY())
			// 	fmt.Println(row)
			// }
		}
	}

	// ******************************
	// Printout final grid
	// ******************************
	finalX := 0
	finalY := 0
	finalGrid := util.NewInfinityGrid("X")
	for tileY := 0; tileY <= maxIndex; tileY++ {
		for y := startTile.grid.GetMaxY(); y >= startTile.grid.GetMinY(); y-- {
			for tileX := 0; tileX <= maxIndex; tileX++ {
				tileID := orgGrid[tileX][tileY]
				if tileID == 0 {
					break
				}
				grid := tiles.asMap[tileID].grid

				row := grid.GetRow(y)
				for _, c := range row {
					finalGrid.Set(c, finalX, finalY)
					finalX++
					// fmt.Print(c)
				}
				// fmt.Print(" ")
			}
			finalY++
			finalX = 0
			// fmt.Println()
		}
		// fmt.Println()
	}

	// finalGrid.FlipH()
	// finalGrid.Rotate(-90)
	finalGrid.Dump()

	seamonsterGrid := util.NewInfinityGridFromFile("seamonster.txt", "X")
	seamonsterGrid.FlipH()
	// seamonsterGrid.Dump()

	// finalGrid.Rotate(90)
	// finalGrid.Rotate(90)
	// finalGrid.Rotate(90)
	// finalGrid.Rotate(90)
	// finalGrid.FlipH()
	// finalGrid.Rotate(90)
	// finalGrid.Rotate(90)
	// finalGrid.Rotate(90)
	// finalGrid.Rotate(90)
	// finalGrid.FlipV()
	// finalGrid.Rotate(90)
	// finalGrid.Rotate(90)
	// finalGrid.Rotate(90)
	// finalGrid.Rotate(90)

	count := findSomeSeaMonsters(finalGrid, seamonsterGrid)
	numSeaMonsterSquares := 0
	seamonsterGrid.VisitAll2D(func(val string, x, y int) {
		if val == "#" {
			numSeaMonsterSquares++
		}
	})
	numSeaMonsterSquares *= count

	totalWaters := 0
	finalGrid.VisitAll2D(func(val string, x, y int) {
		if val == "#" {
			totalWaters++
		}
	})

	return totalWaters - numSeaMonsterSquares
}

func findSomeSeaMonsters(finalGrid, seamonsterGrid *util.InfinityGrid) int {
	rotCount := 0
	flipped := false
	count := 0
	for count == 0 {
		for y := 0; y <= finalGrid.GetMaxY()-seamonsterGrid.GetMaxY(); y++ {
			for x := 0; x <= finalGrid.GetMaxX()-seamonsterGrid.GetMaxX(); x++ {
				if findFullSeaMonster(finalGrid, seamonsterGrid, x, y) {
					count++
				}
			}
		}

		rotCount++
		if rotCount > 3 && flipped {
			panic("not good!")
		}

		if rotCount > 3 {
			finalGrid.FlipV()
			// finalGrid.FlipH() // If panic flip this
			rotCount = 0
			flipped = true
			continue
		}

		finalGrid.Rotate(-90)
	}

	return count
}

func findFullSeaMonster(finalGrid, seamonsterGrid *util.InfinityGrid, x, y int) bool {
	for sy := seamonsterGrid.GetMinY(); sy <= seamonsterGrid.GetMaxY(); sy++ {
		for sx := seamonsterGrid.GetMinX(); sx <= seamonsterGrid.GetMaxX(); sx++ {
			val := seamonsterGrid.Get(sx, sy)
			if val == "#" {
				if finalGrid.Get(x+sx, y+sy) != "#" {
					return false
				}
			}
		}
	}

	return true
}
