package misc

import (
	"math"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
)

// AABB is an axis-aligned-bounding-box typically used for debug visuals
// but can be used for anything.
type AABB struct {
	min api.IPoint
	max api.IPoint
}

// NewAABB constructs a new AABB
func NewAABB() *AABB {
	o := new(AABB)
	o.min = geometry.NewPoint()
	o.max = geometry.NewPoint()
	return o
}

// Min returns the upper-left corner
func (a *AABB) Min() api.IPoint {
	return a.min
}

// Max returns the lower-right corner
func (a *AABB) Max() api.IPoint {
	return a.max
}

// SetBounds sets the bounds based on the provided Mesh
func (a *AABB) SetBounds(vertices []api.IPoint) {
	minx := math.MaxFloat64
	miny := math.MaxFloat64
	maxx := math.SmallestNonzeroFloat64
	maxy := math.SmallestNonzeroFloat64

	for _, v := range vertices {
		minx = math.Min(minx, v.X())
		maxx = math.Max(maxx, v.X())
		miny = math.Min(miny, v.Y())
		maxy = math.Max(maxy, v.Y())
	}

	a.min.SetByComp(minx, miny)
	a.max.SetByComp(maxx, maxy)
}

// Expand expands the current AABB using the provided mesh
func (a *AABB) Expand(vertices []api.IPoint) {
	minx := a.min.X()
	miny := a.min.Y()
	maxx := a.max.X()
	maxy := a.max.Y()

	for _, v := range vertices {
		minx = math.Min(minx, v.X())
		maxx = math.Max(maxx, v.X())
		miny = math.Min(miny, v.Y())
		maxy = math.Max(maxy, v.Y())
	}

	a.min.SetByComp(minx, miny)
	a.max.SetByComp(maxx, maxy)
}
