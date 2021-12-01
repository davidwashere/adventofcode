package day19

import (
	"aoc2020/util"
	"fmt"
	"regexp"
	"strings"
)

func parse(inputfile string) (map[string]string, []string) {
	data, _ := util.ReadFileToStringSlice(inputfile)
	rules := map[string]string{}
	messages := []string{}

	for _, line := range data {
		if strings.TrimSpace(line) == "" {
			// skip blank lines
			continue
		}

		if strings.Index(line, ":") >= 0 {
			// rule
			ruleSp := strings.Split(line, ":")
			rules[strings.TrimSpace(ruleSp[0])] = strings.TrimSpace(ruleSp[1])
		} else {
			// message
			messages = append(messages, strings.TrimSpace(line))
		}
	}

	return rules, messages
}

func match(re *regexp.Regexp, message string) bool {
	return re.MatchString(message)
}

func buildReString(ruleID string, rules map[string]string) string {
	rule := rules[ruleID]
	if strings.Index(rule, "\"") >= 0 {
		ruleS := strings.Split(rule, "\"")
		return ruleS[1]
	}

	ruleS := strings.Split(rule, "|")

	sections := []string{}
	for _, part := range ruleS {
		strs := util.ParseStrs(part)

		reString := ""
		for _, str := range strs {
			reString += buildReString(str, rules)
		}
		sections = append(sections, reString)
	}

	reString := strings.Join(sections, "|")

	return "(" + reString + ")"

}

func part1(inputfile string) int {
	rules, messages := parse(inputfile)

	reString := "^" + buildReString("0", rules) + "$"
	fmt.Println(reString)
	re := regexp.MustCompile(reString)
	result := 0
	for _, message := range messages {
		if match(re, message) {
			fmt.Printf("MATCH: %v\n", message)
			result++
		} else {
			fmt.Printf("  NOT: %v\n", message)
		}
	}

	return result
}

// var LoopCount = 0

func buildReStringInf(maxMessageLen int, ruleID string, rules map[string]string) string {
	rule := rules[ruleID]
	if strings.Index(rule, "\"") >= 0 {
		ruleS := strings.Split(rule, "\"")
		return ruleS[1]
	}

	// if util.IsStringIn(util.ParseStrs(rule), ruleID) {
	// 	LoopCount++
	// }
	ruleSp := strings.Split(rule, "|")

	//  8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	recursiveRule := util.IsStringIn(util.ParseStrs(rule), ruleID)

	if recursiveRule {
		fmt.Println("RRRRUULLLLEE: ", ruleID)
	}
	recursivePartIndex := 0
	recursiveStrIsAtEnd := false
	recursiveStrIsInMiddle := false
	if recursiveRule {
		for i, part := range ruleSp {
			if util.IsStringIn(util.ParseStrs(part), ruleID) {
				recursivePartIndex = i
				strs := util.ParseStrs(part)
				for j, str := range strs {
					if str == ruleID && len(strs)-1 == j {
						recursiveStrIsAtEnd = true
					} else if len(strs) == 3 && j == 1 && str == ruleID {
						recursiveStrIsInMiddle = true
					}
				}
			}
		}
	}

	sections := []string{}
	for i, part := range ruleSp {
		strs := util.ParseStrs(part)

		if recursiveRule && i == recursivePartIndex {
			if recursiveStrIsInMiddle {
				left := buildReStringInf(maxMessageLen, strs[0], rules)
				right := buildReStringInf(maxMessageLen, strs[2], rules)

				repeats := []string{}
				for j := 2; j < maxMessageLen; j++ {
					leftRepeat := strings.Repeat(left, j)
					rightRepeat := strings.Repeat(right, j)
					// repeats = append(repeats, "("+leftRepeat+rightRepeat+")")
					repeats = append(repeats, leftRepeat+rightRepeat)
				}

				sections = append(sections, strings.Join(repeats, "|"))
			}

			// as a result only one section should be preseent
			continue
		}

		reString := ""
		for _, str := range strs {
			tmp := buildReStringInf(maxMessageLen, str, rules)
			reString += tmp
		}
		sections = append(sections, reString)
	}

	if recursiveRule {
		if recursiveStrIsAtEnd {
			// ?: required or get OOM errors
			sections[0] = fmt.Sprintf("(?:%s{1,%d})", sections[0], maxMessageLen)
		}

		// for _, sec := range sections {
		// 	fmt.Println("Recursive Rule: ", sec)
		// }
	}

	reString := strings.Join(sections, "|")

	// ?: required or get OOM errors
	return "(?:" + reString + ")"

}

func part2(inputfile string) int {
	rules, messages := parse(inputfile)

	maxMessageLen := 0
	for _, message := range messages {
		maxMessageLen = util.Max(len(message), maxMessageLen)
	}

	reString := "^" + buildReStringInf(maxMessageLen, "0", rules) + "$"

	// fmt.Println("FINAL REGEX:", reString)
	re := regexp.MustCompile(reString)
	result := 0
	for i, message := range messages {
		fmt.Printf("[%v]: %v  ...  ", i, message)
		if match(re, message) {
			fmt.Printf("MATCH\n")
			result++
		} else {
			fmt.Printf("NOPE\n")
		}
	}

	return result
}
