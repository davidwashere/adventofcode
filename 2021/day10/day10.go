package day10

import (
	"aoc/util"
	"sort"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	op := map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
	push := map[string]bool{"(": true, "[": true, "{": true, "<": true}

	multi := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}

	counts := map[string]int{")": 0, "]": 0, "}": 0, ">": 0}
	for _, line := range data {
		s := util.NewStack[string]()

		for _, cr := range line {
			c := string(cr)

			if push[c] {
				s.Push(c)
				continue
			}

			top := s.Pop()

			if op[top] != c {
				// corrupted line
				counts[c]++
				break
			}
		}
	}

	result := 0

	for k, v := range counts {
		result += (multi[k] * v)
	}

	return result
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	op := map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
	push := map[string]bool{"(": true, "[": true, "{": true, "<": true}

	pts := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}
	scores := []int{}
	for _, line := range data {
		s := util.NewStack[string]()

		corrupt := false
		for _, cr := range line {
			c := string(cr)

			if push[c] {
				s.Push(c)
				continue
			}

			top := s.Pop()
			if op[top] != c {
				corrupt = true
				break
			}
		}

		if corrupt || len(s) == 0 {
			// line is corrupt or complete
			continue
		}

		// if get here line is incomplete
		score := 0
		for len(s) > 0 {
			c := s.Pop()
			score *= 5
			score += pts[c]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}
