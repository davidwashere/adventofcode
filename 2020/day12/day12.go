package day12

import (
	"aoc2020/util"
	"strconv"
)

var vec = util.NewNormalizedVector

type instr struct {
	op  string
	mag int
}

func gimmieInstructions(inputfile string) []instr {
	data, _ := util.ReadFileToStringSlice(inputfile)

	instructions := []instr{}

	for _, line := range data {
		ins := instr{}

		ins.op = line[:1]
		ins.mag, _ = strconv.Atoi(line[1:])

		instructions = append(instructions, ins)
	}

	return instructions
}

func part1(inputfile string) int {
	instructions := gimmieInstructions(inputfile)

	curX := 0
	curY := 0
	dir := vec(1, 0)

	for _, ins := range instructions {
		switch ins.op {
		case "N":
			curY += ins.mag
		case "S":
			curY -= ins.mag
		case "E":
			curX += ins.mag
		case "W":
			curX -= ins.mag
		case "L":
			dir.RotateInt(ins.mag)
		case "R":
			dir.RotateInt(-ins.mag)
		case "F":
			dir.M = ins.mag
			curX, curY = dir.Apply(curX, curY)
		}
	}

	return util.Abs(curX) + util.Abs(curY)
}

func part2(inputfile string) int {
	instructions := gimmieInstructions(inputfile)

	curX := 0
	curY := 0
	dir := vec(10, 1)

	for _, ins := range instructions {
		switch ins.op {
		case "N":
			dir.Y += ins.mag
		case "S":
			dir.Y -= ins.mag
		case "E":
			dir.X += ins.mag
		case "W":
			dir.X -= ins.mag
		case "L":
			dir.RotateInt(ins.mag)
		case "R":
			dir.RotateInt(-ins.mag)
		case "F":
			dir.M = ins.mag
			curX, curY = dir.Apply(curX, curY)
		}
	}

	return util.Abs(curX) + util.Abs(curY)
}
