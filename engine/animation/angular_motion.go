package animation

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
)

// AngularMotion represents rotational motion
type AngularMotion struct {
	from, to float64
	angle    float64
	Motion
}

// NewAngularMotion create an rotational motion object
func NewAngularMotion() api.IMotion {
	o := new(AngularMotion)
	o.InitializeMotion()
	return o
}

// Set sets `from` and `to`
func (m *AngularMotion) Set(from, to interface{}) {
	m.from = from.(float64)
	m.to = to.(float64)
}

// Interpolate performs interpolation between `from` and `to`
func (m *AngularMotion) Interpolate(t float64) interface{} {
	m.angle = maths.Lerp(m.from, m.to, t)

	if m.autoWrap {
		if m.rate > 0.0 {
			if m.angle >= 360 {
				// Wrap range back around
				m.from = 0.0
				m.to = 361.0 - m.angle
				// Calc new angle from the adjusted range
				m.angle = maths.Lerp(m.from, m.to, t)
			}
		} else {
			if m.angle <= 0.0 {
				m.from = 360.0
				m.to = 358.0 - m.angle
				m.angle = maths.Lerp(m.from, m.to, t)
			}
		}
	}

	return m.angle
}

// Update sets a new time window that rendering passes will interpolate
// between.
// dt = milliseconds.
// Only used for non-2D motion
func (m *AngularMotion) Update(dt float64) {
	// During each frame the "from" becomes the current "to"
	m.from = m.to

	// "to" is now moved to the next value
	// divide by time_scale
	m.to += m.rate * (dt / m.timeScale)
}
