package util

import (
	"testing"
)

func TestSumArithmeticProgression(t *testing.T) {
	// `1 + 2 + 3 + 4 + 5 = 15  (length=5, first=1, last=5)` --> `5(1+5)/2  =  30/2  =  15`
	want := 15
	got := SumArithmeticProgression(5, 1, 5)
	vf(t, got, want)
}
