package day16

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

type Valve struct {
	ID       string
	FlowRate int
	Tunnels  []string
}

func (v Valve) String() string {
	return fmt.Sprintf("%v %v %v", v.ID, v.FlowRate, v.Tunnels)
}

func loadFile(inputfile string) (string, map[string]*Valve) {
	data, _ := util.ReadFileToStringSlice(inputfile)
	valves := map[string]*Valve{}
	var first string

	re := regexp.MustCompile(`[A-Z]{2,2}|[0-9]+`) // find '[', ']', or number
	for i, line := range data {
		vals := re.FindAllString(line, -1)
		rate, _ := strconv.Atoi(vals[1])
		v := Valve{
			ID:       vals[0],
			FlowRate: rate,
			Tunnels:  vals[2:],
		}

		if i == 0 {
			first = v.ID
		}

		valves[v.ID] = &v
		fmt.Printf("%+v\n", v)
	}

	return first, valves
}

func part1(inputfile string) int {
	start, valves := loadFile(inputfile)

	minsMax := 2
	// openValves := map[string]bool{}

	// sum := 0
	// cur := start
	// for mins := minsMax; mins > 0; mins-- {
	// 	sum += sumOpen(valves, openValves)

	// 	v := valves[cur]
	// 	if v.FlowRate != 0 && !openValves[cur] {
	// 		openValves[cur] = true
	// 		continue
	// 	}
	// }

	//  1  2  3
	// AA DD CC
	// AA DD AA
	// AA DD EE
	// AA DD CC 20
	// AA DD AA 20
	// AA DD EE 20
	// AA II AA
	// AA II JJ
	// AA II AA
	// AA II JJ

	open := map[string]bool{}

	path := []string{}

	var recur func(cur string, mins int) int
	recur = func(cur string, mins int) int {
		sum := sumOpen(valves, open)
		if mins <= 0 {
			return sum
		}

		v := valves[cur]
		max := 0
		for _, next := range v.Tunnels {
			max = util.Max(max, recur(next, mins-1))
		}

		if !open[cur] {
			open[cur] = true
			for _, next := range v.Tunnels {
				max = util.Max(max, recur(next, mins-2))
			}
		}

		return sum + max
	}

	path = append(path, start)
	result := recur(start, minsMax)
	return result
}

func sumOpen(valves map[string]*Valve, openValves map[string]bool) int {
	sum := 0
	for _, v := range valves {
		if openValves[v.ID] {
			sum += v.FlowRate
		}
	}

	return sum
}

func part2(inputfile string) int {
	return 0
}
