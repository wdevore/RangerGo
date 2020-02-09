package api

// IMesh represents 2D polygon
type IMesh interface {
	// Vertices returns the original vertices
	Vertices() []IPoint

	// AddVertex appends the point to vertices
	AddVertex(x, y float64)

	// Build
	Build()
}
