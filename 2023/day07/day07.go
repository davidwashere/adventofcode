package day07

import (
	"aoc/util"
	"fmt"
	"sort"
	"strings"
)

var ()

type Typ int

const (
	HighCard Typ = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

var cardRanks = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
	"J": 0,
}

type Hand struct {
	cards    string
	cardsMap map[string]int
	bid      int

	_maxQ int
}

func (h *Hand) Compare(oh *Hand, jokers bool) int {
	if h.GetTyp(jokers) > oh.GetTyp(jokers) {
		return 1
	} else if h.GetTyp(jokers) < oh.GetTyp(jokers) {
		return -1
	}

	for i := 0; i < len(h.cards); i++ {
		c := string(h.cards[i])
		oc := string(oh.cards[i])

		cr := cardRanks[c]
		ocr := cardRanks[oc]

		if cr > ocr {
			return 1
		} else if cr < ocr {
			return -1
		}
	}

	return 0
}

func (h *Hand) GetTyp(jokers bool) Typ {

	if h.MaxUnique(jokers) == 1 {
		return FiveOfKind
	}

	if h.MaxUnique(jokers) == 2 {
		if h.MaxQuantity(jokers) == 4 {
			return FourOfKind
		}

		if h.MaxQuantity(jokers) == 3 {
			return FullHouse
		}
	}

	if h.MaxUnique(jokers) == 3 {
		if h.MaxQuantity(jokers) == 3 {
			return ThreeOfKind
		}

		if h.MaxQuantity(jokers) == 2 {
			return TwoPair
		}
	}

	if h.MaxQuantity(jokers) == 2 {
		return OnePair
	}

	return HighCard
}

func (h *Hand) MaxUnique(jokers bool) int {
	if jokers && h.cardsMap["J"] > 0 {
		if len(h.cardsMap)-1 == 0 {
			return 1
		}
		return len(h.cardsMap) - 1
	}

	return len(h.cardsMap)
}

func (h *Hand) MaxQuantity(jokers bool) int {
	if h._maxQ > 0 {
		return h._maxQ
	}

	max := util.MinInt
	for card, q := range h.cardsMap {
		if jokers && card == "J" {
			continue
		}
		max = util.Max(max, q)
	}

	if jokers {
		max += h.cardsMap["J"]
	}

	h._maxQ = max
	return max
}

func (h *Hand) NumJokers() int {
	return h.cardsMap["J"]
}

func load(inputFile string) []*Hand {
	data, _ := util.ReadFileToStringSlice(inputFile)
	hands := []*Hand{}

	for _, line := range data {
		split := strings.Split(line, " ")

		ints := util.ParseInts(split[1])

		bid := ints[len(ints)-1]

		cards := split[0]
		m := map[string]int{}
		for i := 0; i < len(cards); i++ {
			c := string(cards[i])
			m[c]++
		}

		hand := &Hand{
			bid:      bid,
			cards:    cards,
			cardsMap: m,
		}

		hands = append(hands, hand)
	}

	return hands
}

func part1(inputFile string) int {
	hands := load(inputFile)

	jokers := false

	for _, hand := range hands {
		fmt.Printf("%+v - Type: %v\n", hand, hand.GetTyp(jokers))
	}

	sort.Slice(hands, func(i, j int) bool {
		cmp := hands[i].Compare(hands[j], jokers)
		return cmp < 0
	})

	fmt.Println("======")
	for _, hand := range hands {
		fmt.Printf("%+v - Type: %v\n", hand, hand.GetTyp(jokers))
	}

	sum := 0
	for i, hand := range hands {
		sum += ((i + 1) * hand.bid)
		fmt.Printf("%v * %v\n", hand.bid, i+1)
	}

	return sum
}

func part2(inputFile string) int {
	load(inputFile)

	hands := load(inputFile)

	jokers := true

	for _, hand := range hands {
		// if !strings.Contains(hand.cards, "J") {
		// 	continue
		// }
		if hand.cards != "JJJJJ" {
			continue
		}
		typ := hand.GetTyp(jokers)
		fmt.Printf("%+v - Type: %v\n", hand, typ)
	}

	sort.Slice(hands, func(i, j int) bool {
		cmp := hands[i].Compare(hands[j], jokers)
		return cmp < 0
	})

	fmt.Println("======")
	for _, hand := range hands {
		fmt.Printf("%+v - Type: %v\n", hand, hand.GetTyp(jokers))
	}

	sum := 0
	for i, hand := range hands {
		sum += ((i + 1) * hand.bid)
		// fmt.Printf("%v * %v\n", hand.bid, i+1)
	}

	return sum
}
