package api

// ITransform represents the transform properties of an INode
type ITransform interface {
	CalcFilteredTransform(
		excludeTranslation bool,
		excludeRotation bool,
		excludeScale bool,
		aft IAffineTransform)

	// AffineTransform returns this node's transform
	AffineTransform() IAffineTransform

	InverseTransform() IAffineTransform

	SetPosition(x, y float64)
	Position() IPoint

	SetRotation(radian float64)
	Rotation() float64

	SetScale(scale float64)
	Scale() float64

	// Not really useful in this engine.
	// SetNonUniformScale(sx, sy float64)
	// NonUniformScale() float64
}
