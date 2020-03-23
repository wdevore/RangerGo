package api

// IMotion represents animation behaviors
type IMotion interface {
	InitializeMotion()

	Interpolate(t float64) interface{}

	SetRate(rate float64)
	Set(from, to interface{})
	SetAutoWrap(bool)
	SetTimeScale(s float64)

	Update(dt float64)
}
