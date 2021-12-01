package day09

import (
	"aoc2020/util"
)

func part1(inputfile string, bufferSize int) int {
	data, err := util.ReadFileToIntSlice(inputfile)
	util.Check(err)

	buffer := []int{}
	for i := 0; i < bufferSize; i++ {
		buffer = append(buffer, data[i])
	}

	for i := bufferSize; i < len(data); i++ {
		val := data[i]
		if !twoNumsSumToVal(buffer, val) {
			return val
		}
		buffer = append(buffer[1:], val)
	}

	return -1
}

// twoNumsSumToVal returns true if combination of two nums in buffer sum to val, false otherwise
func twoNumsSumToVal(buffer []int, val int) bool {
	for j := 0; j < len(buffer); j++ {
		for k := 1; k < len(buffer); k++ {
			left := buffer[j]
			right := buffer[k]

			if left+right == val && left != right {
				return true
			}
		}
	}
	return false
}

func part2(inputfile string, bufferSize int) int {
	find := part1(inputfile, bufferSize)

	// Parsing file again is inefficient, but hey! it's AoC!
	data, err := util.ReadFileToIntSlice(inputfile)
	util.Check(err)

	for i := 0; i < len(data); i++ {
		sum := data[i]
		for j := i + 1; j < len(data); j++ {
			sum += data[j]

			if sum == find {
				nums := data[i:j]
				min := util.MinAll(nums...)
				max := util.MaxAll(nums...)
				return min + max
			}

			if sum > find {
				break
			}
		}
	}
	return 0
}
