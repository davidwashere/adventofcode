package util

import (
	"fmt"
)

// InfGrid represents an Infinity Grid
//
// An infinity grid allows for coordinates to be set in any
// direction and any dimension
//
// An infinity grid has these features:
// - only consumes memory for values that have been set
// - x,y can be negative, positive, and to infinity (Max Int)
// - coordinates can be set in any dimension (2d, 4d, 100d)
//
type InfGrid struct {
	// Grid data, the first `string` key represents the dimension
	//  i.e. "1" is z=1, "1,2" is z=1, w=2, and so on
	//
	// The two `int` keys represent the `x` and `y` positions in the 2d grid
	data map[string]map[int]map[int]interface{}

	// Default value for positions that are not set, by default will be <nil>
	def interface{}

	// Max and min x and y coordinate positions, applies to every x,y / 2d grid
	xExtents extents
	yExtents extents

	// Min and max coordinates for of every dimension, ie: z min = 0, z max = 100
	dimExtents []extents

	// ---

	// False until a value is set into the grid
	initialized bool

	// When bounds are locked minX, minY, maxX, maxY are not updated dynamically
	// on Set operations
	//
	// Setting values outside the bounds are ignored
	//
	// Getting values outside the bounds returns default value
	boundsLocked bool

	// controls if grid should be flipped horizontally or vertically
	// setting these values will modify behavior of access operations
	flipH bool
	flipV bool

	// degress by which to rotate grid access operations
	deg int
}

type extents struct {
	min int
	max int
}

func NewInfGrid() *InfGrid {
	grid := InfGrid{
		data:     map[string]map[int]map[int]interface{}{},
		yExtents: newExtents(),
		xExtents: newExtents(),
	}
	return &grid
}

func newExtents() extents {
	return extents{
		min: MaxInt,
		max: MinInt,
	}
}

func (g *InfGrid) WithDefaultValue(defaultValue interface{}) *InfGrid {
	g.def = defaultValue
	return g
}

// Set a value to the grid at the provided coordinates and dimensions
func (g *InfGrid) Set(val interface{}, x, y int, dims ...int) {
	x, y = g.applyRotateToCoords(x, y)

	if g.flipH {
		y = g.yExtents.max - y + g.yExtents.min
	}

	if g.flipV {
		x = g.xExtents.max - x + g.xExtents.min
	}

	if !g.boundsLocked {
		g.xExtents.min = Min(g.xExtents.min, x)
		g.xExtents.max = Max(g.xExtents.max, x)
		g.yExtents.min = Min(g.yExtents.min, y)
		g.yExtents.max = Max(g.yExtents.max, y)

		for i, dim := range dims {
			if i > len(g.dimExtents)-1 {
				g.dimExtents = append(g.dimExtents, newExtents())
			}
			g.dimExtents[i].min = Min(g.dimExtents[i].min, dim)
			g.dimExtents[i].max = Max(g.dimExtents[i].max, dim)
		}
	}

	dimKey := _createInfGridDimKey(dims...)
	data := g.data

	if _, ok := data[dimKey]; !ok {
		if g.boundsLocked {
			return // dim doesn't exist, do not create it if bounds locked
		}
		data[dimKey] = map[int]map[int]interface{}{}
	}

	if _, ok := data[dimKey][x]; !ok {
		if g.boundsLocked {
			return // x doesn't exist, do not create it if bounds locked
		}
		data[dimKey][x] = map[int]interface{}{}
	}

	data[dimKey][x][y] = val
	g.initialized = true
}

// applyRotateToCoords will manipulate x, y coords based on current rotation
// will only manipulate coords of rotation is increment of 90
func (g *InfGrid) applyRotateToCoords(x, y int) (int, int) {
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
		return y, g.xExtents.max - x + g.xExtents.min

	} else if Abs(g.deg) == 180 {
		return g.xExtents.max - x + g.xExtents.min, g.yExtents.max - y + g.yExtents.min

	} else if g.deg == 270 || g.deg == -90 {
		return g.yExtents.max - y + g.yExtents.min, x

	}

	return x, y
}

// NumDimensions Returns the number of explicity defined dimensions on top of x and y
// ie: for a 3d grid will return 1, 4d grid will return 2
func (g *InfGrid) NumDimensions() int {
	return len(g.dimExtents)
}

// AddDimension will add another 'dimension' to the Grid
// This can also be done by setting a value with a coordinate in the desired dimension
func (g *InfGrid) AddDimension() {
	extents := newExtents()
	// TODO: Why did i mark these as 0 - 1?
	extents.min = 0
	extents.max = 1
	g.dimExtents = append(g.dimExtents, extents)
}

// Get a value from the grid at the provided coordinates and dimensions, returns the 'default'
// value if the coordinate is outside the extents of the grid
func (g *InfGrid) Get(x, y int, dims ...int) interface{} {
	if !g.initialized {
		return g.def
	}

	x, y = g.applyRotateToCoords(x, y)

	if g.flipH {
		y = g.yExtents.max - y + g.yExtents.min
	}

	if g.flipV {
		x = g.xExtents.max - x + g.xExtents.min
	}

	// If the x,y coord is outside current extents return default value
	if x < g.xExtents.min || x > g.xExtents.max || y < g.yExtents.min || y > g.yExtents.max {
		return g.def
	}

	data := g.data
	dimKey := _createInfGridDimKey(dims...)

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

// GetRow returns the grid 'row' represented by the y cooirdinate and dimensions
func (g *InfGrid) GetRow(y int, dims ...int) []interface{} {
	var row []interface{}
	for x := g.xExtents.min; x <= g.xExtents.max; x++ {
		row = append(row, g.Get(x, y, dims...))
	}

	return row
}

// GetCol returns the grid 'col' represented by the x cooirdinate and dimensions
func (g *InfGrid) GetCol(x int, dims ...int) []interface{} {
	var col []interface{}
	for y := g.yExtents.min; y <= g.yExtents.max; y++ {
		col = append(col, g.Get(x, y, dims...))
	}

	return col
}

// Width returns the width of the grid, negative extents are taken into consideration
// for example if there are positions set in the grid at -2,0 and 2,0, the width returned
// would be 4
func (g *InfGrid) Width() int {
	if !g.initialized {
		return 0
	}

	return Abs(g.xExtents.max-g.xExtents.min) + 1
}

// Height returns the height of the grid, negative extents are taken into consideration
// for example if there are positions set in the grid at -2,0 and 2,0, the width returned
// would be 4
func (g *InfGrid) Height() int {
	if !g.initialized {
		return 0
	}

	return Abs(g.yExtents.max-g.yExtents.min) + 1
}

// getAllDims will produce a list of all combinations (permutations) of all dimensions so that they
// be interated through - for example, if valid dimensional coordinates are z=1, z=2, w=1, y=2, the
// output would be the following to represent all the different dimenions
//   z=1, w=1
//   z=1, w=2
//   z=2, w=1
//   z=2, w=2
func getAllDims(dimExtents []extents) [][]int {
	allDims := &[][]int{}

	getAllDimsRecur(dimExtents, 0, allDims, []int{})

	return *allDims
}

// getAllDimsRecur is a recursive utility function for getAllDims, see getAllDims description
func getAllDimsRecur(dimExtents []extents, index int, allDims *[][]int, curDim []int) {
	if index == len(dimExtents) {
		*allDims = append(*allDims, curDim)
		return
	}

	for i := dimExtents[index].min; i <= dimExtents[index].max; i++ {
		var newDim []int
		newDim = append(newDim, curDim...)
		newDim = append(newDim, i)

		getAllDimsRecur(dimExtents, index+1, allDims, newDim)
	}
}

// _createInfGridDimKey will create a key based on the provided dimensions that is
// used to pull the appropriate grid from the backing data store of an InfGrid
func _createInfGridDimKey(dims ...int) string {
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

// VisitAll will visit every grid coordinate with extents based on
// the grids current width & height
func (g *InfGrid) VisitAll(visitFunc func(val interface{}, x int, y int, dims ...int)) {
	allDims := getAllDims(g.dimExtents)

	for _, dims := range allDims {
		for y := g.yExtents.min; y <= g.yExtents.max; y++ {
			for x := g.xExtents.min; x <= g.xExtents.max; x++ {
				visitFunc(g.Get(x, y, dims...), x, y, dims...)
			}
		}
	}
}

// VisitAll2D is the same as VisitAll but with all dimensions behind x, y set to 0
func (g *InfGrid) VisitAll2D(visitFunc func(val interface{}, x int, y int)) {
	for y := g.yExtents.min; y <= g.yExtents.max; y++ {
		for x := g.xExtents.min; x <= g.xExtents.max; x++ {
			visitFunc(g.Get(x, y), x, y)
		}
	}
}

// VisitAll3D is the same as VisitAll but with all dimensions behind x, y, and z set to 0
func (g *InfGrid) VisitAll3D(visitFunc func(val interface{}, x int, y int, z int)) {
	if g.NumDimensions() <= 0 {
		return
	}

	for z := g.dimExtents[0].min; z <= g.dimExtents[0].max; z++ {
		for y := g.yExtents.min; y <= g.yExtents.max; y++ {
			for x := g.xExtents.min; x <= g.xExtents.max; x++ {
				visitFunc(g.Get(x, y, z), x, y, z)
			}
		}
	}
}

// VisitAll4D is the same as VisitAll but with all dimensions behind x, y, z, and w set to 0
func (g *InfGrid) VisitAll4D(visitFunc func(val interface{}, x int, y int, z int, w int)) {
	if g.NumDimensions() <= 1 {
		return
	}

	for w := g.dimExtents[1].min; w <= g.dimExtents[1].max; w++ {
		for z := g.dimExtents[0].min; z <= g.dimExtents[0].max; z++ {
			for y := g.yExtents.min; y <= g.yExtents.max; y++ {
				for x := g.xExtents.min; x <= g.xExtents.max; x++ {
					visitFunc(g.Get(x, y, z, w), x, y, z, w)
				}
			}
		}
	}
}

// Grow Will extend the min, max of every grid and dimension by amt
// Useful for expanding the extents when using VisitAll
func (g *InfGrid) Grow(amt int) {
	g.xExtents.min -= amt
	g.xExtents.max += amt
	g.yExtents.min -= amt
	g.yExtents.max += amt

	for i := 0; i < len(g.dimExtents); i++ {
		g.dimExtents[i].min -= amt
		g.dimExtents[i].max += amt
	}
}

// Shrink will reduce the min, max of every grid and dimension by amt
// Data previously set will remain in-tact such that if the grid is later
// expanded the out of bounds data can be retrieved
func (g *InfGrid) Shrink(amt int) {
	g.xExtents.min += amt
	g.xExtents.max -= amt
	g.yExtents.min += amt
	g.yExtents.max -= amt

	for i := 0; i < len(g.dimExtents); i++ {
		g.dimExtents[i].min += amt
		g.dimExtents[i].max -= amt
	}
}

// GetMinX .
func (g *InfGrid) GetMinX() int {
	return g.xExtents.min
}

// GetMinY .
func (g *InfGrid) GetMinY() int {
	return g.yExtents.min
}

// GetMaxX .
func (g *InfGrid) GetMaxX() int {
	return g.xExtents.max
}

// GetMaxY .
func (g *InfGrid) GetMaxY() int {
	return g.yExtents.max
}

// LockBounds locks the bounds of the grid
func (g *InfGrid) LockBounds() {
	g.boundsLocked = true
}

// UnlockBounds unlocks the bounds of the grid
func (g *InfGrid) UnlockBounds() {
	g.boundsLocked = false
}

// SetExtents Sets the max extents for a 2d grid to the specified values
// extents are also increased automatically when using Set
func (g *InfGrid) SetExtents(minX, minY, maxX, maxY int) {
	g.xExtents.min = minX
	g.xExtents.max = maxX
	g.yExtents.min = minY
	g.yExtents.max = maxY
}

// Dump Prints out text representation of grid, assumes each values is a single character
func (g *InfGrid) Dump(dims ...int) {
	if !g.initialized {
		fmt.Println("Grid Not Initialized")
	}

	for y := g.yExtents.max; y >= g.yExtents.min; y-- {
		for x := g.xExtents.min; x <= g.xExtents.max; x++ {
			val := g.Get(x, y, dims...)
			if val == "" {
				val = " "
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}

// TopEdge returns the top 'row' of the grid at dimensions specified
func (g *InfGrid) TopEdge(dims ...int) []interface{} {
	return g.GetRow(g.yExtents.max, dims...)
}

// BottomEdge returns the bottom 'row' of the grid at dimensions specified
func (g *InfGrid) BottomEdge(dims ...int) []interface{} {
	return g.GetRow(g.yExtents.min, dims...)
}

// LeftEdge returns the left 'col' of the grid at dimensions specified
func (g *InfGrid) LeftEdge(dims ...int) []interface{} {
	return g.GetCol(g.xExtents.min, dims...)
}

// RightEdge returns the right 'col' of the grid at dimensions specified
func (g *InfGrid) RightEdge(dims ...int) []interface{} {
	return g.GetCol(g.xExtents.max, dims...)
}

// Edges returns all grid edges at the dimensions specified
func (g *InfGrid) Edges(dims ...int) [][]interface{} {
	var edges [][]interface{}
	return append(edges, g.LeftEdge(), g.RightEdge(), g.TopEdge(), g.BottomEdge())
}

// EdgesFlipped returns all grid edges at the dimensions specified in reverse
// that is edges after the grid has been flipped horizontally and vertically
func (g *InfGrid) EdgesFlipped(dims ...int) [][]interface{} {
	var edges [][]interface{}
	g.FlipH()
	g.FlipV()
	edges = append(edges, g.LeftEdge(), g.RightEdge(), g.TopEdge(), g.BottomEdge())
	g.FlipH()
	g.FlipV()

	return edges
}

// FlipH will flip all grids horizontally (top row becomes bottom row and so on)
func (g *InfGrid) FlipH() {
	g.flipH = !g.flipH
}

// FlipV will flip all grids vertically (left col becomes right col and so on)
func (g *InfGrid) FlipV() {
	g.flipV = !g.flipV
}

// Rotate will rotate the grid by a certain number of degrees, if the deg's provided
// are not an increment of 90 the function will be a noop
func (g *InfGrid) Rotate(deg int) {
	if deg%90 != 0 {
		return
	}

	g.deg = (g.deg + deg) % 360
}

// GetN will return the value north of the given coordinate (y+1), if the coordinate is outside the extents
// of the grid returns the default value
func (g *InfGrid) GetN(x, y int, dims ...int) interface{} {
	return g.Get(x, y+1, dims...)
}

// GetE will return the value east of the given coordinate (x+1), if the coordinate is outside the extents
// of the grid returns the default value
func (g *InfGrid) GetE(x, y int, dims ...int) interface{} {
	return g.Get(x+1, y, dims...)
}

// GetS will return the value south of the given coordinate (y-1), if the coordinate is outside the extents
// of the grid returns the default value
func (g *InfGrid) GetS(x, y int, dims ...int) interface{} {
	return g.Get(x, y-1, dims...)
}

// GetW will return the value west of the given coordinate (x-1), if the coordinate is outside the extents
// of the grid returns the default value
func (g *InfGrid) GetW(x, y int, dims ...int) interface{} {
	return g.Get(x-1, y, dims...)
}

// GetNE will return the value north-east of the given coordinate (x+1, y+1), if the coordinate is outside the extents
// of the grid returns the default value
func (g *InfGrid) GetNE(x, y int, dims ...int) interface{} {
	return g.Get(x+1, y+1, dims...)
}

// GetSE will return the value south-east of the given coordinate (x+1, y-1), if the coordinate is outside the extents
// of the grid returns the default value
func (g *InfGrid) GetSE(x, y int, dims ...int) interface{} {
	return g.Get(x+1, y-1, dims...)
}

// GetSW will return the value south-west of the given coordinate (x-1, y-1), if the coordinate is outside the extents
// of the grid returns the default value
func (g *InfGrid) GetSW(x, y int, dims ...int) interface{} {
	return g.Get(x-1, y-1, dims...)
}

// GetNW will return the value north-west of the given coordinate (x-1, y+1), if the coordinate is outside the extents
// of the grid returns the default value
func (g *InfGrid) GetNW(x, y int, dims ...int) interface{} {
	return g.Get(x-1, y+1, dims...)
}

// GetNESW will return the values north, east, south, and west of the given coordinate, if a coordinate is outside
// the extents of the grid it will be set to the default
func (g *InfGrid) GetNESW(x, y int, dims ...int) []interface{} {

	return []interface{}{
		g.GetN(x, y, dims...),
		g.GetE(x, y, dims...),
		g.GetS(x, y, dims...),
		g.GetW(x, y, dims...),
	}
}

// GetNeighbors will return the values north, east, south, west, north-east, south-east, south-west, and north-west
// of the given coordinate, if a coordinate is outside the extents of the grid it will be set to the default
func (g *InfGrid) GetNeighbors(x, y int, dims ...int) []interface{} {

	return []interface{}{
		g.GetN(x, y, dims...),
		g.GetE(x, y, dims...),
		g.GetS(x, y, dims...),
		g.GetW(x, y, dims...),
		g.GetNE(x, y, dims...),
		g.GetSE(x, y, dims...),
		g.GetSW(x, y, dims...),
		g.GetNW(x, y, dims...),
	}
}

// Delete will delete an item found at coords (set its value to the default), extents are not affected
func (g *InfGrid) Delete(x, y int, dims ...int) interface{} {
	v := g.Get(x, y, dims...)
	g.Set(g.def, x, y, dims...)
	return v
}
