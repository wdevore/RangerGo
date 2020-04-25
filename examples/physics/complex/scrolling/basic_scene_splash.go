package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type sceneSplash struct {
	nodes.Node
	nodes.Scene
	nodes.Transition
}

func newBasicSplashScene(name string, replacement api.INode) api.INode {
	o := new(sceneSplash)
	o.Initialize(name)
	o.SetReplacement(replacement)
	return o
}

func (s *sceneSplash) Build(world api.IWorld) {
	buildBackground(s, world)

	newBasicGameLayer("Game Layer", world, s)
}

// --------------------------------------------------------
// Transitioning
// --------------------------------------------------------

func (s *sceneSplash) TransitionAction() int {
	// Basically this scene never transitions to any node.
	return api.SceneNoAction
}

func buildBackground(s *sceneSplash, world api.IWorld) {
	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	// A solid background
	sol := custom.NewBasicRectangleNode("Background", world, s)
	gs := sol.(*custom.BasicRectangleNode)
	gs.SetBoundsUsingComps(x, y, x+vw, y+vh)
	gs.SetColor(rendering.NewPaletteInt64(rendering.DarkerGray))
}
