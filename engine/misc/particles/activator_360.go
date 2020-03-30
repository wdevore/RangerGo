package particles

import (
	"math/rand"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
)

// Activator360 activates particles in random directions encompassing
// 360 degrees
type Activator360 struct {
	maxLife  float64
	maxSpeed float64
}

// NewActivator360 constructs an activator
func NewActivator360() api.IParticleActivator {
	o := new(Activator360)
	o.maxLife = api.MaxParticleLifetime
	o.maxSpeed = api.MaxParticleSpeed
	return o
}

// Activate configures a particle with a random direction and speed.
func (a *Activator360) Activate(particle api.IParticle, center api.IPoint) {
	direction := rand.Float64() * 360.0
	speed := rand.Float64() * a.maxSpeed

	particle.SetVelocity(direction, speed)

	// The location of where the particle is emitted
	particle.SetPosition(center.X(), center.Y())

	// A random lifetime ranging from 0.0 to max_life
	lifespan := rand.Float64() * (a.maxLife * 1000.0)
	particle.SetLifespan(lifespan)

	color, isColorType := particle.Visual().(api.IColor)

	if isColorType {
		// Change the Red color component if the visual supports the IColor type
		shade := int(maths.Clamp(rand.Float64()*32.0*speed, 0, 255))
		color.SetRed(shade)
	}

	particle.Reset()

	particle.Activate(true)
}
