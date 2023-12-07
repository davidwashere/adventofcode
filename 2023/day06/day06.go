package day06

import (
	"aoc/util"
	"strconv"
	"strings"
)

var (
	races = []Race{}
)

type Race struct {
	timeCap  int
	bestDist int
}

func load(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	tokensTime := util.ParseInts(data[0])
	tokensBest := util.ParseInts(data[1])

	for i := range tokensTime {
		cap := tokensTime[i]
		best := tokensBest[i]

		races = append(races, Race{cap, best})
	}
}

func load2(inputFile string) {
	data, _ := util.ReadFileToStringSlice(inputFile)

	line := strings.Split(data[0], ": ")[1]
	tmp := strings.ReplaceAll(line, " ", "")
	cap, _ := strconv.Atoi(tmp)

	line = strings.Split(data[1], ": ")[1]
	tmp = strings.ReplaceAll(line, " ", "")
	best, _ := strconv.Atoi(tmp)

	races = append(races, Race{cap, best})
}

func part1(inputFile string) int {
	load(inputFile)

	raceWins := []int{}
	for _, race := range races {
		wins := 0
		for speed := 0; speed < race.timeCap; speed++ {
			timeRem := race.timeCap - speed
			dist := timeRem * speed
			if dist > race.bestDist {
				wins++
			}
		}

		raceWins = append(raceWins, wins)
	}

	r := 1
	for _, win := range raceWins {
		r *= win
	}

	return r
}

func part2(inputFile string) int {
	load2(inputFile)

	race := races[0]
	wins := 0
	for speed := 0; speed < race.timeCap; speed++ {
		timeRem := race.timeCap - speed
		dist := timeRem * speed
		if dist > race.bestDist {
			wins++
		}
	}

	return wins
}
