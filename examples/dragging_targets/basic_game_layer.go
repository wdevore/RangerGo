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
	"github.com/wdevore/RangerGo/engine/rendering"
)

type gameLayer struct {
	nodes.Node

	textColor      api.IPalette
	cursorPosition api.IPoint

	o1 api.IPoint
	o2 api.IPoint

	crossNode api.INode

	rectNode api.INode

	// Motion is for rotating cube
	angularMotion api.IMotion

	// Dragging
	drag   api.IDragging
	mx, my int32
}

func newBasicGameLayer(name string, parent api.INode) api.INode {
	o := new(gameLayer)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	return o
}

func (g *gameLayer) Build(world api.IWorld) {
	g.Node.Build(world)

	g.drag = misc.NewDragState()

	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	g.textColor = rendering.NewPaletteInt64(rendering.White)

	g.o1 = geometry.NewPoint()
	g.o2 = geometry.NewPoint()
	g.cursorPosition = geometry.NewPoint()

	hLine := custom.NewLineNode("HLine", world, g)
	n := hLine.(*custom.LineNode)
	n.SetPoints(x, 0.0, -x, 0.0)

	vLine := custom.NewLineNode("VLine", world, g)
	n = vLine.(*custom.LineNode)
	n.SetPoints(0.0, -y, 0.0, y)

	g.rectNode = custom.NewRectangleNode("Orange Rect", world, g)
	rn := g.rectNode.(*custom.RectangleNode)
	rn.SetColor(rendering.NewPaletteInt64(rendering.Orange))
	g.rectNode.SetScale(100.0)
	// g.rectNode.SetRotation(maths.DegreeToRadians * 35.0)
	g.rectNode.SetPosition(100.0, -150.0)

	g.angularMotion = animation.NewAngularMotion()
	// amgle is measured in angular-velocity or "degrees/second"
	g.angularMotion.SetRate(maths.DegreeToRadians * 90.0)

	g.crossNode = custom.NewCrossNode("Cross", world, g)
	g.crossNode.SetScale(30.0)
}

// Update updates the time properties of a node.
func (g *gameLayer) Update(msPerUpdate, secPerUpdate float64) {
	g.angularMotion.Update(msPerUpdate)
}

// Interpolate is used for blending time based properties.
func (g *gameLayer) Interpolate(interpolation float64) {
	value := g.angularMotion.Interpolate(interpolation)
	g.rectNode.SetRotation(value.(float64))
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
		g.drag.SetMotionStateUsing(mx, my, event.GetState(), g.rectNode)

		rn := g.rectNode.(*custom.RectangleNode)

		if g.drag.IsDragging() && rn.PointInside() {
			pos := g.rectNode.Position()
			g.rectNode.SetPosition(pos.X()+g.drag.Delta().X(), pos.Y()+g.drag.Delta().Y())
		}

	} else if event.GetType() == api.IOTypeMouseButtonDown || event.GetType() == api.IOTypeMouseButtonUp {
		mx, my := event.GetMousePosition()
		// On mouse events if state = 1 then dragging
		g.drag.SetButtonStateUsing(mx, my, event.GetButton(), event.GetState(), g.rectNode)
	}

	return false
}
