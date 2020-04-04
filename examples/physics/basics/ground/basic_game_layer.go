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

	circleNode     api.INode
	groundLineNode api.INode

	// Box 2D system
	b2Gravity box2d.B2Vec2
	b2World   box2d.B2World

	b2CircleBody *box2d.B2Body
	b2GroundBody *box2d.B2Body
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

	g.textColor = rendering.NewPaletteInt64(rendering.White)

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

	// -------------------------------------------
	// Visuals for Box2D
	g.circleNode = NewCircleNode("Orange Circle", world, g)
	gr := g.circleNode.(*CircleNode)
	gr.Configure(6, 1.0)
	gr.SetColor(rendering.NewPaletteInt64(rendering.Orange))
	gr.SetScale(3.0)
	gr.SetPosition(100.0, -100.0)

	g.groundLineNode = custom.NewLineNode("Ground", world, g)
	gln := g.groundLineNode.(*custom.LineNode)
	gln.SetColor(rendering.NewPaletteInt64(rendering.White))
	gln.SetPoints(-1.0, 0.0, 1.0, 0.0) // Set by unit coordinates
	gln.SetPosition(76.0+50.0, 0.0)
	gln.SetScale(25.0)

	t := custom.NewRasterTextNode("RasterText", world, g)
	tr := t.(*custom.RasterTextNode)
	tr.SetFontScale(2)
	tr.SetFill(2)
	tr.SetText("Press key to reset.")
	tr.SetPosition(50.0, 50.0)
	tr.SetColor(rendering.NewPaletteInt64(rendering.White))

	buildPhysicsWorld(g)
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

	if g.b2CircleBody.IsActive() {
		pos := g.b2CircleBody.GetPosition()
		g.circleNode.SetPosition(pos.X, pos.Y)

		rot := g.b2CircleBody.GetAngle()
		g.circleNode.SetRotation(rot)
	}
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
			x := 100.0
			y := -100.0
			g.circleNode.SetPosition(x, y)
			g.b2CircleBody.SetTransform(box2d.MakeB2Vec2(x, y), 0.0)
			g.b2CircleBody.SetLinearVelocity(box2d.MakeB2Vec2(0.0, 0.0))
			g.b2CircleBody.SetAngularVelocity(0.0)
		}
	}

	return false
}

// -----------------------------------------------------
// Misc private
// -----------------------------------------------------

func buildPhysicsWorld(g *gameLayer) {
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

	// -------------------------------------------
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody
	bDef.Position.Set(g.circleNode.Position().X(), g.circleNode.Position().Y())

	// An instance of a body to contain Fixtures
	g.b2CircleBody = g.b2World.CreateBody(&bDef)

	// Every Fixture has a shape
	circleShape := box2d.MakeB2CircleShape()
	circleShape.M_p.Set(0.0, 0.0) // Relative to body position
	circleShape.M_radius = g.circleNode.Scale()

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &circleShape
	fd.Density = 1.0
	g.b2CircleBody.CreateFixtureFromDef(&fd) // attach Fixture to body

	// -------------------------------------------
	// The Ground = body + fixture + shape
	bDef.Type = box2d.B2BodyType.B2_staticBody
	bDef.Position.Set(g.groundLineNode.Position().X(), g.groundLineNode.Position().Y())

	g.b2GroundBody = g.b2World.CreateBody(&bDef)

	groundShape := box2d.MakeB2EdgeShape()
	groundShape.Set(box2d.MakeB2Vec2(-g.groundLineNode.Scale(), 0.0), box2d.MakeB2Vec2(g.groundLineNode.Scale(), 0.0))

	fDef := box2d.MakeB2FixtureDef()
	fDef.Shape = &groundShape
	fDef.Density = 1.0
	g.b2GroundBody.CreateFixtureFromDef(&fDef) // attach Fixture to body
}
