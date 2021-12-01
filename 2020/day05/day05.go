package day05

import (
	"aoc2020/util"
	"fmt"
	"sort"
)

// F means "front", B means "back", L means "left", and R means "right".
// 128 rows on the plane, 0 - 127
// each letter tells you which half of a region the given seat is in

func half(low, high int) int {
	return (high - low + 1) / 2
}

func part1(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	maxSeatID := -1
	for _, line := range data {
		low := 0
		high := 127

		for _, c := range line[:7] {
			char := string(c)
			if char == "F" {
				high = high - half(low, high)
			} else if char == "B" {
				low = low + half(low, high)
			}
		}
		row := low
		fmt.Println("Final Row: ", row)

		low = 0
		high = 7
		for _, c := range line[7:] {
			char := string(c)
			if char == "L" {
				high = high - half(low, high)
			} else if char == "R" {
				low = low + half(low, high)
			}
		}

		col := low
		fmt.Println("Final Col: ", col)

		seatID := (row * 8) + col

		if seatID > maxSeatID {
			maxSeatID = seatID
		}

		// fmt.Printf("Row %v, Col %v, SeatID %v\n", row, col, seatID)
	}

	return maxSeatID
}

func part2(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	allSeats := []int{}
	for _, line := range data {
		low := 0
		high := 127

		for _, c := range line[:7] {
			char := string(c)
			half := half(low, high)
			if char == "F" {
				high = high - half
			} else if char == "B" {
				low = low + half
			}
		}
		row := low
		// fmt.Println("Final Row: ", row)

		low = 0
		high = 7
		for _, c := range line[7:] {
			char := string(c)
			half := half(low, high)
			if char == "L" {
				high = high - half
			} else if char == "R" {
				low = low + half
			}
		}

		col := low
		seatID := (row * 8) + col

		allSeats = append(allSeats, seatID)
	}

	sort.Ints(allSeats)

	prev := allSeats[0]
	for _, seat := range allSeats[1:] {
		if seat-prev > 1 {
			fmt.Println("Diff: ", prev, seat)
		}
		prev = seat
	}

	return 0
}
