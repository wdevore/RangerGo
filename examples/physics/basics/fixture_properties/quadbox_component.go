package main

import (
	"fmt"
	"math"

	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// QuadBoxComponent is a bunch of boxes.
type QuadBoxComponent struct {
	anchor api.INode // The parent
	// boxVisual api.INode // The parent of the component

	b2Body *box2d.B2Body

	scale float64
}

// NewQuadBoxComponent constructs a component
func NewQuadBoxComponent(name string, parent api.INode) *QuadBoxComponent {
	o := new(QuadBoxComponent)
	o.anchor = custom.NewAnchorNode("QuadBox", parent.World(), parent)
	return o
}

// Configure component
func (q *QuadBoxComponent) Configure(scale float64, b2World *box2d.B2World) {
	q.scale = scale

	buildQuad(q, b2World)
}

// SetPosition sets component's location.
func (q *QuadBoxComponent) SetPosition(x, y float64) {
	q.anchor.SetPosition(x, y)
	q.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), q.b2Body.GetAngle())
}

// Reset configures the component back to defaults
func (q *QuadBoxComponent) Reset(x, y float64) {
	q.anchor.SetPosition(x, y)
	q.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), 0.0)
	q.b2Body.SetLinearVelocity(box2d.MakeB2Vec2(0.0, 0.0))
	q.b2Body.SetAngularVelocity(0.0)
	// If you waited until the topple completed then the body went to sleep.
	// Thus we need to wake it back up otherwise it would just hang in
	// mid air.
	q.b2Body.SetAwake(true)
}

// Update component
func (q *QuadBoxComponent) Update() {
	if q.b2Body.IsActive() {
		pos := q.b2Body.GetPosition()
		q.anchor.SetPosition(pos.X, pos.Y)

		rot := q.b2Body.GetAngle()
		q.anchor.SetRotation(rot)
	}
}

// SetDensity changes the density of all boxes
func (q *QuadBoxComponent) SetDensity(value float64) {
	for fix := q.b2Body.GetFixtureList(); fix != nil; fix = fix.GetNext() {
		fix.SetDensity(value)
	}
	q.b2Body.ResetMassData()
}

// SetFriction changes the friction of all boxes
func (q *QuadBoxComponent) SetFriction(value float64) {
	for fix := q.b2Body.GetFixtureList(); fix != nil; fix = fix.GetNext() {
		fix.SetFriction(value)
	}
}

// SetRestitution changes the restitution of all boxes
func (q *QuadBoxComponent) SetRestitution(value float64) {
	for fix := q.b2Body.GetFixtureList(); fix != nil; fix = fix.GetNext() {
		fix.SetRestitution(value)
	}
}

// ResetFixtures reset all fixtures
func (q *QuadBoxComponent) ResetFixtures(den, fric, rest float64) {
	for fix := q.b2Body.GetFixtureList(); fix != nil; fix = fix.GetNext() {
		fix.SetDensity(den)
		fix.SetFriction(fric)
		fix.SetRestitution(rest)
	}
	q.b2Body.ResetMassData()
}

func buildQuad(c *QuadBoxComponent, b2World *box2d.B2World) {
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody

	// An instance of a body to contain Fixtures
	c.b2Body = b2World.CreateBody(&bDef)

	// ------------------------------------------------
	// Boxes
	// ------------------------------------------------
	for i := 0.0; i < 4.0; i++ {
		b := custom.NewRectangleNode(fmt.Sprintf("::Box%d", int(i)), c.anchor.World(), c.anchor)
		b.SetPosition(math.Sin(i*maths.DegreeToRadians*90.0)*c.scale, math.Cos(i*maths.DegreeToRadians*90.0)*c.scale)
		b.SetScale(2.0 * c.scale)
		gb := b.(*custom.RectangleNode)
		gb.SetColor(rendering.NewPaletteInt64(rendering.Orange))
		c.anchor.AddChild(b)

		// Every Fixture has a shape
		b2Shape := box2d.MakeB2PolygonShape()
		b2Shape.SetAsBoxFromCenterAndAngle(1.0*c.scale, 1.0*c.scale, box2d.B2Vec2{X: b.Position().X(), Y: b.Position().Y()}, 0.0)

		fd := box2d.MakeB2FixtureDef()
		fd.Shape = &b2Shape
		fd.Density = 1.0
		fd.Friction = i / 4.0
		fd.Restitution = (4.0 - i) / 8.0
		c.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
	}

}
