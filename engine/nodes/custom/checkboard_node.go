package custom

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// CheckerBoardNode is a simple node typically used for backgrounds.
type CheckerBoardNode struct {
	nodes.Node

	oddColor  api.IPalette
	evenColor api.IPalette
}

// NewCheckBoardNode constructs an axis aligned checker board
func NewCheckBoardNode(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(CheckerBoardNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (c *CheckerBoardNode) Build(world api.IWorld) {
	c.Node.Build(world)

	vw, vh := world.ViewSize().Components()
	y := -vh / 2.0
	w := 100.0

	c.oddColor = rendering.NewPaletteInt64(rendering.DarkGray)
	c.evenColor = rendering.NewPaletteInt64(rendering.LightGray)

	// Construct grid of rectangles

	// "pFlip" tracks the current row's initial boolean value so that
	// on the next row we start with the opposite value. This
	// insures that each row is inverted from the previous thus always
	// showing a checkboard pattern.
	flip := false
	pFlip := flip

	for y <= vh {
		pFlip = flip
		x := -vw / 2.0
		for x <= vw {
			r := NewBasicRectangleNode(fmt.Sprintf("::(%d,%d)", int(x), int(y)), c)
			r.Build(world)
			r.SetPosition(x, y)
			r.SetScale(w)
			cr := r.(*BasicRectangleNode)
			if flip {
				cr.SetColor(c.oddColor)
			} else {
				cr.SetColor(c.evenColor)
			}
			flip = !flip
			x += w
		}
		y += w
		flip = !pFlip
	}
}
