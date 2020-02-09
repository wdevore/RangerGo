package rectangles

import (
	"testing"

	"github.com/wdevore/RangerGo/engine/geometry"
)

func TestRunner(t *testing.T) {
	runRects(t)
}

func runRects(t *testing.T) {
	intersect := geometry.NewRectangle()
	if intersect.Min().X() != 0.0 || intersect.Min().Y() != 0.0 {
		t.Fatal("Expected Min(0.0, 0.0)")
	}
	if intersect.Max().X() != 1.0 || intersect.Max().Y() != 1.0 {
		t.Fatal("Expected Max(1.0, 1.0)")
	}
	w, h := intersect.Dimesions()
	if w != 1.0 || h != 1.0 {
		t.Fatal("Expected Dimension (1.0 x 1.0)")
	}

	rectA := geometry.NewRectangleUsing(0.0, 0.0, 10.0, 10.0)
	rectB := geometry.NewRectangleUsing(5.0, 5.0, 20.0, 20.0)

	geometry.Intersect(intersect, rectA, rectB)
	if intersect.Min().X() != 5.0 || intersect.Min().Y() != 5.0 {
		t.Fatal("Expected Min(5.0, 5.0)")
	}
	if intersect.Max().X() != 10.0 || intersect.Max().Y() != 10.0 {
		t.Fatal("Expected Max(10.0, 10.0)")
	}
	w, h = intersect.Dimesions()
	if w != 5.0 || h != 5.0 {
		t.Fatal("Expected Dimension (5.0 x 5.0)")
	}

	intersects := rectA.IntersectsOther(rectB)
	if !intersects {
		t.Fatal("Expected rectangles A and B to intersect")
	}

	intersect = geometry.NewRectangle()
	intersects = intersect.IntersectsOther(rectB)
	if intersects {
		t.Fatal("Expected rectangles 'intersect' and B to NOT intersect")
	}

	p := geometry.NewPointUsing(5.0, 5.0)

	contains := geometry.ContainsPointOther(rectA, p)

	if !contains {
		t.Fatal("Expected point in rectangle A")
	}

	p = geometry.NewPointUsing(25.0, 25.0)

	contains = geometry.ContainsPointOther(rectA, p)

	if contains {
		t.Fatal("Expected point NOT in rectangle A")
	}

	bounds := geometry.NewRectangle()
	geometry.Bounds(bounds, rectA, rectB)

	if bounds.Min().X() != 0.0 || bounds.Min().Y() != 0.0 {
		t.Fatal("Expected Min(0.0, 0.0)")
	}
	if bounds.Max().X() != 20.0 || bounds.Max().Y() != 20.0 {
		t.Fatal("Expected Max(20.0, 20.0)")
	}
	w, h = bounds.Dimesions()
	if w != 20.0 || h != 20.0 {
		t.Fatal("Expected Dimension (20.0 x 20.0)")
	}
}
