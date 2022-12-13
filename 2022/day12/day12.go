package day12

import (
	"aoc/util"
)

type PT struct {
	X   int
	Y   int
	Val string
}

func load(inputfile string) (PT, PT, *util.InfGrid[string]) {
	lines, _ := util.ReadFileToStringSlice(inputfile)

	var start PT
	var end PT
	grid := util.NewInfGrid[string]()

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			c := string(line[x])
			if c == "S" {
				start = PT{x, y, c}
			} else if c == "E" {
				end = PT{x, y, c}
			}

			grid.Set(c, x, y)
		}
	}

	grid.LockBounds()

	return start, end, grid
}

func part1(inputfile string) int {
	start, end, grid := load(inputfile)

	return calcDist(start, end, grid)
}

func calcDist(start, end PT, grid *util.InfGrid[string]) int {
	visited := map[PT]bool{}

	queued := map[PT]bool{}
	queued[start] = true

	ptQueue := util.Queue[PT]{}
	ptQueue.Enqueue(start)

	distQueue := util.Queue[int]{}
	distQueue.Enqueue(0)

	result := 0
	// perhaps change to add to visited when added to the queue
	for !ptQueue.IsEmpty() {
		cur := ptQueue.Dequeue()
		curDist := distQueue.Dequeue()
		visited[cur] = true
		delete(queued, cur)
		// fmt.Println(cur, curDist, ptQueue)

		if cur.Val == end.Val {
			result = curDist
			break
		}

		next := nextLetter(cur.Val)
		grid.VisitOrtho(cur.X, cur.Y, func(val string, x, y int) {
			p := PT{x, y, val}
			if visited[p] || queued[p] {
				return
			}

			lower := val[0] > 'a' && val[0] < cur.Val[0]

			if val == cur.Val || val == next || lower {
				ptQueue.Enqueue(p)
				distQueue.Enqueue(curDist + 1)
				queued[p] = true
			}
		})
	}

	return result
}

func nextLetter(val string) string {
	if val == "S" {
		return "a"
	}

	if val == "z" {
		return "E"
	}

	t := val[0]
	t++

	return string(t)
}

func part2(inputfile string) int {
	_, end, grid := load(inputfile)

	lowestPts := []PT{}
	grid.VisitAll2D(func(val string, x, y int) {
		if val == "a" {
			lowestPts = append(lowestPts, PT{x, y, val})
		}
	})

	min := util.MaxInt
	for _, p := range lowestPts {
		dist := calcDist(p, end, grid)
		if dist > 0 {
			min = util.Min(min, dist)
		}
	}

	return min
}
