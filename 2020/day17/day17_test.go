package day17

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {

	// grids := map[int]*util.Grid{}
	// grids[0] = util.NewGridFromFile("sample.txt", INACTIVE)

	// count := activeNeightbors(grids, coord{1, 1, 0, 0})
	// fmt.Println(count)
}

// func TestPointers(t *testing.T) {
// 	grids := map[int]*util.Grid{}
// 	// grids[0] = util.NewGridFromFile("sample.txt", INACTIVE)

// 	grids[0] = util.NewGridFromFile("sample.txt", INACTIVE)

// 	for z, grid := range grids {
// 		fmt.Printf("[%v] %+v\n", z, grid)
// 	}
// 	fmt.Println()

// 	grids[1] = util.NewGrid(INACTIVE)
// 	// grids[1] = util.NewGrid(INACTIVE)
// 	// g := grids[1]
// 	// g := util.NewGrid(INACTIVE)
// 	grids[1].SetExtents(-5, -5, 5, 5)
// 	grids[1].Set(1, 1, "h")
// 	// g.Set(1, 1, "h")
// 	// grids[1] = g

// 	for z, grid := range grids {
// 		fmt.Printf("[%v] %+v\n", z, grid)
// 	}
// 	fmt.Println()
// }

func TestP1(t *testing.T) {
	got := part1("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 112
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP1_Actual(t *testing.T) {
	got := part1("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 301
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2(t *testing.T) {
	got := part2("sample.txt")
	fmt.Printf("Got: %v\n", got)
	want := 848
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}

func TestP2_Actual(t *testing.T) {
	got := part2("input.txt")
	fmt.Printf("Got: %v\n", got)
	want := 2424
	if got != want {
		t.Errorf("Got %v but want %v", got, want)
	}
}
