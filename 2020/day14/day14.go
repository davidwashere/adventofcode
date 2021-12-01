package day14

import (
	"aoc2020/util"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	mem := map[int]int{}
	mask := ""
	for _, line := range data {
		keyValS := strings.Split(line, " = ")
		key := keyValS[0]
		valS := keyValS[1]

		if key == "mask" {
			mask = valS
			continue
		}

		// Assume mem[]
		key = key[4:]                                // remove mem[
		memAddr, _ := strconv.Atoi(key[:len(key)-1]) // remove ]

		val, _ := strconv.Atoi(valS)

		mem[memAddr] = flipBits(mask, val)

	}
	sumOfMemValues := 0
	for _, val := range mem {
		sumOfMemValues += val
	}
	return sumOfMemValues
}

func flipBits(mask string, num int) int {
	andMaskS := strings.ReplaceAll(mask, "X", "1")
	orMaskS := strings.ReplaceAll(mask, "X", "0")
	andMask, _ := strconv.ParseInt(andMaskS, 2, 64)
	orMask, _ := strconv.ParseInt(orMaskS, 2, 64)

	nnum := int64(num) & andMask
	nnum = nnum | orMask

	return int(nnum)
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	mem := map[int]int{}
	mask := ""
	for _, line := range data {
		keyValS := strings.Split(line, " = ")
		key := keyValS[0]
		valS := keyValS[1]

		if key == "mask" {
			mask = valS
			continue
		}

		key = key[4:]                                // remove mem[
		memAddr, _ := strconv.Atoi(key[:len(key)-1]) // remove ]

		val, _ := strconv.Atoi(valS)

		addresses := perms(mask, flipBits2(mask, memAddr))
		for _, address := range addresses {
			mem[address] = val
		}

	}
	sumOfMemValues := 0
	for _, val := range mem {
		sumOfMemValues += val
	}
	return sumOfMemValues
}

func flipBits2(mask string, num int) int {
	orMaskS := strings.ReplaceAll(mask, "X", "0")
	orMask, _ := strconv.ParseInt(orMaskS, 2, 64)

	andMaskS := strings.ReplaceAll(mask, "0", "1")
	andMaskS = strings.ReplaceAll(andMaskS, "X", "0")
	andMask, _ := strconv.ParseInt(andMaskS, 2, 64)

	nnum := int64(num) | orMask
	nnum = nnum & andMask

	return int(nnum)
}

func perms(mask string, num int) []int {
	numX := float64(strings.Count(mask, "X"))
	max := int(math.Pow(2, numX))

	offsets := []int{}
	for i := len(mask) - 1; i >= 0; i-- {
		if mask[len(mask)-1-i] == 'X' {
			offsets = append(offsets, i)
		}
	}

	sort.Ints(offsets)
	fmt.Println(offsets)

	perms := []int{}

	for i := 0; i < max; i++ {
		orMask := 0
		binS := fmt.Sprintf("%b", i)
		numBinDigits := len(binS)
		for j := 0; j < numBinDigits; j++ {
			lastIndex := len(binS) - 1
			digit, _ := strconv.Atoi(string(binS[lastIndex]))
			binS = binS[:lastIndex]
			if digit == 1 {
				shift := digit << offsets[j]
				orMask = orMask | shift
			}
		}

		result := num | orMask

		fmt.Printf("%b | %b = %b (decimal %d)\n", num, orMask, result, result)

		perms = append(perms, result)

	}

	return perms
}
