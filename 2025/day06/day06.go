package day06

import (
	"aoc/util"
	"strconv"
	"strings"
)

type dayItem struct {
	nums      []int
	numstrs   []string
	maxstrlen int
	op        string
}

func (d dayItem) calc() int {
	result := 0
	for i := 0; i < len(d.nums); i++ {
		num := d.nums[i]
		if i == 0 {
			result = num
			continue
		}

		switch d.op {
		case "+":
			result += num
		case "-":
			result -= num
		case "/":
			result /= num
		case "*":
			result *= num
		}
	}

	return result
}
func (d dayItem) calc2() int {
	result := 0
	actual := []int{}
	for i := d.maxstrlen - 1; i >= 0; i-- {
		sb := strings.Builder{}
		for _, numstr := range d.numstrs {
			c := string(numstr[i])
			if c != " " {
				sb.WriteString(c)
			}
		}
		num, _ := strconv.Atoi(sb.String())
		actual = append(actual, num)
	}

	for i := 0; i < len(actual); i++ {
		num := actual[i]
		if i == 0 {
			result = num
			continue
		}

		switch d.op {
		case "+":
			result += num
		case "-":
			result -= num
		case "/":
			result /= num
		case "*":
			result *= num
		}
	}

	return result
}

var (
	items = []dayItem{}
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	for _, line := range data {
		tokens := util.ParseTokens(line)
		if len(tokens.Ints) > 0 {
			if len(tokens.Ints) > len(items) {
				items = make([]dayItem, len(tokens.Ints))
			}

			for i, num := range tokens.Ints {
				items[i].nums = append(items[i].nums, num)
			}
		} else {
			strs := util.ParseStringsRegex("[+\\-*/]", line)
			for i, op := range strs {
				items[i].op = op
			}
		}
	}

	// cols that have spaces
	cols := map[int]bool{}
	spaceindexes := map[int]int{}
	for _, line := range data {
		for i, c := range line {
			if c == ' ' {
				spaceindexes[i]++
				if spaceindexes[i] == len(data) {
					cols[i] = true
				}
			}
		}
	}

	for i, line := range data {
		if i == len(data)-1 {
			continue
		}
		col := 0
		sb := strings.Builder{}
		for i, c := range line {
			if cols[i] {
				str := sb.String()
				items[col].numstrs = append(items[col].numstrs, str)
				items[col].maxstrlen = max(items[col].maxstrlen, len(str))
				sb.Reset()
				col++
				continue
			}
			sb.WriteRune(c)
		}
		str := sb.String()
		items[col].numstrs = append(items[col].numstrs, str)
		items[col].maxstrlen = max(items[col].maxstrlen, len(str))
	}
}

func part1(inputFile string) int {
	load(inputFile)

	result := 0
	for _, item := range items {
		result += item.calc()
	}

	return result
}

func part2(inputFile string) int {
	load(inputFile)

	result := 0
	for _, item := range items {
		result += item.calc2()
	}
	return result
}
