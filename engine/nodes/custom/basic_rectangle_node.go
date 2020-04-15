package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// BasicRectangleNode is a basic rectangle node.
type BasicRectangleNode struct {
	nodes.Node

	color api.IPalette

	min, max api.IPoint
	o1, o2   api.IPoint
}

// NewBasicRectangleNode constructs a rectangle shaped node
func NewBasicRectangleNode(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(BasicRectangleNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (r *BasicRectangleNode) Build(world api.IWorld) {
	r.Node.Build(world)

	r.min = geometry.NewPointUsing(0.0, 0.0)
	r.max = geometry.NewPointUsing(1.0, 1.0)

	r.o1 = geometry.NewPoint()
	r.o2 = geometry.NewPoint()

	r.color = rendering.NewPaletteInt64(rendering.White)
}

// SetBoundsUsingComps sets the min/max corners
func (r *BasicRectangleNode) SetBoundsUsingComps(minx, miny, maxx, maxy float64) {
	r.min.SetByComp(minx, miny)
	r.max.SetByComp(maxx, maxy)
	r.SetDirty(true)
}

// SetColor sets rectangle color
func (r *BasicRectangleNode) SetColor(color api.IPalette) {
	r.color = color
}

// Draw renders shape
func (r *BasicRectangleNode) Draw(context api.IRenderContext) {
	if r.IsDirty() {
		context.TransformPoint(r.min, r.o1)
		context.TransformPoint(r.max, r.o2)
		r.SetDirty(false)
	}

	context.SetDrawColor(r.color)
	context.RenderAARectangle(r.o1, r.o2, api.FILLED)
}
