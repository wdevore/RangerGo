package points

import (
	"fmt"
	"testing"

	"github.com/wdevore/RangerGo/engine/geometry"
)

func TestRunner(t *testing.T) {
	runPoints(t)
}

func runPoints(t *testing.T) {
	p := geometry.NewPoint()
	fmt.Println(p)

	p.SetByComp(1.0, 2.0)
	fmt.Println(p)
	if p.X() != 1.0 && p.Y() == 2.0 {
		t.Fatal("Expected (1.0, 2.0)")
	}

	p2 := geometry.NewPoint()
	fmt.Println(p2)

	p2.SetByPoint(p)
	fmt.Println(p2)
}
