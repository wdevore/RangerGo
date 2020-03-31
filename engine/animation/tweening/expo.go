package tweening

import (
	"math"

	"github.com/wdevore/RangerGo/api"
)

// ExpoEquation construct
type ExpoEquation struct {
	compute api.Compute
}

// NewExpoEquation constructs a simple linear EaseInOut equation
func NewExpoEquation(easeType int) api.ITweenEquation {
	o := new(ExpoEquation)

	switch easeType {
	case api.EaseIn:
		o.compute = expoEaseIn
	case api.EaseOut:
		o.compute = expoEaseOut
	case api.EaseInOut:
		o.compute = expoEaseInOut
	}

	return o
}

// Compute performs tweening
func (lt *ExpoEquation) Compute(t float64) float64 {
	return lt.compute(t)

}

func expoEaseIn(t float64) float64 {
	if t == 0.0 {
		return 0.0
	}

	return math.Pow(2.0, (10.0 * (t - 1.0)))
}

func expoEaseOut(t float64) float64 {
	if t == 1.0 {
		return 1.0
	}

	return -math.Pow(2.0, -10.0*t) + 1.0
}

func expoEaseInOut(t float64) float64 {
	if t == 0.0 {
		return 0.0
	}

	if t == 1.0 {
		return 1.0
	}

	t = t * 2.0
	if t < 1.0 {
		return 0.5 * math.Pow(2.0, (10.0*(t-1.0)))
	}

	return 0.5 * (-math.Pow(2.0, (-10.0*(t-1.0))) + 2.0)
}
