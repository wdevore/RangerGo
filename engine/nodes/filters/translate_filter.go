package filters

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
)

// TranslateFilter will exclude or "block" rotations and scales from
// propagating to the children (default), but "passes" translations.
type TranslateFilter struct {
	nodes.Node
	world api.IWorld

	// This makes the node an IFilter type
	Filter
}

// NewTranslateFilter constructs a Translate filter.
func NewTranslateFilter(name string, parent api.INode) api.INode {
	o := new(TranslateFilter)
	o.Initialize(name)
	o.SetParent(parent)
	o.initializeFilter()
	return o
}

// Build configures the node
func (t *TranslateFilter) Build(world api.IWorld) {
	t.Node.Build(world)
}

// VisitFilter is special in that it they provide their own implementation.
// Because this is a Translate filter we "filter out" everything
// but the translation component from the immediate parent.
func (t *TranslateFilter) VisitFilter(context api.IRenderContext, interpolation float64) {
	if !t.IsVisible() {
		return
	}

	context.Save()

	children := t.Children()

	for _, child := range children {
		context.Save()

		if t.HasParent() {
			parent := t.Parent()

			// This removes the immediate parent's transform effects
			context.Apply(parent.InverseTransform())

			// Re-introduce only the parent's translation component by
			// excluding Rotation and Scale
			parent.CalcFilteredTransform(false, true, true, t.components)

			// And update context to reflect the exclusion.
			context.Apply(t.components)
		} else {
			fmt.Println("TranslateFilter: node ", t, " has NO parent")
			return
		}

		// Now visit the child with the modified context
		nodes.Visit(child, context, interpolation)

		context.Restore()
	}

	context.Restore()
}
