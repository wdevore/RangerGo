package api

// IAffineTransform represents 2D transforms
type IAffineTransform interface {
	Components() (float64, float64, float64, float64, float64, float64)
	// ToIdentity sets the transform to an identity matrix
	ToIdentity()

	// --------------------------------------------
	// Setters
	// --------------------------------------------

	// SetByComp sets by component
	SetByComp(float64, float64, float64, float64, float64, float64)
	// SetByTransform sets point using another transform
	SetByTransform(IAffineTransform)

	// --------------------------------------------
	// Transforms
	// --------------------------------------------

	// TransformPoint applys affine transform to point
	TransformPoint(IPoint)
	// TransformToPoint applys affine transform to out point, "in" is not modified
	TransformToPoint(in IPoint, out IPoint)
	// TransformToComps applys transform and returns results, "in" is not modified
	TransformToComps(in IPoint) (x, y float64)
	// TransformPolygon

	// --------------------------------------------
	// Mutaters
	// --------------------------------------------

	// MakeTranslate sets the transform to a Translate matrix
	MakeTranslate(x, y float64)
	// Translate mutates "this" matrix using tx,ty
	Translate(tx, ty float64)

	// MakeScale sets the transform to a Scale matrix
	MakeScale(x, y float64)
	// Scale mutates "this" matrix using sx, sy
	Scale(sx, sy float64)

	// MakeRotate sets the transform to a Rotate matrix
	MakeRotate(radians float64)
	// Rotate mutates "this" matrix using radian angle
	Rotate(radians float64)

	// --------------------------------------------
	// Inversions
	// --------------------------------------------

	// Invert (mutates) inverts "this" matrix
	Invert()
	// Invert (non-mutating) inverts "this" matrix and sends to "out"
	InvertTo(out IAffineTransform)
	// Transpose
	// Converts either from or to pre or post multiplication.
	//     a c
	//     b d
	// to
	//     a b
	//     c d
	Transpose()
}
