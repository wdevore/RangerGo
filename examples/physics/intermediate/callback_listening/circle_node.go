package main

import (
	"math"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// CircleNode is a basic vector circle shape.
type CircleNode struct {
	nodes.Node

	color api.IPalette

	segments int
	radius   float64
	polygon  api.IPolygon

	// Rotation marker
	lineColor api.IPalette

	p1 api.IPoint
	p2 api.IPoint
	o1 api.IPoint
	o2 api.IPoint
}

// NewCircleNode constructs a circle shaped node
func NewCircleNode(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(CircleNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (c *CircleNode) Build(world api.IWorld) {
	c.Node.Build(world)

	c.lineColor = rendering.NewPaletteInt64(rendering.White)
	c.color = rendering.NewPaletteInt64(rendering.White)
}

// Configure circle, if radius is 1 then diameter is 2
func (c *CircleNode) Configure(segments int, radius float64) {
	c.p1 = geometry.NewPointUsing(0.5, 0.0)
	c.p2 = geometry.NewPointUsing(1.0, 0.0)
	c.o1 = geometry.NewPoint()
	c.o2 = geometry.NewPoint()

	c.segments = segments // typically 12
	c.radius = radius     // typicall 1.0

	c.polygon = geometry.NewPolygon()

	step := math.Pi / float64(c.segments)

	for i := 0.0; i < 2.0*math.Pi; i += step {
		x := math.Cos(i) * c.radius
		y := math.Sin(i) * c.radius
		c.polygon.AddVertex(x, y)
	}

	c.polygon.Build()
}

// SetRadius sets circle's radius (default = 1.0)
func (c *CircleNode) SetRadius(radius float64) {
	c.radius = radius
}

// SetSegments sets how many segments on the circle (default = 12)
func (c *CircleNode) SetSegments(segments int) {
	c.segments = segments
}

// SetColor sets circle's color (default = white)
func (c *CircleNode) SetColor(color api.IPalette) {
	c.color = color
}

// Draw renders shape
func (c *CircleNode) Draw(context api.IRenderContext) {
	if c.IsDirty() {
		context.TransformPoint(c.p1, c.o1)
		context.TransformPoint(c.p2, c.o2)

		context.TransformPolygon(c.polygon)
		c.SetDirty(false)
	}

	context.SetDrawColor(c.lineColor)
	context.DrawLine(int32(c.o1.X()), int32(c.o1.Y()), int32(c.o2.X()), int32(c.o2.Y()))

	context.SetDrawColor(c.color)
	context.RenderPolygon(c.polygon, api.CLOSED)

}
