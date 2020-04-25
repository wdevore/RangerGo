package main

import (
	"fmt"

	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
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

const (
	objectRightZone = 2000
	objectLeftZone  = 2001
)

type gameLayer struct {
	nodes.Node

	initialPosition api.IPoint

	// --------------------------------------------
	// Physics stuff
	// --------------------------------------------
	// Box 2D system
	b2Gravity box2d.B2Vec2
	b2World   box2d.B2World

	starShipComp *StarShipComponent
	landComp     *BasicLandCompoent
	fenceComp    *FenceComponent

	zoneMan *zoneManager

	// Scrolling
	scrollRing *geometry.Circle
	circleNode api.INode
	softness   float64 // Controls how soft the scroll reacts

	tweenValue float64

	// Debug
	shipPosTxt api.INode
	lpTxt      api.INode
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

	g.zoneMan = newZoneManager(g)
	g.zoneMan.Build(world)
	zoom := g.zoneMan.GetZoom()

	targetSize := 4.0
	groundSize := 45.0
	g.softness = 10.0
	g.initialPosition = geometry.NewPointUsing(0.0, 0.0)

	g.scrollRing = geometry.NewCircle()
	g.scrollRing.SetRadius(15.0)
	g.circleNode = NewCircleNode("ScrollRing", world, g)
	g.circleNode.SetVisible(false)
	gcr := g.circleNode.(*CircleNode)
	gcr.Configure(12, g.scrollRing.Radius())

	g.landComp = NewBasicLandCompoent("LandComp", zoom)
	g.landComp.Configure(2.0, 2.0, 0, 40.0, entityLand, entityTriangle|entityStarShip|entityStarShipRight|entityStarShipLeft, &g.b2World)
	g.landComp.SetColor(rendering.NewPaletteInt64(rendering.Silver))

	g.starShipComp = NewStarShipComponent("StarShip", zoom)
	g.starShipComp.Configure(targetSize/2.0, entityStarShip, entityTriangle|entityBoundary|entityLand, &g.b2World)
	g.starShipComp.Reset(g.initialPosition.X(), g.initialPosition.Y())

	g.fenceComp = NewFenceComponent("Ground", zoom)
	g.fenceComp.Configure(entityBoundary, entityTriangle|entityStarShip|entityStarShipRight|entityStarShipLeft, &g.b2World)
	g.fenceComp.SetColor(rendering.NewPaletteInt64(rendering.LightGray))
	g.fenceComp.SetScale(groundSize)

	g.shipPosTxt = custom.NewRasterTextNode("SPText", g.World(), g)
	tr := g.shipPosTxt.(*custom.RasterTextNode)
	tr.SetFontScale(2)
	tr.SetFill(1)
	tr.SetPosition(15.0, 20.0) // Note these coords are in device-space
	tr.SetColor(rendering.NewPaletteInt64(rendering.White))

	g.lpTxt = custom.NewRasterTextNode("LPText", g.World(), g)
	tr = g.lpTxt.(*custom.RasterTextNode)
	tr.SetFontScale(2)
	tr.SetFill(1)
	tr.SetPosition(15.0, 35.0) // Note these coords are in device-space
	tr.SetColor(rendering.NewPaletteInt64(rendering.White))

	filter := newFilterListener()
	g.b2World.SetContactFilter(filter)
}

// --------------------------------------------------------
// Timing
// --------------------------------------------------------
var nodePoint = geometry.NewPoint()

func (g *gameLayer) Update(msPerUpdate, secPerUpdate float64) {
	// Box2D expects a fractional number of dt not ms/frame which is
	// why I use secPerUpdate.

	// Instruct the world to perform a single step of simulation.
	// It is generally best to keep the time step and iterations fixed.
	g.b2World.Step(secPerUpdate, api.VelocityIterations, api.PositionIterations)

	g.starShipComp.Update()

	// shipPos := geometry.NewPointUsing(g.starShipComp.BodyPosition())
	shipPos := g.starShipComp.Position()

	g.zoneMan.UpdateCheck(shipPos, msPerUpdate)

	nodes.MapNodeToNode(g.starShipComp.hullVisual, g.circleNode, nodePoint, g)

	dRange := g.scrollRing.DistanceFromEdge(nodePoint)

	tr := g.shipPosTxt.(*custom.RasterTextNode)
	tr.SetText(fmt.Sprintf("SP (%2.3f) %s", dRange, shipPos))

	if dRange > 0 {
		gc := g.circleNode.(*CircleNode)
		gc.SetColor(rendering.NewPaletteInt64(rendering.Peach))
		tr = g.lpTxt.(*custom.RasterTextNode)
		tr.SetText(fmt.Sprintf("LP %s", nodePoint))

		ray := maths.NewVectorUsing(nodePoint.X(), nodePoint.Y())
		ray.Normalize()
		ray.Scale(-dRange / g.softness)

		zoom := g.zoneMan.GetZoom()
		gz := zoom.(*custom.ZoomNode)
		gz.TranslateBy(ray.X(), ray.Y())
	} else {
		tr = g.lpTxt.(*custom.RasterTextNode)
		tr.SetText("---")
		gc := g.circleNode.(*CircleNode)
		gc.SetColor(rendering.NewPaletteInt64(rendering.Silver))
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
	if event.GetType() == api.IOTypeMouseMotion {
		// mx, my := event.GetMousePosition()
		// nodes.MapDeviceToView(g.World(), mx, my, g.targetPosition)
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
			case 120: // x = stop
			case 49: // 1
			case 50: // 2
			case 51: // 3
			case 52: // 4
			case 53: // 5
			case 54: // 6
			case 55: // 7
			case 56: // 8
			case 57: // 9
			case 116: // t
			case 121: // y
			case 117: // u
			case 105: // i
			case 111: // o
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
	g.b2Gravity = box2d.MakeB2Vec2(0.0, 5.0) // 0.0, 9.8

	// Construct a world object, which will hold and simulate the rigid bodies.
	g.b2World = box2d.MakeB2World(g.b2Gravity)
}
