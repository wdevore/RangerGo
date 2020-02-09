package polygons

import (
	// "fmt"
	"testing"

	"github.com/wdevore/RangerGo/engine/geometry"
)

func TestRunner(t *testing.T) {
	runPolys(t)
}

func runPolys(t *testing.T) {
	p := geometry.NewPolygon()

	p.AddVertex(10.0, 0.0)
	p.AddVertex(10.0, 10.0)
	p.AddVertex(0.0, 10.0)
	p.AddVertex(0.0, 0.0)
	// fmt.Println(p)

	p0 := geometry.NewPointUsing(2.0, 2.0)
	inside := p.PointInside(p0)

	if !inside {
		t.Fatalf("Expected point %v to be inside", p0)
	}

	p0.SetByComp(15.0, 15.0)
	inside = p.PointInside(p0)
	if inside {
		t.Fatalf("Expected point %v to be outside", p0)
	}

	p = geometry.NewPolygon()

	// .--------------->
	// |     5,5
	// |      .-------
	// |      |      |
	// |      |  in  |  out
	// |      |      |
	// |      -------. 15,15
	// |
	// |
	// v
	p.AddVertex(15.0, 5.0)
	p.AddVertex(15.0, 15.0)
	p.AddVertex(5.0, 15.0)
	p.AddVertex(5.0, 5.0)

	p0.SetByComp(4.99, 4.99)
	inside = p.PointInside(p0)
	if inside {
		t.Fatalf("Expected point %v to be outside", p0)
	}

	p0.SetByComp(10.0, 7.0)
	inside = p.PointInside(p0)
	if !inside {
		t.Fatalf("Expected point %v to be inside", p0)
	}

	p0.SetByComp(5.0, 5.0)
	inside = p.PointInside(p0)
	if !inside {
		t.Fatalf("Expected point %v to be inside", p0)
	}

	// 15,15 is on the right and/or bottom edge which is considered outside
	// the left/top edge is considered inside.
	p0.SetByComp(15.0, 15.0)
	inside = p.PointInside(p0)
	if inside {
		t.Fatalf("Expected point %v to be outside", p0)
	}
}
