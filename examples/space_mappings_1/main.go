package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

var ranger api.IEngine

// ##################################################################
// README
// This example shows basic mapping of mouse coordinates to view-space.
// View-space is 1 level below device-space.
// Note: scene-space, layer-space (aka node-spaces) are different than
// view-space.
// ##################################################################

func init() {
	world := engine.NewWorld("Space mappings #1", 1.5, "..")

	ranger = engine.New(world)

	splash := newBasicSplashScene("Splash", nil)
	splash.Build(world)

	// This example uses the super basic Boot scene that does absolutely nothing.
	boot := custom.NewBasicBootScene("Boot", splash)

	nodes.PrintTree(splash)

	ranger.PushStart(boot)
}

func main() {
	ranger.Configure()

	ranger.Start()

	ranger.End()
}
