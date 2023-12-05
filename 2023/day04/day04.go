package day04

import (
	"aoc/util"
	"strings"
)

var (
	cards = []Card{}
)

type Card struct {
	winners map[int]bool
	nums    []int
	copies  int
}

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	for _, line := range data {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		card := Card{
			winners: make(map[int]bool),
		}

		line = strings.Split(line, ": ")[1]
		winNums := strings.Split(line, " | ")

		for _, winner := range util.ParseInts(winNums[0]) {
			card.winners[winner] = true
		}

		card.nums = append(card.nums, util.ParseInts(winNums[1])...)
		card.copies = 1

		cards = append(cards, card)
	}

}

func part1(inputFile string) int {
	load(inputFile)

	sum := 0
	for _, card := range cards {
		cardSum := 0
		for _, num := range card.nums {
			if card.winners[num] {
				if cardSum == 0 {
					cardSum = 1
				} else {
					cardSum *= 2
				}
			}
		}

		sum += cardSum
	}

	return sum
}

func part2(inputFile string) int {
	load(inputFile)

	for i, card := range cards {
		numWins := 0
		for _, num := range card.nums {
			if card.winners[num] {
				numWins++
			}
		}

		for k := i + 1; k < (i + 1 + numWins); k++ {
			cards[k].copies += card.copies
		}
	}

	sum := 0
	for _, card := range cards {
		sum += card.copies
	}

	return sum
}
