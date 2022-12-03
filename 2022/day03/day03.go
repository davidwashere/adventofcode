package day03

import (
	"aoc/util"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	result := 0
	for _, line := range data {
		itemsFound := map[byte]bool{}

		// fill first sack
		half := len(line) / 2
		for i := 0; i < half; i++ {
			itemsFound[line[i]] = true
		}

		var dupe byte
		for i := half; i < len(line); i++ {
			item := line[i]

			if _, ok := itemsFound[item]; ok {
				dupe = item
				break
			}
		}

		result += priority(dupe)
	}

	return result
}

func priority(item byte) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - 'a' + 1
	}

	return int(item) - 'A' + 27
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	result := 0
	for i := 0; i < len(data); i = i + 3 {
		sack1 := sackToMap(data[i])
		sack2 := sackToMap(data[i+1])

		// loop through third sack in group and find the item that is common in other two sacks
		sack3line := data[i+2]
		for j := 0; j < len(sack3line); j++ {
			item := sack3line[j]

			if sack1[item] && sack2[item] {
				// this is the badge
				result += priority(item)
				break
			}
		}

	}

	return result
}

func sackToMap(sack string) map[byte]bool {
	r := map[byte]bool{}
	for i := 0; i < len(sack); i++ {
		item := sack[i]
		r[item] = true
	}

	return r
}
