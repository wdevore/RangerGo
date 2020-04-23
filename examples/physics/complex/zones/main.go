package main

import (
	"fmt"

	"github.com/wdevore/RangerGo/engine"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

// -------------------------------------------------------------
// Note: this really isn't a physics example, it is a Zones example
// that uses physics for animations
// -------------------------------------------------------------

func main() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Keys:")
	fmt.Println("r = reset star ship")
	fmt.Println("a,d = apply Yaw to star ship")
	fmt.Println("l = apply impulse to star ship")
	fmt.Println("-----------------------------------------------------------")

	world := engine.NewWorld("Zones", 0.12, "../../..")

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
