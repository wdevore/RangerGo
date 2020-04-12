package main

import (
	"math"

	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// RadarSensor is a box
type RadarSensor struct {
	visual api.INode
	b2Body *box2d.B2Body

	scale float64

	categoryBits uint16 // I am a...
	maskBits     uint16 // I can collide with a...
}

// NewRadarSensor constructs a component
func NewRadarSensor(name string, parent api.INode) *RadarSensor {
	o := new(RadarSensor)
	o.visual = NewPieNode(name, parent.World(), parent)
	o.visual.SetID(1003)
	cn := o.visual.(*PieNode)
	cn.SetColor(rendering.NewPaletteInt64(rendering.Yellow))
	cn.Configure(6, 1.0, 0.0, math.Pi/4.0)
	return o
}

// Configure component
func (c *RadarSensor) Configure(scale float64, categoryBits, maskBits uint16, b2World *box2d.B2World) {
	c.scale = scale

	c.categoryBits = categoryBits
	c.maskBits = maskBits

	buildRadarSensor(c, b2World)
}

// SetColor sets the visual's color
func (c *RadarSensor) SetColor(color api.IPalette) {
	gr := c.visual.(*PieNode)
	gr.SetColor(color)
}

// SetPosition sets component's location.
func (c *RadarSensor) SetPosition(x, y float64) {
	c.visual.SetPosition(x, y)
	c.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), c.b2Body.GetAngle())
}

// Update component
func (c *RadarSensor) Update() {
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
func (c *RadarSensor) HandleBeginContact(nodeA, nodeB api.INode) bool {
	n, ok := nodeA.(*PieNode)

	if !ok {
		n, ok = nodeB.(*PieNode)
	}

	if ok {
		n.SetColor(rendering.NewPaletteInt64(rendering.Aqua))
	}

	return false
}

// HandleEndContact processes EndContact events
func (c *RadarSensor) HandleEndContact(nodeA, nodeB api.INode) bool {
	n, ok := nodeA.(*PieNode)

	if !ok {
		n, ok = nodeB.(*PieNode)
	}

	if ok {
		n.SetColor(rendering.NewPaletteInt64(rendering.Yellow))
	}

	return false
}

func buildRadarSensor(c *RadarSensor, b2World *box2d.B2World) {
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody

	// An instance of a body to contain Fixture
	c.b2Body = b2World.CreateBody(&bDef)
	c.b2Body.SetAngularVelocity(45.0 * maths.DegreeToRadians)

	c.visual.SetScale(1.0 * c.scale)

	// Note: +Y points down in Ranger verses Upward in Box2D's GUI.
	vertices := []box2d.B2Vec2{}
	vertices = append(vertices, box2d.B2Vec2{X: 0.0, Y: 0.0})
	segments := 7.0
	for i := 0.0; i < segments; i++ {
		angle := i / 6.0 * (math.Pi / 4.0)
		vertices = append(vertices, box2d.B2Vec2{X: math.Cos(angle) * c.scale, Y: math.Sin(angle) * c.scale})
	}

	// Every Fixture has a shape
	b2Shape := box2d.MakeB2PolygonShape()
	b2Shape.Set(vertices, int(segments)+1)

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2Shape
	// The fixture can't have any density otherwise
	// it would swing away.
	// fd.Density = 1.0
	fd.UserData = c.visual
	fd.IsSensor = true

	fd.Filter.CategoryBits = c.categoryBits
	fd.Filter.MaskBits = c.maskBits

	c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}
