package affinetransforms

import (
	"fmt"
	"math"
	"testing"

	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
)

func TestRunner(t *testing.T) {
	runAFs(t)
}

func runAFs(t *testing.T) {
	af := maths.NewTransform()
	fmt.Println(af)

	fmt.Println("")
	af.SetByComp(1.1, 2.2, 3.3, 4.4, 5.5, 6.6)
	fmt.Println(af)

	fmt.Println("")
	af.ToIdentity()
	fmt.Println(af)

	fmt.Println("")
	af.MakeTranslate(5.0, 10.0)
	fmt.Println(af)

	// Transform
	// (1.000000,2.000000)  <-- from
	// (6.000000,12.000000) <-- to
	p := geometry.NewPointUsing(1.0, 2.0)

	af.TransformPoint(p)
	// fmt.Println(p)
	if p.X() != 6.0 && p.Y() == 12.0 {
		msg := fmt.Sprintf("Expected (6.0, 12.0) Got: %0.8f,%0.8f", p.X(), p.Y())
		t.Fatal(msg)
	}

	// SDL 2D coordinate space is:
	// .------------> +X
	// |
	// |
	// |
	// |
	// |
	// v +Y

	// Thus a +Radian rotation is equal to a CW rotation
	// Rotate +X basis vector CW
	xv := geometry.NewPointUsing(1.0, 0.0)
	rott := maths.NewTransform()

	rott.MakeRotate(maths.DegreeToRadians * 45.0)
	rott.TransformPoint(xv)
	// fmt.Println(xv)
	epx := math.Abs(0.707107 - xv.X())
	epy := math.Abs(0.707107 - xv.Y())
	if epx > maths.Epsilon || epy > maths.Epsilon {
		msg := fmt.Sprintf("Expected (0.707107, 0.707107) Got: %0.8f,%0.8f", xv.X(), xv.Y())
		t.Fatal(msg)
	}

	xv.SetByComp(1.0, 0.0)
	rott.MakeRotate(maths.DegreeToRadians * 90.0)
	rott.TransformPoint(xv)
	// fmt.Println(xv)
	epx = math.Abs(0.0 - xv.X())
	epy = math.Abs(1.0 - xv.Y())
	if epx > maths.Epsilon || epy > maths.Epsilon {
		msg := fmt.Sprintf("Expected (0.0, 1.0) Got: %0.8f,%0.8f", xv.X(), xv.Y())
		t.Fatal(msg)
	}

}
