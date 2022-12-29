package day23

import (
	"aoc/util"
	"fmt"
)

type Elf util.Point

var (
	VectorNorth = util.NewNormalizedVector(0, -1)
	VectorSouth = util.NewNormalizedVector(0, 1)
	VectorEast  = util.NewNormalizedVector(1, 0)
	VectorWest  = util.NewNormalizedVector(-1, 0)

	VectorNorthEast = util.NewNormalizedVector(1, -1)
	VectorSouthEast = util.NewNormalizedVector(1, 1)
	VectorSouthWest = util.NewNormalizedVector(-1, 1)
	VectorNorthWest = util.NewNormalizedVector(-1, -1)
)

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

func load(inputFile string) (*util.InfGrid[bool], []Elf) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	maxX := util.MinInt

	elfs := []Elf{}

	grid := util.NewInfGrid[bool]()
	for y := 0; y < len(data); y++ {
		line := data[y]
		for x := 0; x < len(line); x++ {
			c := string(line[x])

			if c == "#" {
				elfs = append(elfs, Elf{x, y})
				grid.Set(true, x, y)
			}

			maxX = util.Max(maxX, x)
		}
	}

	grid.WithDumpFunc(gridDumpFunc)

	return grid, elfs
}

var (
	checks = map[util.Vector][]util.Vector{
		VectorNorth: {VectorNorth, VectorNorthEast, VectorNorthWest},
		VectorSouth: {VectorSouth, VectorSouthEast, VectorSouthWest},
		VectorWest:  {VectorWest, VectorNorthWest, VectorSouthWest},
		VectorEast:  {VectorEast, VectorNorthEast, VectorSouthEast},
	}
)

func createNewGrid(elfs []Elf) *util.InfGrid[bool] {
	grid := util.NewInfGrid[bool]()

	for _, elf := range elfs {
		grid.Set(true, elf.X, elf.Y)
	}

	grid.WithDumpFunc(gridDumpFunc)
	return grid
}

func areEmpty(grid *util.InfGrid[bool], elf Elf, vecs []util.Vector) bool {
	for _, v := range vecs {
		p := util.Point(elf).Apply(v)
		if grid.Get(p.X, p.Y) {
			// is other elf here
			return false
		}

		// also do bounds check?
	}

	return true
}

func part1(inputFile string) int {
	grid, elfs := load(inputFile)

	grid.FlipH()
	grid.Dump()
	grid.FlipH()
	fmt.Println()

	proposalOrder := []util.Vector{}
	proposalOrder = append(proposalOrder, VectorNorth)
	proposalOrder = append(proposalOrder, VectorSouth)
	proposalOrder = append(proposalOrder, VectorWest)
	proposalOrder = append(proposalOrder, VectorEast)

	// fmt.Println(elfs)
	rounds := 10

	for r := 0; r < rounds; r++ {
		proposals := []util.Point{}
		// gather proposals
		for _, elf := range elfs {

			hasNeighbor := false
			hasProposal := false
			newProposal := util.Point(elf)
			for _, proposal := range proposalOrder {
				if !areEmpty(grid, elf, checks[proposal]) {
					hasNeighbor = true
					continue
				}

				if !hasProposal {
					hasProposal = true
					newProposal = util.Point(elf).Apply(proposal)
				}
			}

			if !hasNeighbor {
				// stay put
				proposals = append(proposals, util.Point(elf))
			} else {
				proposals = append(proposals, newProposal)
			}

		}

		// adjust proposals if overlapping
		proposals = adjustProposals(elfs, proposals)

		// do the moves
		for i := 0; i < len(elfs); i++ {
			elfs[i] = Elf(proposals[i])
		}

		// update grid
		grid = createNewGrid(elfs)

		// change moves order
		t := proposalOrder[0]
		proposalOrder = append(proposalOrder[1:], t)

		// fmt.Printf("===== Rnd %v =====\n", r+1)
		// grid.FlipH()
		// grid.Dump()
		// grid.FlipH()
		// fmt.Println()
	}

	count := 0
	grid.VisitAll2D(func(val bool, x, y int) {
		if !val {
			count++
		}
	})

	return count
}

func adjustProposals(elfs []Elf, proposals []util.Point) []util.Point {
	counts := map[util.Point]int{}
	for _, p := range proposals {
		counts[p]++
	}

	updProposals := make([]util.Point, len(proposals))
	for i := 0; i < len(proposals); i++ {
		pro := proposals[i]
		if counts[pro] > 1 {
			// put the elf back where it started
			updProposals[i] = util.Point(elfs[i])
			continue
		}

		updProposals[i] = proposals[i]
	}

	return updProposals
}

func part2(inputFile string) int {
	grid, elfs := load(inputFile)

	grid.FlipH()
	grid.Dump()
	grid.FlipH()
	fmt.Println()

	proposalOrder := []util.Vector{}
	proposalOrder = append(proposalOrder, VectorNorth)
	proposalOrder = append(proposalOrder, VectorSouth)
	proposalOrder = append(proposalOrder, VectorWest)
	proposalOrder = append(proposalOrder, VectorEast)

	// fmt.Println(elfs)
	r := 0
	for {
		r++
		stayingPut := 0
		proposals := []util.Point{}
		// gather proposals
		for _, elf := range elfs {

			hasNeighbor := false
			hasProposal := false
			newProposal := util.Point(elf)
			for _, proposal := range proposalOrder {
				if !areEmpty(grid, elf, checks[proposal]) {
					hasNeighbor = true
					continue
				}

				if !hasProposal {
					hasProposal = true
					newProposal = util.Point(elf).Apply(proposal)
				}
			}

			if !hasNeighbor {
				// stay put
				stayingPut++
				proposals = append(proposals, util.Point(elf))
			} else {
				proposals = append(proposals, newProposal)
			}

		}

		if stayingPut == len(elfs) {
			return r
		}

		// adjust proposals if overlapping
		proposals = adjustProposals(elfs, proposals)

		// do the moves
		for i := 0; i < len(elfs); i++ {
			elfs[i] = Elf(proposals[i])
		}

		// update grid
		grid = createNewGrid(elfs)

		// change moves order
		t := proposalOrder[0]
		proposalOrder = append(proposalOrder[1:], t)

		// fmt.Printf("===== Rnd %v =====\n", r+1)
		// grid.FlipH()
		// grid.Dump()
		// grid.FlipH()
		// fmt.Println()
	}

	count := 0
	grid.VisitAll2D(func(val bool, x, y int) {
		if !val {
			count++
		}
	})

	return count
}
