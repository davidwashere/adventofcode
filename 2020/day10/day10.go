package day10

import (
	"aoc2020/util"
	"sort"
)

func parseJoltagesAndSort(inputfile string) []int {
	data, err := util.ReadFileToIntSlice(inputfile)
	util.Check(err)

	sort.Ints(data)

	// Add 0 to start
	data = append([]int{0}, data...)

	// Add last 3 jolt hop
	last := data[len(data)-1]
	data = append(data, last+3)

	return data
}

func part1(inputfile string) int {
	data := parseJoltagesAndSort(inputfile)

	hops := map[int]int{}
	for i := 1; i < len(data); i++ {
		cur := data[i]
		prev := data[i-1]
		diff := cur - prev
		hops[diff] = hops[diff] + 1
	}

	return hops[1] * hops[3]
}

func part2(inputfile string) int {
	data := parseJoltagesAndSort(inputfile)

	t := util.NewTree()

	lastIndex := len(data) - 1
	for i := 0; i <= lastIndex; i++ {
		cur := data[i]

		for j := 1; j <= util.Min(lastIndex-i, 3); j++ {
			next := data[i+j]
			if next-cur <= 3 {
				t.AddChild(cur, next)
			}
		}
	}

	count := t.CountPaths(0, data[lastIndex])
	return count
}
