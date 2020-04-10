package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// BigPointNode is 9 pixels square
type BigPointNode struct {
	nodes.Node

	lineColor api.IPalette

	p1 api.IPoint

	o1 api.IPoint
}

// NewBigPointNode constructs a cross shaped node
func NewBigPointNode(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(BigPointNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (b *BigPointNode) Build(world api.IWorld) {
	b.Node.Build(world)

	b.p1 = geometry.NewPoint()

	b.o1 = geometry.NewPoint()

	b.lineColor = rendering.NewPaletteInt64(rendering.White)
}

// SetColor sets line color
func (b *BigPointNode) SetColor(color api.IPalette) {
	b.lineColor = color
}

// SetPoint sets the center position
func (b *BigPointNode) SetPoint(x, y float64) {
	b.p1.SetByComp(x, y)
	b.SetDirty(true)
}

// Draw renders shape
func (b *BigPointNode) Draw(context api.IRenderContext) {
	if b.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformPoint(b.p1, b.o1)
		b.SetDirty(false) // Node is no longer dirty
	}

	context.SetDrawColor(b.lineColor)
	context.DrawBigPoint(int32(b.o1.X()), int32(b.o1.Y()))
}
