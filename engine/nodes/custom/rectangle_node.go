package custom

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// RectangleNode is a basic line node
type RectangleNode struct {
	nodes.Node
	world api.IWorld

	color       api.IPalette
	insideColor api.IPalette
	textColor   api.IPalette

	polygon api.IPolygon

	localPosition api.IPoint
	pointInside   bool
}

// NewRectangleNode constructs a rectangle shaped node
func NewRectangleNode(name string) *RectangleNode {
	o := new(RectangleNode)
	o.Initialize(name)
	return o
}

// Build configures the node
func (r *RectangleNode) Build(world api.IWorld) {
	r.world = world

	r.polygon = geometry.NewPolygon()
	r.polygon.AddVertex(-0.5, -0.5)
	r.polygon.AddVertex(-0.5, 0.5)
	r.polygon.AddVertex(0.5, 0.5)
	r.polygon.AddVertex(0.5, -0.5)

	r.polygon.Build()

	r.localPosition = geometry.NewPoint()

	r.color = rendering.NewPaletteInt64(rendering.White)
	r.insideColor = rendering.NewPaletteInt64(rendering.Red)
	r.textColor = rendering.NewPaletteInt64(rendering.White)
}

// SetColor sets rectangle color
func (r *RectangleNode) SetColor(color api.IPalette) {
	r.color = color
}

// SetBounds sets the min,max of rectangle
func (r *RectangleNode) SetBounds(minx, miny, maxx, maxy float64) {
}

// --------------------------------------------------------
// Timing
// --------------------------------------------------------

// Update is for timing
func (r *RectangleNode) Update(dt float64) {
	// This node rotates itself.
}

// EnterNode called when a node is entering the stage
func (r *RectangleNode) EnterNode(man api.INodeManager) {
	// man.RegisterTarget(r)
	man.RegisterEventTarget(r)
}

// ExitNode called when a node is exiting stage
func (r *RectangleNode) ExitNode(man api.INodeManager) {
	// man.UnRegisterTarget(r)
	man.UnRegisterEventTarget(r)
}

// Draw renders shape
func (r *RectangleNode) Draw(context api.IRenderContext) {
	if r.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformPolygon(r.polygon)
		r.SetDirty(false) // Node is no longer dirty
	}

	if r.pointInside {
		context.SetDrawColor(r.insideColor)
		// Show AABB indictor rectangle
	} else {
		context.SetDrawColor(r.color)
	}
	context.RenderPolygon(r.polygon, api.CLOSED)

	context.SetDrawColor(r.textColor)
	text := fmt.Sprintf("(%2.3f, %2.3f)", r.localPosition.X(), r.localPosition.Y())
	context.DrawText(10.0, 20.0, text, 1, 1, false)

}

// -----------------------------------------------------
// IO events
// -----------------------------------------------------

// Handle events from IO
func (r *RectangleNode) Handle(event api.IEvent) bool {
	if event.GetType() == api.IOTypeMouseMotion {
		mx, my := event.GetMousePosition()

		// This get the local-space coords of the rectangle node.
		nodes.MapDeviceToNode(r.world, mx, my, r, r.localPosition)
		r.pointInside = r.polygon.PointInside(r.localPosition)
	}

	return false
}
