package api

// IParticleSystem represents a particle system
type IParticleSystem interface {
	AddParticle(particle IParticle)
	Update(dt float64)

	SetPosition(x, y float64)
	SetAutoTrigger(bool)
	Activate(bool)

	TriggerOneshot()
	TriggerAt(pos IPoint)
	TriggerExplosion()
}
