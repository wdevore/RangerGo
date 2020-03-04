package api

// ITransform represents the transform properties of an INode
type ITransform interface {
	CalcFilteredTransform()

	Transform() ITransform

	AffineTransform() IAffineTransform
	InverseTransform() IAffineTransform

	SetPosition(x, y float64)
	Position() IPoint

	SetRotation()
	Rotation() float64

	SetScale()
	Scale() IPoint
	SetNonUniformScale()
}
