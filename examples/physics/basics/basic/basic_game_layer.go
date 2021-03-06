package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type gameLayer struct {
	nodes.Node

	mesh api.IMesh

	textColor api.IPalette

	oddColor  api.IPalette
	evenColor api.IPalette

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
	w := 50.0

	g.mesh = geometry.NewMesh()

	// Construct grid of rectangles
	for y <= vh {
		x := -vw / 2.0
		for x <= vw {
			g.mesh.AddVertex(x, y) // top-left
			g.mesh.AddVertex(x+w, y+w)
			x += w
		}
		y += w
	}

	g.mesh.Build()

	hLine := custom.NewLineNode("HLine", world, g)
	n := hLine.(*custom.LineNode)
	n.SetColor(rendering.NewPaletteInt64(rendering.LightPurple))
	n.SetPoints(x, 0.0, -x, 0.0)

	vLine := custom.NewLineNode("VLine", world, g)
	n = vLine.(*custom.LineNode)
	n.SetColor(rendering.NewPaletteInt64(rendering.LightPurple))
	n.SetPoints(0.0, -y, 0.0, y)

	g.circleNode = custom.NewCircleNode("Orange Circle", world, g)
	g.circleNode.SetScale(5.0)
	g.circleNode.SetPosition(0.0, -100.0)
	gr := g.circleNode.(*custom.CircleNode)
	gr.Configure(12, 1.0)
	gr.SetColor(rendering.NewPaletteInt64(rendering.Orange))

	g.textColor = rendering.NewPaletteInt64(rendering.LightNavyBlue)
	g.oddColor = rendering.NewPaletteInt64(rendering.DarkGray)
	g.evenColor = rendering.NewPaletteInt64(rendering.LightGray)

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
	circleShape.M_radius = 5

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

func (g *gameLayer) Draw(context api.IRenderContext) {
	// Transform vertices if anything has changed.
	if g.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformMesh(g.mesh)
		g.SetDirty(false) // Node is no longer dirty
	}

	// Draw background first. The background is a grid of squares.
	context.RenderCheckerBoard(g.mesh, g.oddColor, g.evenColor)

}
