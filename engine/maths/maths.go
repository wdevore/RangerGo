package maths

import "math"

var v1 = NewVector()
var v2 = NewVector()
var v3 = NewVector()

const (
	// Epsilon = 0.00001
	Epsilon = 0.00001 // ~32 bits

	// DegreeToRadians converts to radians, for example, 45.0 * DegreeToRadians = radians
	DegreeToRadians = math.Pi / 180.0
)

// Clamp a value between min/max
func Clamp(value, min, max float64) float64 {
	// int value = (value < 0? 0 : value > 255? 255 : value);
	if value < min {
		return min
	} else if value > max {
		return max
	}

	return value
}
