package day06

import (
	"io/ioutil"
	"strings"
)

func add(m map[byte]int, b byte) {
	m[b]++
}

func sub(m map[byte]int, b byte) {
	m[b]--
	if m[b] == 0 {
		delete(m, b)
	}
}

func findStart(inputfile string, seqLen int) int {
	dataB, _ := ioutil.ReadFile(inputfile)
	data := strings.TrimSpace(string(dataB))

	m := map[byte]int{}
	for i := 0; i < seqLen; i++ {
		b := data[i]
		add(m, b)
	}

	if len(m) == seqLen {
		return seqLen
	}

	for i := seqLen; i < len(data); i++ {
		li := i - seqLen
		sub(m, data[li])
		add(m, data[i])

		if len(m) == seqLen {
			return i + 1
		}
	}

	return -1
}

func part1(inputfile string) int {
	return findStart(inputfile, 4)
}

func part2(inputfile string) int {
	return findStart(inputfile, 14)
}
