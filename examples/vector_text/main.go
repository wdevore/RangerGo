package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

var ranger api.IEngine

// This example shows Vector text on the splash screen

func init() {
	world := engine.NewWorld("Vector Text")

	ranger = engine.New(world)

	splash := newBasicSplashScene("Splash", nil)
	splash.Build(world)
	// nodes.PrintTree(splash)

	// This example use the super basic Boot scene that doesn't absolutely nothing.
	boot := custom.NewBasicBootScene("Boot", splash)

	ranger.PushStart(boot)
}

func main() {
	ranger.Configure()

	ranger.Start()

	ranger.End()
}
