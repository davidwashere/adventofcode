package day16

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

type Valve struct {
	Index    int
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
			Index:    i,
			ID:       vals[0],
			FlowRate: rate,
			Tunnels:  vals[2:],
		}

		if i == 0 {
			first = v.ID
		}

		valves[v.ID] = &v
	}

	return first, valves
}

func part1(inputfile string) int {
	start, valves := loadFile(inputfile)

	// valvesSlice := make([]bool, len(valves))

	minsMax := 30
	openValves := map[string]bool{}
	// meno := map[string]int{}

	var recur func(curID string, mins int, open map[string]bool) int
	recur = func(curID string, mins int, open map[string]bool) int {
		if mins <= 0 {
			return 0
		}

		// key := fmt.Sprintf("%v-%v", mins, open)
		// if v, ok := meno[key]; ok {
		// 	return v
		// }

		cur := valves[curID]
		var max int
		// if this valve isn't open, and has flow rate > 0, calc what it would be like if it was open
		if !open[curID] && cur.FlowRate > 0 {
			openValvesClone := cloneMap(open)
			openValvesClone[curID] = true
			max = recur(curID, mins-1, openValvesClone) + ((mins - 1) * cur.FlowRate)
		}

		for _, tun := range cur.Tunnels {
			t := recur(tun, mins-1, open)
			max = util.Max(t, max)
		}

		// meno[key] = max
		// fmt.Printf("min %v max %v\n", (minsMax - (minsMax - mins)), max)
		return max
	}

	result := recur(start, minsMax, openValves)

	return result
}

func cloneMap(in map[string]bool) map[string]bool {
	m := make(map[string]bool)
	for k, v := range in {
		m[k] = v
	}

	return m
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
