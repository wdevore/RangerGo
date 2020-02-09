package api

// IRectangle represents 2D points
type IRectangle interface {
	// Dimesions returns width, height
	Dimesions() (float64, float64)

	// Min returns the upper left corner
	Min() IPoint
	// Max returns the lower right corner
	Max() IPoint

	// Set sets corners i.e. the dimensions
	Set(minx, miny, maxx, maxy float64)

	// SetAABB

	// Intersets indicates if "this" rectangle intersects "rect"
	IntersectsOther(rect IRectangle) bool

	// ContainsPoint determines if point is within "this" rectangle
	ContainsPoint(p IPoint) bool
}
