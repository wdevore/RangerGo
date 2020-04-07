package main

import (
	"fmt"
	"math"

	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// TrackingComponent is a box
type TrackingComponent struct {
	visual api.INode
	b2Body *box2d.B2Body

	scale float64

	algorithm         int
	trackingAlgorithm int
	stopping          bool
	targetingRate     float64
	targetPosition    api.IPoint
}

// NewTrackingComponent constructs a component
func NewTrackingComponent(name string, parent api.INode) *TrackingComponent {
	o := new(TrackingComponent)
	o.visual = NewTriangleNode(name, parent.World(), parent)
	o.algorithm = 1
	o.trackingAlgorithm = 1
	o.targetingRate = 60.0 // default is slow
	return o
}

// Configure component
func (b *TrackingComponent) Configure(scale float64, b2World *box2d.B2World) {
	b.scale = scale

	b.targetPosition = geometry.NewPoint()

	buildComp(b, b2World)
}

// SetColor sets the visual's color
func (b *TrackingComponent) SetColor(color api.IPalette) {
	gr := b.visual.(*custom.RectangleNode)
	gr.SetColor(color)
}

// SetTargetPosition sets the target position. The position is in
// view-space.
func (b *TrackingComponent) SetTargetPosition(p api.IPoint) {
	b.targetPosition.SetByPoint(p)

	// angle := b.b2Body.GetAngle()
	// ray := box2d.MakeB2Vec2(b.targetPosition.X(), b.targetPosition.Y())

	// Subtract 90 if Triangle defaults to pointing at +X
	// desiredAngle := math.Atan2(-ray.X, ray.Y) - maths.DegreeToRadians*-90.0

	// desiredAngle := math.Atan2(-ray.X, ray.Y)
	// b.b2Body.SetTransform(b.b2Body.GetPosition(), desiredAngle)
}

// SetPosition sets component's location.
func (b *TrackingComponent) SetPosition(x, y float64) {
	b.visual.SetPosition(x, y)
	b.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), b.b2Body.GetAngle())
}

// GetPosition returns the body's position
func (b *TrackingComponent) GetPosition() box2d.B2Vec2 {
	return b.b2Body.GetPosition()
}

// EnableGravity enables/disables gravity for this component
func (b *TrackingComponent) EnableGravity(enable bool) {
	if enable {
		b.b2Body.SetGravityScale(-9.8)
	} else {
		b.b2Body.SetGravityScale(0.0)
	}
}

// Reset configures the component back to defaults
func (b *TrackingComponent) Reset(x, y float64) {
	b.visual.SetPosition(x, y)
	b.b2Body.SetTransform(box2d.MakeB2Vec2(x, y), 0.0)
	b.b2Body.SetLinearVelocity(box2d.MakeB2Vec2(0.0, 0.0))
	b.b2Body.SetAngularVelocity(0.0)
	b.b2Body.SetAwake(true)
}

// SetTrackingAlgo changes the tracking style
func (b *TrackingComponent) SetTrackingAlgo(style int) {
	b.trackingAlgorithm = style
}

// SetTargetingRate changes the how fast the tracking locks on.
func (b *TrackingComponent) SetTargetingRate(rate float64) {
	b.targetingRate = rate
}

// SetVelocityAlgo changes the velocity style
func (b *TrackingComponent) SetVelocityAlgo(style int) {
	b.algorithm = style
}

// Stop set linear velocity to 0
func (b *TrackingComponent) Stop() {
	b.stopping = !b.stopping
	if b.stopping {
		fmt.Println("Stopping...")
	}
}

// MoveLeft applies linear force to box center
func (b *TrackingComponent) MoveLeft(dx float64) {
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
func (b *TrackingComponent) MoveRight(dx float64) {
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
func (b *TrackingComponent) MoveUp(dy float64) {
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
func (b *TrackingComponent) MoveDown(dy float64) {
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
func (b *TrackingComponent) ApplyForce(dirX, dirY float64) {
	b.b2Body.ApplyForce(box2d.B2Vec2{X: dirX, Y: dirY}, b.b2Body.GetWorldCenter(), true)
}

// ApplyImpulse applies linear impulse to box center
func (b *TrackingComponent) ApplyImpulse(dirX, dirY float64) {
	b.b2Body.ApplyLinearImpulse(box2d.B2Vec2{X: dirX, Y: dirY}, b.b2Body.GetWorldCenter(), true)
}

// ApplyImpulseToCorner applies linear impulse to 1,1 box corner
// As the box rotates the 1,1 corner rotates which means impulses
// could change the rotation to either CW or CCW.
func (b *TrackingComponent) ApplyImpulseToCorner(dirX, dirY float64) {
	b.b2Body.ApplyLinearImpulse(box2d.B2Vec2{X: dirX, Y: dirY}, b.b2Body.GetWorldPoint(box2d.B2Vec2{X: 1.0, Y: 1.0}), true)
}

// ApplyTorque applies torgue to box center
func (b *TrackingComponent) ApplyTorque(torgue float64) {
	b.b2Body.ApplyTorque(torgue, true)
}

// ApplyAngularImpulse applies angular impulse to box center
func (b *TrackingComponent) ApplyAngularImpulse(impulse float64) {
	b.b2Body.ApplyAngularImpulse(impulse, true)
}

// Update component
func (b *TrackingComponent) Update() {
	if b.b2Body.IsActive() {
		pos := b.b2Body.GetPosition()
		b.visual.SetPosition(pos.X, pos.Y)

		rot := b.b2Body.GetAngle()
		b.visual.SetRotation(rot)
	}

	bodyPos := b.b2Body.GetPosition()
	ray := box2d.MakeB2Vec2(b.targetPosition.X()-bodyPos.X, b.targetPosition.Y()-bodyPos.Y)
	targetAngle := math.Atan2(-ray.X, ray.Y)

	switch b.trackingAlgorithm {
	case 1:
		// Instant targeting.
		b.b2Body.SetTransform(b.b2Body.GetPosition(), targetAngle)
	case 2:
		// This is slow and constantly overshoots
		// Torque unchecked
		totalRotation := targetAngle - b.b2Body.GetAngle()
		if totalRotation < -180.0*maths.DegreeToRadians {
			totalRotation += 360.0 * maths.DegreeToRadians
		}
		if totalRotation > 180.0*maths.DegreeToRadians {
			totalRotation -= 360.0 * maths.DegreeToRadians
		}
		torque := 10.0
		if totalRotation < 0 {
			torque = -10.0
		}
		b.b2Body.ApplyTorque(torque, true)
	case 3:
		// This is slow but eventually locks on.
		// look ahead more than one time step to adjust the rate
		// at which the correct angle is reached. Here I'm looking ahead
		// 1/3 second forward in time.
		nextAngle := b.b2Body.GetAngle() + b.b2Body.GetAngularVelocity()/b.targetingRate
		totalRotation := targetAngle - nextAngle
		if totalRotation < -180.0*maths.DegreeToRadians {
			totalRotation += 360.0 * maths.DegreeToRadians
		}
		if totalRotation > 180.0*maths.DegreeToRadians {
			totalRotation -= 360.0 * maths.DegreeToRadians
		}
		torque := 10.0
		if totalRotation < 0 {
			torque = -10.0
		}
		b.b2Body.ApplyTorque(torque, true)
	case 4:
		// This locks on instantly but includes some jittering in the
		// last few frames.
		nextAngle := b.b2Body.GetAngle() + b.b2Body.GetAngularVelocity()/b.targetingRate
		totalRotation := targetAngle - nextAngle
		if totalRotation < -180.0*maths.DegreeToRadians {
			totalRotation += 360.0 * maths.DegreeToRadians
		}
		if totalRotation > 180.0*maths.DegreeToRadians {
			totalRotation -= 360.0 * maths.DegreeToRadians
		}
		desiredAngularVelocity := totalRotation * b.targetingRate
		torque := b.b2Body.GetInertia() * desiredAngularVelocity / (1 / b.targetingRate)
		b.b2Body.ApplyTorque(torque, true)
	case 5:
		// This is equivalent as case 4 but without the "time" element
		// This locks on instantly but includes some jittering in the
		// last few frames.
		nextAngle := b.b2Body.GetAngle() + b.b2Body.GetAngularVelocity()/b.targetingRate
		totalRotation := targetAngle - nextAngle
		if totalRotation < -180.0*maths.DegreeToRadians {
			totalRotation += 360.0 * maths.DegreeToRadians
		}
		if totalRotation > 180.0*maths.DegreeToRadians {
			totalRotation -= 360.0 * maths.DegreeToRadians
		}
		desiredAngularVelocity := totalRotation * b.targetingRate
		impulse := b.b2Body.GetInertia() * desiredAngularVelocity
		b.b2Body.ApplyAngularImpulse(impulse, true)
	case 6:
		nextAngle := b.b2Body.GetAngle() + b.b2Body.GetAngularVelocity()/b.targetingRate
		totalRotation := targetAngle - nextAngle
		if totalRotation < -180.0*maths.DegreeToRadians {
			totalRotation += 360.0 * maths.DegreeToRadians
		}
		if totalRotation > 180.0*maths.DegreeToRadians {
			totalRotation -= 360.0 * maths.DegreeToRadians
		}
		desiredAngularVelocity := totalRotation * b.targetingRate
		change := 5.0 * maths.DegreeToRadians //allow 1 degree rotation per time step
		desiredAngularVelocity = math.Min(change, math.Max(-change, desiredAngularVelocity))
		impulse := b.b2Body.GetInertia() * desiredAngularVelocity
		b.b2Body.ApplyAngularImpulse(impulse, true)
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

func buildComp(b *TrackingComponent, b2World *box2d.B2World) {
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_dynamicBody

	// Note: +Y points down in Ranger verses Upward in Box2D's GUI.
	vertices := []box2d.B2Vec2{}
	vertices = append(vertices, box2d.B2Vec2{X: -0.5 * b.scale, Y: -0.5 * b.scale})
	vertices = append(vertices, box2d.B2Vec2{X: 0.0 * b.scale, Y: 0.5 * b.scale})
	vertices = append(vertices, box2d.B2Vec2{X: 0.5 * b.scale, Y: -0.5 * b.scale})

	// An instance of a body to contain Fixture
	b.b2Body = b2World.CreateBody(&bDef)

	b.visual.SetScale(1.0 * b.scale)
	gb := b.visual.(*TriangleNode)
	gb.SetColor(rendering.NewPaletteInt64(rendering.Orange))

	// Every Fixture has a shape
	b2Shape := box2d.MakeB2PolygonShape()
	b2Shape.Set(vertices, 3)

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2Shape
	fd.Density = 1.0
	b.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}
