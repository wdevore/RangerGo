package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

var ranger api.IEngine

func init() {
	world := engine.NewWorld("Zooming", 1.5, "..")

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
