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

	hLine := custom.NewLineNode("HLine", world, g)
	n := hLine.(*custom.LineNode)
	n.SetPoints(x, 0.0, -x, 0.0)

	vLine := custom.NewLineNode("VLine", world, g)
	n = vLine.(*custom.LineNode)
	n.SetPoints(0.0, -y, 0.0, y)

	triangle := custom.NewTriangleNode("Triangle", world, g)
	triangle.SetScale(100.0)
	triangle.SetPosition(-400.0, -200.0)

	cross := custom.NewCrossNode("Cross", world, g)
	cross.Build(world)
	cross.SetScale(100.0)
	cross.SetPosition(-200.0, -200.0)

	circle := custom.NewCircleNode("Circle", world, g)
	circle.SetScale(50.0)
	circle.SetPosition(0.0, -200.0)
}
