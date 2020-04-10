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

type gameLayer struct {
	nodes.Node

	// Box 2D system
	b2Gravity box2d.B2Vec2
	b2World   box2d.B2World

	initialPosition api.IPoint

	trackerComp *TrackingComponent
	boxComp     *BoxComponent

	targetNode     api.INode
	targetPosition api.IPoint
	rayNode        api.INode

	contactPointsNode []api.INode
	normalsNode       []api.INode
	normalLength      float64

	impactNode api.INode

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

	targetSize := 5.0
	groundSize := 50.0
	g.normalLength = 2.0

	g.targetNode = custom.NewCrossNode("Target", world, g)
	g.targetNode.SetScale(5.0)
	g.targetPosition = geometry.NewPoint()

	g.rayNode = custom.NewLineNode("Ray", world, g)
	gl := g.rayNode.(*custom.LineNode)
	gl.SetColor(rendering.NewPaletteInt64(rendering.Lime))

	g.contactPointsNode = append(g.contactPointsNode, custom.NewBigPointNode("Big1", world, g))
	gc := g.contactPointsNode[0].(*custom.BigPointNode)
	gc.SetColor(rendering.NewPaletteInt64(rendering.White))
	g.contactPointsNode = append(g.contactPointsNode, custom.NewBigPointNode("Big2", world, g))
	gc = g.contactPointsNode[1].(*custom.BigPointNode)
	gc.SetColor(rendering.NewPaletteInt64(rendering.Lime))
	g.contactPointsNode = append(g.contactPointsNode, custom.NewBigPointNode("Big3", world, g))
	gc = g.contactPointsNode[2].(*custom.BigPointNode)
	gc.SetColor(rendering.NewPaletteInt64(rendering.LightPurple))
	g.contactPointsNode = append(g.contactPointsNode, custom.NewBigPointNode("Big4", world, g))
	gc = g.contactPointsNode[3].(*custom.BigPointNode)
	gc.SetColor(rendering.NewPaletteInt64(rendering.Olive))

	g.normalsNode = append(g.normalsNode, custom.NewLineNode("Norm1", world, g))
	gn := g.normalsNode[0].(*custom.LineNode)
	gn.SetColor(rendering.NewPaletteInt64(rendering.Peach))
	g.normalsNode = append(g.normalsNode, custom.NewLineNode("Norm2", world, g))
	gn = g.normalsNode[1].(*custom.LineNode)
	gn.SetColor(rendering.NewPaletteInt64(rendering.Yellow))

	g.impactNode = custom.NewLineNode("Impact", world, g)
	gl = g.impactNode.(*custom.LineNode)
	gl.SetColor(rendering.NewPaletteInt64(rendering.Teal))

	g.trackerComp = NewTrackingComponent("TriTrackerComp", g)
	g.trackerComp.Configure(targetSize, &g.b2World)
	g.trackerComp.SetPosition(0.0, 0.0)

	g.boxComp = NewBoxComponent("BoxComp", g)
	g.boxComp.Configure(targetSize/2.0, &g.b2World)
	g.boxComp.SetPosition(10.0, 0.0)

	// Because I am using edges instead of boxes for bounding
	// AND because I am applying velocities explicitly, the triangle
	// has a high potential of exiting the fence area.
	g.fenceComp = NewFenceComponent("Ground", g)
	g.fenceComp.Configure(&g.b2World)
	g.fenceComp.SetColor(rendering.NewPaletteInt64(rendering.White))
	g.fenceComp.SetScale(groundSize)

	t := custom.NewRasterTextNode("RasterText", g.World(), g)
	tr := t.(*custom.RasterTextNode)
	tr.SetFontScale(2)
	tr.SetFill(2)
	tr.SetText("See console for keys")
	tr.SetPosition(25.0, 40.0) // Note these coords are in device-space
	tr.SetColor(rendering.NewPaletteInt64(rendering.White))

	// Contacts
	listener := newContactListener()
	lr := listener.(*contactListener)
	lr.addListener(g.trackerComp)
	lr.addListener(g.boxComp)

	g.b2World.SetContactListener(listener)
}

// ContactListener
type contactListener struct {
	components []api.IContactListener
}

func newContactListener() box2d.B2ContactListenerInterface {
	o := new(contactListener)
	o.components = []api.IContactListener{}
	return o
}

func (c *contactListener) addListener(component api.IContactListener) {
	c.components = append(c.components, component)
}

func (c *contactListener) BeginContact(contact box2d.B2ContactInterface) {
	fixA := contact.GetFixtureA()
	fixB := contact.GetFixtureB()

	dataA := fixA.GetUserData()
	dataB := fixB.GetUserData()

	if dataA != nil && dataB != nil {
		for _, l := range c.components {
			handled := l.HandleBeginContact(dataA.(api.INode), dataB.(api.INode))
			if handled {
				break
			}
		}
	}
}

func (c *contactListener) EndContact(contact box2d.B2ContactInterface) {
	fixA := contact.GetFixtureA()
	fixB := contact.GetFixtureB()

	dataA := fixA.GetUserData()
	dataB := fixB.GetUserData()

	if dataA != nil && dataB != nil {
		for _, l := range c.components {
			handled := l.HandleEndContact(dataA.(api.INode), dataB.(api.INode))
			if handled {
				break
			}
		}
	}
}

func (c *contactListener) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {
	// fmt.Println("PreSolve")

}

func (c *contactListener) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {
	// fmt.Println("PostSolve")

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
	g.boxComp.Update()

	//scanContactList(g)

	gl := g.rayNode.(*custom.LineNode)
	bodyPos := g.trackerComp.GetPosition()
	gl.SetPoints(bodyPos.X, bodyPos.Y, g.targetPosition.X(), g.targetPosition.Y())
}

// This an example of Activity checking for contacts. It is useful in some
// cases.
func scanContactList(g *gameLayer) {
	for contact := g.b2World.GetContactList(); contact != nil; contact = contact.GetNext() {
		if contact.IsTouching() {
			fixA := contact.GetFixtureA()
			fixB := contact.GetFixtureB()

			dataA := fixA.GetUserData()
			dataB := fixB.GetUserData()

			if dataA != nil && dataB != nil {
				if dataA.(api.INode).ID() == 1001 && dataB.(api.INode).ID() == 1000 {
					manA := contact.GetManifold()
					worldMan := box2d.B2WorldManifold{}
					contact.GetWorldManifold(&worldMan)

					pointCnt := manA.PointCount
					for i := 0; i < pointCnt; i++ {
						gc := g.contactPointsNode[i].(*custom.BigPointNode)
						gc.SetPoint(worldMan.Points[i].X, worldMan.Points[i].Y)

						// Normals that show the direction for separation not collision normal.
						gn := g.normalsNode[i].(*custom.LineNode)
						gn.SetPoints(
							worldMan.Points[i].X-g.normalLength*worldMan.Normal.X,
							worldMan.Points[i].Y-g.normalLength*worldMan.Normal.Y,
							worldMan.Points[i].X+g.normalLength*worldMan.Normal.X,
							worldMan.Points[i].Y+g.normalLength*worldMan.Normal.Y)

						vel1 := fixA.GetBody().GetLinearVelocityFromWorldPoint(worldMan.Points[i])
						vel2 := fixB.GetBody().GetLinearVelocityFromWorldPoint(worldMan.Points[i])

						impactVelocity := box2d.B2Vec2{X: vel1.X - vel2.X, Y: vel1.Y - vel2.Y}

						gi := g.impactNode.(*custom.LineNode)
						gi.SetPoints(
							worldMan.Points[i].X+g.normalLength*impactVelocity.X,
							worldMan.Points[i].Y+g.normalLength*impactVelocity.Y,
							worldMan.Points[i].X-g.normalLength*impactVelocity.X,
							worldMan.Points[i].Y-g.normalLength*impactVelocity.Y)
					}

				}
			}
		}
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
			case 97: // a = move left
				g.trackerComp.MoveLeft(5.0)
			case 100: // d = move right
				g.trackerComp.MoveRight(5.0)
			case 115: // s = down
				g.trackerComp.MoveDown(5.0)
			case 119: // w = up
				g.trackerComp.MoveUp(5.0)
			case 102: // f = thrust impulse
				g.trackerComp.Thrust()
			case 120: // x = stop
				g.trackerComp.Stop()
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
	g.b2Gravity = box2d.MakeB2Vec2(0.0, 0.0)

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
