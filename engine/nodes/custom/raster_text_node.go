package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// RasterTextNode is a simple pixel node that only responds to translations.
// Scale and Rotations are not implemented yet.
// Note: The node only draws in Device-space so the text could easly disappear
// if you are not paying attention.
type RasterTextNode struct {
	nodes.Node

	color api.IPalette
	text  string
	scale int
	fill  int

	useDeviceSpace bool

	p1 api.IPoint
	o1 api.IPoint
}

// NewRasterTextNode constructs a cross shaped node
func NewRasterTextNode(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(RasterTextNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (r *RasterTextNode) Build(world api.IWorld) {
	r.Node.Build(world)

	r.p1 = geometry.NewPoint()
	r.o1 = geometry.NewPoint()

	r.scale = 1
	r.fill = 1
	r.text = "?"
	r.useDeviceSpace = true

	r.color = rendering.NewPaletteInt64(rendering.LightGray)
}

// SetText set the text value
func (r *RasterTextNode) SetText(text string) {
	r.text = text
}

// SetColor sets line color
func (r *RasterTextNode) SetColor(color api.IPalette) {
	r.color = color
}

// SetFontScale sets the scale factor of the font not the Node.
func (r *RasterTextNode) SetFontScale(scale int) {
	r.scale = scale
}

// SetFill sets the fill factor of the font.
func (r *RasterTextNode) SetFill(fill int) {
	r.fill = fill
}

// Draw renders shape
func (r *RasterTextNode) Draw(context api.IRenderContext) {
	if r.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformPoint(r.p1, r.o1)
		r.SetDirty(false) // Node is no longer dirty
	}

	context.SetDrawColor(r.color)

	if r.useDeviceSpace {
		context.DrawText(r.Position().X(), r.Position().Y(), r.text, r.scale, r.fill, false)
	} else {
		context.DrawText(r.o1.X(), r.o1.Y(), r.text, r.scale, r.fill, false)
	}
}
