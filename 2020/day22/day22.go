package day22

import (
	"aoc2020/util"
	"fmt"
	"strings"
)

func parsefile(inputfile string) []util.IntQueue {
	data, _ := util.ReadFileToStringSlice(inputfile)
	decks := []util.IntQueue{}
	for _, line := range data {
		if strings.Index(line, ":") > 0 {
			decks = append(decks, util.NewIntQueue())
			continue
		}

		if strings.TrimSpace(line) == "" {
			continue
		}

		ints := util.ParseInts(line)
		decks[len(decks)-1].Enqueue(ints[0])
	}

	return decks

}

func part1(inputfile string) int {
	decks := parsefile(inputfile)

	for _, deck := range decks {
		fmt.Println(deck)
	}

	p1 := decks[0]
	p2 := decks[1]

	rounds := 0
	for !p1.IsEmpty() && !p2.IsEmpty() {
		rounds++
		p1c := p1.Dequeue()
		p2c := p2.Dequeue()

		if p1c > p2c {
			p1.Enqueue(p1c)
			p1.Enqueue(p2c)
		} else {
			p2.Enqueue(p2c)
			p2.Enqueue(p1c)
		}

		// fmt.Printf("-- Round %v --\n", rounds)
		// fmt.Println("P1: ", p1)
		// fmt.Println("P1: ", p2)

	}

	var deck util.IntQueue
	if p2.IsEmpty() {
		deck = p1
	} else {
		deck = p2
	}

	result := 0
	for i := len(deck) - 1; i >= 0; i-- {
		multi := len(deck) - i
		result += deck[i] * multi
	}

	return result
}

type GameT struct {
	num        int
	p1         util.IntQueue
	p2         util.IntQueue
	prevRounds map[string]struct{}
	round      int
}

// NewGame .
func NewGame(num int, p1, p2 util.IntQueue) *GameT {
	return &GameT{
		num:        num,
		p1:         p1,
		p2:         p2,
		prevRounds: map[string]struct{}{},
		round:      0,
	}
}

func getRoundKey(g *GameT) string {
	key := ""
	for _, v := range g.p1 {
		key += fmt.Sprintf("%v,", v)
	}
	key += "|"
	for _, v := range g.p2 {
		key += fmt.Sprintf("%v,", v)
	}

	return key
}

func part2(inputfile string) int {
	decks := parsefile(inputfile)

	for _, deck := range decks {
		fmt.Println(deck)
	}

	gameNum := 1
	rootGame := NewGame(gameNum, decks[0], decks[1])

	gameStack := util.NewStack()
	gameStack.Push(rootGame)

	for !rootGame.p1.IsEmpty() && !rootGame.p2.IsEmpty() {
		curGame := gameStack.Peek().(*GameT)
		curGame.round++
		// fmt.Println(len(gameStack), curGame.round)
		winner := 0
		roundKey := getRoundKey(curGame)
		if _, ok := curGame.prevRounds[roundKey]; ok {
			gameStack.Pop() // ThisGame
			curGame = gameStack.Peek().(*GameT)
			winner = 1
		} else {
			curGame.prevRounds[roundKey] = struct{}{}
		}

		if winner == 0 {
			if curGame.p1.IsEmpty() || curGame.p2.IsEmpty() {
				if curGame.p2.IsEmpty() {
					winner = 1
				} else {
					winner = 2
				}

				gameStack.Pop() // ThisGame
				curGame = gameStack.Peek().(*GameT)
			}
		}

		if winner == 0 {
			p1c := curGame.p1.Front()
			p2c := curGame.p2.Front()

			if len(curGame.p1) > p1c && len(curGame.p2) > p2c {
				gameNum++
				deck1 := append(util.IntQueue{}, curGame.p1[1:p1c+1]...)
				deck2 := append(util.IntQueue{}, curGame.p2[1:p2c+1]...)
				gameStack.Push(NewGame(gameNum, deck1, deck2))
				continue
			}

			if p1c > p2c {
				winner = 1
			} else {
				winner = 2
			}
		}

		if winner == 1 {
			curGame.p1.Enqueue(curGame.p1.Dequeue())
			curGame.p1.Enqueue(curGame.p2.Dequeue())
		} else if winner == 2 {
			curGame.p2.Enqueue(curGame.p2.Dequeue())
			curGame.p2.Enqueue(curGame.p1.Dequeue())
		}

	}

	var deck util.IntQueue
	if rootGame.p2.IsEmpty() {
		deck = rootGame.p1
	} else {
		deck = rootGame.p2
	}

	result := 0
	for i := len(deck) - 1; i >= 0; i-- {
		multi := len(deck) - i
		result += deck[i] * multi
	}
	return result
}
