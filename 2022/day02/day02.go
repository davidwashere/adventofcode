package day02

import (
	"aoc/util"
)

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3

	OutcomeLoss = "X"
	OutcomeDraw = "Y"
	OutcomeWin  = "Z"
)

var decoderRing = map[string]int{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var outcomePts = map[int]int{
	0:  0,
	-1: 3,
	1:  6,
}

var outcomePtsP2 = map[string]int{
	OutcomeDraw: 3,
	OutcomeLoss: 0,
	OutcomeWin:  6,
}

// outcomeMap will contain an outcome mapped to elf choice mapped to your choice
var outcomeMap = map[string]map[int]int{
	OutcomeDraw: {
		Rock:     Rock,
		Paper:    Paper,
		Scissors: Scissors,
	},
	OutcomeLoss: {
		Rock:     Scissors,
		Paper:    Rock,
		Scissors: Paper,
	},
	OutcomeWin: {
		Rock:     Paper,
		Paper:    Scissors,
		Scissors: Rock,
	},
}

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	result := 0
	for _, line := range data {
		strs := util.ParseStrs(line)

		otherElf := strs[0]
		you := strs[1]

		yourChoice := decoderRing[you]

		outcome := decide(decoderRing[otherElf], yourChoice)

		result += outcomePts[outcome] + yourChoice
	}

	return result
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	result := 0
	for _, line := range data {
		strs := util.ParseStrs(line)

		elf := strs[0]
		desiredOutcome := strs[1]

		yourChoice := outcomeMap[desiredOutcome][decoderRing[elf]]
		result += outcomePtsP2[desiredOutcome] + yourChoice

		// yourChoice := decoderRing[you]

		// outcome := decide(decoderRing[otherElf], yourChoice)

		// result += outcomePts[outcome] + yourChoice
	}

	return result
}

// decide returns -1 for tie, 0 if otherElf wins, 1 if you wins
func decide(otherElf, you int) int {
	if otherElf == you {
		return -1
	}

	if otherElf == Rock {
		if you == Paper {
			// you win
			return 1
		}

		if you == Scissors {
			// you lose
			return 0
		}
	}

	if otherElf == Paper {
		if you == Rock {
			// you lose
			return 0
		}

		if you == Scissors {
			// you win
			return 1
		}
	}

	if otherElf == Scissors {
		if you == Rock {
			// you win
			return 1
		}

		if you == Paper {
			// you lose
			return 0
		}
	}

	return -1
}
