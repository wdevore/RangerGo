package custom

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

	p1 api.IPoint
	p2 api.IPoint
	p3 api.IPoint

	o1 api.IPoint
	o2 api.IPoint
	o3 api.IPoint
}

// NewTriangleNode constructs a triangle shaped node
func NewTriangleNode(name string, parent api.INode) *TriangleNode {
	o := new(TriangleNode)
	o.Initialize(name)
	o.SetParent(parent)
	return o
}

// Build configures the node
func (t *TriangleNode) Build(world api.IWorld) {
	t.Node.Build(world)

	t.p1 = geometry.NewPointUsing(-0.5, 0.5)
	t.p2 = geometry.NewPointUsing(0.5, 0.5)
	t.p3 = geometry.NewPointUsing(0.0, -0.314)

	t.o1 = geometry.NewPoint()
	t.o2 = geometry.NewPoint()
	t.o3 = geometry.NewPoint()

	t.color = rendering.NewPaletteInt64(rendering.White)
}

// SetColor sets line color
func (t *TriangleNode) SetColor(color api.IPalette) {
	t.color = color
}

// SetPoints sets the start and end points of the line.
func (t *TriangleNode) SetPoints(x1, y1, x2, y2, x3, y3 float64) {
	t.p1.SetByComp(x1, y1)
	t.p2.SetByComp(x2, y2)
	t.p3.SetByComp(x3, y3)
}

// Draw renders shape
func (t *TriangleNode) Draw(context api.IRenderContext) {
	if t.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformPoint(t.p1, t.o1)
		context.TransformPoint(t.p2, t.o2)
		context.TransformPoint(t.p3, t.o3)
		t.SetDirty(false) // Node is no longer dirty
	}

	context.SetDrawColor(t.color)
	context.DrawLine(int32(t.o1.X()), int32(t.o1.Y()), int32(t.o2.X()), int32(t.o2.Y()))
	context.DrawLine(int32(t.o2.X()), int32(t.o2.Y()), int32(t.o3.X()), int32(t.o3.Y()))
	context.DrawLine(int32(t.o3.X()), int32(t.o3.Y()), int32(t.o1.X()), int32(t.o1.Y()))
}
