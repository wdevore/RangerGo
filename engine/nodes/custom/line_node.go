package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// LineNode is a basic line node
type LineNode struct {
	nodes.Node

	lineColor api.IPalette

	p1 api.IPoint
	p2 api.IPoint

	o1 api.IPoint
	o2 api.IPoint
}

// NewLineNode constructs a cross shaped node
func NewLineNode(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(LineNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (l *LineNode) Build(world api.IWorld) {
	l.Node.Build(world)

	l.p1 = geometry.NewPoint()
	l.p2 = geometry.NewPoint()
	l.o1 = geometry.NewPoint()
	l.o2 = geometry.NewPoint()

	l.lineColor = rendering.NewPaletteInt64(rendering.LightGray)
}

// SetColor sets line color
func (l *LineNode) SetColor(color api.IPalette) {
	l.lineColor = color
}

// SetPoints sets the start and end points of the line.
func (l *LineNode) SetPoints(x1, y1, x2, y2 float64) {
	l.p1.SetByComp(x1, y1)
	l.p2.SetByComp(x2, y2)
}

// Draw renders shape
func (l *LineNode) Draw(context api.IRenderContext) {
	if l.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformPoint(l.p1, l.o1)
		context.TransformPoint(l.p2, l.o2)
		l.SetDirty(false) // Node is no longer dirty
	}

	context.SetDrawColor(l.lineColor)
	context.DrawLine(int32(l.o1.X()), int32(l.o1.Y()), int32(l.o2.X()), int32(l.o2.Y()))
}
