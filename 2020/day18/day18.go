package day18

import (
	"aoc2020/util"
	"regexp"
	"strconv"
	"strings"
)

func part1(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	sum := 0
	for _, line := range data {

		result := calcLine(line)

		sum += result
	}

	return sum
}

func calcLine(line string) int {
	line = strings.ReplaceAll(line, " ", "")
	intRe := regexp.MustCompile("[0-9]+")
	parenRe := regexp.MustCompile("[()]+")
	opRe := regexp.MustCompile("[*+]+")

	numStack := util.NewIntStack()
	opStack := util.NewStringStack()

	for _, cbyte := range line {
		char := string(cbyte)
		if intRe.MatchString(char) {
			// is num
			num, _ := strconv.Atoi(char)
			if numStack.IsEmpty() {
				numStack.Push(num)
			} else {
				if opStack.Peek() == "(" {
					numStack.Push(num)
				} else {
					op := opStack.Pop()
					lastNum := numStack.Pop()
					if op == "*" {
						numStack.Push(num * lastNum)
					} else if op == "+" {
						numStack.Push(num + lastNum)
					}
				}
			}
		} else if parenRe.MatchString(char) {
			// is paran
			if char == "(" {
				opStack.Push("(")
			} else if char == ")" {
				op := opStack.Pop()
				if !opStack.IsEmpty() && (opStack.Peek() == "+" || opStack.Peek() == "*") {
					op = opStack.Pop()
					last := numStack.Pop()
					lastLast := numStack.Pop()
					if op == "*" {
						numStack.Push(last * lastLast)
					} else if op == "+" {
						numStack.Push(last + lastLast)
					}
				}
			}
			// paranStack

		} else if opRe.MatchString(char) {
			opStack.Push(char)
			// is + or -
		}
	}

	if len(numStack) > 1 {
		panic("wtf")
	}

	val := numStack.Pop()

	return val
}

func part2(inputfile string) int {
	data, _ := util.ReadFileToStringSlice(inputfile)

	sum := 0
	for _, line := range data {
		sum += calcLineP2(line)
	}

	return sum
}

func calcLineP2(line string) int {
	line = strings.ReplaceAll(line, " ", "")
	intRe := regexp.MustCompile("[0-9]+")
	parenRe := regexp.MustCompile("[()]+")
	opRe := regexp.MustCompile("[*+]+")

	opStack := util.NewStringStack()
	outStack := util.NewStringStack()

	for _, cbyte := range line {
		char := string(cbyte)
		if intRe.MatchString(char) {
			// is num
			outStack.Push(char)

		} else if parenRe.MatchString(char) {
			// is paran
			if char == "(" {
				opStack.Push(char)
			} else if char == ")" {
				for opStack.Peek() != "(" {
					outStack.Push(opStack.Pop())
				}
				opStack.Pop()
			}

		} else if opRe.MatchString(char) {
			// is + or -
			if opStack.IsEmpty() {
				opStack.Push(char)
				continue
			}
			// + higher then *
			for !opStack.IsEmpty() && (char == "*" && opStack.Peek() == "+") {
				outStack.Push(opStack.Pop())
			}

			opStack.Push(char)
		}
	}

	for !opStack.IsEmpty() {
		outStack.Push(opStack.Pop())
	}

	// fmt.Println(outStack)

	numStack := util.NewIntStack()
	for i := 0; i < len(outStack); i++ {
		char := string(outStack[i])
		if intRe.MatchString(char) {
			num, _ := strconv.Atoi(char)
			numStack.Push(num)
			continue
		}

		if opRe.MatchString(char) {
			right := numStack.Pop()
			left := numStack.Pop()

			if char == "*" {
				numStack.Push(left * right)
			} else if char == "+" {
				numStack.Push(left + right)
			}
		}
	}

	return numStack.Pop()
}
