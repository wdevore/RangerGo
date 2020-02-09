package meshes

import (
	"fmt"
	"testing"

	"github.com/wdevore/RangerGo/engine/geometry"
)

func TestRunner(t *testing.T) {
	runMeshes(t)
}

func runMeshes(t *testing.T) {
	m := geometry.NewMesh()
	fmt.Println(m)

	m.AddVertex(1.0, 2.0)
	fmt.Println(m)

	m.AddVertex(3.0, 4.0)
	fmt.Println(m)

	m.Build()
}
