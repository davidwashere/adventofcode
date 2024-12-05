package day05

import (
	"aoc/util"
)

var (
	rules   = [][]int{}
	updates = [][]int{}
	nodes   = map[int]node{}
)

type node struct {
	val      int
	children map[int]bool
	parents  map[int]bool
}

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	// data, _ := util.ReadFileToIntSlice(inputFile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputFile, "\n\n")
	// grid := util.NewInfinityGridFromFile(inputfile, ".")

	onRules := true
	for _, line := range data {
		if line == "" {
			onRules = false
			continue
		}

		// tokens := util.ParseTokens(line)
		ints := util.ParseInts(line)
		// strs := util.ParseStrs(line)
		// words := util.ParseWords(line)

		if onRules {
			rules = append(rules, ints)
			continue
		}

		updates = append(updates, ints)
	}

	for _, rule := range rules {
		l := rule[0]
		r := rule[1]

		if _, ok := nodes[l]; !ok {
			nodes[l] = node{val: l, children: map[int]bool{}, parents: map[int]bool{}}
		}
		if _, ok := nodes[r]; !ok {
			nodes[r] = node{val: r, children: map[int]bool{}, parents: map[int]bool{}}
		}
		nodes[l].children[r] = true
		nodes[r].parents[l] = true
	}

	// fmt.Printf("Rules: %v\n", rules)
	// fmt.Printf("Updes: %v\n", updates)
	// fmt.Printf("Nodes: %v\n", nodes)
}

func part1(inputFile string) int {
	load(inputFile)

	sum := 0
	for _, pages := range updates {
		bad := false
		for i := 0; i < len(pages); i++ {
			l := pages[i]
			for j := i + 1; j < len(pages); j++ {
				r := pages[j]
				if !nodes[l].children[r] {
					bad = true
					i = len(pages)
					break
				}
			}
		}

		if !bad {
			sum += pages[(len(pages) / 2)]
		}
	}

	return sum
}

func part2(inputFile string) int {
	load(inputFile)

	sum := 0
	for _, pages := range updates {
		bad := false
		for i := 0; i < len(pages); i++ {
			l := pages[i]
			for j := i + 1; j < len(pages); j++ {
				r := pages[j]
				if !nodes[l].children[r] {
					bad = true
					i = len(pages)
					break
				}
			}
		}

		if !bad {
			// sum += pages[(len(pages) / 2)]
			continue
		}

		// reorder

		for i := 0; i < len(pages); i++ {
			for j := 0; j < len(pages)-i-1; j++ {
				l := pages[j]
				r := pages[j+1]

				if nodes[l].parents[r] {
					// nodes in wrong order
					pages[j] = r
					pages[j+1] = l
				}
			}
		}

		sum += pages[(len(pages) / 2)]
	}

	return sum
}
