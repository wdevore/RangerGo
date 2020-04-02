package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/animation"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
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

	// Motion is for rotating triangle
	triangleMotion api.IMotion

	anchorMotion api.IMotion

	triangle      api.INode
	greenRectNode api.INode
	anchor        api.INode
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

	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	g.textColor = rendering.NewPaletteInt64(rendering.Red)

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

	g.triangle = NewTriangleNode("Triangle", g)
	g.triangle.Build(world)
	g.triangle.SetScale(100.0)
	g.triangle.SetPosition(-100.0, -100.0)
	g.triangle.SetRotation(maths.DegreeToRadians * 20.0)

	// Create Filter and set Triangle as parent
	filter := filters.NewTransformFilter("Filter", g.triangle)
	filter.Build(world)

	// Create an Anchor with the Filter as parent
	g.anchor = custom.NewAnchorNode("Anchor", filter)
	g.anchor.Build(world)

	// Create green rectangle with the Anchor as parent
	g.greenRectNode = custom.NewRectangleNode("Green Rect", g.anchor)
	g.greenRectNode.Build(world)
	grn := g.greenRectNode.(*custom.RectangleNode)
	grn.SetColor(rendering.NewPaletteInt64(rendering.Green))
	g.greenRectNode.SetScale(10.0)
	g.greenRectNode.SetPosition(100.0, 0.0)

	// amgle is measured in angular-velocity or "degrees/second"
	g.triangleMotion = animation.NewAngularMotion()
	g.triangleMotion.SetRate(maths.DegreeToRadians * 45.0)

	g.anchorMotion = animation.NewAngularMotion()
	g.anchorMotion.SetRate(maths.DegreeToRadians * (-90.0 - 45.0))
}

// Update updates the time properties of a node.
func (g *gameLayer) Update(msPerUpdate, secPerUpdate float64) {
	g.triangleMotion.Update(msPerUpdate)
	g.anchorMotion.Update(msPerUpdate)
}

// Interpolate is used for blending time based properties.
func (g *gameLayer) Interpolate(interpolation float64) {
	value := g.triangleMotion.Interpolate(interpolation)
	g.triangle.SetRotation(value.(float64))

	value = g.anchorMotion.Interpolate(interpolation)
	g.anchor.SetRotation(value.(float64))
}

// -----------------------------------------------------
// Node lifecycles
// -----------------------------------------------------

// EnterNode called when a node is entering the stage
func (g *gameLayer) EnterNode(man api.INodeManager) {
	// Register this node such that we get Update events above.
	man.RegisterTarget(g)
}

// ExitNode called when a node is exiting stage
func (g *gameLayer) ExitNode(man api.INodeManager) {
	man.UnRegisterTarget(g)
}
