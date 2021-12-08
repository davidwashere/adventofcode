package util

// SumArithmeticProgression sums progressive values that follow an
// arithmetic progression (constant increment)
//
// `1 + 2 + 3 + 4 + 5 = 15`
// ```
// first = first num
// last = last num
// length = n  = num of nums
//
// length (first + last)
// ---------------------
//           2
// ```
// ref: https://en.wikipedia.org/wiki/Arithmetic_progression#Sum
func SumArithmeticProgression(length, first, last int) int {
	return length * (first + last) / 2
}
