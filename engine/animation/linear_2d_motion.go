package animation

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
)

// Linear2DMotion represents vector interpolations
type Linear2DMotion struct {
	from, to api.IVector
	p        api.IVector

	Motion
}

// NewLinear2DMotion create a motion object
func NewLinear2DMotion() api.IMotion {
	o := new(Linear2DMotion)
	o.InitializeMotion()
	o.p = maths.NewVector()
	return o
}

// Set sets `from` and `to`
func (m *Linear2DMotion) Set(from, to interface{}) {
	m.from = from.(api.IVector)
	m.to = to.(api.IVector)
}

// Interpolate performs interpolation between `from` and `to`
func (m *Linear2DMotion) Interpolate(t float64) interface{} {
	maths.LerpVectors(m.from, m.to, m.p, t)
	return m.p
}

// Update sets a new time window that rendering passes will interpolate
// between.
// dt = milliseconds.
// Only used for non-2D motion
// ###### This may not be correct AND hasn't been tested #######
func (m *Linear2DMotion) Update(dt float64) {
	// During each frame the "from" becomes the current "to"
	m.from = m.to

	// "to" is now moved to the next value
	// divide by time_scale
	n := m.rate * (dt / m.timeScale)
	m.to.Add(n*m.to.X(), n*m.to.Y())
}
