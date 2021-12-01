package day13

import (
	"aoc2020/util"
	"strconv"
	"strings"
)

type bus struct {
	busID  int
	offset int
}

func parseFile(inputfile string) (int, []bus) {
	data, _ := util.ReadFileToStringSlice(inputfile)

	earlistTs, _ := strconv.Atoi(data[0])

	buses := []bus{}

	split := strings.Split(data[1], ",")
	for i, it := range split {
		busID, err := strconv.Atoi(it)
		if err != nil {
			continue
		}

		buses = append(buses, bus{busID, i})
	}

	// During initial part2 attempts, optimized by putting the 'biggest' busID first
	// to increment faster
	// sort.SliceStable(buses, func(i, j int) bool {
	// 	return buses[i].busID < buses[j].busID
	// })

	return earlistTs, buses
}

func part1(inputfile string) int {
	earliestTs, buses := parseFile(inputfile)

	cur := earliestTs
	for {
		for _, bus := range buses {
			mod := cur % bus.busID

			if mod == 0 { // Bus is departing
				minsWait := cur - earliestTs
				return bus.busID * minsWait
			}
		}

		cur++
	}
}

func part2(inputfile string) int {
	earliest, buses := parseFile(inputfile)

	lastBusIndex := len(buses) - 1
	start := buses[0].busID - buses[0].offset // The current 'start' timestamp
	inc := buses[0].busID                     // The current 'increment' to increase 'start' by when buses aren't aligned
	locked := 0                               // The largest bus index that is in 'sync'
	for {
		for i, bus := range buses {
			mod := (start + bus.offset) % bus.busID

			if mod != 0 {
				break
			}

			if i > locked {
				// At this point i and i-1 are in 'sync', so increase
				// increment so they remain in sync
				locked++
				inc = inc * bus.busID
			}

			if i == lastBusIndex && start >= earliest {
				return start
			}
		}

		start += inc
	}
}
