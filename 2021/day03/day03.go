package day03

import (
	"aoc/util"
	"strconv"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	ones := make([]int, len(data[0]))

	for _, line := range data {
		for i, char := range line {
			if string(char) == "1" {
				ones[i]++
			}
		}
	}

	gamma := ""
	epsilon := ""
	for _, numOnes := range ones {
		numZeros := len(data) - numOnes
		if numOnes > numZeros {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}

	}

	gammaInt, _ := strconv.ParseUint(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseUint(epsilon, 2, 64)

	return int(gammaInt * epsilonInt)
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	oxygenEligible := data
	co2Eligible := data

	for i := 0; i < len(data[0]); i++ {
		oxygenDigitToKeep := digitToKeep(oxygenEligible, i, '1')
		oxygenEligible = keepers(oxygenEligible, oxygenDigitToKeep, i)

		co2DigitToKeep := digitToKeep(co2Eligible, i, '0')
		co2Eligible = keepers(co2Eligible, co2DigitToKeep, i)
	}

	oxy, _ := strconv.ParseUint(oxygenEligible[0], 2, 64)
	co2, _ := strconv.ParseUint(co2Eligible[0], 2, 64)

	return int(oxy * co2)
}

func digitToKeep(list []string, pos int, retGEOnes byte) byte {
	ones := numOnesAtPos(list, pos)
	zeros := len(list) - ones
	if ones >= zeros {
		return retGEOnes
	}

	if retGEOnes == '0' {
		return '1'
	}

	return '0'

}

func keepers(list []string, digitToKeep byte, pos int) []string {
	if len(list) > 1 {
		newList := []string{}
		for _, line := range list {
			if line[pos] == digitToKeep {
				newList = append(newList, line)
			}
		}

		list = newList
	}

	return list
}

func numOnesAtPos(list []string, pos int) int {
	count := 0
	for _, line := range list {
		if string(line[pos]) == "1" {
			count++
		}
	}

	return count
}
