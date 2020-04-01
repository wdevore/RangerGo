package main

import (
	"fmt"
	"time"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine"
)

var ranger api.IEngine

func init() {
	world := engine.NewWorld("RangerGo", "..")

	ranger = engine.New(world)
}

func main() {
	fmt.Println("Bare minimum")

	w, h := ranger.DisplaySize()
	fmt.Printf("Display size: %d x %d\n", w, h)

	ranger.Configure()

	ranger.Start()

	time.Sleep(time.Millisecond * 3000)

	ranger.End()
}
