package day15

import (
	"aoc/util"
	"fmt"
	"sort"
)

type Range struct {
	low int
	hi  int
}

type Sensor struct {
	pos           util.Point
	closestBeacon util.Point
	dist          int
}

func (s *Sensor) PrepDist() {
	s.dist = s.pos.DistOrtho(s.closestBeacon)
}

func (s Sensor) String() string {
	return fmt.Sprintf("%v-%v(%v)", s.pos, s.closestBeacon, s.dist)
}

func loadData(inputfile string) ([]*Sensor, int, int) {
	data, _ := util.ReadFileToStringSlice(inputfile)

	minX := util.MaxInt
	maxX := util.MinInt

	sensors := []*Sensor{}
	for _, line := range data {
		ints := util.ParseIntsNeg(line)
		sensor := Sensor{
			pos:           util.NewPoint(ints[0], ints[1]),
			closestBeacon: util.NewPoint(ints[2], ints[3]),
		}
		sensor.PrepDist()

		minX = util.Min(minX, sensor.pos.X-sensor.dist)
		maxX = util.Max(maxX, sensor.pos.X+sensor.dist)

		sensors = append(sensors, &sensor)
	}

	return sensors, minX, maxX
}

func part1(inputfile string, y int) int {
	// A grid shouldn't be necessary
	// just need to know each point and its 'strength' (distance to beacon)
	sensors, minX, maxX := loadData(inputfile)

	ranges := []Range{}
	for _, sensor := range sensors {
		sY := sensor.pos.Y
		distToY := util.Abs(sY - y)
		rem := sensor.dist - distToY
		if rem <= 0 {
			continue
		}

		ranges = append(ranges, Range{low: sensor.pos.X - rem, hi: sensor.pos.X + rem})

	}

	ranges = consolidateRanges(ranges)

	fmt.Println(minX, maxX)

	sum := 0
	for _, r := range ranges {
		sum += r.hi - r.low
	}

	return sum
}

// Return a new set of ranges, none overlapping
func consolidateRanges(in []Range) []Range {
	ranges := make([]Range, len(in))
	copy(ranges, in)

	// sort by lowest low

	i := 0
	for {
		if i+1 > len(ranges)-1 {
			break
		}

		sort.Slice(ranges, func(i, j int) bool {
			return ranges[i].low < ranges[j].low
		})

		cur := ranges[i]
		next := ranges[i+1]

		if cur.hi >= next.low {
			r := Range{cur.low, util.Max(cur.hi, next.hi)}
			ranges = append(ranges[0:i+1], ranges[i+2:]...)
			ranges[i] = r
			continue
		}
		i++
	}

	return ranges

}

func part2(inputfile string, max int) int {
	sensors, _, _ := loadData(inputfile)

	doit := func(x, y int) int {
		return (x * 4000000) + y
	}

	for y := 0; y <= max; y++ {
		ranges := []Range{}
		for _, sensor := range sensors {
			sY := sensor.pos.Y
			distToY := util.Abs(sY - y)
			rem := sensor.dist - distToY
			if rem < 0 {
				continue
			}

			ranges = append(ranges, Range{low: sensor.pos.X - rem, hi: sensor.pos.X + rem})

		}
		ranges = consolidateRanges(ranges)

		// beacon on left end of this row
		if len(ranges) > 1 {
			if ranges[1].low-ranges[0].hi > 1 {
				x := ranges[0].hi + 1
				return doit(x, y)
			}
		} else {
			// only one range
			low := ranges[0].low
			hi := ranges[0].hi
			if low > 0 {
				return doit(low-1, y)
			} else if hi < max {
				return doit(hi+1, y)
			}
		}

	}

	return -1
}
