package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// TriangleComponent is a triangle physics object
type TriangleComponent struct {
	visual api.INode
	b2Body *box2d.B2Body

	scale float64
}

// NewTriangleComponent constructs a component
func NewTriangleComponent(name string, parent api.INode) *TriangleComponent {
	o := new(TriangleComponent)
	o.visual = NewTriangleNode(name, parent.World(), parent)
	return o
}

// Configure component
func (t *TriangleComponent) Configure(scale float64, b2World *box2d.B2World) {
	t.scale = scale

	buildTri(t, b2World)
}

// SetColor sets the visual's color
func (t *TriangleComponent) SetColor(color api.IPalette) {
	gr := t.visual.(*TriangleNode)
	gr.SetColor(color)
}

// SetPosition sets component's location.
func (t *TriangleComponent) SetPosition(x, y float64) {
	t.visual.SetPosition(x, y)
	t.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), t.b2Body.GetAngle())
}

// EnableGravity enables/disables gravity for this component
func (t *TriangleComponent) EnableGravity(enable bool) {
	if enable {
		t.b2Body.SetGravityScale(-9.8)
	} else {
		t.b2Body.SetGravityScale(0.0)
	}
}

// Reset configures the component back to defaults
func (t *TriangleComponent) Reset(x, y float64) {
	t.visual.SetPosition(x, y)
	t.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), 0.0)
	t.b2Body.SetLinearVelocity(box2d.MakeB2Vec2(0.0, 0.0))
	t.b2Body.SetAngularVelocity(0.0)
	t.b2Body.SetAwake(true)
}

// ApplyForce applies linear force to comp center
func (t *TriangleComponent) ApplyForce(dirX, dirY float64) {
	t.b2Body.ApplyForce(box2d.B2Vec2{X: dirX, Y: dirY}, t.b2Body.GetWorldCenter(), true)
}

// ApplyImpulse applies linear impulse to box center
func (t *TriangleComponent) ApplyImpulse(dirX, dirY float64) {
	t.b2Body.ApplyLinearImpulse(box2d.B2Vec2{X: dirX, Y: dirY}, t.b2Body.GetWorldCenter(), true)
}

// ApplyImpulseToCorner applies linear impulse to 1,1 comp corner
// As the box rotates the 1,1 corner rotates which means impulses
// could change the rotation to either CW or CCW.
func (t *TriangleComponent) ApplyImpulseToCorner(dirX, dirY float64) {
	t.b2Body.ApplyLinearImpulse(box2d.B2Vec2{X: dirX, Y: dirY}, t.b2Body.GetWorldPoint(box2d.B2Vec2{X: 1.0, Y: 1.0}), true)
}

// ApplyTorque applies torgue to comp center
func (t *TriangleComponent) ApplyTorque(torgue float64) {
	t.b2Body.ApplyTorque(torgue, true)
}

// ApplyAngularImpulse applies angular impulse to box center
func (t *TriangleComponent) ApplyAngularImpulse(impulse float64) {
	t.b2Body.ApplyAngularImpulse(impulse, true)
}

// Update component
func (t *TriangleComponent) Update() {
	if t.b2Body.IsActive() {
		pos := t.b2Body.GetPosition()
		t.visual.SetPosition(pos.X, pos.Y)

		rot := t.b2Body.GetAngle()
		t.visual.SetRotation(rot)
	}
}

func buildTri(t *TriangleComponent, b2World *box2d.B2World) {
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody

	// Note: +Y points down in Ranger verses Upward in Box2D's GUI.
	vertices := []box2d.B2Vec2{}

	// Sync the visual's vertices to this physic object
	tr := t.visual.(*TriangleNode)
	verts := tr.Polygon().Mesh().Vertices()

	vertices = append(vertices, box2d.B2Vec2{X: verts[0].X() * t.scale, Y: verts[0].Y() * t.scale})
	vertices = append(vertices, box2d.B2Vec2{X: verts[1].X() * t.scale, Y: verts[1].Y() * t.scale})
	vertices = append(vertices, box2d.B2Vec2{X: verts[2].X() * t.scale, Y: verts[2].Y() * t.scale})

	// An instance of a body to contain Fixture
	t.b2Body = b2World.CreateBody(&bDef)

	t.visual.SetScale(1.0 * t.scale)
	gb := t.visual.(*TriangleNode)
	gb.SetColor(rendering.NewPaletteInt64(rendering.Orange))

	// Every Fixture has a shape
	b2Shape := box2d.MakeB2PolygonShape()
	b2Shape.Set(vertices, 3)

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2Shape
	fd.Density = 1.0
	t.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}
