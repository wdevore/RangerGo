package api

// IVector represents 2D vectors
// that have direction and magnitude
type IVector interface {
	Components() (float64, float64)
	// X sets the x component
	X() float64
	// Y sets the y component
	Y() float64
	// SetByComp sets by component
	SetByComp(x, y float64)
	// SetByPoint sets point using another point
	SetByPoint(IPoint)
	// SetByVector sets point using another vector
	SetByVector(IVector)

	SetByAngle(radians float64)

	// Length returns the square root length
	Length() float64
	// LengthSqr return the squared length
	LengthSqr() float64

	// Scale vector by s
	Scale(s float64)

	// Add offsets a this vector
	Add(x, y float64)
	// Sub offsets a this vector
	Sub(x, y float64)

	AddV(IVector)
	SubV(IVector)

	// Div vector by d
	Div(d float64)

	// AngleX computes the angle (radians) between this vector and v.
	AngleX(v IVector) float64

	// Normalize normalizes this vector, if the vector is zero
	// then nothing happens
	Normalize()

	// SetDirection set the direction of the vector, however,
	// it will erase the magnitude
	SetDirection(radians float64)

	// CrossCW computes the cross-product faster in the CW direction
	CrossCW()
	// CrossCCW computes the cross-product faster in the CCW direction
	CrossCCW()
}
