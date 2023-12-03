package day03

import (
	"aoc/util"
	"regexp"
	"strconv"
	"strings"
)

var (
	grid  *util.InfGrid[string]
	numRe = regexp.MustCompile("[0-9]+")
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	y := 0
	grid = util.NewInfGrid[string]()
	grid.WithDefaultValue(".")
	for _, line := range data {
		for x := 0; x < len(line); x++ {
			c := string(line[x])
			if c == "." {
				continue
			}
			grid.Set(c, x, y)
		}
		y++
	}
}

func part1(inputFile string) int {
	load(inputFile)

	numOrDotRe := regexp.MustCompile("[0-9.]+")

	parts := []int{}
	for y := 0; y <= grid.GetMaxY(); y++ {
		line := strings.Join(grid.GetRow(y), "")
		matches := numRe.FindAllStringIndex(line, -1)

		for _, match := range matches {
			num, _ := strconv.Atoi(line[match[0]:match[1]])

			loX := match[0] - 1
			hiX := match[1]

		Outer:
			for tx := loX; tx <= hiX; tx++ {
				for ty := y - 1; ty <= y+1; ty++ {
					if !numOrDotRe.MatchString(grid.Get(tx, ty)) {
						parts = append(parts, num)
						break Outer
					}
				}
			}
		}
	}

	sum := 0
	for _, part := range parts {
		sum += part
	}

	return sum
}

func part2(inputFile string) int {
	load(inputFile)

	sum := 0
	grid.VisitAll2D(func(val string, x, y int) {
		if val != "*" {
			return
		}

		nums := []int{}
		for ty := y - 1; ty <= y+1; ty++ {
			line := strings.Join(grid.GetRow(ty), "")

			matches := numRe.FindAllStringIndex(line, -1)
			for _, match := range matches {
				// determine if any number borders the '*'
				num, _ := strconv.Atoi(line[match[0]:match[1]])
				for tx := x - 1; tx <= x+1; tx++ {
					if tx >= match[0] && tx < match[1] {
						nums = append(nums, num)
						break
					}
				}
			}
		}

		if len(nums) == 2 {
			sum += (nums[0] * nums[1])
		}
	})

	return sum
}
