package day21

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strconv"
)

type MonkeyMath struct {
	l  string
	op string
	r  string
}

func load(inputFile string) (map[string]int, map[string]MonkeyMath) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	re := regexp.MustCompile(`[a-z]{4,4}|[+-/*]{1,1}|[0-9]+`) // find '[', ']', or number

	monkeysYell := map[string]int{}
	monkeysMath := map[string]MonkeyMath{}

	for _, line := range data {
		vals := re.FindAllString(line, -1)
		id := vals[0]
		vals = vals[1:]
		if len(vals) == 1 {
			v, _ := strconv.Atoi(vals[0])
			monkeysYell[id] = v
			continue
		}

		monkeysMath[id] = MonkeyMath{vals[0], vals[1], vals[2]}
	}

	return monkeysYell, monkeysMath
}

func part1(inputFile string) int {
	monkeysYell, monkeysMath := load(inputFile)

	var recur func(cur string) int
	recur = func(cur string) int {
		if v, ok := monkeysYell[cur]; ok {
			return v
		}

		m := monkeysMath[cur]

		l := recur(m.l)
		r := recur(m.r)

		var res int
		switch m.op {
		case "+":
			res = l + r
		case "-":
			res = l - r
		case "*":
			res = l * r
		case "/":
			res = l / r
		}

		return res
	}

	return recur("root")
}

// hasParent finds if child is a child of possibleParent
func hasParent(child, possibleParent string, reverseMonkeysMath map[string]string) bool {
	cur := child
	// fmt.Printf("hasParent [c:%v][p:%v] ", child, possibleParent)
	for {
		v, ok := reverseMonkeysMath[cur]
		if !ok {
			// fmt.Println("NA")
			return false
		}

		// fmt.Printf("%v > ", v)

		if v == possibleParent {
			// fmt.Println("FOUND")
			return true
		}

		cur = v
	}
}

func part2(inputFile string) int {
	monkeysYell, monkeysMath := load(inputFile)
	delete(monkeysYell, "humn")

	for {
		updated := false
		for k, v := range monkeysMath {
			if _, ok := monkeysYell[v.l]; ok {
				if _, ok := monkeysYell[v.r]; ok {
					l := monkeysYell[v.l]
					r := monkeysYell[v.r]
					res := doOp(v.op, l, r)
					monkeysYell[k] = res
					delete(monkeysMath, k)
					updated = true
				}
			}
		}
		if !updated {
			break
		}
	}

	// holds a l or r monkey to their parent
	reverseMonkeysMath := map[string]string{}
	for k, v := range monkeysMath {
		l := v.l
		r := v.r

		if _, ok := reverseMonkeysMath[l]; ok {
			panic("uh oh L")
		}

		if _, ok := reverseMonkeysMath[r]; ok {
			panic("uh oh R")
		}
		reverseMonkeysMath[l] = k
		reverseMonkeysMath[r] = k
	}

	cur := monkeysMath["root"]
	left, lok := monkeysYell[cur.l]
	right, rok := monkeysYell[cur.r]

	var val int
	if lok {
		val = left
		cur = monkeysMath[cur.r]
	} else if rok {
		val = right
		cur = monkeysMath[cur.l]
	} else {
		panic("both roots can't be yells")
	}

	fmt.Printf("need to yell = %v\n", val)

	for {
		left, lok := monkeysYell[cur.l]
		right, rok := monkeysYell[cur.r]

		if lok && rok {
			panic("nothing to solve")
		}

		if lok {
			val = doOppositeRight(cur.op, left, val)
			if cur.r == "humn" {
				return val

			}
			cur = monkeysMath[cur.r]
		} else {
			val = doOpposite(cur.op, right, val)
			if cur.l == "humn" {
				return val

			}
			cur = monkeysMath[cur.l]
		}
	}
}

func doOpposite(op string, num, total int) int {
	res := 0
	switch op {
	case "+":
		res = total - num
	case "-":
		res = total + num
	case "*":
		res = total / num
	case "/":
		res = total * num
	}

	return res
}

func doOppositeRight(op string, num, total int) int {
	res := 0
	switch op {
	case "+":
		res = total - num
	case "-":
		res = num - total
	case "*":
		res = total / num
	case "/":
		res = num / total
	}

	return res
}

func doOp(op string, l, r int) int {
	res := 0
	switch op {
	case "+":
		res = l + r
	case "-":
		res = l - r
	case "*":
		res = l * r
	case "/":
		res = l / r
	}

	return res
}
