package tweening

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/maths"
)

// ##################################################################
// Note: This was/is an experiment to see if I could implement
// basic easing animation sand understand them--mission sort of accomplished.
//
// Nonetheless, I recommend you use : github.com/tanema/gween instead as
// is it more feature complete with more easings.
// ##################################################################

// Tween performs tween animations on values
type Tween struct {
	equation api.ITweenEquation

	// Parameters
	duration float64
	begin    float64
	end      float64

	// Controls
	elapsed   float64
	change    float64
	accumTime float64
}

// NewTween constructs a new tweener
func NewTween(begin, end, duration float64, equation, style int) api.ITween {
	o := new(Tween)
	o.duration = duration
	o.begin = begin
	o.end = end
	o.change = end - begin

	switch equation {
	case api.EquationLinear:
		o.equation = NewLinearEquation(style)
	case api.EquationExpo:
		o.equation = NewExpoEquation(style)
	case api.EquationQuad:
		o.equation = NewQuadEquation(style)
	}

	return o
}

// Update interpolates values at each dt (ms) = ms-per-frame ~= 33.333ms
func (t *Tween) Update(dt float64) (value float64, isFinished bool) {
	step := dt / t.duration

	t.elapsed += dt

	t.accumTime += step

	interpolation := maths.Clamp(t.equation.Compute(t.accumTime), 0.0, 1.0)

	// fmt.Println(t.accumTime, " ", step, " ", t.elapsed, " ", interpolation)

	isFinished = t.elapsed >= t.duration

	value = interpolation*t.change + t.begin

	return value, isFinished
}

// Elapsed returns how much time has passed since animation began.
// It should eventually equal duration.
func (t *Tween) Elapsed() float64 {
	return t.elapsed
}

// Reset tween
func (t *Tween) Reset() {
	t.accumTime = 0.0
	t.elapsed = 0.0
}
