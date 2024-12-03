package day02

import (
	"aoc/util"
)

var (
	reports [][]int
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	// data, _ := util.ReadFileToIntSlice(inputFile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputFile, "\n\n")
	// grid := util.NewInfinityGridFromFile(inputfile, ".")

	for _, line := range data {
		tokens := util.ParseTokens(line)
		// ints := util.ParseInts(line)
		// strs := util.ParseStrs(line)
		// words := util.ParseWords(line)
		reports = append(reports, tokens.Ints)
	}
	// fmt.Println(reports)
}

func part1(inputFile string) int {
	load(inputFile)

	safe := 0
	for _, levels := range reports {

		increasing := false
		for i := 0; i < len(levels)-1; i++ {
			cur := levels[i]
			nxt := levels[i+1]

			if i == 0 {
				if cur-nxt > 0 {
					increasing = true
				} else if cur-nxt < 0 {
					increasing = false
				} else {
					// not safe ?
					break
				}
			}

			if increasing && cur-nxt <= 0 {
				// invalid level
				break
			} else if !increasing && cur-nxt >= 0 {
				// invalid level
				break
			}

			hi := max(cur, nxt)
			lo := min(cur, nxt)
			diff := hi - lo

			if diff < 1 || diff > 3 {
				// not safe
				break
			}

			if i == len(levels)-2 {
				safe++
			}
		}
	}

	return safe
}

func part2(inputFile string) int {
	load(inputFile)

	skip := -1
	safeCount := 0
	for li := 0; li < len(reports); li++ {
		levels := reports[li]

		if skip >= 0 {
			levels = append(levels[:skip], levels[skip+1:]...)
		}
		if skip >= len(reports) {
			panic("skip too hi")
			continue
		}

		safe := false
		increasing := false
		init := false
		for i := 0; i < len(levels)-1; i++ {
			if i == skip {
				continue
			}
			cur := levels[i]
			nxt := levels[i+1]

			if !init {
				if cur-nxt > 0 {
					increasing = true
				} else if cur-nxt < 0 {
					increasing = false
				} else {
					// not safe ?
					break
				}
				init = true
			}

			if increasing && cur-nxt <= 0 {
				// invalid level
				break
			} else if !increasing && cur-nxt >= 0 {
				// invalid level
				break
			}

			hi := max(cur, nxt)
			lo := min(cur, nxt)
			diff := hi - lo

			if diff < 1 || diff > 3 {
				// not safe
				break
			}

			if i == len(levels)-2 {
				safe = true
			}
		}

		if safe {
			safeCount++
		} else {
			skip++
			if skip >= len(levels)-1 {
				skip = -1
				continue
			}
			li--
		}

	}

	return safeCount
}
