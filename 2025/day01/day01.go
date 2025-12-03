package day01

import (
	"aoc/util"
	"fmt"
)

type singleStep struct {
	dir  string
	dist int
}

var (
	steps = []singleStep{}
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	// data, _ := util.ReadFileToIntSlice(inputFile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputFile, "\n\n")
	// grid := util.NewInfinityGridFromFile(inputfile, ".")

	for _, line := range data {
		tokens := util.ParseTokens(line)
		// ints := util.ParseInts(line)
		// strs := util.ParseStrs(line)
		// words := util.ParseWords(line)
		step := singleStep{
			dir:  string(line[0]),
			dist: tokens.Ints[0],
		}

		steps = append(steps, step)

		// fmt.Println(tokens)
	}
	fmt.Println(steps)
}

func part1(inputFile string) int {
	load(inputFile)

	max := 99
	pos := 50
	count := 0
	for i, step := range steps {
		if step.dir == "L" {
			pos = (pos - step.dist) % (max + 1)
			if pos < 0 {
				pos = (max + 1) + pos
			}
		} else {
			pos = (pos + step.dist) % (max + 1)
		}

		fmt.Printf("[%d] %d\n", i, pos)
		if pos == 0 {
			count++
		}
	}

	return count
}

func part2(inputFile string) int {
	load(inputFile)

	max := 100
	pos := 50
	count := 0
	for _, step := range steps {
		start := pos
		// every 100 dist = guaranteed 1 click
		if step.dist >= max {
			count += step.dist / int(100)
		}
		step.dist = step.dist % max

		if step.dir == "L" {
			pos = pos - step.dist

			if pos == 0 {
				count++
			} else if pos < 0 {
				if start != 0 {
					count++
				}

				pos = max + pos // wraps around (pos is -, so this is effectively max - pos)
			}
		} else {
			pos = pos + step.dist

			if pos > 99 {
				count++
				pos = pos - max
			}
		}

		fmt.Printf("[%s%d] %d - clicks: %d\n", step.dir, step.dist, pos, count)
	}

	return count

}
