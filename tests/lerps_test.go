package lerps

import (
	// "fmt"

	"fmt"
	"testing"

	"github.com/wdevore/RangerGo/engine/maths"
)

func TestRunner(t *testing.T) {
	runLerps(t)
	runLinears(t)
	runLerpVectors(t)
}

func runLerps(t *testing.T) {
	lp := maths.Lerp(2.0, 4.0, 0.5)

	if lp != 3.0 {
		t.Fatalf("Expected 3.0")
	}

	lp = maths.Lerp(2.0, 4.0, 0.0)

	if lp != 2.0 {
		t.Fatalf("Expected 2.0")
	}

	lp = maths.Lerp(2.0, 4.0, 1.0)
	if lp != 4.0 {
		t.Fatalf("Expected 4.0")
	}

	lp = maths.Lerp(2.0, 4.0, 0.25)
	if lp != 2.5 {
		t.Fatalf("Expected 2.5")
	}
}

func runLinears(t *testing.T) {
	lp := maths.Linear(2.0, 4.0, 3.0)
	if lp != 0.5 {
		t.Fatalf("Expected 0.5")
	}

	lp = maths.Linear(2.0, 4.0, 2.5)
	if lp != 0.25 {
		t.Fatalf("Expected 0.25")
	}

	lp = maths.Linear(-2.0, 2.0, 0.0)
	if lp != 0.5 {
		t.Fatalf("Expected 0.5")
	}

	lp = maths.Linear(-2.0, 2.0, 0.0)
	if lp != 0.5 {
		t.Fatalf("Expected 0.5")
	}

	lp = maths.Linear(-2.0, 2.0, -2.0)
	if lp != 0.0 {
		t.Fatalf("Expected 0.0")
	}

	lp = maths.Linear(-2.0, 2.0, 2.0)
	if lp != 1.0 {
		t.Fatalf("Expected 1.0")
	}

	lp = maths.Linear(-2.0, -1.0, -2.0)
	if lp != 0.0 {
		t.Fatalf("Expected 0.0")
	}

	lp = maths.Linear(-2.0, -1.0, -1.0)
	if lp != 1.0 {
		t.Fatalf("Expected 1.0")
	}

	lp = maths.Linear(-2.0, -1.0, -1.5)
	if lp != 0.5 {
		t.Fatalf("Expected 0.5")
	}

	lp = maths.Linear(-1.0, -2.0, -1.5)
	if lp != 0.5 {
		t.Fatalf("Expected 0.5")
	}
}

func runLerpVectors(t *testing.T) {
	v1 := maths.NewVectorUsing(0.0, 0.0)
	v2 := maths.NewVectorUsing(1.0, 1.0)
	o := maths.NewVector()

	maths.LerpVectors(v1, v2, o, 0.5)
	if o.X() != 0.5 && o.Y() != 0.5 {
		t.Fatalf("Expected (0.5, 0.5)")
	}

	maths.LerpVectors(v1, v2, o, 1.0)
	if o.X() != 1.0 && o.Y() != 1.0 {
		t.Fatalf("Expected (1.0, 1.0)")
	}
	fmt.Println(o)
}
