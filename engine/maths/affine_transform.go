package maths

import (
	"fmt"
	"math"

	"github.com/wdevore/RangerGo/api"
)

// A minified affine transform.
//  column major (form used by this class)
//     x'   |a c tx|   |x|
//     y' = |b d ty| x |y|               <=== Post multiply
//     1    |0 0  1|   |1|
//  or
//  Row major
//                           |a  b   0|
//     |x' y' 1| = |x y 1| x |c  d   0|  <=== Pre multiply
//                           |tx ty  1|

type affineTransform struct {
	a, b, c, d float64
	tx, ty     float64
}

// NewTransform constructs an Identity Affine Transform matrix
func NewTransform() api.IAffineTransform {
	o := new(affineTransform)
	o.ToIdentity()
	return o
}

// ----------------------------------------------------------
// Methods
// ----------------------------------------------------------

func (at *affineTransform) ToIdentity() {
	at.a = 1.0
	at.d = 1.0
	at.b = 0.0
	at.c = 0.0
	at.tx = 0.0
	at.ty = 0.0
}

func (at *affineTransform) Components() (a, b, c, d, tx, ty float64) {
	return at.a, at.b, at.c, at.d, at.tx, at.ty
}

func (at *affineTransform) SetByComp(a, b, c, d, tx, ty float64) {
	at.a = a
	at.b = b
	at.c = c
	at.d = d
	at.tx = tx
	at.ty = ty
}

func (at *affineTransform) SetByTransform(t api.IAffineTransform) {
	at.a, at.b, at.c, at.d, at.tx, at.ty = t.Components()
}

func (at *affineTransform) TransformPoint(p api.IPoint) {
	p.SetByComp(
		(at.a*p.X())+(at.c*p.Y())+at.tx,
		(at.b*p.X())+(at.d*p.Y())+at.ty)
}

func (at *affineTransform) TransformToPoint(in api.IPoint, out api.IPoint) {
	out.SetByComp(
		(at.a*in.X())+(at.c*in.Y())+at.tx,
		(at.b*in.X())+(at.d*in.Y())+at.ty)
}

func (at *affineTransform) TransformCompToPoint(x, y float64, out api.IPoint) {
	out.SetByComp(
		(at.a*x)+(at.c*y)+at.tx,
		(at.b*x)+(at.d*y)+at.ty)
}

func (at *affineTransform) TransformToComps(in api.IPoint) (x, y float64) {
	return (at.a * in.X()) + (at.c * in.Y()) + at.tx, (at.b * in.X()) + (at.d * in.Y()) + at.ty
}

func (at *affineTransform) Translate(x, y float64) {
	at.tx += (at.a * x) + (at.c * y)
	at.ty += (at.b * x) + (at.d * y)
}

func (at *affineTransform) MakeTranslate(x, y float64) {
	at.SetByComp(1.0, 0.0, 0.0, 1.0, x, y)
}

func (at *affineTransform) MakeTranslateUsingPoint(p api.IPoint) {
	at.SetByComp(1.0, 0.0, 0.0, 1.0, p.X(), p.Y())
}

func (at *affineTransform) Scale(sx, sy float64) {
	at.a *= sx
	at.b *= sx
	at.c *= sy
	at.d *= sy
}

func (at *affineTransform) MakeScale(sx, sy float64) {
	at.SetByComp(sx, 0.0, 0.0, sy, 0.0, 0.0)
}

func (at *affineTransform) GetPsuedoScale() float64 {
	return at.a
}

// Concatinate a rotation (radians) onto this transform.
//
// Rotation is just a matter of perspective. A CW rotation can be seen as
// CCW depending on what you are talking about rotating. For example,
// if the coordinate system is thought as rotating CCW then objects are
// seen as rotating CW, and that is what the 2x2 matrix below represents.
//
// It is also the frame of reference we use. In this library +Y axis is downward
//     |cos  -sin|   object appears to rotate CW.
//     |sin   cos|
//
// In the matrix below the object appears to rotate CCW.
//     |cos  sin|
//     |-sin cos|
//
//     |a  c|    |cos  -sin|
//     |b  d|  x |sin   cos|
//
// If Y axis is downward (default for SDL and Image) then:
// +angle yields a CW rotation
// -angle yeilds a CCW rotation.
//
// else
// -angle yields a CW rotation
// +angle yeilds a CCW rotation.
func (at *affineTransform) Rotate(radians float64) {
	rsin := math.Sin(radians)
	rcos := math.Cos(radians)
	a := at.a
	b := at.b
	c := at.c
	d := at.d

	at.a = a*rcos + c*rsin
	at.b = b*rcos + d*rsin
	at.c = c*rcos - a*rsin
	at.d = d*rcos - b*rsin
}

func (at *affineTransform) MakeRotate(radians float64) {
	rsin := math.Sin(radians)
	rcos := math.Cos(radians)
	at.a = rcos
	at.b = rsin
	at.c = -rsin
	at.d = rcos
	at.tx = 0
	at.ty = 0
}

// MultiplyPre performs: n = n * m
func MultiplyPre(m api.IAffineTransform, n api.IAffineTransform) {
	na, nb, nc, nd, _, _ := n.Components()
	ma, mb, mc, md, mtx, mty := m.Components()

	n.SetByComp(
		na*ma+nb*mc,
		na*mb+nb*md,
		nc*ma+nd*mc,
		nc*mb+nd*md,
		(na*mtx)+(nc*mty)+mtx,
		(nb*mtx)+(nd*mty)+mty)
}

// MultiplyPost performs: n = m * n
func MultiplyPost(m api.IAffineTransform, n api.IAffineTransform) {
	na, nb, nc, nd, ntx, nty := n.Components()
	ma, mb, mc, md, _, _ := m.Components()

	n.SetByComp(
		ma*na+mb*nc,
		ma*nb+mb*nd,
		mc*na+md*nc,
		mc*nb+md*nd,
		(ma*ntx)+(mc*nty)+ntx,
		(mb*ntx)+(md*nty)+nty)
}

// Multiply performs: out = m * n
func Multiply(m api.IAffineTransform, n api.IAffineTransform, out api.IAffineTransform) {
	na, nb, nc, nd, ntx, nty := n.Components()
	ma, mb, mc, md, mtx, mty := m.Components()

	out.SetByComp(
		ma*na+mb*nc,
		ma*nb+mb*nd,
		mc*na+md*nc,
		mc*nb+md*nd,
		(mtx*na)+(mty*nc)+ntx,
		(mtx*nb)+(mty*nd)+nty)
}

func (at *affineTransform) Invert() {
	a := at.a
	b := at.b
	c := at.c
	d := at.d
	tx := at.tx
	ty := at.ty

	determinant := 1.0 / (a*d - b*c)

	at.a = determinant * d
	at.b = -determinant * b
	at.c = -determinant * c
	at.d = determinant * a
	at.tx = determinant * (c*ty - d*tx)
	at.ty = determinant * (b*tx - a*ty)
}

func (at *affineTransform) InvertTo(out api.IAffineTransform) {
	determinant := 1.0 / (at.a*at.d - at.b*at.c)
	out.SetByComp(
		determinant*at.d,
		-determinant*at.b,
		-determinant*at.c,
		determinant*at.a,
		determinant*(at.c*at.ty-at.d*at.tx),
		determinant*(at.b*at.tx-at.a*at.ty))
}

func (at *affineTransform) Transpose() {
	c := at.c
	at.c = at.b
	at.b = c
}

func (at affineTransform) String() string {
	return fmt.Sprintf("|%3.3f,%3.3f,%3.3f|\n|%3.3f,%3.3f,%3.3f|", at.a, at.b, at.tx, at.c, at.d, at.ty)
}
