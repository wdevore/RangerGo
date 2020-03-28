package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type gameLayer struct {
	nodes.Node

	api.IRender

	textColor      api.IPalette
	cursorPosition api.IPoint

	o1 api.IPoint
	o2 api.IPoint
}

func newBasicGameLayer(name string, parent api.INode) api.INode {
	o := new(gameLayer)
	o.Initialize(name)
	o.SetParent(parent)
	return o
}

func (g *gameLayer) Build(world api.IWorld) {
	g.Node.Build(world)

	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	g.textColor = rendering.NewPaletteInt64(rendering.White)

	g.o1 = geometry.NewPoint()
	g.o2 = geometry.NewPoint()
	g.cursorPosition = geometry.NewPoint()

	hLine := custom.NewLineNode("HLine", g)
	hLine.Build(world)
	hLine.SetPoints(x, 0.0, -x, 0.0)
	g.AddChild(hLine)

	vLine := custom.NewLineNode("VLine", g)
	vLine.Build(world)
	vLine.SetPoints(0.0, -y, 0.0, y)
	g.AddChild(vLine)

	triangle := custom.NewTriangleNode("Triangle", g)
	triangle.Build(world)
	triangle.SetScale(100.0)
	triangle.SetPosition(-500.0, -300.0)
	g.AddChild(triangle)

	cross := custom.NewCrossNode("Cross", g)
	cross.Build(world)
	cross.SetScale(100.0)
	cross.SetPosition(-400.0, -300.0)
	g.AddChild(cross)

	circle := custom.NewCircleNode("Circle", g)
	circle.Build(world)
	circle.SetScale(50.0)
	circle.SetPosition(-250.0, -300.0)
	g.AddChild(circle)
}

// -----------------------------------------------------
// Visuals
// -----------------------------------------------------

func (g *gameLayer) Draw(context api.IRenderContext) {
}
