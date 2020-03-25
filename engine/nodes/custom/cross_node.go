package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// CrossNode is a basic node shaped like "+"
type CrossNode struct {
	nodes.Node

	color api.IPalette

	h1  api.IPoint
	h2  api.IPoint
	oh1 api.IPoint
	oh2 api.IPoint

	v1  api.IPoint
	v2  api.IPoint
	ov1 api.IPoint
	ov2 api.IPoint
}

// NewCrossNode constructs a cross shaped node
func NewCrossNode(name string, parent api.INode) *CrossNode {
	o := new(CrossNode)
	o.Initialize(name)
	o.SetParent(parent)
	return o
}

// Build configures the node
func (c *CrossNode) Build(world api.IWorld) {
	c.Node.Build(world)

	c.h1 = geometry.NewPointUsing(-0.5, 0.0)
	c.h2 = geometry.NewPointUsing(0.5, 0.0)
	c.oh1 = geometry.NewPoint()
	c.oh2 = geometry.NewPoint()

	c.v1 = geometry.NewPointUsing(0.0, -0.5)
	c.v2 = geometry.NewPointUsing(0.0, 0.5)
	c.ov1 = geometry.NewPoint()
	c.ov2 = geometry.NewPoint()

	c.color = rendering.NewPaletteInt64(rendering.White)
}

// SetColor sets line color
func (c *CrossNode) SetColor(color api.IPalette) {
	c.color = color
}

// Draw renders shape
func (c *CrossNode) Draw(context api.IRenderContext) {
	if c.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformPoint(c.h1, c.oh1)
		context.TransformPoint(c.h2, c.oh2)

		context.TransformPoint(c.v1, c.ov1)
		context.TransformPoint(c.v2, c.ov2)

		c.SetDirty(false) // Node is no longer dirty
	}

	context.SetDrawColor(c.color)
	context.DrawLineUsing(c.ov1, c.ov2)
	context.DrawLineUsing(c.oh1, c.oh2)
}
