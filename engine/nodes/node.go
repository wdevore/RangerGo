package nodes

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
)

type node struct {
	id      int
	name    string
	visible bool
	dirty   bool

	parent api.INode
}

// Called by base objects from their Initialize
func (n *node) Initialize(id int, name string) {
	n.id = id
	n.name = name
	n.visible = true
	n.dirty = true
}

func (n *node) Visit(context api.IRenderContext, interpolation float64) {
	if !n.IsVisible() {
		return
	}

	context.Save()

	// Because position and angles are dependent
	// on lerping we perform interpolation first.
	n.Interpolate(interpolation)

	aft := n.CalcTransform()

	context.Apply(aft)

	children := n.Children()

	if children != nil {
		n.Draw(context)

		for _, child := range children {
			child.Visit(context, interpolation)
		}
	} else {
		n.Draw(context)
	}

	context.Restore()
}

func (n *node) IsVisible() bool {
	return n.visible
}

func (n *node) Interpolate(interpolation float64) {
	fmt.Println("Node Interpolate")
}

func (n *node) IsDirty() bool {
	return n.dirty
}

func (n *node) Draw(context api.IRenderContext) {
}

// -----------------------------------------------------
// ITransform defaults
// -----------------------------------------------------

func (n *node) Transform() api.ITransform {
	return nil
}

func (n *node) CalcTransform() api.IAffineTransform {
	tr := n.Transform()

	aft := tr.AffineTransform()

	if n.IsDirty() {
		pos := tr.Position()
		aft.MakeTranslate(pos.X(), pos.Y())

		rot := tr.Rotation()
		if rot != 0.0 {
			aft.Rotate(rot)
		}

		s := tr.Scale()
		if s.X() != 1.0 || s.Y() != 1.0 {
			aft.Scale(s.X(), s.Y())
		}

		// Invert...
		inv := tr.InverseTransform()
		aft.InvertTo(inv)
	}

	return aft
}

var p = geometry.NewPoint()

func (n *node) Position() api.IPoint {
	// User should override
	return p
}

func (n *node) Rotation() float64 {
	// User should override
	return 0.0
}

// -------------------------------------------------------------------
// INodeGroup implementations
// -------------------------------------------------------------------

func (n *node) Children() []api.INode {
	return nil
}

// -------------------------------------------------------------------
// Misc
// -------------------------------------------------------------------

func (n node) String() string {
	return fmt.Sprintf("|'%s' (%d)|", n.name, n.id)
}
