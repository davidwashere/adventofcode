package util

// RangeOverlaps returns true if any part of r1 or r2 overlap,
// assumes len(r1) == len(r2) == 2
func RangeOverlaps(r1, r2 []int) bool {
	if r1[0] == r2[0] || r1[1] == r2[1] {
		// if they start or end on the same number, is overlap
		return true
	}

	if r1[0] < r2[0] {
		// r1 has lower start, compare end of r1 to beginning of r2
		return r1[1] >= r2[0]
	}

	// else if r1[0] > r2[0]
	// r2 has lower start, compare end of r2 to beginning of r1
	return r2[1] >= r1[0]
}

// RangeFullyContains returns true if either r1 or r2 full contains the other,
// assumes len(r1) == len(r2) == 2
func RangeFullyContains(r1, r2 []int) bool {
	if r1[0] == r2[0] || r1[1] == r2[1] {
		// if they start or end on the same number, one must fully contain the other
		return true
	} else if r1[0] <= r2[0] {
		return r1[1] >= r2[1]
	} else {
		// r2 start less then r1 start
		return r2[1] >= r1[1]
	}
}
