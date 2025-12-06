package day05

import (
	"aoc/util"
	"fmt"
	"sort"
)

type dayItem struct {
	start int
	end   int
}

func (d dayItem) Includes(num int) bool {
	return num >= d.start && num <= d.end
}

func AnyIncludes(items []dayItem, num int) bool {
	for _, item := range items {
		if item.Includes(num) {
			return true
		}
	}

	return false
}

var (
	items = []dayItem{}
	fresh = []int{}
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	// data, _ := util.ReadFileToIntSlice(inputFile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputFile, "\n\n")
	// grid := util.NewInfinityGridFromFile(inputfile, ".")

	for _, line := range data {
		// tokens := util.ParseTokens(line)
		ints := util.ParseInts(line)
		// strs := util.ParseStrs(line)
		// words := util.ParseWords(line)

		if len(ints) == 2 {
			items = append(items, dayItem{
				start: ints[0],
				end:   ints[1],
			})
			continue
		}

		if len(ints) == 1 {
			fresh = append(fresh, ints[0])
		}

	}
	// fmt.Println(items)
	// fmt.Println(fresh)
}

func part1(inputFile string) int {
	load(inputFile)

	result := 0
	for _, f := range fresh {
		if AnyIncludes(items, f) {
			result++
		}
	}

	return result
}

func part2(inputFile string) int {
	load(inputFile)

	// reduce the ranges into a canocial range

	sort.Slice(items, func(i, j int) bool {
		l := items[i]
		r := items[j]

		return l.start < r.start
	})

	for {
		changed := false
	outer:
		for i := 0; i < len(items)-1; i++ {
			for k := i + 1; k < len(items); k++ {
				left := items[i]
				right := items[k]

				newItem := dayItem{}
				if right.start <= left.end+1 {

					newItem.start = min(left.start, right.start)
					newItem.end = max(right.end, left.end)
					items[i] = newItem
					items = append(items[:k], items[k+1:]...)
					changed = true
					break outer
				}
			}
		}

		if !changed {
			break
		}
	}

	result := 0
	for _, item := range items {
		adding := (item.end - item.start + 1)
		fmt.Printf("[%d-%d] (+%d)\n", item.start, item.end, adding)
		result += adding
	}

	return result
}
