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
	fmt.Println("w,a,s,d apply linear velocity")
	fmt.Println("1,2,3 changes the velocity algorithm")
	fmt.Println("4,5,6,7,8,9 changes tracking algorithm")
	fmt.Println("y,u,i,o changes targeting rate from: 10(fast), 20, 40, 60(slow)")
	// fmt.Println("s apply immediate torge to yellow box")
	fmt.Println("r resets everything")
	fmt.Println("-----------------------------------------------------------")

	world := engine.NewWorld("Target tracking", 0.25, "../../..")

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
