package day08

import (
	"aoc/util"
	"aoc/util/disjointset"
	"sort"
)

type dayItem struct {
	id    int
	pt    util.Point3
	dists map[int]float64
}

var (
	items = []dayItem{}
	grid  = util.NewInfGrid[string]().WithDefaultValue(".")
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	for i, line := range data {
		tokens := util.ParseTokens(line)

		items = append(items, dayItem{
			id: i,
			pt: util.NewPoint3(
				tokens.Ints[0],
				tokens.Ints[1],
				tokens.Ints[2],
			),
			dists: map[int]float64{},
		})
	}
}

func part1(inputFile string, connections int) int {
	load(inputFile)

	type edge struct {
		p1   int
		p2   int
		dist float64
	}
	edges := []edge{}

	// sort coord sets by distance
	for i := 0; i < len(items)-1; i++ {
		for k := i + 1; k < len(items); k++ {
			l := items[i]
			r := items[k]
			dist := l.pt.Dist(r.pt)
			edges = append(edges, edge{
				p1:   i,
				p2:   k,
				dist: dist,
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		l := edges[i]
		r := edges[j]

		return l.dist < r.dist
	})

	ds := disjointset.New[int]()
	for i, edge := range edges {
		if i == connections {
			break
		}
		ds.Union(edge.p1, edge.p2)
	}

	// List all graphs
	allgraphs := ds.GetAllGraphs()
	sizes := []int{}
	for _, v := range allgraphs {
		sizes = append(sizes, len(v))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func part2(inputFile string) int {
	load(inputFile)

	type edge struct {
		p1   int
		p2   int
		dist float64
	}
	edges := []edge{}

	// sort coord sets by distance
	for i := 0; i < len(items)-1; i++ {
		for k := i + 1; k < len(items); k++ {
			l := items[i]
			r := items[k]
			dist := l.pt.Dist(r.pt)
			edges = append(edges, edge{
				p1:   i,
				p2:   k,
				dist: dist,
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		l := edges[i]
		r := edges[j]

		return l.dist < r.dist
	})

	ds := disjointset.New[int]()
	for _, edge := range edges {
		ds.Union(edge.p1, edge.p2)
		numnodes := ds.NumNodes()
		numgraphs := ds.NumGraphs()
		if numnodes == len(items) && numgraphs == 1 {

			return items[edge.p1].pt.X * items[edge.p2].pt.X
		}
	}

	return 0
}
