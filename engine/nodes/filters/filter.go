package filters

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
)

// #############################################################################
// Warning: Filters can interfer with things like dragging which
// rely on parent transform properties.
// If you need hiearchial dragging then skip using Filter and just
// manually manage the child transform properties relative to their parent.
// #############################################################################

// Filter is the base property of Filter nodes
type Filter struct {
	// The node's immediate parent translation components
	components api.IAffineTransform

	// What to exclude from the parent
	excludeTranslation bool
	excludeRotation    bool
	excludeScale       bool
}

func (t *Filter) initializeFilter() {
	t.components = maths.NewTransform()
}

// InheritAll causes the filter to pass all of the parent's transform
// properties: Translate, Rotation and Scale.
func (t *Filter) InheritAll() {
	t.excludeTranslation = false
	t.excludeRotation = false
	t.excludeScale = false
}

// InheritOnlyRotation causes the filter to pass only the parent's rotational
// property.
func (t *Filter) InheritOnlyRotation() {
	t.excludeTranslation = true
	t.excludeRotation = false
	t.excludeScale = true
}

// InheritOnlyScale causes the filter to pass only the parent's scale
// property.
func (t *Filter) InheritOnlyScale() {
	t.excludeTranslation = true
	t.excludeRotation = true
	t.excludeScale = false
}

// InheritOnlyTranslation causes the filter to pass only the parent's translation
// property.
func (t *Filter) InheritOnlyTranslation() {
	t.excludeTranslation = false
	t.excludeRotation = true
	t.excludeScale = true
}

// InheritRotationAndTranslation causes the filter to pass the parent's translation
// and rotational properties.
func (t *Filter) InheritRotationAndTranslation() {
	t.excludeTranslation = false
	t.excludeRotation = false
	t.excludeScale = true
}
