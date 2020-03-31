package tweening

import (
	"github.com/wdevore/RangerGo/api"
)

// QuadEquation construct
type QuadEquation struct {
	compute api.Compute
}

// NewQuadEquation constructs a simple linear EaseInOut equation
func NewQuadEquation(easeType int) api.ITweenEquation {
	o := new(QuadEquation)

	switch easeType {
	case api.EaseIn:
		o.compute = quadEaseIn
	case api.EaseOut:
		o.compute = quadEaseOut
	case api.EaseInOut:
		o.compute = quadEaseInOut
	}

	return o
}

// Compute performs tweening
func (lt *QuadEquation) Compute(t float64) float64 {
	return lt.compute(t)

}

func quadEaseIn(t float64) float64 {
	return t * t
}

func quadEaseOut(t float64) float64 {
	return -t * (t - 2.0)
}

func quadEaseInOut(t float64) float64 {
	t = t * 2.0
	if t < 1.0 {
		return 0.5 * t * t
	}

	return -0.5 * ((t-1.0)*(t-2.0) - 1.0)
}
