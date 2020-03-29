package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/animation"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
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

	// Motion is for rotating triangle
	angularMotion api.IMotion

	triangle *TriangleNode
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

	g.textColor = rendering.NewPaletteInt64(rendering.Red)

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

	g.triangle = NewTriangleNode("Triangle", g)
	g.triangle.Build(world)
	g.triangle.SetScale(100.0)
	g.triangle.SetPosition(-100.0, -100.0)
	g.triangle.SetRotation(maths.DegreeToRadians * 20.0)
	g.AddChild(g.triangle)

	g.angularMotion = animation.NewAngularMotion()
	// amgle is measured in angular-velocity or "degrees/second"
	g.angularMotion.SetRate(maths.DegreeToRadians * 45.0)

}

// Update updates the time properties of a node.
func (g *gameLayer) Update(dt float64) {
	g.angularMotion.Update(dt)
}

// Interpolate is used for blending time based properties.
func (g *gameLayer) Interpolate(interpolation float64) {
	value := g.angularMotion.Interpolate(interpolation)
	g.triangle.SetRotation(value.(float64))
}

// -----------------------------------------------------
// Node lifecycles
// -----------------------------------------------------

// EnterNode called when a node is entering the stage
func (g *gameLayer) EnterNode(man api.INodeManager) {
	man.RegisterTarget(g)
}

// ExitNode called when a node is exiting stage
func (g *gameLayer) ExitNode(man api.INodeManager) {
	man.UnRegisterTarget(g)
}
