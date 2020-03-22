package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

var ranger api.IEngine

// ##################################################################
// README
// This example show mapping from mouse-space to node-space where the
// node is a rotated rectangle.
// Also, the background has been moved to the scene simply to make
// the layer less clutter for this example.
// ##################################################################

func init() {
	world := engine.NewWorld("Space mappings #2")

	ranger = engine.New(world)

	splash := newBasicSplashScene("Splash", nil)
	splash.Build(world)

	// This example uses the super basic Boot scene that does absolutely nothing.
	boot := custom.NewBasicBootScene("Boot", splash)

	// nodes.PrintTree(splash)

	ranger.PushStart(boot)
}

func main() {
	ranger.Configure()

	ranger.Start()

	ranger.End()
}
