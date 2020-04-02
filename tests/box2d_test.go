package lerps

import (
	// "fmt"

	"fmt"
	"testing"

	"github.com/ByteArena/box2d"
)

func TestRunner(t *testing.T) {
	// Define the gravity vector.
	gravity := box2d.MakeB2Vec2(0.0, -10.0)

	// Construct a world object, which will hold and simulate the rigid bodies.
	world := box2d.MakeB2World(gravity)

	fmt.Println(world)
}
