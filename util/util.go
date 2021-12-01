package util

import (
	"fmt"
	"os"
	"strconv"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinAll(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	min := Min(nums[0], nums[1])
	if len(nums) == 2 {
		return min
	}

	for i := 2; i < len(nums); i++ {
		min = Min(min, nums[i])
	}

	return min
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxAll(nums ...int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	max := Max(nums[0], nums[1])
	if len(nums) == 2 {
		return max
	}

	for i := 2; i < len(nums); i++ {
		max = Max(max, nums[i])
	}

	return max
}

func Check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// AllInMap will return true if all elements of slice are found in map, false otherwise
func AllInMap(slice []string, m map[string]string) bool {
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			return false
		}
	}

	return true
}

// IsHex will return true if the string represents valid hex uint64 (ignores case)
func IsHex(val string) bool {
	_, err := strconv.ParseUint(val, 16, 64)
	if err != nil {
		return false
	}
	return true
}
