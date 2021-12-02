package day02

import (
	"aoc/util"
	"strconv"
	"strings"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	x := 0
	y := 0

	for _, line := range data {
		lineS := strings.Split(line, " ")

		dir := lineS[0]
		mag, _ := strconv.Atoi(lineS[1])

		switch dir {
		case "forward":
			x += mag
		case "up":
			y -= mag
		case "down":
			y += mag
		}
	}

	return x * y
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	x := 0
	y := 0
	aim := 0

	for _, line := range data {
		lineS := strings.Split(line, " ")

		dir := lineS[0]
		mag, _ := strconv.Atoi(lineS[1])

		switch dir {
		case "forward":
			x += mag
			y += (aim * mag)
		case "up":
			aim -= mag
		case "down":
			aim += mag
		}
	}

	return x * y
}
