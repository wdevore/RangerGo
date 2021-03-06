package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// TriangleNode is a basic triangle
type TriangleNode struct {
	nodes.Node

	color api.IPalette

	polygon api.IPolygon
}

// NewTriangleNode constructs a triangle shaped node whos tip points
// along the +Y axis.
func NewTriangleNode(name string, world api.IWorld, parent api.INode) *TriangleNode {
	o := new(TriangleNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (t *TriangleNode) Build(world api.IWorld) {
	t.Node.Build(world)

	t.polygon = geometry.NewPolygon()
	t.polygon.AddVertex(-0.5, -0.5)
	t.polygon.AddVertex(0.0, 0.5)
	t.polygon.AddVertex(0.5, -0.5)

	t.polygon.Build()

	t.color = rendering.NewPaletteInt64(rendering.White)
}

// Polygon returns the internal polygon mesh
func (t *TriangleNode) Polygon() api.IPolygon {
	return t.polygon
}

// SetColor sets line color
func (t *TriangleNode) SetColor(color api.IPalette) {
	t.color = color
}

// SetPoints sets the edge points of the triangle
func (t *TriangleNode) SetPoints(x1, y1, x2, y2, x3, y3 float64) {
	t.polygon.SetVertex(x1, y1, 0)
	t.polygon.SetVertex(x2, y2, 1)
	t.polygon.SetVertex(x3, y3, 2)
}

// Draw renders shape
func (t *TriangleNode) Draw(context api.IRenderContext) {
	if t.IsDirty() {
		context.TransformPolygon(t.polygon)

		t.SetDirty(false)
	}

	context.SetDrawColor(t.color)
	context.RenderPolygon(t.polygon, api.CLOSED)
}
