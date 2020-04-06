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
	fmt.Println("d increases the density by 1.0")
	fmt.Println("f increases friction by 0.1 to a max of 1.0")
	fmt.Println("r increases restitution by 0.1 to a max of 1.0")
	fmt.Println("z resets density, friction, restitution to: 1.0, 0.0, 0.0")
	fmt.Println("-----------------------------------------------------------")

	world := engine.NewWorld("Fixture properties", "../../..")

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
