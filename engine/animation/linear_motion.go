package animation

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
)

// LinearMotion represents zero dimension interpolations
type LinearMotion struct {
	from, to float64
	angle    float64
	Motion
}

// NewLinearMotion create a motion object
func NewLinearMotion() api.IMotion {
	o := new(LinearMotion)
	o.InitializeMotion()
	return o
}

// Set sets `from` and `to`
func (m *LinearMotion) Set(from, to interface{}) {
	m.from = from.(float64)
	m.to = to.(float64)
}

// Interpolate performs interpolation between `from` and `to`
func (m *LinearMotion) Interpolate(t float64) interface{} {
	m.angle = maths.Lerp(m.from, m.to, t)
	return m.angle
}

// Update sets a new time window that rendering passes will interpolate
// between.
// dt = milliseconds.
// Only used for non-2D motion
func (m *LinearMotion) Update(dt float64) {
	// During each frame the "from" becomes the current "to"
	m.from = m.to

	// "to" is now moved to the next value
	// divide by time_scale
	m.to += m.rate * (dt / m.timeScale)
}
