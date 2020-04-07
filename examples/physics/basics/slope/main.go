package main

import (
	"github.com/wdevore/RangerGo/engine"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

func main() {
	world := engine.NewWorld("Basic Slope", 0.25, "../../..")

	ranger := engine.New(world)

	splash := newBasicSplashScene("Splash", nil)
	splash.Build(world)

	boot := custom.NewBasicBootScene("Boot", splash)

	nodes.PrintTree(splash)

	ranger.PushStart(boot)

	ranger.Configure()

	ranger.Start()

	ranger.End()
}
