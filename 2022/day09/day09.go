package day09

import (
	"aoc/util"
	"fmt"
	"strconv"
)

var (
	dirToVec = map[string]util.Vector{
		"R": util.VectorEast,
		"L": util.VectorWest,
		"U": util.VectorNorth,
		"D": util.VectorSouth,
	}
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	head := util.NewPoint(0, 0)
	tail := util.NewPoint(0, 0)

	grid := util.NewInfGrid[bool]()
	for _, line := range data {
		strs := util.ParseStrs(line)

		dir := dirToVec[strs[0]]
		numSteps, _ := strconv.Atoi(strs[1])

		fmt.Println(strs)
		for i := 0; i < numSteps; i++ {
			newHead := head.Apply(dir)
			if newHead.DistOrthoDiag(tail) > 1 {
				tail.X = head.X
				tail.Y = head.Y
			}
			head = newHead

			// Record tail visit
			grid.Set(true, tail.X, tail.Y)
			fmt.Printf("H: %v T: %v\n", head, tail)
		}
	}

	return grid.Len()
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	rope := make([]util.Point, 10)
	last := len(rope) - 1

	grid := util.NewInfGrid[string]()
	grid.WithDefaultValue(".")
	for _, line := range data {
		strs := util.ParseStrs(line)

		dir := dirToVec[strs[0]]
		numSteps, _ := strconv.Atoi(strs[1])

		for i := 0; i < numSteps; i++ {
			newHead := rope[0].Apply(dir)
			rope[0].X = newHead.X
			rope[0].Y = newHead.Y

			for j := 1; j < len(rope); j++ {
				prev := rope[j-1]
				cur := rope[j]

				if cur.DistOrthoDiag(prev) <= 1 {
					break
				}

				thisDir := cur.TowardVector(prev)
				newKnot := cur.Apply(thisDir)

				if newKnot == cur {
					break
				}

				rope[j].X = newKnot.X
				rope[j].Y = newKnot.Y
			}

			grid.Set("X", rope[last].X, rope[last].Y)
		}
	}

	return grid.Len()
}

// func ropeToString(rope []util.Point) string {
// 	r := ""
// 	for i, p := range rope {
// 		if i > 0 {
// 			r += "-"
// 		}
// 		if i == 0 {
// 			r += "H"
// 		} else {
// 			r += fmt.Sprintf("%v", i)
// 		}
// 		r += p.String()
// 	}

// 	return r
// }
