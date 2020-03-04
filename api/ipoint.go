package api

// IPoint represents 2D points
type IPoint interface {
	// Components returns the x,y parts
	Components() (float64, float64)

	// ComponentsAsInt32 return x,y parts for render context
	ComponentsAsInt32() (int32, int32)

	// X sets the x component
	X() float64
	// Y sets the y component
	Y() float64
	// SetByComp sets by component
	SetByComp(x, y float64)
	// SetByPoint sets point using another point
	SetByPoint(IPoint)
}
