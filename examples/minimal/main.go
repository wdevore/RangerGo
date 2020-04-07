package main

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine"
)

var ranger api.IEngine

// This example just runs to basic scenes and then exits.
// Each scene pauses for a few seconds.

func init() {
	world := engine.NewWorld("RangerGo", 1.5, "..")

	ranger = engine.New(world)

	splash := newBasicSplashScene("Splash", nil)
	splash.Build(world)

	boot := newBasicBootScene("Boot", splash)

	ranger.PushStart(boot)
}

func main() {
	fmt.Println("Minimal")

	ranger.Configure()

	ranger.Start()

	ranger.End()
}
