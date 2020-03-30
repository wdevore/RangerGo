package api

// IParticle represents a particle
type IParticle interface {
	Update(dt float64)

	IsActive() bool
	Activate(bool)
	Reset()

	SetPosition(x, y float64)
	GetPosition() IPoint
	SetLifespan(duration float64)

	SetVelocity(direction, speed float64)

	Visual() INode
}
