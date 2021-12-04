package day04

import (
	"aoc/util"
	"strconv"
	"strings"
)

type board struct {
	// row, col
	grid [][]int
}

func NewBoard() *board {
	return &board{
		grid: [][]int{},
	}
}

func parseNumsToPull(line string) []int {
	numsToPull := []int{}
	nums := strings.Split(line, ",")

	for _, numStr := range nums {
		num, err := strconv.Atoi(numStr)
		util.Check(err)

		numsToPull = append(numsToPull, num)
	}

	return numsToPull
}

func parseBoards(data []string) []*board {
	boards := []*board{}

	b := NewBoard()
	boards = append(boards, b)
	row := 0
	for _, line := range data {
		if len(line) == 0 {
			b = NewBoard()
			boards = append(boards, b)
			row = 0
			continue
		}

		lineS := strings.Split(line, " ")
		for _, numS := range lineS {
			if len(numS) == 0 {
				continue
			}
			num, err := strconv.Atoi(numS)
			util.Check(err)

			if len(b.grid) < row+1 {
				b.grid = append(b.grid, []int{})
			}

			b.grid[row] = append(b.grid[row], num)
		}
		row++
	}

	return boards
}

func isWinnerByRows(grid [][]int, pulled map[int]bool) bool {
	numRows := len(grid)
	numCols := len(grid[0])
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			cur := grid[row][col]
			if !pulled[cur] {
				break
			}

			if col == numCols-1 {
				return true
			}
		}
	}

	return false
}

func isWinnerByCols(grid [][]int, pulled map[int]bool) bool {
	numRows := len(grid)
	numCols := len(grid[0])
	for col := 0; col < numCols; col++ {
		for row := 0; row < numRows; row++ {
			cur := grid[row][col]
			if !pulled[cur] {
				break
			}

			if row == numRows-1 {
				return true
			}
		}
	}

	return false
}

func sumRemaining(grid [][]int, pulled map[int]bool) int {
	numRows := len(grid)
	numCols := len(grid[0])

	sum := 0
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			cur := grid[row][col]
			if !pulled[cur] {
				sum += cur
			}
		}
	}

	return sum
}

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	numsToPull := parseNumsToPull(data[0])
	boards := parseBoards(data[2:])

	winningBoard := -1

	pulled := map[int]bool{}
	pulled[numsToPull[0]] = true
	pulled[numsToPull[1]] = true
	pulled[numsToPull[2]] = true
	pulled[numsToPull[3]] = true
	pullPos := 3
	for {
		if winningBoard >= 0 || pullPos >= len(numsToPull) {
			break
		}

		pullPos++
		pulled[numsToPull[pullPos]] = true

		for bi, b := range boards {
			if isWinnerByRows(b.grid, pulled) || isWinnerByCols(b.grid, pulled) {
				winningBoard = bi
			}

		}
	}

	sum := sumRemaining(boards[winningBoard].grid, pulled)

	return sum * numsToPull[pullPos]
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	numsToPull := parseNumsToPull(data[0])
	boards := parseBoards(data[2:])

	winningBoards := map[int]bool{}
	winningBoard := -1

	pulled := map[int]bool{}
	pulled[numsToPull[0]] = true
	pulled[numsToPull[1]] = true
	pulled[numsToPull[2]] = true
	pulled[numsToPull[3]] = true
	pullPos := 3
	for {
		if winningBoard >= 0 || pullPos >= len(numsToPull) {
			break
		}

		pullPos++
		pulled[numsToPull[pullPos]] = true

		for bi, b := range boards {
			if winningBoards[bi] {
				continue
			}

			if isWinnerByRows(b.grid, pulled) || isWinnerByCols(b.grid, pulled) {
				if len(winningBoards) == len(boards)-1 {
					winningBoard = bi
				} else {
					winningBoards[bi] = true
				}
			}
		}
	}

	sum := sumRemaining(boards[winningBoard].grid, pulled)

	return sum * numsToPull[pullPos]
}
