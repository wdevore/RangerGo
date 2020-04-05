package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// PolygonNode is a basic polygon type node. Default is
// no points present.
type PolygonNode struct {
	nodes.Node
	mx, my int32

	color       api.IPalette
	insideColor api.IPalette

	polygon api.IPolygon

	localPosition api.IPoint
	pointInside   bool
}

// NewPolygonNode constructs a node
func NewPolygonNode(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(PolygonNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (r *PolygonNode) Build(world api.IWorld) {
	r.Node.Build(world)

	r.polygon = geometry.NewPolygon()

	r.localPosition = geometry.NewPoint()

	r.color = rendering.NewPaletteInt64(rendering.White)
	r.insideColor = rendering.NewPaletteInt64(rendering.Red)
}

// SetColor sets rectangle color
func (r *PolygonNode) SetColor(color api.IPalette) {
	r.color = color
}

// AddVertex add a vertex.
// Set "complete" to "true" on the last vertex added.
func (r *PolygonNode) AddVertex(x, y float64, complete bool) {
	r.polygon.AddVertex(x, y)

	if complete {
		r.polygon.Build()
	}
}

// --------------------------------------------------------
// Lifecycle
// --------------------------------------------------------

// EnterNode called when a node is entering the stage
func (r *PolygonNode) EnterNode(man api.INodeManager) {
	man.RegisterEventTarget(r)
}

// ExitNode called when a node is exiting stage
func (r *PolygonNode) ExitNode(man api.INodeManager) {
	man.UnRegisterEventTarget(r)
}

// Draw renders shape
func (r *PolygonNode) Draw(context api.IRenderContext) {
	if r.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformPolygon(r.polygon)
		r.SetDirty(false) // Node is no longer dirty
	}

	// This get the local-space coords of the rectangle node.
	nodes.MapDeviceToNode(r.mx, r.my, r, r.localPosition)
	r.pointInside = r.polygon.PointInside(r.localPosition)

	if r.pointInside {
		context.SetDrawColor(r.insideColor)
	} else {
		context.SetDrawColor(r.color)
	}
	context.RenderPolygon(r.polygon, api.CLOSED)
}

// PointInside returns status
func (r *PolygonNode) PointInside() bool {
	return r.pointInside
}

// -----------------------------------------------------
// IO events
// -----------------------------------------------------

// Handle events from IO
func (r *PolygonNode) Handle(event api.IEvent) bool {
	// fmt.Println(event)
	if event.GetType() == api.IOTypeMouseMotion {
		r.mx, r.my = event.GetMousePosition()
	}

	return false
}
