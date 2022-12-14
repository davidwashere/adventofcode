package day13

import (
	"aoc/util"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type PacketPair struct {
	LeftRaw  string
	RightRaw string
	Left     []string
	Right    []string
}

func (p PacketPair) String() string {
	return fmt.Sprintf("L %v - R %v", p.Left, p.Right)
}

func parseLine(input string) []string {
	re := regexp.MustCompile(`[\]\[]{1,1}|[0-9]+`) // find '[', ']', or number
	res := []string{}
	for _, val := range re.FindAllString(input, -1) {
		val = strings.TrimSpace(val)
		if val != "" {
			res = append(res, val)
		}
	}

	// fmt.Printf("%v == ", input)
	// for _, tok := range res {
	// 	fmt.Printf("-%v- ", tok)
	// }
	// fmt.Println()
	return res
}

func loadData(inputfile string) []*PacketPair {
	pairs := []*PacketPair{}

	data, _ := util.ReadFileToStringSlice(inputfile)

	for i := 0; i < len(data); i = i + 3 {
		pair := new(PacketPair)

		pair.LeftRaw = data[i]
		pair.RightRaw = data[i+1]

		pair.Left = parseLine(pair.LeftRaw)
		pair.Right = parseLine(pair.RightRaw)

		pairs = append(pairs, pair)
	}

	return pairs
}

func part1(inputfile string) int {
	pairs := loadData(inputfile)

	// true == correct order, false = wrong order
	pairResults := make([]bool, len(pairs))
	for i, pair := range pairs {
		correct := correctOrder(pair.Left, pair.Right)
		pairResults[i] = correct
	}

	result := 0
	for i, correct := range pairResults {
		if correct {
			result += (i + 1)
		}
	}

	return result
}

func correctOrder(leftO, rightO []string) bool {
	left := make([]string, len(leftO))
	copy(left, leftO)

	right := make([]string, len(rightO))
	copy(right, rightO)

	li := 0
	ri := 0

	for li < len(left) && ri < len(right) {
		l := left[li]
		r := right[ri]

		if l == "]" || r == "]" {
			if l == "]" && r != "]" {
				// left runs out first
				return true
			} else if l != "]" && r == "]" {
				// right runs out first
				return false
			}

			li++
			ri++
		}

		if l == "[" || r == "[" {
			if l == "[" {
				li++
			}
			if r == "[" {
				ri++
			}

			// handle when one is a list and the other is an int
			if l == "[" && util.IsInt(r) {
				// right is int, convert it to fake list
				insertAt := ri + 1
				right = append(right[:insertAt+1], right[insertAt:]...)
				right[insertAt] = "]"

			} else if r == "[" && util.IsInt(l) {
				insertAt := li + 1
				left = append(left[:insertAt+1], left[insertAt:]...)
				left[insertAt] = "]"
			}
			continue
		}

		if util.IsInt(l) && util.IsInt(r) {
			if l == r {
				li++
				ri++
				continue
			}

			lNum, _ := strconv.Atoi(l)
			rNum, _ := strconv.Atoi(r)
			if lNum < rNum {
				return true
			} else if lNum > rNum {
				return false
			}
		}
	}

	return false
}

func part2(inputfile string) int {
	pairs := loadData(inputfile)

	// hack, put all pairs into a big list
	all := [][]string{}
	div1 := []string{"[", "[", "2", "]", "]"}
	div2 := []string{"[", "[", "6", "]", "]"}
	all = append(all, div1)
	all = append(all, div2)
	for _, pair := range pairs {
		all = append(all, pair.Left)
		all = append(all, pair.Right)
	}

	// bubble sort
	for i := 0; i < len(all); i++ {
		for j := 0; j < len(all)-i-1; j++ {
			l := all[j]
			r := all[j+1]
			if correctOrder(l, r) {
				continue
			}

			all[j], all[j+1] = all[j+1], all[j]
		}
	}

	result := 1
	for i, t := range all {
		if reflect.DeepEqual(t, div1) || reflect.DeepEqual(t, div2) {
			fmt.Printf("divider at: %v\n", i+1)
			result *= i + 1
		}
	}

	return result
}
