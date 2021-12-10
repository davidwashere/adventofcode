package util

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

// Set .
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

	dimKey := _createDimKey(dims...)
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

// Get .
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
