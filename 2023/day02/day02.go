package day02

import (
	"aoc/util"
	"strings"
)

var (
	games []map[string]int
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	games = []map[string]int{}
	for _, line := range data {
		sp := strings.Split(line, ": ")
		if len(sp) < 2 {
			continue
		}

		line = sp[1]
		game := map[string]int{}

		pulls := strings.Split(line, "; ")

		for _, pull := range pulls {
			cubes := strings.Split(pull, ", ")
			for _, cube := range cubes {
				tokens := util.ParseTokens(cube)

				num := tokens.Ints[0]
				color := tokens.Words[0]
				if num > game[color] {
					game[color] = num
				}
			}

		}

		games = append(games, game)
	}
}

func part1(inputFile string) int {
	load(inputFile)

	colorLimits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	sum := 0
	for i, game := range games {
		good := true
		for color, limit := range colorLimits {
			if game[color] > limit {
				good = false
				break
			}
		}

		if good {
			sum += (i + 1)
		}
	}

	return sum
}

func part2(inputFile string) int {
	load(inputFile)

	sum := 0
	for _, game := range games {
		power := 1
		for _, limit := range game {
			power *= limit
		}

		sum += power
	}

	return sum
}
