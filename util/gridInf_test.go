package util

import (
	"fmt"
	"testing"
)

func TestInfinityGridThings(t *testing.T) {
	got := _createDimKey(1, 3, 7)
	want := "1,3,7"
	vf(t, got, want)
}

func TestInfinityGridDimKey(t *testing.T) {
	vf(t, _createDimKey(1, 3, 7), "1,3,7")
	vf(t, _createDimKey(), "")
}

func TestInfinityGridSetGet(t *testing.T) {
	grid := NewInfinityGrid(".")
	// 2D
	grid.Set("a", 0, 0)
	grid.Set("b", 1, 1)
	grid.Set("c", -1, -1)

	vf(t, grid.Get(0, 0), "a")
	vf(t, grid.Get(1, 1), "b")
	vf(t, grid.Get(-1, -1), "c")

	// 3D
	grid.Set("x", 0, 0, 1)
	grid.Set("z", -1, -1, -1)
	grid.Set("G", 6, 6, 0)

	vf(t, grid.Get(0, 0, 0), "a")
	vf(t, grid.Get(1, 1, 0), "b")
	vf(t, grid.Get(-1, -1, 0), "c")
	vf(t, grid.Get(0, 0, 1), "x")
	vf(t, grid.Get(-1, -1, -1), "z")
	vf(t, grid.Get(6, 6, 0), "G")
}

func TestInfinityGridWidth(t *testing.T) {
	g := NewInfinityGrid(".")
	vf(t, g.Width(), 0)

	g.Set("A", 0, 0)
	g.Set("C", 1, 1)
	g.Set("B", -1, -1)
	g.Set("D", -1, 1)
	vf(t, g.Width(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", 2, 0)
	g.Set("B", 3, 0)
	g.Set("C", 4, 0)
	vf(t, g.Width(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", -2, 0)
	g.Set("B", -3, 0)
	g.Set("C", -4, 0)
	vf(t, g.Width(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", -2, 0, -1)
	g.Set("B", -3, 0, 1)
	g.Set("C", -4, 0, 20)
	vf(t, g.Width(), 3)
}

func TestInfinityGridHeight(t *testing.T) {
	g := NewInfinityGrid(".")
	vf(t, g.Height(), 0)

	g.Set("A", 0, 0)
	g.Set("C", 1, 1)
	g.Set("B", 0, -1)
	vf(t, g.Height(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", 0, 2)
	g.Set("B", 0, 3)
	g.Set("C", 0, 4)
	vf(t, g.Height(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", 0, -2)
	g.Set("B", 0, -3)
	g.Set("C", 0, -4)
	vf(t, g.Height(), 3)

	g = NewInfinityGrid(".")
	g.Set("A", 0, -2, -1)
	g.Set("B", 0, -3, 1)
	g.Set("C", 0, -4, 20)
	vf(t, g.Height(), 3)
}

func TestInfinityGridAllDims(t *testing.T) {
	dimMinMax := []infinityGridMinMax{
		infinityGridMinMax{-1, 1},
	}

	allDims := calcAllDims(dimMinMax)

	vf(t, len(allDims), 3)
	vf(t, allDims[0][0], -1)
	vf(t, allDims[1][0], 0)
	vf(t, allDims[2][0], 1)

	dimMinMax = []infinityGridMinMax{
		infinityGridMinMax{-1, 1},
		infinityGridMinMax{-1, 1},
	}

	allDims = calcAllDims(dimMinMax)
	vf(t, len(allDims), 9)
	vf(t, allDims[0][0], -1)
	vf(t, allDims[0][1], -1)

	vf(t, allDims[1][0], -1)
	vf(t, allDims[1][1], 0)

	vf(t, allDims[2][0], -1)
	vf(t, allDims[2][1], 1)

	vf(t, allDims[7][0], 1)
	vf(t, allDims[7][1], 0)

	vf(t, allDims[8][0], 1)
	vf(t, allDims[8][1], 1)
}

func TestInfinityGridLockBounds(t *testing.T) {
	g := NewInfinityGrid(".")
	g.Set("A", 0, 0)
	g.LockBounds()
	g.Set("C", 1, 1)
	g.Set("B", -1, -1)

	g.VisitAll2D(func(val string, x, y int) {
		vf(t, val, "A")
	})

	vf(t, g.Height(), 1)
	vf(t, g.Width(), 1)

	g.UnlockBounds()
	g.Set("C", 1, 1)
	g.Set("B", -1, -1)

	vf(t, g.Height(), 3)
	vf(t, g.Width(), 3)
}

func TestInfinityGridSet2DExtents(t *testing.T) {
	g := NewInfinityGrid(".")
	g.Set("A", 0, 0)
	g.Set("B", 3, 3)
	g.SetExtents(-5, -5, 5, 5)
	vf(t, g.Height(), 11)
	vf(t, g.Width(), 11)

	g.SetExtents(0, 0, 2, 2)
	vf(t, g.Get(3, 3), ".")
	g.SetExtents(0, 0, 4, 4)
	vf(t, g.Get(3, 3), "B")
}

func TestInfinityGridFlips(t *testing.T) {
	g := NewInfinityGrid(".")
	// B    C
	//
	//
	// A    D
	g.Set("A", 0, 0)
	g.Set("B", 0, 5)
	g.Set("C", 5, 5)
	g.Set("D", 5, 0)

	g.FlipH()
	// A    D
	//
	//
	// B    C
	vf(t, g.Get(0, 0), "B")
	vf(t, g.Get(0, 5), "A")
	vf(t, g.Get(5, 5), "D")
	vf(t, g.Get(5, 0), "C")

	// back to original
	g.FlipH()
	// B    C
	//
	//
	// A    D
	vf(t, g.Get(0, 0), "A")
	vf(t, g.Get(0, 5), "B")
	vf(t, g.Get(5, 5), "C")
	vf(t, g.Get(5, 0), "D")

	g.FlipV()
	// C    B
	//
	//
	// D    A
	vf(t, g.Get(0, 0), "D")
	vf(t, g.Get(0, 5), "C")
	vf(t, g.Get(5, 5), "B")
	vf(t, g.Get(5, 0), "A")

	// back to original
	g.FlipV()
	// B    C
	//
	//
	// A    D
	vf(t, g.Get(0, 0), "A")
	vf(t, g.Get(0, 5), "B")
	vf(t, g.Get(5, 5), "C")
	vf(t, g.Get(5, 0), "D")

	// DoubleFlip
	g.FlipV()
	g.FlipH()
	// D    A
	//
	//
	// C    B
	vf(t, g.Get(0, 0), "C")
	vf(t, g.Get(0, 5), "D")
	vf(t, g.Get(5, 5), "A")
	vf(t, g.Get(5, 0), "B")
}

func TestInfinityGridRotate(t *testing.T) {
	g := NewInfinityGrid(".")
	// B    C
	//
	//
	// A    D
	g.Set("A", 0, 0)
	g.Set("B", 0, 5)
	g.Set("C", 5, 5)
	g.Set("D", 5, 0)

	g.Rotate(90)
	// C    D
	//
	//
	// B    A
	vf(t, g.Get(0, 0), "B")
	vf(t, g.Get(0, 5), "C")
	vf(t, g.Get(5, 5), "D")
	vf(t, g.Get(5, 0), "A")

	g.Rotate(90)
	// D    A
	//
	//
	// C    B
	vf(t, g.Get(0, 0), "C")
	vf(t, g.Get(0, 5), "D")
	vf(t, g.Get(5, 5), "A")
	vf(t, g.Get(5, 0), "B")

	g.Rotate(90)
	// A    B
	//
	//
	// D    C
	vf(t, g.Get(0, 0), "D")
	vf(t, g.Get(0, 5), "A")
	vf(t, g.Get(5, 5), "B")
	vf(t, g.Get(5, 0), "C")

	g.Rotate(90)
	// B    C
	//
	//
	// A    D
	vf(t, g.Get(0, 0), "A")
	vf(t, g.Get(0, 5), "B")
	vf(t, g.Get(5, 5), "C")
	vf(t, g.Get(5, 0), "D")

	g.Rotate(-180)
	// D    A
	//
	//
	// C    B
	vf(t, g.Get(0, 0), "C")
	vf(t, g.Get(0, 5), "D")
	vf(t, g.Get(5, 5), "A")
	vf(t, g.Get(5, 0), "B")
}

func TestInfinityGridDirectionalEdges(t *testing.T) {
	g := NewInfinityGrid(".")
	// BC
	// AD
	g.Set("A", 0, 0)
	g.Set("B", 0, 1)
	g.Set("C", 1, 1)
	g.Set("D", 1, 0)

	vf(t, g.TopEdge(), "BC")
	vf(t, g.BottomEdge(), "AD")
	vf(t, g.LeftEdge(), "AB")
	vf(t, g.RightEdge(), "DC")

	g.FlipH()
	g.FlipV()

	// DA
	// CB
	vf(t, g.TopEdge(), "DA")
	vf(t, g.BottomEdge(), "CB")
	vf(t, g.LeftEdge(), "CD")
	vf(t, g.RightEdge(), "BA")

	g.FlipH()
	g.FlipV()
	g.Rotate(-90)
	// AB
	// DC
	vf(t, g.TopEdge(), "AB")
	vf(t, g.BottomEdge(), "DC")
	vf(t, g.LeftEdge(), "DA")
	vf(t, g.RightEdge(), "CB")

	g.FlipH()
	g.FlipV()
	// CD
	// BA
	vf(t, g.TopEdge(), "CD")
	vf(t, g.BottomEdge(), "BA")
	vf(t, g.LeftEdge(), "BC")
	vf(t, g.RightEdge(), "AD")

	g.Rotate(-90)
	// BC
	// AD
	vf(t, g.TopEdge(), "BC")
	vf(t, g.BottomEdge(), "AD")
	vf(t, g.LeftEdge(), "AB")
	vf(t, g.RightEdge(), "DC")
}

func Test(t *testing.T) {
	g := NewInfinityGrid(".")
	g.Set("A", 0, 0)
	g.Set("B", 0, 1)
	g.Set("C", 1, 1)
	g.Set("D", 1, 0)

	// BC
	// AD
	left := g.LeftEdge()
	vf(t, left, "AB")

	g.Rotate(-90)
	// AB
	// DC
	left = g.LeftEdge()
	vf(t, left, "DA")

	g.Rotate(-90)
	// DA
	// CB
	left = g.LeftEdge()
	vf(t, left, "CD")

	g.Rotate(-90)
	// CD
	// BA
	left = g.LeftEdge()
	vf(t, left, "BC")

	fmt.Println("Flip")
	g.FlipH()
	g.FlipV()

	// AB
	// DC
	left = g.LeftEdge()
	vf(t, left, "DA")

	g.Rotate(-90)
	// DA
	// CB
	left = g.LeftEdge()
	vf(t, left, "CD")

	g.Rotate(-90)
	// CD
	// BA
	left = g.LeftEdge()
	vf(t, left, "BC")

	g.Rotate(-90)
	// BC
	// AD
	left = g.LeftEdge()
	vf(t, left, "AB")
}
