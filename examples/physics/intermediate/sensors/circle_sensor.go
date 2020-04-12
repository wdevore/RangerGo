package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// CircleSensor is a box
type CircleSensor struct {
	visual api.INode
	b2Body *box2d.B2Body

	scale float64

	categoryBits uint16 // I am a...
	maskBits     uint16 // I can collide with a...
}

// NewCircleSensor constructs a component
func NewCircleSensor(name string, parent api.INode) *CircleSensor {
	o := new(CircleSensor)
	o.visual = NewCircleNode(name, parent.World(), parent)
	o.visual.SetID(1003)
	cn := o.visual.(*CircleNode)
	cn.SetColor(rendering.NewPaletteInt64(rendering.Blue))
	cn.Configure(12, 1.0)
	return o
}

// Configure component
func (c *CircleSensor) Configure(scale float64, categoryBits, maskBits uint16, b2World *box2d.B2World) {
	c.scale = scale

	c.categoryBits = categoryBits
	c.maskBits = maskBits

	buildCircleSensor(c, b2World)
}

// SetColor sets the visual's color
func (c *CircleSensor) SetColor(color api.IPalette) {
	gr := c.visual.(*CircleNode)
	gr.SetColor(color)
}

// SetPosition sets component's location.
func (c *CircleSensor) SetPosition(x, y float64) {
	c.visual.SetPosition(x, y)
	c.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), c.b2Body.GetAngle())
}

// Update component
func (c *CircleSensor) Update() {
	if c.b2Body.IsActive() {
		pos := c.b2Body.GetPosition()
		c.visual.SetPosition(pos.X, pos.Y)

		rot := c.b2Body.GetAngle()
		c.visual.SetRotation(rot)
	}
}

// ------------------------------------------------------
// Physics feedback
// ------------------------------------------------------

// HandleBeginContact processes BeginContact events
func (c *CircleSensor) HandleBeginContact(nodeA, nodeB api.INode) bool {
	n, ok := nodeA.(*CircleNode)

	if !ok {
		n, ok = nodeB.(*CircleNode)
	}

	if ok {
		n.SetColor(rendering.NewPaletteInt64(rendering.Aqua))
	}

	return false
}

// HandleEndContact processes EndContact events
func (c *CircleSensor) HandleEndContact(nodeA, nodeB api.INode) bool {
	n, ok := nodeA.(*CircleNode)

	if !ok {
		n, ok = nodeB.(*CircleNode)
	}

	if ok {
		n.SetColor(rendering.NewPaletteInt64(rendering.Blue))
	}

	return false
}

func buildCircleSensor(c *CircleSensor, b2World *box2d.B2World) {
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody

	// An instance of a body to contain Fixture
	c.b2Body = b2World.CreateBody(&bDef)

	c.visual.SetScale(1.0 * c.scale)

	// Every Fixture has a shape
	b2Shape := box2d.MakeB2CircleShape()
	b2Shape.SetRadius(c.scale)

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2Shape
	fd.Density = 1.0
	fd.UserData = c.visual
	fd.IsSensor = true

	fd.Filter.CategoryBits = c.categoryBits
	fd.Filter.MaskBits = c.maskBits

	c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}
