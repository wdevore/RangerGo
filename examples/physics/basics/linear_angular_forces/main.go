package main

import (
	"fmt"

	"github.com/wdevore/RangerGo/engine"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

func main() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Keys:")
	fmt.Println("d apply gradual force upwards to orange box")
	fmt.Println("f apply immediate force upwards to blue box")
	fmt.Println("z apply immediate force upwards to the (1,1) lime box's corner")
	fmt.Println("a apply gradual torge to purple box")
	fmt.Println("s apply immediate torge to yellow box")
	fmt.Println("r resets everything")
	fmt.Println("-----------------------------------------------------------")

	world := engine.NewWorld("Linear/Angular forces", "../../..")

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
