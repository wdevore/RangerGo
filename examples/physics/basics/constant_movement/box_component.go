package main

import (
	"fmt"
	"math"

	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// BoxComponent is a box
type BoxComponent struct {
	visual api.INode
	b2Body *box2d.B2Body

	scale float64

	algorithm int
	stopping  bool
}

// NewBoxComponent constructs a component
func NewBoxComponent(name string, parent api.INode) *BoxComponent {
	o := new(BoxComponent)
	o.visual = custom.NewRectangleNode(name, parent.World(), parent)
	o.algorithm = 1
	return o
}

// Configure component
func (b *BoxComponent) Configure(scale float64, b2World *box2d.B2World) {
	b.scale = scale

	buildBox(b, b2World)
}

// SetColor sets the visual's color
func (b *BoxComponent) SetColor(color api.IPalette) {
	gr := b.visual.(*custom.RectangleNode)
	gr.SetColor(color)
}

// SetPosition sets component's location.
func (b *BoxComponent) SetPosition(x, y float64) {
	b.visual.SetPosition(x, y)
	b.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), b.b2Body.GetAngle())
}

// EnableGravity enables/disables gravity for this component
func (b *BoxComponent) EnableGravity(enable bool) {
	if enable {
		b.b2Body.SetGravityScale(-9.8)
	} else {
		b.b2Body.SetGravityScale(0.0)
	}
}

// Reset configures the component back to defaults
func (b *BoxComponent) Reset(x, y float64) {
	b.visual.SetPosition(x, y)
	b.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), 0.0)
	b.b2Body.SetLinearVelocity(box2d.MakeB2Vec2(0.0, 0.0))
	b.b2Body.SetAngularVelocity(0.0)
	b.b2Body.SetAwake(true)
}

// SetVelocityAlgo changes the velocity style
func (b *BoxComponent) SetVelocityAlgo(style int) {
	b.algorithm = style
}

// Stop set linear velocity to 0
func (b *BoxComponent) Stop() {
	b.stopping = !b.stopping
	if b.stopping {
		fmt.Println("Stopping...")
	}
}

// MoveLeft applies linear force to box center
func (b *BoxComponent) MoveLeft(dx float64) {
	velocity := b.b2Body.GetLinearVelocity()
	switch b.algorithm {
	case 1:
		velocity.X -= dx
	case 2:
		velocity.X = math.Max(velocity.X-0.1, -5.0)
	}
	b.b2Body.SetLinearVelocity(velocity)
}

// MoveRight applies linear force to box center
func (b *BoxComponent) MoveRight(dx float64) {
	velocity := b.b2Body.GetLinearVelocity()
	switch b.algorithm {
	case 1:
		velocity.X += dx
	case 2:
		velocity.X = math.Max(velocity.X+0.1, 5.0)
	}
	b.b2Body.SetLinearVelocity(velocity)
}

// MoveUp applies linear force to box center
func (b *BoxComponent) MoveUp(dy float64) {
	velocity := b.b2Body.GetLinearVelocity()
	switch b.algorithm {
	case 1:
		velocity.Y -= dy
	case 2:
		velocity.Y = math.Max(velocity.Y-0.1, -5.0)
	}
	b.b2Body.SetLinearVelocity(velocity)
}

// MoveDown applies linear force to box center
func (b *BoxComponent) MoveDown(dy float64) {
	velocity := b.b2Body.GetLinearVelocity()
	switch b.algorithm {
	case 1:
		velocity.Y += dy
	case 2:
		velocity.Y = math.Max(velocity.Y+0.1, 5.0)
	}
	b.b2Body.SetLinearVelocity(velocity)
}

// ApplyForce applies linear force to box center
func (b *BoxComponent) ApplyForce(dirX, dirY float64) {
	b.b2Body.ApplyForce(box2d.B2Vec2{X: dirX, Y: dirY}, b.b2Body.GetWorldCenter(), true)
}

// ApplyImpulse applies linear impulse to box center
func (b *BoxComponent) ApplyImpulse(dirX, dirY float64) {
	b.b2Body.ApplyLinearImpulse(box2d.B2Vec2{X: dirX, Y: dirY}, b.b2Body.GetWorldCenter(), true)
}

// ApplyImpulseToCorner applies linear impulse to 1,1 box corner
// As the box rotates the 1,1 corner rotates which means impulses
// could change the rotation to either CW or CCW.
func (b *BoxComponent) ApplyImpulseToCorner(dirX, dirY float64) {
	b.b2Body.ApplyLinearImpulse(box2d.B2Vec2{X: dirX, Y: dirY}, b.b2Body.GetWorldPoint(box2d.B2Vec2{X: 1.0, Y: 1.0}), true)
}

// ApplyTorque applies torgue to box center
func (b *BoxComponent) ApplyTorque(torgue float64) {
	b.b2Body.ApplyTorque(torgue, true)
}

// ApplyAngularImpulse applies angular impulse to box center
func (b *BoxComponent) ApplyAngularImpulse(impulse float64) {
	b.b2Body.ApplyAngularImpulse(impulse, true)
}

// Update component
func (b *BoxComponent) Update() {
	if b.b2Body.IsActive() {
		pos := b.b2Body.GetPosition()
		b.visual.SetPosition(pos.X, pos.Y)

		rot := b.b2Body.GetAngle()
		b.visual.SetRotation(rot)
	}

	if b.stopping {
		velocity := b.b2Body.GetLinearVelocity()
		switch b.algorithm {
		case 1: // hard
			velocity.X = 0
			velocity.Y = 0
		case 2: // soft
			velocity.X *= 0.98
			velocity.Y *= 0.98
		}
		b.b2Body.SetLinearVelocity(velocity)
	}
}

func buildBox(b *BoxComponent, b2World *box2d.B2World) {
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody

	// An instance of a body to contain Fixture
	b.b2Body = b2World.CreateBody(&bDef)

	b.visual.SetScale(2.0 * b.scale)
	gb := b.visual.(*custom.RectangleNode)
	gb.SetColor(rendering.NewPaletteInt64(rendering.Orange))

	// Every Fixture has a shape
	b2Shape := box2d.MakeB2PolygonShape()
	b2Shape.SetAsBoxFromCenterAndAngle(1.0*b.scale, 1.0*b.scale, box2d.B2Vec2{X: b.visual.Position().X(), Y: b.visual.Position().Y()}, 0.0)

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2Shape
	fd.Density = 1.0
	b.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}
