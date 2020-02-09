package api

// ILine represents 2D lines
type ILine interface {
	Components() (IPoint, IPoint)
	// SetP1 sets Start end point
	SetP1(float64, float64)
	// SetP2 sets End end point
	SetP2(float64, float64)
	// SetByComp sets by component
	SetByComp(x1, y1, x2, y2 float64)
	// SetByPoint sets point using another point
	SetByLine(ILine)
}
