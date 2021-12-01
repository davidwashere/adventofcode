package util

import (
	"fmt"
	"strings"
)

// InfinityGrid grid with infinite dimensions in all directions
// 2d, 3d, ... 11d, doesn't matter (memory constrained)
type InfinityGrid struct {
	// All Data Points
	// The first key `string` will represent the dimensions behind 2d
	//   e.g. "1" might mean z = 1, "1,2" might mean z=1, w=2
	// The two int keys are the `x` and `y` values of of the grid for that dimension
	data map[string]map[int]map[int]string

	// Default value
	def string

	// Max Extents for every `grid`
	xMinMax infinityGridMinMax
	yMinMax infinityGridMinMax

	// MinMax of every dimension
	dimMinMax []infinityGridMinMax

	// False until a value is set into the grid
	initialized bool

	// When bounds are locked minX, minY, maxX, maxY are not updated dynamically
	// Setting values outside the bounds are ignored
	// Getting values outside the bounds returns default value
	BoundsLocked bool

	// Flips will result in coordinate translations
	flipH bool
	flipV bool

	// Rotating the grid results in coordinate translations
	deg int
}

type infinityGridMinMax struct {
	min int
	max int
}

func newInfinityGridMinMax() infinityGridMinMax {
	return infinityGridMinMax{
		min: MaxInt,
		max: MinInt,
	}
}

// NewInfinityGrid .
func NewInfinityGrid(defaultValue string) *InfinityGrid {
	return &InfinityGrid{
		data:    map[string]map[int]map[int]string{},
		def:     defaultValue,
		xMinMax: newInfinityGridMinMax(),
		yMinMax: newInfinityGridMinMax(),
	}
}

// NewInfinityGridFromFile .
func NewInfinityGridFromFile(filename string, defaultValue string) *InfinityGrid {
	grid := NewInfinityGrid(defaultValue)

	x := 0
	y := 0
	err := ParseFileAsString(filename, func(line string) {
		for _, char := range line {
			c := string(char)
			grid.Set(c, x, y)
			x++
		}
		y++
		x = 0
	})

	Check(err)

	return grid
}

// NewInfinityGridFromSlice .
func NewInfinityGridFromSlice(data []string, defaultValue string) *InfinityGrid {
	grid := NewInfinityGrid(defaultValue)

	x := 0
	y := 0
	for _, line := range data {
		for _, char := range line {
			grid.Set(string(char), x, y)
			x++
		}
		y++
		x = 0
	}

	return grid
}

func _createDimKey(dims ...int) string {
	result := ""

	if len(dims) == 0 {
		return ""
	}

	lastNonZeroIndex := -1
	for i := len(dims) - 1; i >= 0; i-- {
		if dims[i] != 0 {
			lastNonZeroIndex = i
			break
		}
	}

	for i := 0; i <= lastNonZeroIndex; i++ {
		result += fmt.Sprintf("%v", dims[i])
		if i != lastNonZeroIndex {
			result += ","
		}
	}

	return result
}

// GetDimensions Returns the number of explicity defined dimensions on top of x and y
// ie: for a 3d grid, GetDimensions will return 1, 4d grid will return 2
func (g *InfinityGrid) GetDimensions() int {
	return len(g.dimMinMax)
}

// AddDimension will add another 'dimension' to the Grid
// This can also be done by setting a value with a coordinate in the desired dimension
func (g *InfinityGrid) AddDimension() {
	minMax := newInfinityGridMinMax()
	minMax.min = 0
	minMax.max = 1
	g.dimMinMax = append(g.dimMinMax, minMax)
}

// applyRotateToCoords will manipulate coords x, y coords based on current rotation
func (g *InfinityGrid) applyRotateToCoords(x, y int) (int, int) {
	// rad := float64(g.deg) * (math.Pi / 180.0)

	// ca := math.Cos(rad)
	// sa := math.Sin(rad)
	// tx := math.Round(ca*float64(x) - sa*float64(y))
	// ty := math.Round(sa*float64(x) + ca*float64(y))

	// x = int(tx)
	// y = int(ty)

	if g.deg == 0 {
		return x, y

	} else if g.deg == 90 || g.deg == -270 {
		return y, g.xMinMax.max - x + g.xMinMax.min

	} else if Abs(g.deg) == 180 {
		return g.xMinMax.max - x + g.xMinMax.min, g.yMinMax.max - y + g.yMinMax.min

	} else if g.deg == 270 || g.deg == -90 {
		return g.yMinMax.max - y + g.yMinMax.min, x

	}

	return x, y
}

// Set .
func (g *InfinityGrid) Set(val string, x, y int, dims ...int) {
	x, y = g.applyRotateToCoords(x, y)

	if g.flipH {
		y = g.yMinMax.max - y + g.yMinMax.min
	}

	if g.flipV {
		x = g.xMinMax.max - x + g.yMinMax.min
	}

	if !g.BoundsLocked {
		g.xMinMax.min = Min(g.xMinMax.min, x)
		g.xMinMax.max = Max(g.xMinMax.max, x)
		g.yMinMax.min = Min(g.yMinMax.min, y)
		g.yMinMax.max = Max(g.yMinMax.max, y)

		for i, dim := range dims {
			if i > len(g.dimMinMax)-1 {
				g.dimMinMax = append(g.dimMinMax, newInfinityGridMinMax())
			}
			g.dimMinMax[i].min = Min(g.dimMinMax[i].min, dim)
			g.dimMinMax[i].max = Max(g.dimMinMax[i].max, dim)
		}
	}

	dimKey := _createDimKey(dims...)
	data := g.data

	if _, ok := data[dimKey]; !ok {
		if g.BoundsLocked {
			return // dim doesn't exist, do not create it if bounds locked
		}
		data[dimKey] = map[int]map[int]string{}
	}

	if _, ok := data[dimKey][x]; !ok {
		if g.BoundsLocked {
			return // x doesn't exist, do not create it if bounds locked
		}
		data[dimKey][x] = map[int]string{}
	}

	data[dimKey][x][y] = val
	g.initialized = true
}

// Get .
func (g *InfinityGrid) Get(x, y int, dims ...int) string {
	if !g.initialized {
		return g.def
	}

	x, y = g.applyRotateToCoords(x, y)

	if g.flipH {
		y = g.yMinMax.max - y + g.yMinMax.min
	}

	if g.flipV {
		x = g.xMinMax.max - x + g.xMinMax.min
	}

	// If the x,y coord is outside current extents return default value
	if x < g.xMinMax.min || x > g.xMinMax.max || y < g.yMinMax.min || y > g.yMinMax.max {
		return g.def
	}

	data := g.data
	dimKey := _createDimKey(dims...)

	if _, ok := data[dimKey]; !ok {
		return g.def
	}

	if _, ok := data[dimKey][x]; !ok {
		return g.def
	}

	if _, ok := data[dimKey][x][y]; !ok {
		return g.def
	}

	return data[dimKey][x][y]
}

// GetRow .
func (g *InfinityGrid) GetRow(y int, dims ...int) []string {
	var row []string
	for x := g.xMinMax.min; x <= g.xMinMax.max; x++ {
		row = append(row, g.Get(x, y, dims...))
	}

	return row
}

// GetCol .
func (g *InfinityGrid) GetCol(x int, dims ...int) []string {
	var col []string
	for y := g.yMinMax.min; y <= g.yMinMax.max; y++ {
		col = append(col, g.Get(x, y, dims...))
	}

	return col
}

// Width .
func (g *InfinityGrid) Width() int {
	if !g.initialized {
		return 0
	}

	return Abs(g.xMinMax.max-g.xMinMax.min) + 1
}

// Height .
func (g *InfinityGrid) Height() int {
	if !g.initialized {
		return 0
	}

	return Abs(g.yMinMax.max-g.yMinMax.min) + 1
}

// VisitAll will visit every grid coordinate with extents based on
// grids current width & height
func (g *InfinityGrid) VisitAll(visitFunc func(val string, x int, y int, dims ...int)) {
	allDims := calcAllDims(g.dimMinMax)

	for _, dims := range allDims {
		for y := g.yMinMax.min; y <= g.yMinMax.max; y++ {
			for x := g.xMinMax.min; x <= g.xMinMax.max; x++ {
				visitFunc(g.Get(x, y, dims...), x, y, dims...)
			}
		}
	}
}

// VisitAll2D will assume all dimensions behind x, y are 0
func (g *InfinityGrid) VisitAll2D(visitFunc func(val string, x int, y int)) {
	for y := g.yMinMax.min; y <= g.yMinMax.max; y++ {
		for x := g.xMinMax.min; x <= g.xMinMax.max; x++ {
			visitFunc(g.Get(x, y), x, y)
		}
	}
}

// VisitAll3D will assume all dimensions behind x, y, z are 0
func (g *InfinityGrid) VisitAll3D(visitFunc func(val string, x int, y int, z int)) {
	if g.GetDimensions() <= 0 {
		return
	}

	for z := g.dimMinMax[0].min; z <= g.dimMinMax[0].max; z++ {
		for y := g.yMinMax.min; y <= g.yMinMax.max; y++ {
			for x := g.xMinMax.min; x <= g.xMinMax.max; x++ {
				visitFunc(g.Get(x, y, z), x, y, z)
			}
		}
	}
}

// VisitAll4D will assume all dimensions behind x, y, z, w are 0
func (g *InfinityGrid) VisitAll4D(visitFunc func(val string, x int, y int, z int, w int)) {
	if g.GetDimensions() <= 1 {
		return
	}

	for w := g.dimMinMax[1].min; w <= g.dimMinMax[1].max; w++ {
		for z := g.dimMinMax[0].min; z <= g.dimMinMax[0].max; z++ {
			for y := g.yMinMax.min; y <= g.yMinMax.max; y++ {
				for x := g.xMinMax.min; x <= g.xMinMax.max; x++ {
					visitFunc(g.Get(x, y, z, w), x, y, z, w)
				}
			}
		}
	}
}

func calcAllDims(dimMinMax []infinityGridMinMax) [][]int {
	allDims := &[][]int{}

	calcAllDimsRecur(dimMinMax, 0, allDims, []int{})

	return *allDims
}

func calcAllDimsRecur(dimMinMax []infinityGridMinMax, index int, allDims *[][]int, curDim []int) {
	if index == len(dimMinMax) {
		*allDims = append(*allDims, curDim)
		return
	}

	for i := dimMinMax[index].min; i <= dimMinMax[index].max; i++ {
		var newDim []int
		newDim = append(newDim, curDim...)
		newDim = append(newDim, i)

		calcAllDimsRecur(dimMinMax, index+1, allDims, newDim)
	}
}

// Grow Will extend the min, max of every grid and dimension by amt
// Useful for expanding the extents when using VisitAll
func (g *InfinityGrid) Grow(amt int) {
	g.xMinMax.min -= amt
	g.xMinMax.max += amt
	g.yMinMax.min -= amt
	g.yMinMax.max += amt

	for i := 0; i < len(g.dimMinMax); i++ {
		g.dimMinMax[i].min -= amt
		g.dimMinMax[i].max += amt
	}
}

// Shrink will reduce the min, max of every grid and dimension by amt
// Data previously set will remain in-tact such that if the grid is later
// expanded the out of bounds data can be retrieved
func (g *InfinityGrid) Shrink(amt int) {
	g.xMinMax.min += amt
	g.xMinMax.max -= amt
	g.yMinMax.min += amt
	g.yMinMax.max -= amt

	for i := 0; i < len(g.dimMinMax); i++ {
		g.dimMinMax[i].min += amt
		g.dimMinMax[i].max -= amt
	}
}

// GetMinX .
func (g *InfinityGrid) GetMinX() int {
	return g.xMinMax.min
}

// GetMinY .
func (g *InfinityGrid) GetMinY() int {
	return g.yMinMax.min
}

// GetMaxX .
func (g *InfinityGrid) GetMaxX() int {
	return g.xMinMax.max
}

// GetMaxY .
func (g *InfinityGrid) GetMaxY() int {
	return g.yMinMax.max
}

// LockBounds locks the bounds of the grid
func (g *InfinityGrid) LockBounds() {
	g.BoundsLocked = true
}

// UnlockBounds unlocks the bounds of the grid
func (g *InfinityGrid) UnlockBounds() {
	g.BoundsLocked = false
}

// SetExtents Sets the max extents for a 2d grid to the specified values
// extents are also increased automatically when using Set
func (g *InfinityGrid) SetExtents(minX, minY, maxX, maxY int) {
	g.xMinMax.min = minX
	g.xMinMax.max = maxX
	g.yMinMax.min = minY
	g.yMinMax.max = maxY
}

// Dump Prints out text representation of grid, assumes each values is a single character
func (g *InfinityGrid) Dump(dims ...int) {
	if !g.initialized {
		fmt.Println("Grid Not Initialized")
	}

	for y := g.yMinMax.max; y >= g.yMinMax.min; y-- {
		for x := g.xMinMax.min; x <= g.xMinMax.max; x++ {
			val := g.Get(x, y, dims...)
			if val == "" {
				val = " "
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}

// TopEdge .
func (g *InfinityGrid) TopEdge(dims ...int) string {
	return strings.Join(g.GetRow(g.yMinMax.max, dims...), "")
}

// BottomEdge .
func (g *InfinityGrid) BottomEdge(dims ...int) string {
	return strings.Join(g.GetRow(g.yMinMax.min, dims...), "")
}

// LeftEdge .
func (g *InfinityGrid) LeftEdge(dims ...int) string {
	return strings.Join(g.GetCol(g.xMinMax.min, dims...), "")
}

// RightEdge .
func (g *InfinityGrid) RightEdge(dims ...int) string {
	return strings.Join(g.GetCol(g.xMinMax.max, dims...), "")
}

// Edges returns all grid edges
func (g *InfinityGrid) Edges(dims ...int) []string {
	var edges []string
	return append(edges, g.LeftEdge(), g.RightEdge(), g.TopEdge(), g.BottomEdge())
}

// EdgesFlipped returns all grid edges
func (g *InfinityGrid) EdgesFlipped(dims ...int) []string {
	var edges []string
	g.FlipH()
	g.FlipV()
	edges = append(edges, g.LeftEdge(), g.RightEdge(), g.TopEdge(), g.BottomEdge())
	g.FlipH()
	g.FlipV()

	return edges
}

// FlipH .
func (g *InfinityGrid) FlipH() {
	g.flipH = !g.flipH
}

// FlipV .
func (g *InfinityGrid) FlipV() {
	g.flipV = !g.flipV
}

// Rotate only rotates in increments of 90 are accepted
func (g *InfinityGrid) Rotate(deg int) {
	g.deg = (g.deg + deg) % 360
}
