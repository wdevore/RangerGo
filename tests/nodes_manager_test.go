package nodesmanager

import (
	// "fmt"

	"fmt"
	"testing"

	"github.com/wdevore/RangerGo/engine"

	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

func TestRunner(t *testing.T) {
	runBasicScenes(t)
}

func runBasicNodes(t *testing.T) {
	world := engine.NewWorld("NodeMangler")

	nm := nodes.NewNodeManager(world)

	// Note: CrossNode is not a IScene, but I use it to do simple tests
	// of the node manager
	crn := custom.NewCrossNode()
	crn.Initialize("CrossNode")

	nm.PreVisit()

	scenesPresent := nm.Visit(0.0)

	if scenesPresent {
		t.Fatal("Didn't expect any scenes present.")
	}

	nm.PushNode(crn)

	crn2 := custom.NewCrossNode()
	crn2.Initialize("CrossNode2")
	nm.PushNode(crn2)

	crn3 := custom.NewCrossNode()
	crn3.Initialize("CrossNode3")
	nm.PushNode(crn3)

	fmt.Println("Node stack:--------------")
	fmt.Println(nm)
	fmt.Println("-------------------------")

	nm.PopNode()
	fmt.Println("Node stack:--------------")
	fmt.Println(nm)
	fmt.Println("-------------------------")
}

func runBasicScenes(t *testing.T) {
	// This method tests basic nodemanager Scene management.

	world := engine.NewWorld("NodeMangler")

	nm := nodes.NewNodeManager(world)

	splash := custom.NewBasicSplashScene("Splash", nil)

	boot := custom.NewBasicBootScene("Boot", splash)
	nm.PushNode(boot)

	scenesPresent := nm.Visit(0.0)

	if !scenesPresent {
		t.Fatal("Expected scenes present.")
	}

	fmt.Println("Node stack:--------------")
	fmt.Println(nm)
	fmt.Println("-------------------------")

	scenesPresent = nm.Visit(0.0)

	if !scenesPresent {
		t.Fatal("Expected scenes present.")
	}

	fmt.Println("Node stack:--------------")
	fmt.Println(nm)
	fmt.Println("-------------------------")
}
