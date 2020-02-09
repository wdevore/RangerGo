package api

// IPolygon represents 2D lines
type IPolygon interface {
	// AddVertex appends the point to vertices
	AddVertex(x, y float64)

	// Build
	Build()

	PointInside(p IPoint) bool
}
