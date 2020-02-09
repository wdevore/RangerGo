package lines

import (
	"fmt"
	"testing"

	"github.com/wdevore/RangerGo/engine/geometry"
)

func TestRunner(t *testing.T) {
	runLines(t)
}

func runLines(t *testing.T) {
	l := geometry.NewLine()
	fmt.Println(l)

	l.SetP1(1.0, 2.0)
	p1, p2 := l.Components()
	if p1.X() != 1.0 && p1.Y() != 2.0 {
		t.Fatal("Expected Start point = (1.0, 2.0)")
	}
	if p2.X() != 0.0 && p2.Y() != 0.0 {
		t.Fatal("Expected End point = (0.0, 0.0)")
	}

	l.SetP2(3.0, 4.0)
	p1, p2 = l.Components()
	if p1.X() != 1.0 && p1.Y() != 2.0 {
		t.Fatal("Expected Start point = (1.0, 2.0)")
	}
	if p2.X() != 3.0 && p2.Y() != 4.0 {
		t.Fatal("Expected End point = (3.0, 4.0)")
	}

	l.SetByComp(5.0, 6.0, 7.0, 8.0)
	p1, p2 = l.Components()
	if p1.X() != 5.0 && p1.Y() != 6.0 {
		t.Fatal("Expected Start point = (5.0, 6.0)")
	}
	if p2.X() != 7.0 && p2.Y() != 8.0 {
		t.Fatal("Expected End point = (7.0, 8.0)")
	}

	l2 := geometry.NewLine()
	l2.SetByLine(l)
	p1, p2 = l2.Components()
	if p1.X() != 5.0 && p1.Y() != 6.0 {
		t.Fatal("Expected Start point = (5.0, 6.0)")
	}
	if p2.X() != 7.0 && p2.Y() != 8.0 {
		t.Fatal("Expected End point = (7.0, 8.0)")
	}
}
