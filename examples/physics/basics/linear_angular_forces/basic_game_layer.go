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

	// Box 2D system
	b2Gravity box2d.B2Vec2
	b2World   box2d.B2World

	initialPosition api.IPoint

	orangeBoxComp *BoxComponent
	blueBoxComp   *BoxComponent
	limeBoxComp   *BoxComponent
	purpleBoxComp *BoxComponent
	yellowBoxComp *BoxComponent

	flatComp *GroundComponent

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

	initializePhysics(g)

	buildBackground(g, world)

	boxSize := 2.0
	groundSize := 100.0

	g.orangeBoxComp = NewBoxComponent("OrangeBoxComp", g)
	g.orangeBoxComp.Configure(boxSize, &g.b2World)
	g.orangeBoxComp.SetPosition(-20.0, -5.0)

	g.blueBoxComp = NewBoxComponent("BlueBoxComp", g)
	g.blueBoxComp.Configure(boxSize, &g.b2World)
	g.blueBoxComp.SetColor(rendering.NewPaletteInt64(rendering.SoftBlue))
	g.blueBoxComp.SetPosition(0.0, -5.0)

	g.limeBoxComp = NewBoxComponent("LimeBoxComp", g)
	g.limeBoxComp.Configure(boxSize, &g.b2World)
	g.limeBoxComp.SetColor(rendering.NewPaletteInt64(rendering.Lime))
	g.limeBoxComp.SetPosition(20.0, -5.0)

	g.purpleBoxComp = NewBoxComponent("PurpleBoxComp", g)
	g.purpleBoxComp.Configure(boxSize, &g.b2World)
	g.purpleBoxComp.EnableGravity(false)
	g.purpleBoxComp.SetColor(rendering.NewPaletteInt64(rendering.LightPurple))
	g.purpleBoxComp.SetPosition(40.0, -5.0)

	g.yellowBoxComp = NewBoxComponent("YellowBoxComp", g)
	g.yellowBoxComp.Configure(boxSize, &g.b2World)
	g.yellowBoxComp.EnableGravity(false)
	g.yellowBoxComp.SetColor(rendering.NewPaletteInt64(rendering.Yellow))
	g.yellowBoxComp.SetPosition(60.0, -5.0)

	g.flatComp = NewGroundComponent("Ground", g)
	g.flatComp.Configure(&g.b2World)
	g.flatComp.SetColor(rendering.NewPaletteInt64(rendering.White))
	g.flatComp.SetPosition(0.0, 0.0)
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

	g.orangeBoxComp.Update()
	g.blueBoxComp.Update()
	g.limeBoxComp.Update()
	g.purpleBoxComp.Update()
	g.yellowBoxComp.Update()
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
		// fmt.Println(event.GetKeyCode())
		if event.GetState() == 1 {
			switch event.GetKeyCode() {
			case 100: // d
				g.orangeBoxComp.ApplyForce(0.0, -500.0)
			case 102: // f
				g.blueBoxComp.ApplyImpulse(0.0, -200.0)
			case 122: // z
				g.limeBoxComp.ApplyImpulseToCorner(0.0, -200.0)
			case 97: // a
				g.purpleBoxComp.ApplyTorque(150.0)
			case 115: // s
				g.yellowBoxComp.ApplyAngularImpulse(50.0)
			case 114: // r
				// Reset node and body properties
				g.orangeBoxComp.Reset(-20.0, -5.0)
				g.blueBoxComp.Reset(0.0, -5.0)
				g.limeBoxComp.Reset(20.0, -5.0)
				g.purpleBoxComp.Reset(40.0, -5.0)
				g.yellowBoxComp.Reset(60.0, -5.0)
			}
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
