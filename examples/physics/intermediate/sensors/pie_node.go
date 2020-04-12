package main

import (
	"math"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// PieNode is a basic vector circle shape.
type PieNode struct {
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

// NewPieNode constructs a circle shaped node
func NewPieNode(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(PieNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (c *PieNode) Build(world api.IWorld) {
	c.Node.Build(world)

	c.lineColor = rendering.NewPaletteInt64(rendering.White)
	c.color = rendering.NewPaletteInt64(rendering.Yellow)
}

// Configure pie, if radius is 1 then diameter is 2
func (c *PieNode) Configure(segments int, radius, startAngle, radianSize float64) {
	c.p1 = geometry.NewPointUsing(0.75*math.Cos(math.Pi/8.0), 0.75*math.Sin(math.Pi/8.0))
	c.p2 = geometry.NewPointUsing(1.0*math.Cos(math.Pi/8.0), 1.0*math.Sin(math.Pi/8.0))
	c.o1 = geometry.NewPoint()
	c.o2 = geometry.NewPoint()

	c.segments = segments // typically 12
	c.radius = radius     // typically 1.0

	c.polygon = geometry.NewPolygon()

	step := radianSize / float64(c.segments)
	c.polygon.AddVertex(0.0, 0.0)

	for i := startAngle; i <= startAngle+radianSize; i += step {
		x := math.Cos(i) * c.radius
		y := math.Sin(i) * c.radius
		c.polygon.AddVertex(x, y)
	}

	c.polygon.Build()
}

// SetRadius sets circle's radius (default = 1.0)
func (c *PieNode) SetRadius(radius float64) {
	c.radius = radius
}

// SetSegments sets how many segments on the circle (default = 12)
func (c *PieNode) SetSegments(segments int) {
	c.segments = segments
}

// SetColor sets circle's color (default = white)
func (c *PieNode) SetColor(color api.IPalette) {
	c.color = color
}

// Draw renders shape
func (c *PieNode) Draw(context api.IRenderContext) {
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
