package particles

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
)

// Particle is the base object of a particle system.
type Particle struct {
	elapsed  float64
	lifespan float64
	position api.IPoint
	velocity api.IVelocity

	active bool

	// Visual representation
	node api.INode
}

// NewParticle constructs a new particle
func NewParticle(visual api.INode) api.IParticle {
	o := new(Particle)

	o.active = false
	o.elapsed = 0.0
	o.lifespan = 0.0
	o.velocity = maths.NewVelocity()
	o.position = geometry.NewPoint()

	o.node = visual

	return o
}

// SetPosition sets particles initial position
func (p *Particle) SetPosition(x, y float64) {
	p.position.SetByComp(x, y)
	p.node.SetPosition(x, y)
}

// GetPosition gets the particle's current position
func (p *Particle) GetPosition() api.IPoint {
	return p.position
}

// SetLifespan sets how long the particle lives
func (p *Particle) SetLifespan(duration float64) {
	p.lifespan = duration
}

// Visual gets the current INode assigned to this particle
func (p *Particle) Visual() api.INode {
	return p.node
}

// Activate changes the particle's state
func (p *Particle) Activate(active bool) {
	p.active = active
	p.node.SetVisible(active)
}

// IsActive indicates if the particle is alive
func (p *Particle) IsActive() bool {
	return p.active
}

// SetVelocity changes the velocity
func (p *Particle) SetVelocity(angle, speed float64) {
	p.velocity.SetDirectionByAngle(angle)
	p.velocity.SetMagnitude(speed)
}

// Update changes the particle's state based on time
func (p *Particle) Update(dt float64) {
	p.elapsed += dt

	p.active = p.elapsed < p.lifespan

	// Update particle's position as long as the particle is active.
	if p.active {
		p.velocity.ApplyToPoint(p.position)
		p.node.SetPosition(p.position.X(), p.position.Y())
	}
}

// Reset resets the particle
func (p *Particle) Reset() {
	p.active = false
	p.elapsed = 0.0
	p.node.SetVisible(p.active)
}
