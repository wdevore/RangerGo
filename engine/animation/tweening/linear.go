package tweening

import "github.com/wdevore/RangerGo/api"

// LinearEquation construct
type LinearEquation struct {
	compute api.Compute
}

// NewLinearEquation constructs a simple linear EaseInOut equation
func NewLinearEquation(easeType int) api.ITweenEquation {
	o := new(LinearEquation)

	o.compute = linearEaseInOut

	return o
}

// Compute performs tweening
func (lt *LinearEquation) Compute(t float64) float64 {
	return lt.compute(t)

}

func linearEaseInOut(t float64) float64 {
	return t
}
