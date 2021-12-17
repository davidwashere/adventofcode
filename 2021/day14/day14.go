package day14

import (
	"aoc/util"
)

type rule struct {
	vals   []string
	insert string
}

type chg struct {
	r   rule
	num int
}

func loadFile(inputfile string) (string, []rule) {
	data, _ := util.ReadFileToStringSlice(inputfile)

	template := data[0]

	rules := []rule{}
	for i := 2; i < len(data); i++ {
		line := data[i]
		strs := util.ParseStrs(line)
		vals := []string{string(strs[0][0]), string(strs[0][1])}
		insert := strs[1]
		r := rule{vals: vals, insert: insert}
		rules = append(rules, r)
	}

	return template, rules
}

func peek(m map[string]map[string]int, k1, k2 string) int {
	if _, ok := m[k1]; !ok {
		return 0
	} else if _, ok := m[k1][k2]; !ok {
		return 0
	}

	return m[k1][k2]
}

func pop(m map[string]map[string]int, k1, k2 string) int {
	val := peek(m, k1, k2)

	delete(m[k1], k2)

	return val
}

func ins(m map[string]map[string]int, k1, k2 string, val int) {
	if _, ok := m[k1]; !ok {
		m[k1] = map[string]int{}
	}

	m[k1][k2] += val
}

func sub(m map[string]map[string]int, k1, k2 string, val int) {
	// assume keys exist
	m[k1][k2] -= val
}

func run(inputfile string, steps int) int {
	template, rules := loadFile(inputfile)
	// fmt.Println(template, rules)

	// N -> N -> 3
	// map keys represent a 'pair' of elements
	// int is the number of those 'pairs' in the polymer
	polymer := map[string]map[string]int{}
	counts := map[string]int{}

	for i := 1; i < len(template); i++ {
		k1 := string(template[i-1])
		k2 := string(template[i])

		ins(polymer, k1, k2, 1)
		counts[k2]++
	}

	counts[string(template[0])]++

	for steps > 0 {
		steps--
		chgs := []chg{}

		for _, rule := range rules {
			k1 := rule.vals[0]
			k2 := rule.vals[1]

			old := peek(polymer, k1, k2)

			if old > 0 {
				chgs = append(chgs, chg{r: rule, num: old})
			}
		}

		for _, chg := range chgs {
			k1 := chg.r.vals[0]
			k2 := chg.r.vals[1]
			kNew := chg.r.insert
			num := chg.num

			counts[kNew] += num

			if peek(polymer, k1, k2) == num {
				pop(polymer, k1, k2)
			} else {
				sub(polymer, k1, k2, num)
			}

			ins(polymer, k1, kNew, num)
			ins(polymer, kNew, k2, num)
		}
	}

	min := util.MaxInt
	max := util.MinInt
	for _, v := range counts {
		min = util.Min(min, v)
		max = util.Max(max, v)
	}

	return max - min
}

func part1(inputfile string) int {
	return run(inputfile, 10)
}

func part2(inputfile string) int {
	return run(inputfile, 40)
}
