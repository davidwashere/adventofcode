package util

import (
	"testing"
)

func TestPerms(t *testing.T) {
	sample := []string{"a", "b", "c"}

	got := [][]string{}
	PermsString(sample, true, func(perm []string) {
		got = append(got, perm)
	})

	vf(t, len(got), 27)
	vf(t, got[0][0], "a")
	vf(t, got[0][1], "a")
	vf(t, got[0][2], "a")

	vf(t, got[1][0], "a")
	vf(t, got[1][1], "a")
	vf(t, got[1][2], "b")

	vf(t, got[25][0], "c")
	vf(t, got[25][1], "c")
	vf(t, got[25][2], "b")

	vf(t, got[26][0], "c")
	vf(t, got[26][1], "c")
	vf(t, got[26][2], "c")

	got = [][]string{}
	PermsString(sample, false, func(perm []string) {
		got = append(got, perm)
	})

	vf(t, len(got), 6)
	vf(t, got[0][0], "a")
	vf(t, got[0][1], "b")
	vf(t, got[0][2], "c")

	vf(t, got[1][0], "a")
	vf(t, got[1][1], "c")
	vf(t, got[1][2], "b")

	vf(t, got[2][0], "b")
	vf(t, got[2][1], "a")
	vf(t, got[2][2], "c")

	vf(t, got[3][0], "b")
	vf(t, got[3][1], "c")
	vf(t, got[3][2], "a")

	vf(t, got[4][0], "c")
	vf(t, got[4][1], "a")
	vf(t, got[4][2], "b")

	vf(t, got[5][0], "c")
	vf(t, got[5][1], "b")
	vf(t, got[5][2], "a")
}

func TestCharPerms(t *testing.T) {
	sample := "abc"

	got := []string{}
	PermsChar(sample, true, func(perm string) {
		got = append(got, perm)
	})

	vf(t, len(got), 27)
	vf(t, got[0], "aaa")
	vf(t, got[1], "aab")
	vf(t, got[25], "ccb")
	vf(t, got[26], "ccc")

	got = []string{}
	PermsChar(sample, false, func(perm string) {
		got = append(got, perm)
	})

	vf(t, len(got), 6)
	vf(t, got[0], "abc")
	vf(t, got[1], "acb")
	vf(t, got[2], "bac")
	vf(t, got[3], "bca")
	vf(t, got[4], "cab")
	vf(t, got[5], "cba")
}

func TestCharPermsOfLen(t *testing.T) {
	sample := "abc"

	got := []string{}
	PermsCharOfLen(sample, 2, true, func(perm string) {
		got = append(got, perm)
	})

	vf(t, len(got), 9)
	vf(t, got[0], "aa")
	vf(t, got[1], "ab")
	vf(t, got[2], "ac")
	vf(t, got[3], "ba")
	vf(t, got[4], "bb")
	vf(t, got[5], "bc")
	vf(t, got[6], "ca")
	vf(t, got[7], "cb")
	vf(t, got[8], "cc")

	got = []string{}
	PermsCharOfLen(sample, 2, false, func(perm string) {
		got = append(got, perm)
	})

	vf(t, len(got), 6)
	vf(t, got[0], "ab")
	vf(t, got[1], "ac")
	vf(t, got[2], "ba")
	vf(t, got[3], "bc")
	vf(t, got[4], "ca")
	vf(t, got[5], "cb")
}

func TestPermsOfLen(t *testing.T) {
	sample := []string{"a", "b", "c"}

	got := [][]string{}
	PermsStringOfLen(sample, 2, true, func(perm []string) {
		got = append(got, perm)
	})

	vf(t, len(got), 9)
	vf(t, got[0][0], "a")
	vf(t, got[0][1], "a")
	vf(t, got[1][0], "a")
	vf(t, got[1][1], "b")
	vf(t, got[2][0], "a")
	vf(t, got[2][1], "c")
	vf(t, got[3][0], "b")
	vf(t, got[3][1], "a")
	vf(t, got[4][0], "b")
	vf(t, got[4][1], "b")
	vf(t, got[5][0], "b")
	vf(t, got[5][1], "c")
	vf(t, got[6][0], "c")
	vf(t, got[6][1], "a")
	vf(t, got[7][0], "c")
	vf(t, got[7][1], "b")
	vf(t, got[8][0], "c")
	vf(t, got[8][1], "c")

	got = [][]string{}
	PermsStringOfLen(sample, 2, false, func(perm []string) {
		got = append(got, perm)
	})

	vf(t, len(got), 6)
	vf(t, got[0][0], "a")
	vf(t, got[0][1], "b")
	vf(t, got[1][0], "a")
	vf(t, got[1][1], "c")
	vf(t, got[2][0], "b")
	vf(t, got[2][1], "a")
	vf(t, got[3][0], "b")
	vf(t, got[3][1], "c")
	vf(t, got[4][0], "c")
	vf(t, got[4][1], "a")
	vf(t, got[5][0], "c")
	vf(t, got[5][1], "b")
}
