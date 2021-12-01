package day14

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestBit(t *testing.T) {
	num := 11

	fmt.Printf("%b\n", num)

	num = 0b100

	fmt.Printf("%b\n", num)

	num = 1 << 2

	fmt.Printf("%b\n", num)

	var mask int
	mask = 0b1101
	num = 0b1011
	// num = 11
	num = mask & num
	fmt.Printf("%b\n", num)

	maskS := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"
	nnum := int64(11)
	andMaskS := strings.ReplaceAll(maskS, "X", "1")
	orMaskS := strings.ReplaceAll(maskS, "X", "0")
	andMask, _ := strconv.ParseInt(andMaskS, 2, 64)
	orMask, _ := strconv.ParseInt(orMaskS, 2, 64)

	nnum = int64(nnum) & andMask
	nnum = int64(nnum) | orMask
	fmt.Printf("%b\n", nnum)

	fmt.Printf("%b\n", flipBits(maskS, 11))

}

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 0
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 0
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestP2(t *testing.T) {
	got := part2("sample2.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 0
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	// want := 0
	// if got != want {
	// 	t.Errorf("Got %v but want %v", got, want)
	// }
}

func TestFlipBits2(t *testing.T) {
	maskS := "00X000000000000000000000000000X1001X"
	fmt.Printf("%b\n", flipBits2(maskS, 42))
}

func TestPerms(t *testing.T) {
	maskS := "000000000000000000000000000000X1001X"
	num := flipBits2(maskS, 42)
	result := perms(maskS, num)
	fmt.Println(result)
}
