package api

// IPoint represents 2D points
type IPoint interface {
	Components() (float64, float64)
	// X sets the x component
	X() float64
	// Y sets the y component
	Y() float64
	// SetByComp sets by component
	SetByComp(x, y float64)
	// SetByPoint sets point using another point
	SetByPoint(IPoint)
}
