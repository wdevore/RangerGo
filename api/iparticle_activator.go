package api

const (
	// MaxParticleLifetime is a default lifetime
	MaxParticleLifetime = 1.0
	// MaxParticleSpeed is a good default starting value
	MaxParticleSpeed = 20.0
)

// IParticleActivator activates particles
type IParticleActivator interface {
	Activate(particle IParticle, center IPoint)
}
