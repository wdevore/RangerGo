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

	zoomStepSize float64

	// State management
	mx, my    int32
	zoomPoint api.IPoint
}

// NewZoomNode constructs a zooming node
func NewZoomNode(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(ZoomNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (z *ZoomNode) Build(world api.IWorld) {
	z.Node.Build(world)

	z.zoomStepSize = 0.1

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

// SetStepSize sets the sensitivity of the zoom. If the view area
// is very tight then you want smaller values so that zooming
// doesn't jump by "glides"
func (z *ZoomNode) SetStepSize(size float64) {
	z.zoomStepSize = size
}

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

// ScaleTo sets the scale absolutely
func (z *ZoomNode) ScaleTo(s float64) {
	z.zoom.SetScale(s)
	z.RippleDirty(true)
}

// ZoomScale returns the zoom's current scale value
func (z *ZoomNode) ZoomScale() float64 {
	return z.zoom.PsuedoScale()
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
	z.zoom.ZoomBy(z.zoomStepSize, z.zoomStepSize)
	z.RippleDirty(true)
}

// ZoomOut zooms outward making things smaller
func (z *ZoomNode) ZoomOut() {
	z.zoom.ZoomBy(-z.zoomStepSize, -z.zoomStepSize)
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
