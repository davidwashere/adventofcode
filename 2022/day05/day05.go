package day05

import (
	"aoc/util"
	"fmt"
)

type stack struct {
	first *dll
	last  *dll
}

func (s stack) String() string {
	r := ""
	cur := s.first.next
	for len(cur.val) > 0 {
		r += cur.val
		cur = cur.next
	}

	return r
}

type dll struct {
	val  string
	next *dll
	prev *dll
}

func createStacks(num int) []*stack {
	stacks := make([]*stack, num)
	for i, s := range stacks {
		s = new(stack)
		s.first = new(dll)
		s.last = new(dll)
		s.first.next = s.last
		s.last.prev = s.first
		stacks[i] = s
	}

	return stacks
}

func part1(inputfile string) string {
	data, _ := util.ReadFileToStringSlice(inputfile)
	numStacks := len(util.ParseStrs(data[0]))
	stacks := createStacks(numStacks)

	stacksDone := false
	for i := 0; i < len(data); i++ {
		line := data[i]
		if len(line) == 0 {
			stacksDone = true
			fmt.Println(stacks)
			continue
		}

		// process instructions
		if stacksDone {
			ints := util.ParseInts(line)

			qty := ints[0]
			from := ints[1] - 1
			to := ints[2] - 1

			for j := 0; j < qty; j++ {
				// grab and remove top of from
				item := stacks[from].last.prev
				item.prev.next = stacks[from].last
				stacks[from].last.prev = item.prev

				// add to to
				stacks[to].last.prev.next = item
				item.prev = stacks[to].last.prev
				item.next = stacks[to].last
				stacks[to].last.prev = item
			}
			continue
		}

		// Do setup
		strs := util.ParseStrs(line)
		for i, stackItem := range strs {
			if stackItem[0] >= '0' && stackItem[0] <= '9' {
				// is number, ignore it
				continue
			}

			stack := stacks[i]
			newDLL := new(dll)
			newDLL.val = stackItem

			newDLL.prev = stack.first
			newDLL.next = stack.first.next

			newDLL.next.prev = newDLL
			newDLL.prev.next = newDLL
		}

	}

	result := ""
	for _, s := range stacks {
		v := s.last.prev.val
		result += v
	}

	return result
}

func part2(inputfile string) string {
	data, _ := util.ReadFileToStringSlice(inputfile)
	numStacks := len(util.ParseStrs(data[0]))
	stacks := createStacks(numStacks)

	stacksDone := false
	for i := 0; i < len(data); i++ {
		line := data[i]
		if len(line) == 0 {
			stacksDone = true
			fmt.Println(stacks)
			continue
		}

		// process instructions
		if stacksDone {
			ints := util.ParseInts(line)

			qty := ints[0]
			from := ints[1] - 1
			to := ints[2] - 1

			// the last item in the qty to move
			l := stacks[from].last.prev
			f := l
			for j := 1; j < qty; j++ {
				f = f.prev
			}

			// remove from from
			f.prev.next = l.next
			l.next.prev = f.prev

			// Add to to
			f.prev = stacks[to].last.prev
			f.prev.next = f

			l.next = stacks[to].last
			l.next.prev = l

			continue
		}

		// Do setup
		strs := util.ParseStrs(line)
		for i, stackItem := range strs {
			if stackItem[0] >= '0' && stackItem[0] <= '9' {
				// is number, ignore it
				continue
			}

			stack := stacks[i]
			newDLL := new(dll)
			newDLL.val = stackItem

			newDLL.prev = stack.first
			newDLL.next = stack.first.next

			newDLL.next.prev = newDLL
			newDLL.prev.next = newDLL
		}

	}

	result := ""
	for _, s := range stacks {
		v := s.last.prev.val
		result += v
	}

	return result
}
