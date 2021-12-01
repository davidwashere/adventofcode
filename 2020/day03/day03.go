package day03

import (
	"aoc2020/util"
	"fmt"
)

func part1(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	incX := 3
	incY := 1

	curX := 3
	curY := 1

	maxWidth := len(data[0])
	maxHeight := len(data)

	trees := 0

	for curY < maxHeight {
		row := data[curY]
		char := string(row[curX])

		if char == "#" {
			trees++
		}

		curX = (curX + incX) % maxWidth
		curY += incY
		fmt.Println(incX, incY, curX, curY, maxWidth, maxHeight)
	}

	return trees
}

// Pair .
type Pair struct {
	incX  int
	incY  int
	trees int
}

func part2(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	pairs := []*Pair{
		&Pair{1, 1, 0},
		&Pair{3, 1, 0},
		&Pair{5, 1, 0},
		&Pair{7, 1, 0},
		&Pair{1, 2, 0},
	}

	maxWidth := len(data[0])
	maxHeight := len(data)

	for _, pair := range pairs {
		curY := pair.incY
		curX := pair.incX
		incY := pair.incY
		incX := pair.incX
		for curY < maxHeight {
			row := data[curY]
			char := string(row[curX])

			if char == "#" {
				pair.trees = pair.trees + 1
			}

			curX = (curX + incX) % maxWidth
			curY += incY
		}

	}

	var result int
	result = 1
	for _, pair := range pairs {
		result = result * pair.trees
	}

	return result
}
