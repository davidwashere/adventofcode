package day23

import (
	"aoc2020/util"
	"fmt"
	"strconv"
)

type cupsT []int

func cupIndex(cups []int, cup int) int {
	for i := 0; i < len(cups); i++ {
		if cups[i] == cup {
			return i
		}
	}
	return -1
}

func nextDest(cups []int, picks []int, curCup int) int {
	next := curCup - 1
	if next < 1 {
		next = len(cups) + len(picks)
	}

	for util.IsIntIn(picks, next) {
		next = next - 1
		if next < 1 {
			next = len(cups) + len(picks)
		}

	}

	return next
}

func part1(input string) string {
	cups := cupsT{}
	for _, char := range input {
		num, _ := strconv.Atoi(string(char))
		cups = append(cups, num)
	}

	numCups := len(cups)
	curCup := cups[0]
	for i := 0; i < 100; i++ {
		curCupIndex := cupIndex(cups, curCup)
		// fmt.Printf("cups: ")
		// for i, cup := range cups {
		// 	if i == curCupIndex {
		// 		if i == 0 {
		// 			fmt.Printf("(%v) ", cup)
		// 		} else if i == numCups-1 {
		// 			fmt.Printf("(%v) ", cup)
		// 		} else {
		// 			fmt.Printf("(%v) ", cup)
		// 		}
		// 	} else {
		// 		fmt.Printf("%v  ", cup)
		// 	}
		// }
		// fmt.Println()

		cup1I := (curCupIndex + 1) % numCups
		cup2I := (cup1I + 1) % numCups
		cup3I := (cup2I + 1) % numCups

		picks := []int{}
		picks = append(picks, cups[cup1I])
		picks = append(picks, cups[cup2I])
		picks = append(picks, cups[cup3I])

		// fmt.Printf("pick up: ")
		// for i, pick := range picks {
		// 	fmt.Printf("%v", pick)
		// 	if i < len(picks)-1 {
		// 		fmt.Printf(", ")
		// 	}
		// }
		// fmt.Println()

		cups = util.RemoveIntFromSlice(cups, picks[0])
		cups = util.RemoveIntFromSlice(cups, picks[1])
		cups = util.RemoveIntFromSlice(cups, picks[2])

		dest := nextDest(cups, picks, curCup)
		// fmt.Printf("destination: %v\n", dest)
		// fmt.Println()
		i := cupIndex(cups, dest)
		i++

		picks = append(picks, cups[i:]...)
		cups = append(cups[:i], picks...)

		curCupIndex = cupIndex(cups, curCup)
		curCup = cups[(curCupIndex+1)%numCups]
	}

	result := ""
	ind := cupIndex(cups, 1)
	for i := 1; i < len(cups); i++ {
		result += fmt.Sprintf("%v", cups[(ind+i)%len(cups)])
	}

	// for _, line := range data {
	// 	tokens := util.ParseTokens(line)
	// 	// ints := util.ParseInts(line)
	// 	// strs := util.ParseStrs(line)
	// 	// words := util.ParseWords(line)

	// 	fmt.Println(tokens)
	// }

	return result
}

func part2(input string) int {
	cups := map[int]int{}

	firstNum, _ := strconv.Atoi(string(input[0]))
	prevNum := firstNum
	for i := 1; i < len(input); i++ {
		num, _ := strconv.Atoi(string(input[i]))
		cups[prevNum] = num
		prevNum = num
	}

	for i := len(input) + 1; i <= 1000000; i++ {
		cups[prevNum] = i
		prevNum = i
	}

	cups[prevNum] = firstNum

	curCup := firstNum

	for i := 0; i < 10000000; i++ {
		// for i := 0; i < 100; i++ {
		if i%1000000 == 0 {
			fmt.Printf("[%v]\n", i)
		}

		c1 := cups[curCup]
		c2 := cups[c1]
		c3 := cups[c2]

		cups[curCup] = cups[c3]

		// dest
		dest := curCup - 1
		for {
			if dest < 1 {
				dest = len(cups)
			}

			if dest == c1 || dest == c2 || dest == c3 {
				dest = dest - 1
				continue
			}

			break
		}

		last := cups[dest]
		cups[dest] = c1
		cups[c3] = last
		curCup = cups[curCup]
	}

	c1 := cups[1]
	c2 := cups[c1]
	fmt.Println(c1, c2)

	return c1 * c2
}
