package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
)

// CircleComponent represents both the visual and physic components
type CircleComponent struct {
	visual api.INode

	b2Body    *box2d.B2Body
	b2Shape   box2d.B2CircleShape
	b2Fixture *box2d.B2Fixture
}

// NewCircleComponent constructs a circle component
func NewCircleComponent(name string, parent api.INode) *CircleComponent {
	o := new(CircleComponent)
	o.visual = NewCircleNode(name, parent.World(), parent)
	return o
}

// Configure component
func (c *CircleComponent) Configure(segments int, b2World *box2d.B2World) {
	gr := c.visual.(*CircleNode)
	gr.Configure(segments, 1.0)

	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody

	// An instance of a body to contain Fixtures
	c.b2Body = b2World.CreateBody(&bDef)

	// Every Fixture has a shape
	c.b2Shape = box2d.MakeB2CircleShape()

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &c.b2Shape
	fd.Density = 1.0
	c.b2Fixture = c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}

// SetColor sets the visual's color
func (c *CircleComponent) SetColor(color api.IPalette) {
	gr := c.visual.(*CircleNode)
	gr.SetColor(color)
}

// SetRadius sets circle's radius
func (c *CircleComponent) SetRadius(radius float64) {
	c.visual.SetScale(radius)

	// To change a shape's property we need to destroy the old
	// fixture and create a new based on the new radius.
	c.b2Shape.SetRadius(radius) // Set new radius

	c.b2Body.DestroyFixture(c.b2Fixture)

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &c.b2Shape
	fd.Density = 1.0
	c.b2Fixture = c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}

// SetPosition sets component's location.
func (c *CircleComponent) SetPosition(x, y float64) {
	c.visual.SetPosition(x, y)
	c.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), c.b2Body.GetAngle())
}

// Reset configures the component back to defaults
func (c *CircleComponent) Reset() {
	x := 100.0
	y := -100.0
	c.visual.SetPosition(x, y)
	c.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), 0.0)
	c.b2Body.SetLinearVelocity(box2d.MakeB2Vec2(0.0, 0.0))
	c.b2Body.SetAngularVelocity(0.0)
}

// Update component
func (c *CircleComponent) Update() {
	if c.b2Body.IsActive() {
		pos := c.b2Body.GetPosition()
		c.visual.SetPosition(pos.X, pos.Y)

		rot := c.b2Body.GetAngle()
		c.visual.SetRotation(rot)
	}
}
