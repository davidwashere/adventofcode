package day02

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type productRange struct {
	start int
	end   int
}

var (
	ranges = []productRange{}
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	// data, _ := util.ReadFileToIntSlice(inputFile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputFile, "\n\n")
	// grid := util.NewInfinityGridFromFile(inputfile, ".")

	for _, line := range data {
		rawRanges := strings.Split(line, ",")

		for _, r := range rawRanges {
			tokens := util.ParseTokens(r)
			ranges = append(ranges, productRange{
				tokens.Ints[0],
				tokens.Ints[1],
			})

		}
	}
	fmt.Println(ranges)
}

func part1(inputFile string) int {
	load(inputFile)

	result := 0
	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			str := strconv.Itoa(i)
			if len(str)%2 != 0 {
				continue // skip ranges that cannot be split exactly
			}

			mid := len(str) / 2
			l := str[0:mid]
			r := str[mid:]
			if l == r {
				result += i
			}
		}
	}

	return result
}

func part2(inputFile string) int {
	load(inputFile)

	already := map[int]bool{}
	result := 0
	for _, r := range ranges {
		// fmt.Printf("\n== RANGE %d -> %d\n", r.start, r.end)
		for i := r.start; i <= r.end; i++ {
			str := strconv.Itoa(i)
			if len(str) == 1 {
				continue
			}
			factors := Factors(len(str))

			for _, f := range factors {
				// grab f characters
				base := str[0:f]
				good := true
				for li := f; li <= len(str)-f; li += f {
					ri := li + f
					if base != str[li:ri] {
						good = false
						break
					}
				}
				if good {
					if _, ok := already[i]; ok {
						panic(fmt.Sprintf("NOOO %d", i))
					}
					result += i
					already[i] = true
					fmt.Printf("num: %d, base: %s, sum: %d\n", i, base, result)
					break
				}
			}
		}
	}

	return result
}

func Factors(num int) []int {
	if num == 0 {
		return []int{} // Factors are typically not defined for 0 in this context.
	}

	var factors []int
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
			if i*i != num { // Avoid adding the same factor twice for perfect squares
				f := num / i
				if f != num {
					factors = append(factors, num/i)
				}
			}

		}
	}

	sort.Slice(factors, func(i, j int) bool {
		return factors[i] > factors[j] // Descending order
	})

	return factors
}
