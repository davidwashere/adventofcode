package day25

import (
	"aoc2020/util"
	"fmt"
	"strconv"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	cardPubKey, _ := strconv.Atoi(data[0])
	doorPubKey, _ := strconv.Atoi(data[1])

	loops := 0
	val := 1
	subNum := 7
	for {
		loops++
		val = val * subNum
		val = val % 20201227

		if val == cardPubKey || val == doorPubKey {
			break
		}
	}
	fmt.Printf("Loops: %v, Key: %v\n", loops, val)

	keyToUse := 0
	if val == doorPubKey {
		// have doors loop size
		keyToUse = cardPubKey
	} else {
		// have cards loop size
		keyToUse = doorPubKey
	}

	val = 1
	subNum = keyToUse
	for i := 0; i < loops; i++ {
		val = val * subNum
		val = val % 20201227
	}

	return val
}

func part2(inputfile string) int {
	return 0
}
