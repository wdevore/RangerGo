package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/misc"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// AABBNode can be used for several things: encompassing objects with
// visual rectangles for debugging or simply for visuals.
// AABB is usually used relative to a node's parent which means it
// used the transformed vertices not the local-space vertices.
type AABBNode struct {
	nodes.Node

	color api.IPalette

	aabb *misc.AABB

	o1 api.IPoint
	o2 api.IPoint
}

// NewAABBNode constructs a axis aligned bounding box node
func NewAABBNode(name string, parent api.INode) *AABBNode {
	o := new(AABBNode)
	o.Initialize(name)
	o.SetParent(parent)
	return o
}

// Build configures the node
func (a *AABBNode) Build(world api.IWorld) {
	a.Node.Build(world)

	a.aabb = misc.NewAABB()

	a.o1 = geometry.NewPoint()
	a.o2 = geometry.NewPoint()

	a.color = rendering.NewPaletteInt64(rendering.Red)
}

// SetColor sets line color
func (a *AABBNode) SetColor(color api.IPalette) {
	a.color = color
}

// SetBounds sets the bounds based on the provided Mesh
func (a *AABBNode) SetBounds(mesh api.IMesh) {
	a.aabb.SetBounds(mesh.Vertices())
}

// Draw renders shape
func (a *AABBNode) Draw(context api.IRenderContext) {
	if a.IsDirty() {
		context.TransformPoint(a.aabb.Min(), a.o1)
		context.TransformPoint(a.aabb.Max(), a.o2)
		a.SetDirty(false)
	}

	context.SetDrawColor(a.color)
	context.RenderAARectangle(a.o1, a.o2, api.OUTLINED)
}
