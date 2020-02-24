package api

// IZoomTransform represents 2D zoom transform
type IZoomTransform interface {
	// GetTransform updates and returns the internal transform.
	GetTransform() IAffineTransform

	// Update modifies the internal transform state based on current values.
	Update()

	// SetPosition is an absolute position. Typically you would use TranslateBy.
	SetPosition(x, y float64)

	// SetScale sets the scale based on the current scale value making
	// this a relative scale.
	SetScale(scale float64)

	// SetAt sets the center zoom point.
	SetAt(x, y float64)

	// ZoomBy performs a relative zoom based on the current scale/zoom.
	ZoomBy(dx, dy float64)

	// TranslateBy is a relative positional translation.
	TranslateBy(dx, dy float64)
}
