package maths

import (
	"github.com/wdevore/RangerGo/api"
)

type zoomTransform struct {
	// An optional (occasionally used) translation.
	position api.IVector

	// The zoom factor generally incremented in small steps.
	// For example, 0.1
	scale api.IVector

	// The focal point where zooming occurs
	zoomAt api.IVector

	// A "running" accumulating transform
	accTransform api.IAffineTransform

	// A transform that includes position translation.
	transform api.IAffineTransform
}

// NewZoomTransform constructs an zoom transform object
func NewZoomTransform() api.IZoomTransform {
	o := new(zoomTransform)
	o.position = NewVector()
	o.scale = NewVectorUsing(1.0, 1.0)
	o.zoomAt = NewVector()
	o.accTransform = NewTransform()
	o.transform = NewTransform()
	return o
}

// ----------------------------------------------------------
// Methods
// ----------------------------------------------------------

func (zt *zoomTransform) GetTransform() api.IAffineTransform {
	zt.Update()
	return zt.transform
}

func (zt *zoomTransform) Update() {
	// Accumulate zoom transformations.
	// acc_transform is an intermediate accumulative matrix used for tracking the current zoom target.
	zt.accTransform.Translate(zt.zoomAt.X(), zt.zoomAt.Y())
	zt.accTransform.Scale(zt.scale.X(), zt.scale.Y())
	zt.accTransform.Translate(-zt.zoomAt.X(), -zt.zoomAt.Y())

	// We reset Scale because acc_transform is accumulative and has "captured" the information.
	zt.scale.SetByComp(1.0, 1.0)

	// We want to leave acc_transform solely responsible for zooming.
	// "transform" is the final matrix.
	zt.transform.SetByTransform(zt.accTransform)

	// Tack on translation. Note: we don't append it, but concat it into a separate matrix.
	zt.transform.Translate(zt.position.X(), zt.position.Y())
}

func (zt *zoomTransform) SetPosition(x, y float64) {
	zt.position.SetByComp(x, y)
}

func (zt *zoomTransform) ZoomBy(dx, dy float64) {
	zt.scale.Add(dx, dy)
}

func (zt *zoomTransform) TranslateBy(dx, dy float64) {
	zt.position.Add(dx, dy)
}

func (zt *zoomTransform) Scale() float64 {
	return zt.scale.X()
}

func (zt *zoomTransform) PsuedoScale() float64 {
	return zt.accTransform.GetPsuedoScale()
}

func (zt *zoomTransform) SetScale(scale float64) {
	zt.Update()

	// We use dimensional analysis to set the scale. Remember we can't
	// just set the scale absolutely because acc_transform is an accumulating matrix.
	// We have to take its current value and compute a new value based
	// on the passed in value.

	// Also, I can use acc_transform.a because I don't allow rotations for zooms,
	// so the diagonal components correctly represent the matrix's current scale.
	// And because I only perform uniform scaling I can safely use just the "a" element.
	scaleFactor := scale / zt.accTransform.GetPsuedoScale()

	zt.scale.SetByComp(scaleFactor, scaleFactor)
}

func (zt *zoomTransform) SetAt(x, y float64) {
	zt.zoomAt.SetByComp(x, y)
}
