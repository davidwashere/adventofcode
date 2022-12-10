package util

import (
	"testing"
)

func TestInfGridThings(t *testing.T) {
	got := _createInfGridDimKey(1, 3, 7)
	want := "1,3,7"
	vf(t, got, want)
}

func TestInfGridDimKey(t *testing.T) {
	vf(t, _createInfGridDimKey(1, 3, 7), "1,3,7")
	vf(t, _createInfGridDimKey(), "")
}

func TestInfGridSetGet(t *testing.T) {
	grid := NewInfGrid()
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

func TestInfGridWidth(t *testing.T) {
	g := NewInfGrid()
	vf(t, g.Width(), 0)

	g.Set("A", 0, 0)
	g.Set("C", 1, 1)
	g.Set("B", -1, -1)
	g.Set("D", -1, 1)
	vf(t, g.Width(), 3)

	g = NewInfGrid()
	g.Set("A", 2, 0)
	g.Set("B", 3, 0)
	g.Set("C", 4, 0)
	vf(t, g.Width(), 3)

	g = NewInfGrid()
	g.Set("A", -2, 0)
	g.Set("B", -3, 0)
	g.Set("C", -4, 0)
	vf(t, g.Width(), 3)

	g = NewInfGrid()
	g.Set("A", -2, 0, -1)
	g.Set("B", -3, 0, 1)
	g.Set("C", -4, 0, 20)
	vf(t, g.Width(), 3)
}

func TestInfGridHeight(t *testing.T) {
	g := NewInfGrid()
	vf(t, g.Height(), 0)

	g.Set("A", 0, 0)
	g.Set("C", 1, 1)
	g.Set("B", 0, -1)
	vf(t, g.Height(), 3)

	g = NewInfGrid()
	g.Set("A", 0, 2)
	g.Set("B", 0, 3)
	g.Set("C", 0, 4)
	vf(t, g.Height(), 3)

	g = NewInfGrid()
	g.Set("A", 0, -2)
	g.Set("B", 0, -3)
	g.Set("C", 0, -4)
	vf(t, g.Height(), 3)

	g = NewInfGrid()
	g.Set("A", 0, -2, -1)
	g.Set("B", 0, -3, 1)
	g.Set("C", 0, -4, 20)
	vf(t, g.Height(), 3)
}

func TestInfGridAllDims(t *testing.T) {
	dimMinMax := []extents{
		{-1, 1},
	}

	allDims := getAllDims(dimMinMax)

	vf(t, len(allDims), 3)
	vf(t, allDims[0][0], -1)
	vf(t, allDims[1][0], 0)
	vf(t, allDims[2][0], 1)

	dimMinMax = []extents{
		{-1, 1},
		{-1, 1},
	}

	allDims = getAllDims(dimMinMax)
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

func TestInfGridLockBounds(t *testing.T) {
	g := NewInfGrid()
	g.Set("A", 0, 0)
	g.LockBounds()
	g.Set("C", 1, 1)
	g.Set("B", -1, -1)

	g.VisitAll2D(func(val interface{}, x, y int) {
		v := val.(string)
		vf(t, v, "A")
	})

	vf(t, g.Height(), 1)
	vf(t, g.Width(), 1)

	g.UnlockBounds()
	g.Set("C", 1, 1)
	g.Set("B", -1, -1)

	vf(t, g.Height(), 3)
	vf(t, g.Width(), 3)
}

func TestInfGridSet2DExtents(t *testing.T) {
	g := NewInfGrid()
	g.Set("A", 0, 0)
	g.Set("B", 3, 3)
	g.SetExtents(-5, -5, 5, 5)
	vf(t, g.Height(), 11)
	vf(t, g.Width(), 11)

	g.SetExtents(0, 0, 2, 2)
	vf(t, g.Get(3, 3), nil)
	g.SetExtents(0, 0, 4, 4)
	vf(t, g.Get(3, 3), "B")
}

func TestInfGridFlips(t *testing.T) {
	g := NewInfGrid()
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

func TestInfGridRotate(t *testing.T) {
	g := NewInfGrid()
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

func conv(in []interface{}) string {
	out := ""

	for _, v := range in {
		out += v.(string)
	}

	return out
}

func TestInfGridDirectionalEdges(t *testing.T) {
	g := NewInfGrid()
	// BC
	// AD
	g.Set("A", 0, 0)
	g.Set("B", 0, 1)
	g.Set("C", 1, 1)
	g.Set("D", 1, 0)

	vf(t, conv(g.TopEdge()), "BC")
	vf(t, conv(g.BottomEdge()), "AD")
	vf(t, conv(g.LeftEdge()), "AB")
	vf(t, conv(g.RightEdge()), "DC")

	g.FlipH()
	g.FlipV()

	// DA
	// CB
	vf(t, conv(g.TopEdge()), "DA")
	vf(t, conv(g.BottomEdge()), "CB")
	vf(t, conv(g.LeftEdge()), "CD")
	vf(t, conv(g.RightEdge()), "BA")

	g.FlipH()
	g.FlipV()
	g.Rotate(-90)
	// AB
	// DC
	vf(t, conv(g.TopEdge()), "AB")
	vf(t, conv(g.BottomEdge()), "DC")
	vf(t, conv(g.LeftEdge()), "DA")
	vf(t, conv(g.RightEdge()), "CB")

	g.FlipH()
	g.FlipV()
	// CD
	// BA
	vf(t, conv(g.TopEdge()), "CD")
	vf(t, conv(g.BottomEdge()), "BA")
	vf(t, conv(g.LeftEdge()), "BC")
	vf(t, conv(g.RightEdge()), "AD")

	g.Rotate(-90)
	// BC
	// AD
	vf(t, conv(g.TopEdge()), "BC")
	vf(t, conv(g.BottomEdge()), "AD")
	vf(t, conv(g.LeftEdge()), "AB")
	vf(t, conv(g.RightEdge()), "DC")
}

func TestVisitOrtho(t *testing.T) {
	g := NewInfGrid()
	// BC
	// AD
	g.Set("A", 0, 0)
	g.Set("B", 0, 1)
	g.Set("C", 1, 1)
	g.Set("D", 1, 0)

	g.VisitOrtho(0, 0, func(val interface{}, x, y int) {
		if x == 0 && y == 1 {
			vf(t, val, "B")
			return
		}

		if x == 1 && y == 0 {
			vf(t, val, "D")
			return
		}

		if (x == 0 && y == -1) || (x == -1 && y == 0) {
			vf(t, val, nil)
			return
		}

		t.Errorf("Unexpected visit x=%v y=%v val=%v", x, y, val)
	})

	g.LockBounds()

	g.VisitOrtho(0, 0, func(val interface{}, x, y int) {
		if x == 0 && y == 1 {
			vf(t, val, "B")
			return
		}

		if x == 1 && y == 0 {
			vf(t, val, "D")
			return
		}

		t.Errorf("Unexpected visit x=%v y=%v val=%v", x, y, val)
	})
}

func TestVisitDiag(t *testing.T) {
	g := NewInfGrid()
	// BC
	// AD
	g.Set("A", 0, 0)
	g.Set("B", 0, 1)
	g.Set("C", 1, 1)
	g.Set("D", 1, 0)

	g.VisitDiag(0, 0, func(val interface{}, x, y int) {
		if x == 1 && y == 1 {
			vf(t, val, "C")
			return
		}

		if (x == 1 && y == -1) || (x == -1 && y == -1) || (x == -1 && y == 1) {
			vf(t, val, nil)
			return
		}

		t.Errorf("Unexpected visit x=%v y=%v val=%v", x, y, val)
	})

	g.LockBounds()

	g.VisitDiag(0, 0, func(val interface{}, x, y int) {
		if x == 1 && y == 1 {
			vf(t, val, "C")
			return
		}

		t.Errorf("Unexpected visit x=%v y=%v val=%v", x, y, val)
	})
}

func TestLen(t *testing.T) {
	g := NewInfGrid()
	g.Set("A", 0, 0)
	g.Set("B", 5, 5)
	g.Set("C", 0, 5)

	// g.WithDefaultValue(".")
	// g.Dump()
	want := 3

	got := g.Len()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	// overwrite a value, len should not change
	g.Set("D", 0, 5)

	got = g.Len()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	// delete a value and len should be reduced
	want = 2
	g.Delete(0, 5)
	got = g.Len()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	// try another dimension
	g.Set("Z", 5, 5, 1)
	// there should only be one coordinate in this dimension
	want = 1
	got = g.Len(1) // use z=1

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

}
