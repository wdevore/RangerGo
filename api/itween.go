package api

const (
	// EaseNoMeaning indicates that some equations don't have a meaningful
	// value, for example, Linear
	EaseNoMeaning = -1
	// EaseIn eases in
	EaseIn = 0
	// EaseOut eases out
	EaseOut = 1
	// EaseInOut eases both directions
	EaseInOut = 2

	// EquationLinear Linear equations
	EquationLinear = 0
	// EquationExpo Exponential equations
	EquationExpo = 1
	// EquationQuad Quadratic equations
	EquationQuad = 2
)

// Compute computes time
type Compute func(float64) float64

// ITweenEquation represents tween equation
type ITweenEquation interface {
	Compute(float64) float64
}

// ITween represent tween behaviours
type ITween interface {
	Update(dt float64) (value float64, isFinished bool)

	Elapsed() float64

	Reset()
}
