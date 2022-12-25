package day20

import (
	"aoc/util"
	"fmt"
)

type Num struct {
	val  int
	prev *Num
	next *Num
}

func (n Num) String() string {
	return fmt.Sprintf("%v", n.val)
}

func load(inputFile string) []int {
	data, _ := util.ReadFileToStringSlice(inputFile)

	nums := []int{}
	for _, line := range data {
		ints := util.ParseIntsNeg(line)
		nums = append(nums, ints[0])
	}

	return nums
}

func buildLLNoMap(nums []int) (*Num, []*Num) {
	numL := []*Num{}

	var zero *Num

	for _, n := range nums {
		node := new(Num)
		node.val = n
		if n == 0 {
			zero = node
		}
		numL = append(numL, node)
	}

	for i := 0; i < len(numL)-1; i++ {
		cur := numL[i]
		next := numL[i+1]

		cur.next = next
		next.prev = cur
	}

	last := len(numL) - 1
	numL[0].prev = numL[last]
	numL[last].next = numL[0]

	return zero, numL
}

func part1(inputFile string) int {
	nums := load(inputFile)
	zeroPt, numL := buildLLNoMap(nums)

	for _, n := range numL {
		pt := n
		cur := n

		if pt.val == 0 {
			continue
		}

		pt.prev.next = pt.next
		pt.next.prev = pt.prev

		if pt.val > 0 {
			// moving it 'right'
			moves := pt.val
			for i := 0; i < moves; i++ {
				cur = cur.next
			}
		} else if pt.val < 0 {
			// moving it left
			moves := util.Abs(pt.val)
			for i := 0; i <= moves; i++ {
				cur = cur.prev
			}
		}

		pt.next = cur.next
		pt.prev = cur
		cur.next = pt
		pt.next.prev = pt
	}

	cur := zeroPt
	sum := 0
	for i := 0; i < 3000; i++ {
		cur = cur.next
		if (i+1)%1000 == 0 {
			sum += cur.val
		}
	}

	return sum
}

func part2(inputFile string) int {
	nums := load(inputFile)
	zeroPt, numL := buildLLNoMap(nums)

	for _, n := range numL {
		n.val = n.val * 811589153
	}

	for mix := 0; mix < 10; mix++ {

		for _, n := range numL {
			pt := n
			cur := n

			if pt.val == 0 {
				continue
			}

			pt.prev.next = pt.next
			pt.next.prev = pt.prev

			if pt.val > 0 {
				// moving it 'right'
				moves := pt.val % (len(numL) - 1)
				for i := 0; i < moves; i++ {
					cur = cur.next
				}
			} else if pt.val < 0 {
				// moving it left
				moves := util.Abs(pt.val) % (len(numL) - 1)
				for i := 0; i <= moves; i++ {
					cur = cur.prev
				}
			}

			pt.next = cur.next
			pt.prev = cur
			cur.next = pt
			pt.next.prev = pt
		}
	}

	cur := zeroPt
	sum := 0
	for i := 0; i < 3000; i++ {
		cur = cur.next
		if (i+1)%1000 == 0 {
			sum += cur.val
		}
	}

	return sum
}
