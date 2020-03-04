package api

// IRectangle represents 2D points
type IRectangle interface {
	// Dimesions returns width, height
	Dimesions() (float64, float64)

	// DimesionsAsInt32 returns width, height used for render context
	DimesionsAsInt32() (int32, int32)

	// Min returns the upper left corner
	Min() IPoint
	// Max returns the lower right corner
	Max() IPoint

	// Set sets corners i.e. the dimensions
	Set(minx, miny, maxx, maxy float64)

	// SetBounds set the min/max based on vertex array
	SetBounds(vertices []IPoint)

	// Expand expands the rectangle if needed based on array.
	Expand(vertices []IPoint)

	// Intersets indicates if "this" rectangle intersects "rect"
	IntersectsOther(rect IRectangle) bool

	// ContainsPoint determines if point is within "this" rectangle
	ContainsPoint(p IPoint) bool
}
