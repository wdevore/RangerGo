package geometry

import (
	"fmt"
	"math"

	"github.com/wdevore/RangerGo/api"
)

type rectangle struct {
	// Top-left corner
	min api.IPoint
	// Bottom-right corner
	max api.IPoint

	width  float64
	height float64
}

// NewRectangle constructs a new IRectangle
func NewRectangle() api.IRectangle {
	o := new(rectangle)
	o.min = NewPoint()
	o.max = NewPointUsing(1.0, 1.0)
	o.width = 1.0
	o.height = 1.0
	return o
}

// NewRectangleUsing constructs a new IRectangle using components
func NewRectangleUsing(minx, miny, maxx, maxy float64) api.IRectangle {
	o := new(rectangle)
	o.min = NewPointUsing(minx, miny)
	o.max = NewPointUsing(maxx, maxy)
	o.width = maxx - minx
	o.height = maxy - miny
	return o
}

// NewRectangleUsingPoints constructs a new IRectangle using components
func NewRectangleUsingPoints(min, max api.IPoint) api.IRectangle {
	o := new(rectangle)
	o.min = NewPointUsing(min.X(), min.Y())
	o.max = NewPointUsing(max.X(), max.Y())
	o.width = max.X() - min.X()
	o.height = max.Y() - min.Y()
	return o
}

func (r *rectangle) Dimesions() (w, h float64) {
	return r.width, r.height
}

func (r *rectangle) Min() api.IPoint {
	return r.min
}

func (r *rectangle) Max() api.IPoint {
	return r.max
}

func (r *rectangle) Set(minx, miny, maxx, maxy float64) {
	r.min.SetByComp(minx, miny)
	r.max.SetByComp(maxx, maxy)
	r.width = maxx - minx
	r.height = maxy - miny
}

// Intersect computes the intersection of rect A and B
func Intersect(intersect, rectA, rectB api.IRectangle) {
	minA := rectA.Min()
	maxA := rectA.Max()
	minB := rectB.Min()
	maxB := rectB.Max()

	x0 := math.Max(minA.X(), minB.X())
	x1 := math.Min(maxA.X(), maxB.X())

	if x0 <= x1 {
		y0 := math.Max(minA.Y(), minB.Y())
		y1 := math.Min(maxA.Y(), maxB.Y())

		if y0 <= y1 {
			intersect.Set(x0, y0, x1, y1)
		}
	}
}

func (r *rectangle) IntersectsOther(rect api.IRectangle) bool {
	min := rect.Min()
	max := rect.Max()
	w, h := rect.Dimesions()

	return (r.min.X() <= min.X()+w &&
		min.X() <= r.min.X()+r.width &&
		r.max.Y() <= max.Y()+h &&
		max.Y() <= r.max.Y()+r.height)
}

// Intersects indicates if two rectangles intersect
func Intersects(rectA, rectB api.IRectangle) bool {
	minA := rectA.Min()
	maxA := rectA.Max()
	minB := rectB.Min()
	maxB := rectB.Max()
	wa, ha := rectA.Dimesions()
	wb, hb := rectB.Dimesions()

	return (minA.X() <= minB.X()+wb &&
		minB.X() <= minA.X()+wa &&
		maxA.Y() <= maxB.Y()+hb &&
		maxB.Y() <= maxA.Y()+ha)
}

// Bounds computes a new rectangle that completely
// contains A and B
func Bounds(bounds, rectA, rectB api.IRectangle) {
	right := math.Max(rectA.Max().X(), rectB.Max().X())
	bottom := math.Max(rectA.Max().Y(), rectB.Max().Y())
	left := math.Min(rectA.Min().X(), rectB.Min().X())
	top := math.Min(rectA.Min().Y(), rectB.Min().Y())

	bounds.Set(left, top, right, bottom)
}

func (r *rectangle) ContainsPoint(p api.IPoint) bool {
	return p.X() >= r.Min().X() &&
		p.X() <= r.Max().X() &&
		p.Y() >= r.Min().Y() &&
		p.Y() <= r.Max().Y()
}

// ContainsPointOther determines if point is within the rectangle
func ContainsPointOther(rect api.IRectangle, p api.IPoint) bool {
	return rect.ContainsPoint(p)
}

func (r rectangle) String() string {
	return fmt.Sprintf("Min %v, Max %v, %3.3f x %3.3f", r.min, r.max, r.width, r.height)
}
