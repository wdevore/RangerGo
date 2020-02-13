package maths

import "github.com/wdevore/RangerGo/api"

// Lerp returns a the value between min and max given t = 0->1
func Lerp(min, max, t float64) float64 {
	return min*(1.0-t) + max*t
}

var v1 = NewVector()
var v2 = NewVector()

// LerpVectors lerps two vectors into the "out" vector.
// As per:
// https://gamedev.stackexchange.com/questions/18615/how-do-i-linearly-interpolate-between-two-vectors
func LerpVectors(min, max, out api.IVector, t float64) {
	ScaleBy(min, (1.0 - t), v1)
	ScaleBy(max, t, v2)
	Add(v1, v2, out)
}

// Linear returns 0->1 for a "value" between min and max.
// Generally used to map from view-space to unit-space
func Linear(min, max, value float64) float64 {
	if max < min {
		tmp := max
		max = min
		min = tmp
	}

	if min < 0.0 {
		return 1.0 - (value-max)/(min-max)
	}

	return (value - min) / (max - min)
}
