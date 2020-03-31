package custom

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// VectorTextNode is a node that render text using vector fonts
type VectorTextNode struct {
	nodes.Node

	world api.IWorld

	text      string
	textColor api.IPalette

	mesh *geometry.Mesh
}

// NewVectorTextNode constructs a text node
func NewVectorTextNode(world api.IWorld, parent api.INode) *VectorTextNode {
	o := new(VectorTextNode)
	o.Build(world)
	parent.AddChild(o)

	o.textColor = rendering.NewPaletteInt64(rendering.White)
	o.mesh = geometry.NewMesh()
	o.world = world
	return o
}

// SetText sets the text of node
func (v *VectorTextNode) SetText(text string) {
	v.text = text
	v.SetDirty(true)
}

// ReBuild reconstructs the internal mesh based on text
func (v *VectorTextNode) ReBuild() {
	// Use glyph properties to adjust char location.
	xpos := 0.0

	font := v.world.VectorFont()
	for _, c := range v.text {
		vertices := font.Glyph(byte(c))

		for i := 0; i < len(vertices); i += 2 {
			v.mesh.AddVertex(vertices[i]+xpos, vertices[i+1])
		}

		xpos += font.HorizontalOffset()
	}

	v.mesh.Build()
}

// Draw renders shape
func (v *VectorTextNode) Draw(context api.IRenderContext) {
	if v.IsDirty() {
		v.ReBuild()
		context.TransformMesh(v.mesh)
		v.SetDirty(false) // Node is no longer dirty
	}

	context.SetDrawColor(v.textColor)
	context.RenderLines(v.mesh)
}

func (v VectorTextNode) String() string {
	return fmt.Sprintf("%s = '%s'", v.Node, v.text)
}
