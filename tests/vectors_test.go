package vectors

import (
	"fmt"
	"math"
	"testing"

	"github.com/wdevore/RangerGo/engine/maths"
)

func TestRunner(t *testing.T) {
	runVectors(t)
}

func runVectors(t *testing.T) {
	v := maths.NewVector()
	fmt.Println(v)

	v1 := maths.NewVector()
	v2 := maths.NewVectorUsing(1.0, 1.0)
	distance := maths.Distance(v1, v2)

	ep := math.Abs(1.41421 - distance)
	if ep > maths.Epsilon {
		msg := fmt.Sprintf("Expected 1.414213 Got: %0.8f", distance)
		t.Fatal(msg)
	}

	v3 := maths.NewVectorUsing(1.0, 0.0)
	v4 := maths.NewVectorUsing(1.0, 1.0)

	angle := v3.AngleX(v4) // / maths.DegreeToRadians

	ep = math.Abs(0.7853981 - angle) // = 45 degrees
	if ep > maths.Epsilon {
		msg := fmt.Sprintf("Expected 0.7853981 radians Got: %0.8f", angle)
		t.Fatal(msg)
	}

	v5 := maths.NewVectorUsing(2.0, 2.0)
	v5.Normalize()
	epx := math.Abs(0.707106 - v5.X())
	epy := math.Abs(0.707106 - v5.Y())
	if epx > maths.Epsilon || epy > maths.Epsilon {
		msg := fmt.Sprintf("Expected 0.707106,0.707106  Got: %v", v5)
		t.Fatal(msg)
	}

	v6 := maths.NewVectorUsing(1.0, 0.0)
	v7 := maths.NewVectorUsing(0.0, 1.0)

	angle = maths.Angle(v6, v7)
	ep = math.Abs(1.5707963 - angle)
	if ep > maths.Epsilon {
		msg := fmt.Sprintf("Expected 1.5707963  Got: %0.8f", angle)
		t.Fatal(msg)
	}

	v7 = maths.NewVectorUsing(-1.0, 0.0)

	angle = maths.Angle(v6, v7)
	// fmt.Println(angle) // / maths.DegreeToRadians
	ep = math.Abs(3.141592 - angle)
	if ep > maths.Epsilon {
		msg := fmt.Sprintf("Expected 3.141592  Got: %0.8f", angle)
		t.Fatal(msg)
	}
}
