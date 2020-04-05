package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// SeaSawComponent represents both the visual and physic components.
// It has three children: circle, box and polygon.
//
// When the component drops and hits the ground it will topple over
// towards the rectangle because it has a higher mass value than the
// circle.
// Depending on how much you scale the polygon it will take longer or
// shorter to topple over.
type SeaSawComponent struct {
	polyVisual   api.INode // The parent of the component
	circleVisual api.INode
	rectVisual   api.INode

	b2Body *box2d.B2Body

	scale float64
}

// NewSeaSawComponent constructs a component
func NewSeaSawComponent(name string, parent api.INode) *SeaSawComponent {
	o := new(SeaSawComponent)

	world := parent.World()
	o.polyVisual = custom.NewPolygonNode("PolyFix", world, parent)
	o.circleVisual = NewCircleNode("CircleFix", world, o.polyVisual)
	o.rectVisual = custom.NewRectangleNode("RectFix", world, o.polyVisual)

	return o
}

// Configure component
func (c *SeaSawComponent) Configure(scale float64, b2World *box2d.B2World) {
	c.scale = scale

	buildPolygon(c, b2World)
	buildCircle(c)
	buildRectangle(c)
}

// SetPosition sets component's location.
func (c *SeaSawComponent) SetPosition(x, y float64) {
	c.polyVisual.SetPosition(x, y)
	c.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), c.b2Body.GetAngle())
}

// Reset configures the component back to defaults
func (c *SeaSawComponent) Reset(x, y float64) {
	c.polyVisual.SetPosition(x, y)
	c.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), 0.0)
	c.b2Body.SetLinearVelocity(box2d.MakeB2Vec2(0.0, 0.0))
	c.b2Body.SetAngularVelocity(0.0)
	// You waited until the topple completed then the body went to sleep.
	// Thus we need to wake it back up otherwise it would just hang in
	// mid air.
	c.b2Body.SetAwake(true)
}

// Update component
func (c *SeaSawComponent) Update() {
	if c.b2Body.IsActive() {
		pos := c.b2Body.GetPosition()
		c.polyVisual.SetPosition(pos.X, pos.Y)

		rot := c.b2Body.GetAngle()
		c.polyVisual.SetRotation(rot)
	}
}

func buildPolygon(c *SeaSawComponent, b2World *box2d.B2World) {
	// ------------------------------------------------
	// Polygon
	// ------------------------------------------------
	gp := c.polyVisual.(*custom.PolygonNode)
	gp.SetColor(rendering.NewPaletteInt64(rendering.Orange))

	// Note: +Y points down in Ranger verses Upward in Box2D's GUI.
	vertices := []box2d.B2Vec2{}
	vertices = append(vertices, box2d.B2Vec2{X: -1.0 * c.scale, Y: -2.0 * c.scale})
	vertices = append(vertices, box2d.B2Vec2{X: -1.0 * c.scale, Y: 0.0 * c.scale})
	vertices = append(vertices, box2d.B2Vec2{X: 0.0 * c.scale, Y: 3.0 * c.scale})
	vertices = append(vertices, box2d.B2Vec2{X: 1.0 * c.scale, Y: 0.0 * c.scale})
	vertices = append(vertices, box2d.B2Vec2{X: 1.0 * c.scale, Y: -1.0 * c.scale})

	ply := c.polyVisual.(*custom.PolygonNode)
	ply.AddVertex(vertices[0].X, vertices[0].Y, false)
	ply.AddVertex(vertices[1].X, vertices[1].Y, false)
	ply.AddVertex(vertices[2].X, vertices[2].Y, false)
	ply.AddVertex(vertices[3].X, vertices[3].Y, false)
	ply.AddVertex(vertices[4].X, vertices[4].Y, true)

	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody

	// An instance of a body to contain Fixtures
	c.b2Body = b2World.CreateBody(&bDef)

	// Every Fixture has a shape
	b2PolyShape := box2d.MakeB2PolygonShape()
	b2PolyShape.Set(vertices, 5)

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2PolyShape
	fd.Density = 1.0
	c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}

func buildCircle(c *SeaSawComponent) {
	// ------------------------------------------------
	// Circle
	// ------------------------------------------------
	gc := c.circleVisual.(*CircleNode)
	gc.Configure(6, 1.0)
	gc.SetColor(rendering.NewPaletteInt64(rendering.Lime))
	c.circleVisual.SetPosition(-15.0, 0.0)
	c.circleVisual.SetScale(c.scale)

	// Every Fixture has a shape
	b2CircleShape := box2d.MakeB2CircleShape()
	b2CircleShape.SetRadius(1.0 * c.scale)
	b2CircleShape.M_p.Set(c.circleVisual.Position().X(), c.circleVisual.Position().Y()) // Relative to body

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2CircleShape
	fd.Density = 1.0
	c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}

func buildRectangle(c *SeaSawComponent) {
	// ------------------------------------------------
	// Box
	// ------------------------------------------------
	gr := c.rectVisual.(*custom.RectangleNode)
	gr.SetScale(2.0 * c.scale)
	gr.SetColor(rendering.NewPaletteInt64(rendering.Aqua))
	c.rectVisual.SetPosition(15.0, 0.0)

	// Every Fixture has a shape
	b2RectShape := box2d.MakeB2PolygonShape() // 2x2 rectangle
	b2RectShape.SetAsBoxFromCenterAndAngle(1.0*c.scale, 1.0*c.scale, box2d.B2Vec2{X: c.rectVisual.Position().X(), Y: c.rectVisual.Position().Y()}, 0.0)

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2RectShape
	fd.Density = 1.0
	c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}
