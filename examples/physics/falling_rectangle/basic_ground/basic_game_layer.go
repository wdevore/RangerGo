package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type gameLayer struct {
	nodes.Node

	textColor api.IPalette

	circleNode api.INode

	// Box 2D system
	b2Gravity box2d.B2Vec2
	b2World   box2d.B2World

	b2CircleBody *box2d.B2Body
}

func newBasicGameLayer(name string) api.INode {
	o := new(gameLayer)
	o.Initialize(name)
	return o
}

func (g *gameLayer) Build(world api.IWorld) {
	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	g.textColor = rendering.NewPaletteInt64(rendering.LightNavyBlue)

	custom.NewCheckBoardNode("CheckerBoard", world, g)

	hLine := custom.NewLineNode("HLine", world, g)
	n := hLine.(*custom.LineNode)
	n.SetColor(rendering.NewPaletteInt64(rendering.LightPurple))
	n.SetPoints(x, 0.0, -x, 0.0)

	vLine := custom.NewLineNode("VLine", world, g)
	n = vLine.(*custom.LineNode)
	n.SetColor(rendering.NewPaletteInt64(rendering.LightPurple))
	n.SetPoints(0.0, -y, 0.0, y)

	// Visual for Box2D
	g.circleNode = custom.NewCircleNode("Orange Circle", world, g)
	gr := g.circleNode.(*custom.CircleNode)
	gr.SetColor(rendering.NewPaletteInt64(rendering.Orange))
	g.circleNode.SetScale(0.1 * api.PTM)
	g.circleNode.SetPosition(0.0, -200.0)

	// --------------------------------------------
	// Box 2d configuration
	// --------------------------------------------

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

	// A body def used to create body
	bd := box2d.MakeB2BodyDef()
	bd.Type = box2d.B2BodyType.B2_dynamicBody
	bd.Position.Set(g.circleNode.Position().X(), g.circleNode.Position().Y())

	// An instance of a body to contain the Fixtures
	g.b2CircleBody = g.b2World.CreateBody(&bd)

	// Every Fixture has a shape
	circleShape := box2d.MakeB2CircleShape()
	circleShape.M_p.Set(0.0, 0.0) // Relative to body position
	circleShape.M_radius = 0.1

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &circleShape
	fd.Density = 10.0
	g.b2CircleBody.CreateFixtureFromDef(&fd) // attach Fixture to body
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

	pos := g.b2CircleBody.GetPosition()
	if g.b2CircleBody.IsActive() {
		g.circleNode.SetPosition(pos.X, pos.Y)
	}
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

	g.b2World.Destroy()
}

// -----------------------------------------------------
// Visuals
// -----------------------------------------------------

// func (g *gameLayer) Draw(context api.IRenderContext) {
// 	if g.IsDirty() {
// 		g.SetDirty(false)
// 	}
// }
