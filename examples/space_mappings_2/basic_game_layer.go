package main

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type gameLayer struct {
	nodes.Node
	world api.IWorld

	textColor      api.IPalette
	cursorPosition api.IPoint

	o1 api.IPoint
	o2 api.IPoint

	crossNode *custom.CrossNode

	rectNode *custom.RectangleNode
}

func newBasicGameLayer(name string) api.INode {
	o := new(gameLayer)
	o.Initialize(name)
	return o
}

func (g *gameLayer) Build(world api.IWorld) {
	g.world = world

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

	g.rectNode = custom.NewRectangleNode("Orange Rect")
	g.rectNode.Build(world)
	g.rectNode.SetColor(rendering.NewPaletteInt64(rendering.Orange))
	g.rectNode.SetScale(100.0)
	g.rectNode.SetRotation(maths.DegreeToRadians * 35.0)
	g.rectNode.SetPosition(100.0, -150.0)
	g.AddChild(g.rectNode)

	g.crossNode = custom.NewCrossNode("Cross", g)
	g.crossNode.Build(world)
	g.crossNode.SetScale(30.0)
	g.AddChild(g.crossNode)
}

// -----------------------------------------------------
// Node lifecycles
// -----------------------------------------------------

// EnterNode called when a node is entering the stage
func (g *gameLayer) EnterNode(man api.INodeManager) {
	man.RegisterEventTarget(g)
}

// ExitNode called when a node is exiting stage
func (g *gameLayer) ExitNode(man api.INodeManager) {
	man.UnRegisterEventTarget(g)
}

// -----------------------------------------------------
// Visuals
// -----------------------------------------------------

func (g *gameLayer) Draw(context api.IRenderContext) {
	// Transform vertices if anything has changed.
	if g.IsDirty() {
		// Transform this node's vertices using the context
		g.SetDirty(false) // Node is no longer dirty
	}

	context.SetDrawColor(g.textColor)
	text := fmt.Sprintf("(%d, %d)", int(g.cursorPosition.X()), int(g.cursorPosition.Y()))
	context.DrawText(10.0, 10.0, text, 1, 1, false)

	g.Node.Draw(context)
}

// -----------------------------------------------------
// IO events
// -----------------------------------------------------

func (g *gameLayer) Handle(event api.IEvent) bool {
	if event.GetType() == api.IOTypeMouseMotion {
		mx, my := event.GetMousePosition()
		nodes.MapDeviceToView(g.world, mx, my, g.cursorPosition)

		g.crossNode.SetPosition(g.cursorPosition.X(), g.cursorPosition.Y())
	}

	return false
}
