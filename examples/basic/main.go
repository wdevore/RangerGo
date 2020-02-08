package main

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine"
)

var ranger api.IEngine

func init() {
	ranger = engine.New(1024, 600)
}

func main() {
	fmt.Println("Bare minimum")
	w, h := ranger.DisplaySize()
	fmt.Printf("Display size: %d x %d\n", w, h)
}
