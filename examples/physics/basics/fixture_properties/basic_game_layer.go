package main

import (
	"fmt"
	"math"

	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// In this example we have a parent node "SeaSaw" that has three
// children: circle, box and polygon.
type gameLayer struct {
	nodes.Node

	// Box 2D system
	b2Gravity box2d.B2Vec2
	b2World   box2d.B2World

	initialPosition api.IPoint

	// Three Fixtures attach to one body
	quadBoxComp *QuadBoxComponent
	flatComp    *GroundComponent

	density     float64
	friction    float64
	restitution float64
}

func newBasicGameLayer(name string, world api.IWorld, parent api.INode) api.INode {
	o := new(gameLayer)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

func (g *gameLayer) Build(world api.IWorld) {
	g.Node.Build(world)

	g.initialPosition = geometry.NewPointUsing(0.0, -40.0)

	initializePhysics(g)

	buildBackground(g, world)

	boxSize := 4.0
	groundSize := 100.0

	g.quadBoxComp = NewQuadBoxComponent("QuadBoxComp", g)
	g.quadBoxComp.Configure(boxSize, &g.b2World)
	g.quadBoxComp.SetPosition(g.initialPosition.X(), g.initialPosition.Y())

	g.flatComp = NewGroundComponent("Ground", g)
	g.flatComp.Configure(&g.b2World)
	g.flatComp.SetColor(rendering.NewPaletteInt64(rendering.White))
	g.flatComp.SetPosition(0.0, 0.0)
	g.flatComp.SetRotation(maths.DegreeToRadians * -5.0)
	g.flatComp.SetScale(groundSize)

	t := custom.NewRasterTextNode("RasterText", g.World(), g)
	tr := t.(*custom.RasterTextNode)
	tr.SetFontScale(2)
	tr.SetFill(2)
	tr.SetText("See console for keys")
	tr.SetPosition(25.0, 40.0) // Note these coords are in device-space
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

	g.quadBoxComp.Update()
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
		fmt.Println(event.GetKeyCode())
		if event.GetState() == 1 {
			switch event.GetKeyCode() {
			case 100: // d = density
				g.density += 1.0
				fmt.Println("Setting density to: ", g.density)
				g.quadBoxComp.SetDensity(g.density)
			case 102: // f = friction
				fmt.Println("Setting friction to: ", g.friction)
				g.quadBoxComp.SetFriction(g.friction)
				g.friction += 0.1
				g.friction = math.Min(g.friction, 1.0)
			case 114: // r = restitution
				fmt.Println("Setting restitution to: ", g.restitution)
				g.quadBoxComp.SetDensity(g.restitution)
				g.restitution += 0.1
				g.restitution = math.Min(g.restitution, 1.0)
			case 122: // z = reset fixtures
				fmt.Println("Resetting fixtures")
				g.density = 1.0
				g.friction = 0.0
				g.restitution = 0.0
				g.quadBoxComp.ResetFixtures(1.0, 0.0, 0.0)
			}
			// Reset node and body properties
			g.quadBoxComp.Reset(g.initialPosition.X(), g.initialPosition.Y())
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
