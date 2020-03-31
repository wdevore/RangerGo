package main

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/animation"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/misc"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/nodes/filters"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type gameLayer struct {
	nodes.Node

	textColor      api.IPalette
	cursorPosition api.IPoint

	o1 api.IPoint
	o2 api.IPoint

	crossNode api.INode

	orangeRectNode api.INode
	greenRectNode  api.INode

	// Motion is for rotating cube
	angularMotion api.IMotion

	// Dragging
	drag   api.IDragging
	mx, my int32

	// Zooming
	zoom api.INode
}

func newBasicGameLayer(name string, parent api.INode) api.INode {
	o := new(gameLayer)
	o.Initialize(name)
	o.SetParent(parent)
	return o
}

func (g *gameLayer) Build(world api.IWorld) {
	g.Node.Build(world)

	g.drag = misc.NewDragState()

	g.zoom = custom.NewZoomNode("ZoomNode", g)
	g.zoom.Build(world)

	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	g.textColor = rendering.NewPaletteInt64(rendering.White)

	g.o1 = geometry.NewPoint()
	g.o2 = geometry.NewPoint()
	g.cursorPosition = geometry.NewPoint()

	hLine := custom.NewLineNode("HLine", g)
	hLine.Build(world)
	n := hLine.(*custom.LineNode)
	n.SetPoints(x, 0.0, -x, 0.0)

	vLine := custom.NewLineNode("VLine", g)
	vLine.Build(world)
	n = vLine.(*custom.LineNode)
	n.SetPoints(0.0, -y, 0.0, y)

	// Rectangle scale-space
	rectScale := 100.0

	g.orangeRectNode = custom.NewRectangleNode("Orange Rect", g.zoom)
	g.orangeRectNode.Build(world)
	gor := g.orangeRectNode.(*custom.RectangleNode)
	gor.SetColor(rendering.NewPaletteInt64(rendering.Orange))
	g.orangeRectNode.SetScale(rectScale)
	// g.rectNode.SetRotation(maths.DegreeToRadians * 35.0)
	g.orangeRectNode.SetPosition(100.0, -150.0)

	// Add Filter to remove parent's (aka Orange rectangle) Scale
	filter := filters.NewTransformFilter("TransformFilter", g.orangeRectNode)
	filter.Build(world)

	g.greenRectNode = custom.NewRectangleNode("Green Rect", filter)
	g.greenRectNode.Build(world)
	grr := g.greenRectNode.(*custom.RectangleNode)
	grr.SetColor(rendering.NewPaletteInt64(rendering.Green))
	g.greenRectNode.SetScale(10.0)
	g.greenRectNode.SetPosition(100.0, 0.0)

	g.angularMotion = animation.NewAngularMotion()
	// amgle is measured in angular-velocity or "degrees/second"
	g.angularMotion.SetRate(maths.DegreeToRadians * 90.0)

	g.crossNode = custom.NewCrossNode("Cross", g)
	g.crossNode.Build(world)
	g.crossNode.SetScale(30.0)

	g.AddChild(g.zoom)
}

// Update updates the time properties of a node.
func (g *gameLayer) Update(dt float64) {
	g.angularMotion.Update(dt)
}

// Interpolate is used for blending time based properties.
func (g *gameLayer) Interpolate(interpolation float64) {
	value := g.angularMotion.Interpolate(interpolation)
	g.orangeRectNode.SetRotation(value.(float64))
}

// -----------------------------------------------------
// Node lifecycles
// -----------------------------------------------------

// EnterNode called when a node is entering the stage
func (g *gameLayer) EnterNode(man api.INodeManager) {
	man.RegisterTarget(g)
	// We want the mouse events so the node can track the mouse.
	man.RegisterEventTarget(g)
}

// ExitNode called when a node is exiting stage
func (g *gameLayer) ExitNode(man api.INodeManager) {
	man.UnRegisterTarget(g)
	man.UnRegisterEventTarget(g)
}

// -----------------------------------------------------
// Visuals
// -----------------------------------------------------

func (g *gameLayer) Draw(context api.IRenderContext) {
	context.SetDrawColor(g.textColor)
	text := fmt.Sprintf("(%d, %d)", int(g.cursorPosition.X()), int(g.cursorPosition.Y()))
	context.DrawText(10.0, 10.0, text, 1, 1, false)
}

// -----------------------------------------------------
// IO events
// -----------------------------------------------------

func (g *gameLayer) Handle(event api.IEvent) bool {
	if event.GetType() == api.IOTypeMouseMotion {
		mx, my := event.GetMousePosition()
		nodes.MapDeviceToView(g.World(), mx, my, g.cursorPosition)

		g.crossNode.SetPosition(g.cursorPosition.X(), g.cursorPosition.Y())

		// Because the Layer and parent Scene have no transformation between
		// each other we could also pass "g" instead of "g.rectNode".
		// Passing "g" would cause SetMotion...() to use g's parent which
		// is SplashScene verses rectangle node's parent which is GameLayer.
		// However, to be explicit I pass "g.rectNode"
		g.drag.SetMotionStateUsing(mx, my, event.GetState(), g.orangeRectNode)

		gor := g.orangeRectNode.(*custom.RectangleNode)
		if g.drag.IsDragging() && gor.PointInside() {
			pos := g.orangeRectNode.Position()
			g.orangeRectNode.SetPosition(pos.X()+g.drag.Delta().X(), pos.Y()+g.drag.Delta().Y())
		}

	} else if event.GetType() == api.IOTypeMouseButtonDown || event.GetType() == api.IOTypeMouseButtonUp {
		mx, my := event.GetMousePosition()
		// On mouse events if state = 1 then dragging
		g.drag.SetButtonStateUsing(mx, my, event.GetButton(), event.GetState(), g.orangeRectNode)
	}

	return false
}
