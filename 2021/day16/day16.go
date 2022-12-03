package day16

import (
	"aoc/util"
	"fmt"
	"strings"
)

var key = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	input := data[0]

	bin := ""
	for _, cB := range input {
		c := string(cB)
		b := key[c]
		bin = fmt.Sprintf("%v%v", bin, b)
		// fmt.Printf("%v", b)
	}

	bin = strings.TrimRight(bin, "0")

	v := bin[:3]
	t := bin[3:6]

	fmt.Println(v, t)

	result := 0

	return result
}

func part2(inputfile string) int {
	return 0
}
