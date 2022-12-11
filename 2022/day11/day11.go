package day11

import (
	"aoc/util"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items          util.BigIntQueue
	test           *big.Int
	onTrue         int
	onFalse        int
	inspectedCount int
	reducer        *big.Int
	opMultiply     bool
	lcm            *big.Int

	// true if use old value for left in op
	opLeftOld bool
	// if opLeftOld = false, this is the value to use for op
	opLeft *big.Int

	// true if use old value for right in op
	opRightOld bool

	// if opRightOld = false, this is the value to use for op
	opRight *big.Int
}

func NewMonkey(reducer *big.Int) *Monkey {
	m := new(Monkey)
	m.reducer = reducer

	return m
}

func (m *Monkey) SetOp(op []string) {
	if op[0] == "old" {
		m.opLeftOld = true
	} else {
		m.opLeftOld = false
		t, _ := strconv.Atoi(op[0])
		m.opLeft = big.NewInt(int64(t))
	}

	if op[2] == "old" {
		m.opRightOld = true
	} else {
		m.opRightOld = false
		t, _ := strconv.Atoi(op[2])
		m.opRight = big.NewInt(int64(t))
	}

	if op[1] == "*" {
		m.opMultiply = true
	} else if op[1] == "+" {
		m.opMultiply = false
	}
}

func (m *Monkey) Do(monkeys []*Monkey) {
	for !m.items.IsEmpty() {
		item := m.items.Dequeue()

		var left *big.Int
		var right *big.Int

		if m.opLeftOld {
			left = new(big.Int).Set(item)
		} else {
			left = m.opLeft
		}

		if m.opRightOld {
			right = new(big.Int).Set(item)
		} else {
			right = m.opRight
		}

		if m.opMultiply {
			item.Mul(left, right)
		} else {
			item.Add(left, right)
		}

		// bored
		item.Div(item, m.reducer)

		if m.lcm != nil {
			item.Mod(item, m.lcm)
		}

		// test
		mod := big.NewInt(0)
		if mod.Mod(item, m.test).Cmp(big.NewInt(0)) == 0 {
			monkeys[m.onTrue].items.Enqueue(item)
		} else {
			monkeys[m.onFalse].items.Enqueue(item)
		}

		m.inspectedCount++
	}
}

func (m Monkey) String() string {
	return fmt.Sprintf("Items %v Test %v True %v False %v", m.items, m.test, m.onTrue, m.onFalse)
}

func gimmieMonkeys(inputfile string, reducer int) []*Monkey {
	data, _ := util.ReadFileToStringSlice(inputfile)

	r := big.NewInt(int64(reducer))

	cur := NewMonkey(r)

	monkeys := []*Monkey{}
	monkeys = append(monkeys, cur)
	for _, line := range data {
		if len(line) == 0 {
			cur = NewMonkey(r)
			monkeys = append(monkeys, cur)
			continue
		}

		tokens := util.ParseTokens(line)

		if tokens.Words[0] == "Monkey" {
			continue
		}

		if tokens.Words[0] == "Starting" {
			items := []*big.Int{}
			for _, t := range tokens.Ints {
				items = append(items, big.NewInt(int64(t)))
			}
			cur.items = append(cur.items, items...)
		} else if tokens.Words[0] == "Operation" {
			equation := strings.Split(line, " = ")[1]
			op := strings.Split(equation, " ")
			cur.SetOp(op)
		} else if tokens.Words[0] == "Test" {
			cur.test = big.NewInt(int64(tokens.Ints[0]))
		} else if tokens.Words[1] == "true" {
			cur.onTrue = tokens.Ints[0]
		} else if tokens.Words[1] == "false" {
			cur.onFalse = tokens.Ints[0]
		}
	}

	// for i, monkey := range monkeys {
	// 	fmt.Printf("[%v] %v\n", i, monkey)
	// }

	return monkeys

}

func part1(inputfile string) int {
	monkeys := gimmieMonkeys(inputfile, 3)

	rounds := 20
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.Do(monkeys)
		}
	}

	inspectedCounts := []int{}
	for _, monkey := range monkeys {
		inspectedCounts = append(inspectedCounts, monkey.inspectedCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectedCounts)))

	return inspectedCounts[0] * inspectedCounts[1]
}

func part2(inputfile string, rounds int) int {
	monkeys := gimmieMonkeys(inputfile, 1)

	// part2 hack
	// "find another way to keep your worry levels manageable"
	// find a divisor to reduce the worry level by that will not throw the math off but
	// keep the numbers from getting too high
	lcm := big.NewInt(1)
	for _, monkey := range monkeys {
		lcm.Mul(lcm, monkey.test)
	}

	for _, monkey := range monkeys {
		monkey.lcm = lcm
	}

	fmt.Println(lcm)

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			monkey.Do(monkeys)
		}
	}

	fmt.Println("After round", rounds)
	for j, monkey := range monkeys {
		fmt.Printf("  [%v] %v\n", j+1, monkey.inspectedCount)
	}

	inspectedCounts := []int{}
	for _, monkey := range monkeys {
		inspectedCounts = append(inspectedCounts, monkey.inspectedCount)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspectedCounts)))

	return inspectedCounts[0] * inspectedCounts[1]
}
