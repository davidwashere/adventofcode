package day08

import (
	"aoc/util"
	"fmt"
	"strings"
)

var ()

type graph struct {
	nodes map[string]node
	insts string
}
type node struct {
	id string
	l  string
	r  string
}

func (n node) String() string {
	return fmt.Sprintf("%v (%v, %v)", n.id, n.l, n.r)
}

func load(inputFile string) *graph {
	data, _ := util.ReadFileToStringSlice(inputFile)

	g := &graph{
		nodes: make(map[string]node),
	}
	for i, line := range data {
		if i == 0 {
			g.insts = line
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		strs := util.ParseStrs(line)
		g.nodes[strs[0]] = node{id: strs[0], l: strs[1], r: strs[2]}

		// fmt.Println(strs)
	}

	return g
}

func part1(inputFile string) int {
	g := load(inputFile)

	steps := 0
	cur := "AAA"
	instPos := 0
	for {
		if cur == "ZZZ" {
			break
		}
		n := g.nodes[cur]
		dir := string(g.insts[instPos])

		if dir == "L" {
			cur = n.l
		} else {
			cur = n.r
		}

		instPos = (instPos + 1) % len(g.insts)
		steps++
	}

	return steps
}

func part2(inputFile string) int {
	g := load(inputFile)

	curs := []string{}
	for k, _ := range g.nodes {
		if strings.HasSuffix(k, "A") {
			curs = append(curs, k)
		}
	}

	lcms := []int{}
	for _, cur := range curs {

		steps := 0
		instPos := 0
		for {
			if strings.HasSuffix(cur, "Z") {
				lcms = append(lcms, steps)
				break
			}
			dir := string(g.insts[instPos])
			n := g.nodes[cur]

			if dir == "L" {
				cur = n.l
			} else {
				cur = n.r
			}

			instPos = (instPos + 1) % len(g.insts)
			steps++
		}
	}

	return util.LeastCommonMultiple(lcms[0], lcms[1], lcms[2:]...)
}
