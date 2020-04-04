package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type gameLayer struct {
	nodes.Node

	// Box 2D system
	b2Gravity box2d.B2Vec2
	b2World   box2d.B2World

	circleComp   *CircleComponent
	slopeCWComp  *GroundComponent
	slopeCCWComp *GroundComponent
	slopeCWComp2 *GroundComponent
	flatComp     *GroundComponent
}

func newBasicGameLayer(name string) api.INode {
	o := new(gameLayer)
	o.Initialize(name)
	return o
}

func (g *gameLayer) Build(world api.IWorld) {
	initializePhysics(g)

	buildBackground(g, world)

	g.circleComp = NewCircleComponent("CircleComp", g)
	g.circleComp.Configure(6, &g.b2World)
	g.circleComp.SetColor(rendering.NewPaletteInt64(rendering.Orange))
	g.circleComp.SetPosition(50.0, -100.0)
	g.circleComp.SetRadius(3.0)

	g.slopeCWComp = NewGroundComponent("SlopeCWComp", g)
	g.slopeCWComp.Configure(&g.b2World)
	g.slopeCWComp.SetColor(rendering.NewPaletteInt64(rendering.White))
	g.slopeCWComp.SetPosition(50.0, -50.0)
	g.slopeCWComp.SetRotation(maths.DegreeToRadians * 45.0)
	g.slopeCWComp.SetScale(25.0)

	g.slopeCCWComp = NewGroundComponent("SlopeCCWComp", g)
	g.slopeCCWComp.Configure(&g.b2World)
	g.slopeCCWComp.SetColor(rendering.NewPaletteInt64(rendering.White))
	g.slopeCCWComp.SetPosition(100.0, 0.0)
	g.slopeCCWComp.SetRotation(maths.DegreeToRadians * -45.0)
	g.slopeCCWComp.SetScale(25.0)

	g.slopeCWComp2 = NewGroundComponent("SlopeCWComp2", g)
	g.slopeCWComp2.Configure(&g.b2World)
	g.slopeCWComp2.SetColor(rendering.NewPaletteInt64(rendering.White))
	g.slopeCWComp2.SetPosition(50.0, 50.0)
	g.slopeCWComp2.SetRotation(maths.DegreeToRadians * 45.0)
	g.slopeCWComp2.SetScale(25.0)

	g.flatComp = NewGroundComponent("FlatComp", g)
	g.flatComp.Configure(&g.b2World)
	g.flatComp.SetColor(rendering.NewPaletteInt64(rendering.White))
	g.flatComp.SetPosition(75.0, 75.0)
	g.flatComp.SetScale(25.0)

	t := custom.NewRasterTextNode("RasterText", g.World(), g)
	tr := t.(*custom.RasterTextNode)
	tr.SetFontScale(2)
	tr.SetFill(2)
	tr.SetText("Press a key to reset.")
	tr.SetPosition(50.0, 50.0) // Note these coords are in device-space
	tr.SetColor(rendering.NewPaletteInt64(rendering.White))
}

// --------------------------------------------------------
// Timing
// --------------------------------------------------------

func (g *gameLayer) Update(msPerUpdate, secPerUpdate float64) {
	// Box2D expects a fractional number of dt not ms/frame which is
	// why I use secPerUpdate.

	// Instruct the world to perform a single step of simulation.
	// It is generally best to keep the time step and iterations fixed.
	g.b2World.Step(secPerUpdate, api.VelocityIterations, api.PositionIterations)

	g.circleComp.Update()
}

// -----------------------------------------------------
// Node lifecycles
// -----------------------------------------------------

// EnterNode called when a node is entering the stage
func (g *gameLayer) EnterNode(man api.INodeManager) {
	man.RegisterTarget(g)
	// Register for IO events so we can detect keyboard clicks
	man.RegisterEventTarget(g)
}

// ExitNode called when a node is exiting stage
func (g *gameLayer) ExitNode(man api.INodeManager) {
	man.UnRegisterTarget(g)
	man.UnRegisterEventTarget(g)

	g.b2World.Destroy()
}

// -----------------------------------------------------
// IO events
// -----------------------------------------------------

func (g *gameLayer) Handle(event api.IEvent) bool {
	if event.GetType() == api.IOTypeKeyboard {
		if event.GetState() == 1 {
			// Reset node and body properties
			g.circleComp.Reset()
		}
	}

	return false
}

// -----------------------------------------------------
// Misc private
// -----------------------------------------------------

func initializePhysics(g *gameLayer) {
	// Define the gravity vector.
	// Ranger's coordinate space is defined as:
	// .--------> +X
	// |
	// |
	// |
	// v +Y
	// Thus gravity is specified as positive for downward motion.
	g.b2Gravity = box2d.MakeB2Vec2(0.0, 9.8)

	// Construct a world object, which will hold and simulate the rigid bodies.
	g.b2World = box2d.MakeB2World(g.b2Gravity)
}

func buildBackground(g *gameLayer, world api.IWorld) {
	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	cb := custom.NewCheckBoardNode("CheckerBoard", world, g)
	cbr := cb.(*custom.CheckerBoardNode)
	cbr.Configure(25.0)

	hLine := custom.NewLineNode("HLine", world, g)
	n := hLine.(*custom.LineNode)
	n.SetColor(rendering.NewPaletteInt64(rendering.LightPurple))
	n.SetPoints(x, 0.0, -x, 0.0)

	vLine := custom.NewLineNode("VLine", world, g)
	n = vLine.(*custom.LineNode)
	n.SetColor(rendering.NewPaletteInt64(rendering.LightPurple))
	n.SetPoints(0.0, -y, 0.0, y)
}
