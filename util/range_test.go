package util

import "testing"

func TestRangeOverlaps(t *testing.T) {
	tt := []struct {
		r1   []int
		r2   []int
		want bool
	}{
		{[]int{2, 4}, []int{6, 8}, false},
		{[]int{2, 3}, []int{4, 5}, false},
		{[]int{5, 7}, []int{7, 9}, true},
		{[]int{2, 8}, []int{3, 7}, true},
		{[]int{6, 6}, []int{4, 6}, true},
	}

	for _, test := range tt {
		got := RangeOverlaps(test.r1, test.r2)
		if got != test.want {
			t.Errorf("got %v want %v for %v %v", got, test.want, test.r1, test.r2)
		}
	}
}

func TestRangeFullyContains(t *testing.T) {
	tt := []struct {
		r1   []int
		r2   []int
		want bool
	}{
		{[]int{2, 4}, []int{6, 8}, false},
		{[]int{2, 8}, []int{3, 7}, true},
		{[]int{6, 6}, []int{4, 6}, true},
		{[]int{8, 8}, []int{8, 82}, true},
	}

	for _, test := range tt {
		got := RangeFullyContains(test.r1, test.r2)
		if got != test.want {
			t.Errorf("got %v want %v for %v %v", got, test.want, test.r1, test.r2)
		}
	}
}
