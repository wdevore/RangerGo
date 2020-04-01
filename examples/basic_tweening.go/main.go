package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

var ranger api.IEngine

// Note: This uses Ranger's minimalistic tweening framework.
// It is suggested that you use github.com/tanema/gween

func init() {
	world := engine.NewWorld("Basic tweening", "..")

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
