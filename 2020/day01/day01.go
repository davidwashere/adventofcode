package day01

import "aoc2020/util"

func part1(inputfile string) int {
	data, err := util.ReadFileToIntSlice(inputfile)
	util.Check(err)

	for i := 0; i < len(data); i++ {
		left := data[i]

		for j := i + 1; j < len(data); j++ {
			right := data[j]

			if left+right == 2020 {
				return left * right
			}
		}
	}

	return -1
}

func part2(inputfile string) int {
	data, err := util.ReadFileToIntSlice(inputfile)
	util.Check(err)

	for i := 0; i < len(data); i++ {
		left := data[i]

		for j := i + 1; j < len(data); j++ {
			right := data[j]

			for k := i + 2; k < len(data); k++ {
				mid := data[k]

				if left+right+mid == 2020 {
					return left * right * mid
				}
			}

		}
	}

	return -1
}
