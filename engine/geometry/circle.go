package geometry

import (
	"math"

	"github.com/wdevore/RangerGo/api"
)

// Circle is a circle with behaviours
type Circle struct {
	radius float64

	center api.IPoint
}

// NewCircle creates a circle
func NewCircle() *Circle {
	o := new(Circle)
	o.radius = 1.0
	o.center = NewPoint()
	return o
}

// SetRadius of circle
func (c *Circle) SetRadius(radius float64) {
	c.radius = radius
}

// Radius of circle
func (c *Circle) Radius() float64 {
	return c.radius
}

// SetCenter of circle
func (c *Circle) SetCenter(x, y float64) {
	c.center.SetByComp(x, y)
}

// Center of circle
func (c *Circle) Center() api.IPoint {
	return c.center
}

// PointInside checks point to radius
func (c *Circle) PointInside(point api.IPoint) bool {
	distance := c.DistanceFromCenter(point)
	return distance < c.radius
}

// DistanceFromCenter returns distance from point to circle center
func (c *Circle) DistanceFromCenter(point api.IPoint) float64 {
	dx := c.center.X() - point.X()
	dy := c.center.Y() - point.Y()
	return math.Sqrt(dx*dx + dy*dy)
}

// DistanceFromEdge returns distance from point to circle edge
// if <= 0 then on inside edge.
func (c *Circle) DistanceFromEdge(point api.IPoint) float64 {
	dx := c.center.X() - point.X()
	dy := c.center.Y() - point.Y()
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance - c.radius
}
