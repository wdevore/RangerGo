package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/nodes"
)

// ZoomNode provides zooming at parent nodes
type ZoomNode struct {
	nodes.Node
	zoom api.IZoomTransform

	// State management
	mx, my    int32
	zoomPoint api.IPoint
	// wheelDirection int // 0 = not active, 1 = zoom-in, -1 = zoom-out

}

// NewZoomNode constructs a zooming node
func NewZoomNode(name string, parent api.INode) api.INode {
	o := new(ZoomNode)
	o.Initialize(name)
	o.SetParent(parent)
	return o
}

// Build configures the node
func (z *ZoomNode) Build(world api.IWorld) {
	z.Node.Build(world)

	z.zoom = maths.NewZoomTransform()
	z.zoomPoint = geometry.NewPoint()
}

// EnterNode called when a node is entering the stage
func (z *ZoomNode) EnterNode(man api.INodeManager) {
	// We want the mouse events so the node can track the mouse.
	man.RegisterEventTarget(z)
}

// ExitNode called when a node is exiting stage
func (z *ZoomNode) ExitNode(man api.INodeManager) {
	man.UnRegisterEventTarget(z)
}

// --------------------------------------------------------
// Zooming
// --------------------------------------------------------

// SetPosition sets the zooms position and ripples to children
// mouse is located.
func (z *ZoomNode) SetPosition(x, y float64) {
	z.zoom.SetPosition(x, y)
	z.RippleDirty(true)
}

// SetFocalPoint sets the epi center of zoom
func (z *ZoomNode) SetFocalPoint(x, y float64) {
	z.zoom.SetAt(x, y)
	z.RippleDirty(true)
}

// ZoomBy is relative zooming using deltas
func (z *ZoomNode) ZoomBy(dx, dy float64) {
	z.zoom.ZoomBy(dx, dy)
	z.RippleDirty(true)
}

// TranslateBy is relative translation
func (z *ZoomNode) TranslateBy(dx, dy float64) {
	z.zoom.TranslateBy(dx, dy)
	z.RippleDirty(true)
}

// ZoomIn zooms inward making things bigger
func (z *ZoomNode) ZoomIn() {
	z.zoom.ZoomBy(0.1, 0.1)
	z.RippleDirty(true)
}

// ZoomOut zooms outward making things smaller
func (z *ZoomNode) ZoomOut() {
	z.zoom.ZoomBy(-0.1, -0.1)
	z.RippleDirty(true)
}

// --------------------------------------------------------
// Transforms
// --------------------------------------------------------

// CalcTransform : zoom nodes manage their own transform differently.
func (z *ZoomNode) CalcTransform() api.IAffineTransform {
	if z.IsDirty() {
		z.zoom.Update()
		z.SetDirty(false)
	}

	return z.zoom.GetTransform()
}

// -----------------------------------------------------
// IO events
// -----------------------------------------------------

// Handle processes IO events
func (z *ZoomNode) Handle(event api.IEvent) bool {
	if event.GetType() == api.IOTypeMouseMotion {
		z.mx, z.my = event.GetMousePosition()
		nodes.MapDeviceToNode(z.mx, z.my, z, z.zoomPoint)
		z.SetFocalPoint(z.zoomPoint.X(), z.zoomPoint.Y())
	} else if event.GetType() == api.IOTypeMouseWheel {
		_, dir := event.GetMouseRelMovement()
		if dir == 1 {
			z.ZoomIn()
		} else {
			z.ZoomOut()
		}
	}

	return false
}
