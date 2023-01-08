package day16

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Valve struct {
	Index    int
	ID       string
	FlowRate int
	Tunnels  map[string]bool
}

func (v Valve) String() string {
	return fmt.Sprintf("%v %v %v", v.ID, v.FlowRate, v.Tunnels)
}

func load(inputFile string) map[string]*Valve {
	data, _ := util.ReadFileToStringSlice(inputFile)
	valves := map[string]*Valve{}

	re := regexp.MustCompile(`[A-Z]{2,2}|[0-9]+`) // find '[', ']', or number
	for i, line := range data {
		vals := re.FindAllString(line, -1)
		tuns := map[string]bool{}
		for _, tun := range vals[2:] {
			tuns[tun] = true
		}
		rate, _ := strconv.Atoi(vals[1])
		v := Valve{
			Index:    i,
			ID:       vals[0],
			FlowRate: rate,
			Tunnels:  tuns,
		}

		valves[v.ID] = &v
	}

	return valves
}

func copyMap(in map[string]bool) map[string]bool {
	newMap := map[string]bool{}

	for k, v := range in {
		newMap[k] = v
	}

	return newMap
}

func mapKey(vOrdered []string, open map[string]bool) string {
	key := ""

	for _, v := range vOrdered {
		if open[v] {
			key += "1,"
		} else {
			key += "0,"
		}
	}

	return key
}

func buildFlowingValves(in map[string]*Valve) map[string]int {
	r := map[string]int{}
	for k, v := range in {
		if v.FlowRate > 0 {
			r[k] = v.FlowRate
		}
	}

	return r
}

// buildOrderedFlowingValves result will be used to generate a consistent map key
func buildOrderedFlowingValves(in map[string]int) []string {
	r := []string{}
	for k := range in {
		r = append(r, k)
	}

	return r
}

func buildDistMap(in map[string]*Valve) map[string]map[string]int {
	r := map[string]map[string]int{}

	for x := range in {
		for y := range in {

			if _, ok := r[x]; !ok {
				r[x] = map[string]int{}
			}

			// test if y is in the tunnels for x
			if _, ok := in[x].Tunnels[y]; ok {
				r[x][y] = 1
			} else {
				r[x][y] = util.MaxInt
			}
		}
	}

	for k := range r {
		for i := range r {
			for j := range r {
				if r[i][k] == util.MaxInt || r[k][j] == util.MaxInt {
					continue
				}

				r[i][j] = util.Min(r[i][j], r[i][k]+r[k][j])
			}
		}
	}

	return r
}

func part1(inputFile string) int {
	valves := load(inputFile)
	minsMax := 30

	flowingValves := buildFlowingValves(valves)
	distMap := buildDistMap(valves)

	maxes := map[string]int{}
	_ = maxes

	max := 0
	var visit func(src string, mins int, open map[string]bool, flow int)
	visit = func(src string, mins int, open map[string]bool, flow int) {
		max = util.Max(max, flow)

		for dst, pressure := range flowingValves {
			updMins := mins - distMap[src][dst] - 1
			if open[dst] || updMins <= 0 {
				continue
			}

			updOpen := copyMap(open)
			updOpen[dst] = true
			updFlow := (updMins * pressure) + flow

			visit(dst, updMins, updOpen, updFlow)
		}

	}

	visit("AA", minsMax, map[string]bool{}, 0)
	return max
}

func part2(inputFile string) int {
	valves := load(inputFile)
	minsMax := 26

	flowingValves := buildFlowingValves(valves)
	flowingValvesOrdered := buildOrderedFlowingValves(flowingValves)
	distMap := buildDistMap(valves)

	maxes := map[string]int{}

	var visit func(src string, mins int, open map[string]bool, flow int)
	visit = func(src string, mins int, open map[string]bool, flow int) {
		openKey := mapKey(flowingValvesOrdered, open)
		maxes[openKey] = util.Max(maxes[openKey], flow)

		for dst, pressure := range flowingValves {
			updMins := mins - distMap[src][dst] - 1
			if open[dst] || updMins <= 0 {
				continue
			}

			updOpen := copyMap(open)
			updOpen[dst] = true
			updFlow := (updMins * pressure) + flow

			visit(dst, updMins, updOpen, updFlow)
		}

	}

	visit("AA", minsMax, map[string]bool{}, 0)

	max := 0
	for k, v := range maxes {
		for k2, v2 := range maxes {
			// only if didn't visit same spots
			if !isOverlap(k, k2) {
				max = util.Max(max, v+v2)

			}
		}
	}

	return max
}

func isOverlap(k1, k2 string) bool {
	k1s := strings.Split(k1, ",")
	k2s := strings.Split(k2, ",")

	for i := 0; i < len(k1s)-1; i++ {

		if k1s[i] == "1" && k2s[i] == "1" {
			return true
		}
	}

	return false
}
