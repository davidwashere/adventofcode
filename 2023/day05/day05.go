package day05

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

var (
	seeds = []int{}
	maps  = []*Map{}
)

type Map struct {
	ranges []*Range
}

func (m *Map) ToDest(src int) int {
	for _, rng := range m.ranges {
		l := rng.srcStart
		r := rng.srcStart + rng.length - 1

		if src < l || src > r {
			continue
		}

		diff := src - l

		return rng.dstStart + diff
	}

	return src
}

type Range struct {
	dstStart int
	srcStart int
	length   int
}

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	for i, line := range data {
		line = strings.TrimSpace(line)
		tokens := util.ParseTokens(line)

		if i == 0 {
			seeds = tokens.Ints
			continue
		}

		if line == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			maps = append(maps, &Map{})
			continue
		}

		r := &Range{tokens.Ints[0], tokens.Ints[1], tokens.Ints[2]}
		m := maps[len(maps)-1]
		m.ranges = append(m.ranges, r)
	}
}

func part1(inputFile string) int {
	load(inputFile)

	min := util.MaxInt
	for _, src := range seeds {
		for _, m := range maps {
			src = m.ToDest(src)
		}

		min = util.Min(min, src)
	}

	return min
}

type workerResult struct {
	id  string
	min int
}

func part2(inputFile string) int {
	load(inputFile)
	min := util.MaxInt
	workers := 0

	results := make(chan workerResult)
	for i := 0; i < len(seeds); i = i + 2 {
		start := seeds[i]
		end := seeds[i] + seeds[i+1] - 1

		workers++
		id := strconv.Itoa(start)
		go calcMin(id, start, end, results)
	}

	for val := range results {
		min = util.Min(min, val.min)
		fmt.Printf("Worker %v done, returned %v (min: %v)\n", val.id, val.min, min)
		workers--
		if workers == 0 {
			close(results)
		}
	}

	return min
}

func calcMin(id string, start, end int, result chan workerResult) {
	min := util.MaxInt
	for j := start; j <= end; j++ {
		src := j
		for _, m := range maps {
			src = m.ToDest(src)
		}
		min = util.Min(min, src)
	}

	result <- workerResult{id, min}
}
