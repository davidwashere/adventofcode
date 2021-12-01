package day15

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseFile(inputfile string) []int {
	dataB, _ := ioutil.ReadFile(inputfile)

	raw := string(dataB)
	dataS := strings.Split(raw, ",")

	data := []int{}
	for _, item := range dataS {
		i, _ := strconv.Atoi(item)
		data = append(data, i)
	}

	return data
}

func part1(inputfile string) int {
	return doit(inputfile, 2020)
}

func part2(inputfile string) int {
	return doit(inputfile, 30000000)
}

func doit(inputfile string, turns int) int {
	initialNums := parseFile(inputfile)

	mem := map[int]int{}
	prev := 0
	turn := 0
	for _, num := range initialNums {
		turn++
		prev = num
		mem[num] = turn
	}

	for turn = turn + 1; turn <= turns; turn++ {
		lastTurn, ok := mem[prev]
		mem[prev] = turn - 1
		if ok {
			prev = turn - 1 - lastTurn // cur turn, -1 to get prev turn, - lastTurn # seen
		} else {
			prev = 0
		}

		if turns >= 10000000 {
			if turn%10000000 == 0 {
				fmt.Printf("[%v] %v\n", turn, prev)
			}
		}

	}

	return prev
}
