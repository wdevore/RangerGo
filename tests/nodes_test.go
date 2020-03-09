package nodes

import (
	// "fmt"

	"fmt"
	"testing"

	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

func TestRunner(t *testing.T) {
	runBasic(t)
}

func runBasic(t *testing.T) {
	crn := custom.NewCrossNode()
	crn.Initialize("CrossNode")
	fmt.Println(crn)
	if crn.ID() != 0 {
		t.Logf("Node ID wrong")
	}

	crn2 := custom.NewCrossNode()
	crn2.Initialize("CrossNode2")
	fmt.Println(crn2)
	if crn2.ID() != 1 {
		t.Fatal("Node ID wrong")
	}

	crn.AddChild(crn2)

	crn3 := custom.NewCrossNode()
	crn3.Initialize("CrossNode3")
	crn.AddChild(crn3)

	nodes.PrintTree(crn)

	crn.SetPosition(1.0, 2.0)

	crn.Draw(nil)

}
