package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// CircleComponent is a box
type CircleComponent struct {
	visual api.INode
	b2Body *box2d.B2Body

	scale float64

	categoryBits uint16 // I am a...
	maskBits     uint16 // I can collide with a...
}

// NewCircleComponent constructs a component
func NewCircleComponent(name string, parent api.INode) *CircleComponent {
	o := new(CircleComponent)
	o.visual = NewCircleNode(name, parent.World(), parent)
	o.visual.SetID(1003)
	cn := o.visual.(*CircleNode)
	cn.Configure(12, 1.0)
	return o
}

// Configure component
func (c *CircleComponent) Configure(scale float64, categoryBits, maskBits uint16, b2World *box2d.B2World) {
	c.scale = scale

	c.categoryBits = categoryBits
	c.maskBits = maskBits

	buildCircle(c, b2World)
}

// SetColor sets the visual's color
func (c *CircleComponent) SetColor(color api.IPalette) {
	gr := c.visual.(*CircleNode)
	gr.SetColor(color)
}

// SetPosition sets component's location.
func (c *CircleComponent) SetPosition(x, y float64) {
	c.visual.SetPosition(x, y)
	c.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), c.b2Body.GetAngle())
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

// ------------------------------------------------------
// Physics control
// ------------------------------------------------------

// EnableGravity enables/disables gravity for this component
func (c *CircleComponent) EnableGravity(enable bool) {
	if enable {
		c.b2Body.SetGravityScale(-9.8)
	} else {
		c.b2Body.SetGravityScale(0.0)
	}
}

// Reset configures the component back to defaults
func (c *CircleComponent) Reset(x, y float64) {
	c.visual.SetPosition(x, y)
	c.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), 0.0)
	c.b2Body.SetLinearVelocity(box2d.MakeB2Vec2(0.0, 0.0))
	c.b2Body.SetAngularVelocity(0.0)
	c.b2Body.SetAwake(true)
}

// ApplyForce applies linear force to box center
func (c *CircleComponent) ApplyForce(dirX, dirY float64) {
	c.b2Body.ApplyForce(box2d.B2Vec2{X: dirX, Y: dirY}, c.b2Body.GetWorldCenter(), true)
}

// ApplyImpulse applies linear impulse to box center
func (c *CircleComponent) ApplyImpulse(dirX, dirY float64) {
	c.b2Body.ApplyLinearImpulse(box2d.B2Vec2{X: dirX, Y: dirY}, c.b2Body.GetWorldCenter(), true)
}

// ApplyImpulseToCorner applies linear impulse to 1,1 box corner
// As the box rotates the 1,1 corner rotates which means impulses
// could change the rotation to either CW or CCW.
func (c *CircleComponent) ApplyImpulseToCorner(dirX, dirY float64) {
	c.b2Body.ApplyLinearImpulse(box2d.B2Vec2{X: dirX, Y: dirY}, c.b2Body.GetWorldPoint(box2d.B2Vec2{X: 1.0, Y: 1.0}), true)
}

// ApplyTorque applies torgue to box center
func (c *CircleComponent) ApplyTorque(torgue float64) {
	c.b2Body.ApplyTorque(torgue, true)
}

// ApplyAngularImpulse applies angular impulse to box center
func (c *CircleComponent) ApplyAngularImpulse(impulse float64) {
	c.b2Body.ApplyAngularImpulse(impulse, true)
}

// ------------------------------------------------------
// Physics feedback
// ------------------------------------------------------

// HandleBeginContact processes BeginContact events
func (c *CircleComponent) HandleBeginContact(nodeA, nodeB api.INode) bool {
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
func (c *CircleComponent) HandleEndContact(nodeA, nodeB api.INode) bool {
	n, ok := nodeA.(*CircleNode)

	if !ok {
		n, ok = nodeB.(*CircleNode)
	}

	if ok {
		n.SetColor(rendering.NewPaletteInt64(rendering.Orange))
	}

	return false
}

func buildCircle(c *CircleComponent, b2World *box2d.B2World) {
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody

	// An instance of a body to contain Fixture
	c.b2Body = b2World.CreateBody(&bDef)

	c.visual.SetScale(1.0 * c.scale)
	gb := c.visual.(*CircleNode)
	gb.SetColor(rendering.NewPaletteInt64(rendering.Orange))

	// Every Fixture has a shape
	b2Shape := box2d.MakeB2CircleShape()
	b2Shape.SetRadius(c.scale)

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2Shape
	fd.Density = 1.0
	fd.UserData = c.visual

	fd.Filter.CategoryBits = c.categoryBits
	fd.Filter.MaskBits = c.maskBits

	c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}
