package maths

import (
	"fmt"
	"math"

	"github.com/wdevore/RangerGo/api"
)

type vector struct {
	x, y float64
}

// NewVector constructs a new IVector
func NewVector() api.IVector {
	o := new(vector)
	return o
}

// NewVectorUsing constructs a new IVector using components
func NewVectorUsing(x, y float64) api.IVector {
	o := new(vector)
	o.x = x
	o.y = y
	return o
}

func (v *vector) Components() (x, y float64) {
	return v.x, v.y
}

func (v *vector) X() float64 {
	return v.x
}

func (v *vector) Y() float64 {
	return v.y
}

func (v *vector) SetByComp(x, y float64) {
	v.x = x
	v.y = y
}

func (v *vector) SetByPoint(ip api.IPoint) {
	v.x = ip.X()
	v.y = ip.Y()
}

func (v *vector) SetByVector(ip api.IVector) {
	v.x = ip.X()
	v.y = ip.Y()
}

func (v *vector) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *vector) LengthSqr() float64 {
	return v.x*v.x + v.y*v.y
}

// Add performs: out = v1 + v2
func Add(v1, v2, out api.IVector) {
	out.SetByComp(v1.X()+v2.X(), v1.Y()+v2.Y())
}

// Sub performs: out = v1 - v2
func Sub(v1, v2, out api.IVector) {
	out.SetByComp(v1.X()-v2.X(), v1.Y()-v2.Y())
}

func (v *vector) Scale(s float64) {
	v.x = v.x * s
	v.y = v.y * s
}

// ScaleBy performs: out = v * s
func ScaleBy(v api.IVector, s float64, out api.IVector) {
	out.SetByComp(v.X()*s, v.Y()*s)
}

func (v *vector) Div(d float64) {
	v.x = v.x / d
	v.y = v.y / d
}

var tmp = NewVector()

// Distance between two vectors
func Distance(v1, v2 api.IVector) float64 {
	Sub(v1, v2, tmp)
	return tmp.Length()
}

func (v *vector) AngleX(vo api.IVector) float64 {
	return math.Atan2(vo.Y(), vo.X())
}

func (v *vector) Normalize() {
	len := v.Length()
	if len != 0.0 {
		v.Div(len)
	}
}

// Dot computes the dot-product between the vectors
func Dot(v1, v2 api.IVector) float64 {
	return v1.X()*v2.X() + v1.Y()*v2.Y()
}

// Cross computes the cross-product of two vectors
func Cross(v1, v2 api.IVector) float64 {
	return v1.X()*v2.Y() - v1.Y()*v2.X()
}

var tmp2 = NewVector()

// Angle computes the angle in radians between two vector directions
func Angle(v1, v2 api.IVector) float64 {
	tmp.SetByVector(v1)
	tmp2.SetByVector(v2)

	tmp.Normalize()  // a2
	tmp2.Normalize() // b2

	angle := math.Atan2(Cross(tmp, tmp2), Dot(tmp, tmp2))

	if math.Abs(angle) < Epsilon {
		return 0.0
	}

	return angle
}

func (v vector) String() string {
	return fmt.Sprintf("<%0.3f,%0.3f>", v.x, v.y)
}
