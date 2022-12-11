package day10

import (
	"aoc/util"
	"fmt"
	"strconv"
	"strings"
)

var specialCycles = map[int]bool{
	20:  true,
	60:  true,
	100: true,
	140: true,
	180: true,
	220: true,
}

type Op interface {
	Process(cycle int, args ...string)
}

// ---
type noopOp struct {
}

func NewNoopOp() *noopOp {
	return &noopOp{}
}

func (n *noopOp) Process(cycle int, args ...string) {}

// ---
type addxOp struct {
	mods map[int]int
}

func NewAddxOp(mods map[int]int) *addxOp {
	return &addxOp{
		mods: mods,
	}
}

func (a *addxOp) Process(cycle int, args ...string) {
	inc, _ := strconv.Atoi(args[0])
	a.mods[cycle+1] = inc
}

func part1(inputfile string) int {
	lines, _ := util.ReadFileToStringSlice(inputfile)

	// mods holds a cycle # to a change to X
	mods := map[int]int{}

	opMap := map[string]Op{
		"noop": NewNoopOp(),
		"addx": NewAddxOp(mods),
	}

	result := 0
	x := 1
	cycle := 0
	lineIndex := -1
	for {
		cycle += 1

		if specialCycles[cycle] {
			strength := cycle * x
			result += strength
			fmt.Printf("[%v] x=%v strength=%v\n", cycle, x, strength)
		}

		fmt.Printf("[%v] x=%v -> ", cycle, x)

		if inc, ok := mods[cycle]; ok {
			x += inc
			fmt.Printf("%v\n", x)
			continue
		}

		lineIndex++
		if lineIndex == len(lines) {
			break
		}
		line := lines[lineIndex]

		parts := strings.Split(line, " ")
		op := opMap[parts[0]]
		op.Process(cycle, parts[1:]...)

		fmt.Printf("%v    %v\n", x, line)

	}

	return result
}

func part2(inputfile string) int {
	lines, _ := util.ReadFileToStringSlice(inputfile)

	// mods holds a cycle # to a change to X
	mods := map[int]int{}

	opMap := map[string]Op{
		"noop": NewNoopOp(),
		"addx": NewAddxOp(mods),
	}

	result := 0
	x := 1
	cycle := 0
	lineIndex := -1
	for {
		cycle += 1

		// Print Start
		pos := (cycle - 1) % 40
		if pos == 0 {
			fmt.Println()
		}

		c := " "
		if pos >= x-1 && pos <= x+1 {
			c = "#"
		}

		fmt.Print(c)
		// Print Stop

		if inc, ok := mods[cycle]; ok {
			x += inc
			continue
		}

		lineIndex++
		if lineIndex == len(lines) {
			break
		}
		line := lines[lineIndex]

		parts := strings.Split(line, " ")
		op := opMap[parts[0]]
		op.Process(cycle, parts[1:]...)
	}

	return result
}
