package main

import (
	"fmt"

	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

const (
	entityBoundary      = uint16(1)
	entityCircle        = uint16(1 << 2)
	entityTriangle      = uint16(1 << 3)
	entityRectangle     = uint16(1 << 4)
	entityStarShip      = uint16(1 << 5)
	entityStarShipRight = uint16(1 << 6)
	entityStarShipLeft  = uint16(1 << 7)
	entityLand          = uint16(1 << 8)
)

type gameLayer struct {
	nodes.Node

	// Box 2D system
	b2Gravity box2d.B2Vec2
	b2World   box2d.B2World

	initialPosition api.IPoint

	trackerComp  *TrackingComponent
	starShipComp *StarShipComponent
	landComp     *BasicLandCompoent

	targetNode     api.INode
	targetPosition api.IPoint
	rayNode        api.INode

	fenceComp *FenceComponent

	velocityAlg int

	density     float64
	friction    float64
	restitution float64

	debug int
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

	targetSize := 4.0
	groundSize := 45.0
	g.initialPosition = geometry.NewPointUsing(0.0, 30.0)

	g.targetNode = custom.NewCrossNode("Target", world, g)
	g.targetNode.SetScale(5.0)
	g.targetPosition = geometry.NewPoint()

	g.rayNode = custom.NewLineNode("Ray", world, g)
	gl := g.rayNode.(*custom.LineNode)
	gl.SetColor(rendering.NewPaletteInt64(rendering.Lime))

	g.landComp = NewBasicLandCompoent("LandComp", g)
	g.landComp.Configure(2.0, 2.0, 00, 40.0, entityLand, entityTriangle|entityStarShip|entityStarShipRight|entityStarShipLeft, &g.b2World)
	g.landComp.SetColor(rendering.NewPaletteInt64(rendering.Silver))

	g.trackerComp = NewTrackingComponent("TriTrackerComp", g)
	g.trackerComp.Configure(targetSize, entityTriangle, entityStarShip|entityBoundary|entityLand, &g.b2World)
	g.trackerComp.SetPosition(50.0, -20.0)

	g.starShipComp = NewStarShipComponent("StarShip", g)
	g.starShipComp.Configure(targetSize/2.0, entityStarShip, entityTriangle|entityBoundary|entityLand, &g.b2World)
	g.starShipComp.Reset(g.initialPosition.X(), g.initialPosition.Y())

	g.fenceComp = NewFenceComponent("Ground", g)
	g.fenceComp.Configure(entityBoundary, entityTriangle|entityStarShip|entityStarShipRight|entityStarShipLeft, &g.b2World)
	g.fenceComp.SetColor(rendering.NewPaletteInt64(rendering.Silver))
	g.fenceComp.SetScale(groundSize)

	t := custom.NewRasterTextNode("RasterText", g.World(), g)
	tr := t.(*custom.RasterTextNode)
	tr.SetFontScale(2)
	tr.SetFill(2)
	tr.SetText("See console for keys")
	tr.SetPosition(15.0, 20.0) // Note these coords are in device-space
	tr.SetColor(rendering.NewPaletteInt64(rendering.White))

	filter := newFilterListener()
	g.b2World.SetContactFilter(filter)
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

	g.trackerComp.Update()
	g.starShipComp.Update()

	gl := g.rayNode.(*custom.LineNode)
	bodyPos := g.trackerComp.GetPosition()
	gl.SetPoints(bodyPos.X, bodyPos.Y, g.targetPosition.X(), g.targetPosition.Y())
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
	if event.GetType() == api.IOTypeMouseMotion {
		mx, my := event.GetMousePosition()
		nodes.MapDeviceToView(g.World(), mx, my, g.targetPosition)

		g.targetNode.SetPosition(g.targetPosition.X(), g.targetPosition.Y())
		g.trackerComp.SetTargetPosition(g.targetPosition)

		gl := g.rayNode.(*custom.LineNode)
		bodyPos := g.trackerComp.GetPosition()
		gl.SetPoints(bodyPos.X, bodyPos.Y, g.targetPosition.X(), g.targetPosition.Y())
	} else if event.GetType() == api.IOTypeKeyboard {
		// fmt.Println(event.GetKeyCode())
		if event.GetState() == 1 {
			switch event.GetKeyCode() {
			case 97: // a = CCW torque
				g.starShipComp.EnableYaw(true, -3.0)
			case 100: // d = CW torgue
				g.starShipComp.EnableYaw(true, 3.0)
			case 108: // l = impulse thrust
				g.starShipComp.SetThrust(true)
			case 119: // w = create

			case 102: // f = thrust impulse
				g.trackerComp.Thrust()
			case 120: // x = stop
			case 49: // 1
				fmt.Println("Set velocity algorithm to: 1")
				g.trackerComp.SetVelocityAlgo(1)
			case 50: // 2
				fmt.Println("Set velocity algorithm to: 2")
				g.trackerComp.SetVelocityAlgo(2)
			case 51: // 3
				fmt.Println("Set velocity algorithm to: 3")
				g.trackerComp.SetVelocityAlgo(3)
			case 52: // 4
				fmt.Println("Set tracking algorithm to: 1")
				g.trackerComp.SetTrackingAlgo(1)
			case 53: // 5
				fmt.Println("Set tracking algorithm to: 2")
				g.trackerComp.SetTrackingAlgo(2)
			case 54: // 6
				fmt.Println("Set tracking algorithm to: 3")
				g.trackerComp.SetTrackingAlgo(3)
			case 55: // 7
				fmt.Println("Set tracking algorithm to: 4")
				g.trackerComp.SetTrackingAlgo(4)
			case 56: // 8
				fmt.Println("Set tracking algorithm to: 5")
				g.trackerComp.SetTrackingAlgo(5)
			case 57: // 9
				fmt.Println("Set tracking algorithm to: 6")
				g.trackerComp.SetTrackingAlgo(6)
			case 116: // t
				fmt.Println("Set targeting rate to: 5(fast)")
				g.trackerComp.SetTargetingRate(5)
			case 121: // y
				fmt.Println("Set targeting rate to: 10")
				g.trackerComp.SetTargetingRate(10)
			case 117: // u
				fmt.Println("Set targeting rate to: 20")
				g.trackerComp.SetTargetingRate(20)
			case 105: // i
				fmt.Println("Set targeting rate to: 40")
				g.trackerComp.SetTargetingRate(40)
			case 111: // o
				fmt.Println("Set targeting rate to: 60(slow)")
				g.trackerComp.SetTargetingRate(60)
			case 114: // r
				// Reset node and body properties
				g.starShipComp.Reset(g.initialPosition.X(), g.initialPosition.Y())
			}
		} else {
			switch event.GetKeyCode() {
			case 108: // l = impulse thrust
				g.starShipComp.SetThrust(false)
			case 97: // a = CCW torque
				g.starShipComp.EnableYaw(false, 0.0)
			case 100: // d = CW torgue
				g.starShipComp.EnableYaw(false, 0.0)
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

	// A solid background
	sol := custom.NewBasicRectangleNode("Background", world, g)
	gs := sol.(*custom.BasicRectangleNode)
	gs.SetBoundsUsingComps(x, y, x+vw, y+vh)
	gs.SetColor(rendering.NewPaletteInt64(rendering.DarkerGray))

	hLine := custom.NewLineNode("HLine", world, g)
	n := hLine.(*custom.LineNode)
	n.SetColor(rendering.NewPaletteInt64(rendering.LightGray))
	n.SetPoints(x, 0.0, -x, 0.0)

	vLine := custom.NewLineNode("VLine", world, g)
	n = vLine.(*custom.LineNode)
	n.SetColor(rendering.NewPaletteInt64(rendering.LightGray))
	n.SetPoints(0.0, -y, 0.0, y)
}
