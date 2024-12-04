package day03

import (
	"aoc/util"
	"fmt"
	"regexp"
	"strings"
)

var (
	raw    string
	tokens = []string{
		"mul",
		"don't()",
		"do()",
	}
)

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)
	raw = strings.Join(data, "")
	// data, _ := util.ReadFileToIntSlice(inputFile)
	// data, _ := util.ReadFileToStringSliceWithDelim(inputFile, "\n\n")
	// grid := util.NewInfinityGridFromFile(inputfile, ".")

	// for _, line := range data {
	// 	tokens := util.ParseTokens(line)
	// 	// ints := util.ParseInts(line)
	// 	// strs := util.ParseStrs(line)
	// 	// words := util.ParseWords(line)

	// 	fmt.Println(tokens)
	// }
}

func reset(poses []int) {
	for i := 0; i < len(poses); i++ {
		poses[i] = 0
	}
}

func part1(inputFile string) int {
	load(inputFile)

	re := regexp.MustCompile(`^\([0-9]{1,3},[0-9]{1,3}\)`)

	tokenPoses := make([]int, len(tokens))

	sum := 0
	for i := 0; i < len(raw); i++ {
		c := raw[i]
		for j := 0; j < len(tokens); j++ {
			if tokens[j][tokenPoses[j]] == c {
				tokenPoses[j]++
				if tokenPoses[j] == len(tokens[j]) {
					// token match found, preprocess next stuff
					t := raw[i+1:]
					str := re.FindString(t)
					if str != "" {
						tokenPoses[j] = 0
						fmt.Printf("Yippie: %s\n", str)
						tokens := util.ParseInts(str)
						sum += tokens[0] * tokens[1]
					}
					reset(tokenPoses)
					i += len(str)
					break
				}
			}
		}
	}

	return sum
}

func part2(inputFile string) int {
	load(inputFile)

	// re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	re := regexp.MustCompile(`^\([0-9]{1,3},[0-9]{1,3}\)`)

	tokenPoses := make([]int, len(tokens))

	enabled := true
	sum := 0
	for i := 0; i < len(raw); i++ {
		c := raw[i]
		for j := 0; j < len(tokens); j++ {
			if tokens[j][tokenPoses[j]] == c {
				tokenPoses[j]++
				if tokenPoses[j] == len(tokens[j]) {
					if j == 0 && enabled {
						// token match found, preprocess next stuff
						t := raw[i+1:]
						str := re.FindString(t)
						if str != "" {
							tokenPoses[j] = 0
							fmt.Printf("Yippie: %s\n", str)
							tokens := util.ParseInts(str)
							sum += tokens[0] * tokens[1]
						}
						i += len(str)
					}

					if j == 1 {
						enabled = false
					}

					if j == 2 {
						enabled = true
					}

					reset(tokenPoses)
					break
				}
			}
		}
	}

	return sum
}
