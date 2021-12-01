package day07

import (
	"aoc2020/util"
	"strconv"
	"strings"
)

type rule struct {
	bagName string
	bags    []baggy
}

type baggy struct {
	name string
	num  int
}

func parseRule(line string) rule {
	lineSplit := strings.Split(line, "bags contain")
	bagName := strings.TrimSpace(lineSplit[0])
	contains := strings.TrimSpace(lineSplit[1])

	bags := []baggy{}

	if strings.HasPrefix(contains, "no") {
		return rule{bagName, bags}
	}

	canHoldS := strings.Split(contains, ",")
	for _, val := range canHoldS {
		val = val[:strings.Index(val, "bag")] // remove 'bag*' from end
		val = strings.TrimSpace(val)

		valS := strings.SplitN(val, " ", 2)

		num, err := strconv.Atoi(valS[0])
		util.Check(err)

		bag := baggy{valS[1], num}
		bags = append(bags, bag)
	}

	return rule{bagName, bags}
}

func part1(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	r := util.NewTree()

	for _, line := range data {
		if strings.HasPrefix(line, "shiny gold") {
			continue
		}

		rule := parseRule(line)
		for _, bag := range rule.bags {
			r.AddChild(rule.bagName, bag.name)
		}
	}

	result := len(r.GetAllUniqueParents("shiny gold"))
	return result
}

func part2(inputfile string) int {
	data, err := util.ReadFileToStringSlice(inputfile)
	util.Check(err)

	r := util.NewTree()
	for _, line := range data {
		rule := parseRule(line)

		for _, bag := range rule.bags {
			for i := 0; i < bag.num; i++ {
				r.AddChild(rule.bagName, bag.name)
			}
		}
	}

	result := len(r.GetAllChildren("shiny gold"))

	return result
}
