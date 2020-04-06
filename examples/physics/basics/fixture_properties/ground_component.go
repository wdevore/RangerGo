package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

// GroundComponent represents both the visual and physic components
type GroundComponent struct {
	visual api.INode

	b2Body    *box2d.B2Body
	b2Shape   box2d.B2EdgeShape
	b2Fixture *box2d.B2Fixture
}

// NewGroundComponent constructs a component
func NewGroundComponent(name string, parent api.INode) *GroundComponent {
	o := new(GroundComponent)
	o.visual = custom.NewLineNode(name, parent.World(), parent)
	gln := o.visual.(*custom.LineNode)
	gln.SetPoints(-1.0, 0.0, 1.0, 0.0) // Set by unit coordinates
	return o
}

// Configure component
func (c *GroundComponent) Configure(b2World *box2d.B2World) {
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_staticBody

	// An instance of a body to contain Fixtures
	c.b2Body = b2World.CreateBody(&bDef)

	// Every Fixture has a shape
	c.b2Shape = box2d.MakeB2EdgeShape()

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &c.b2Shape
	fd.Density = 1.0
	c.b2Fixture = c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}

// SetColor sets the visual's color
func (c *GroundComponent) SetColor(color api.IPalette) {
	gr := c.visual.(*custom.LineNode)
	gr.SetColor(color)
}

// SetScale sets the component's length
func (c *GroundComponent) SetScale(scale float64) {
	c.visual.SetScale(scale)

	// To change a shape's property we need to destroy the old
	// fixture and create a new one based on the new value.
	c.b2Body.DestroyFixture(c.b2Fixture)

	c.b2Shape.Set(box2d.MakeB2Vec2(-scale, 0.0), box2d.MakeB2Vec2(scale, 0.0))

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &c.b2Shape
	fd.Density = 1.0
	c.b2Fixture = c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}

// SetRotation sets the component's orientation
func (c *GroundComponent) SetRotation(angle float64) {
	c.visual.SetRotation(angle)
	c.b2Body.SetTransform(c.b2Body.GetPosition(), angle)
}

// SetPosition sets component's location.
func (c *GroundComponent) SetPosition(x, y float64) {
	c.visual.SetPosition(x, y)
	c.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), c.b2Body.GetAngle())
}

// Update component
// func (c *GroundComponent) Update() {
// 	// Static objects rarely have any "dynamics" so update required
// }
