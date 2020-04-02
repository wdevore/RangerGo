package main

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type gameLayer struct {
	nodes.Node

	backgroundColor api.IPalette
	backgroundMin   api.IPoint
	backgroundMax   api.IPoint

	textColor api.IPalette

	o1 api.IPoint
	o2 api.IPoint

	viewPoint api.IPoint
}

func newBasicGameLayer(name string) api.INode {
	o := new(gameLayer)
	o.Initialize(name)
	return o
}

func (g *gameLayer) Build(world api.IWorld) {
	g.Node.Build(world)

	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	g.backgroundMin = geometry.NewPointUsing(x, y)
	g.backgroundMax = geometry.NewPointUsing(x+vw, y+vh)

	g.backgroundColor = rendering.NewPaletteInt64(rendering.DarkGray)
	g.textColor = rendering.NewPaletteInt64(rendering.White)

	g.o1 = geometry.NewPoint()
	g.o2 = geometry.NewPoint()
	g.viewPoint = geometry.NewPoint()

	hLine := custom.NewLineNode("HLine", g)
	hLine.Build(world)
	n := hLine.(*custom.LineNode)
	n.SetPoints(x, 0.0, -x, 0.0)

	vLine := custom.NewLineNode("VLine", g)
	vLine.Build(world)
	n = vLine.(*custom.LineNode)
	n.SetPoints(0.0, -y, 0.0, y)
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
		context.TransformPoint(g.backgroundMin, g.o1)
		context.TransformPoint(g.backgroundMax, g.o2)
		g.SetDirty(false) // Node is no longer dirty
	}

	// Draw background first.
	context.SetDrawColor(g.backgroundColor)
	context.RenderAARectangle(g.o1, g.o2, api.FILLED)

	context.SetDrawColor(g.textColor)
	text := fmt.Sprintf("(%d, %d)", int(g.viewPoint.X()), int(g.viewPoint.Y()))
	context.DrawText(10.0, 10.0, text, 1, 1, false)

	g.Node.Draw(context)
}

// -----------------------------------------------------
// IO events
// -----------------------------------------------------

func (g *gameLayer) Handle(event api.IEvent) bool {
	if event.GetType() == api.IOTypeMouseMotion {
		mx, my := event.GetMousePosition()
		nodes.MapDeviceToView(g.World(), mx, my, g.viewPoint)
	}

	return false
}
