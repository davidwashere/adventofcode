package day25

import (
	"aoc/util"
	"fmt"
	"math"
)

var (
	digitMap = map[byte]int{
		'2': 2,
		'1': 1,
		'0': 0,
		'-': -1,
		'=': -2,
	}

	snafMap = map[int]string{
		0: "0",
		1: "1",
		2: "2",
		3: "=",
		4: "-",
	}
)

func snafuToDec(snafu string) int {
	size := len(snafu)
	// l = 5
	// i = 0 > 4
	// i = 1 > 3
	// i = 2 > 2
	// i = 3 > 1
	// i = 4 > 0

	total := 0
	for i := 0; i < len(snafu); i++ {
		exp := (size - 1 - i)
		multi := int(math.Pow(float64(5), float64(exp)))
		dig := digitMap[snafu[i]]
		total += multi * dig
	}

	return total
}

func decToSnafu(dec int) string {

	carry := false
	snafu := ""
	for dec >= 0 {
		quo := dec / 5
		rem := dec % 5

		if carry {
			rem++

			if rem == 5 {
				quo++
				rem = 0
			}
		}

		dig := snafMap[rem]
		snafu = fmt.Sprintf("%v%v", dig, snafu)
		dec = quo
		carry = rem >= 3

		if dec == 0 && !carry {
			break
		}
	}

	// 0 = 0
	// 1 = 1
	// 2 = 2
	// 3 = 1=
	// 4 = 1-
	// 5 = 10
	// 6 = 11
	// 7 = 12
	// 8 = 2=
	// 9 = 2-
	// 10 = 20
	// 11 = 21
	// 12 = 22
	// 13 = 1==
	// 14 = 1=-
	// 15 = 1=0
	// 16 = 1=1
	// 17 = 1=2
	// 18 = 1-=
	// 19 = 1--
	// 20 = 110
	return snafu
}

func part1(inputFile string) string {
	data, _ := util.ReadFileToStringSlice(inputFile)
	sum := 0
	for _, line := range data {
		sum += snafuToDec(line)
	}

	return decToSnafu(sum)
}

func part2(inputFile string) string {
	// load(inputFile)

	return ""
}
