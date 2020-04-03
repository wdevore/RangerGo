package filters

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/nodes"
)

// TransformFilter by default passes Rotation and Translation components,
// which is the most common use case.
type TransformFilter struct {
	nodes.Node
	world api.IWorld

	// This makes the node an IFilter type
	Filter
}

// NewTransformFilter constructs a default transform filter. Default
// is to inherit both Rotation and Translation.
func NewTransformFilter(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(TransformFilter)
	o.Initialize(name)
	o.SetParent(parent)
	o.initializeFilter()
	o.InheritRotationAndTranslation()
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (t *TransformFilter) Build(world api.IWorld) {
	t.Node.Build(world)
}

// Visit is special in that it they provide their own implementation
func (t *TransformFilter) Visit(context api.IRenderContext, interpolation float64) {
	if !t.IsVisible() {
		return
	}

	context.Save()

	children := t.Children()

	for _, child := range children {
		context.Save()

		if t.HasParent() {
			parent := t.Parent()

			// Re-introduce only the parent's components as defined by exclusion flags.
			parent.CalcFilteredTransform(t.excludeTranslation,
				t.excludeRotation,
				t.excludeScale, t.components)

			// Combine using pre-multiply
			// "parent.InverseTransform" removes the immediate parent's transform effects
			maths.Multiply(t.components, parent.InverseTransform(), t.AffineTransform())

			// Merge them with the current context.
			context.Apply(t.AffineTransform())
		} else {
			fmt.Println("TransformFilter: node ", t, " has NO parent")
			return
		}

		// Now visit the child with the modified context
		nodes.Visit(child, context, interpolation)

		context.Restore()
	}

	context.Restore()
}
