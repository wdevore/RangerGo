package api

// IVelocity represents the direction and magnitude of a vector
// Velocity's direction is alway defined relative to the +X axis.
// Default direction is +X axis.
type IVelocity interface {
	// SetMin sets minimum magnitude
	SetMin(float64)
	// SetMax sets maximum magnitude
	SetMax(float64)
	// SetMinMax sets min/maximum magnitude
	SetMinMax(float64, float64)
	// SetMagnitude set the current magnitude directly
	SetMagnitude(float64)
	// Magnitude returns the current magnitude
	Magnitude() float64
	// Range returns the min/max magnitude range
	Range() (float64, float64)
	// SetDirection set the direction vector component
	SetDirectionByVector(IVector)
	SetDirectionByAngle(radians float64)

	// Direction returns the current direction component
	Direction() IVector

	// ConstrainMagnitude enables/disables the limiting of magnitude
	// to within the min/max range
	ConstrainMagnitude(bool)

	ApplyToPoint(point IPoint)
}
