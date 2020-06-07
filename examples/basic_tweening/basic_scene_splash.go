package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type sceneSplash struct {
	nodes.Node
	nodes.Scene
	nodes.Transition

	backgroundColor api.IPalette
	backgroundMin   api.IPoint
	backgroundMax   api.IPoint

	o1 api.IPoint
	o2 api.IPoint
}

func newBasicSplashScene(name string, replacement api.INode) api.INode {
	o := new(sceneSplash)
	o.Initialize(name)
	o.SetReplacement(replacement)
	return o
}

func (s *sceneSplash) Build(world api.IWorld) {
	layer := newBasicGameLayer("Game Layer", s)
	layer.Build(world)
	s.AddChild(layer)

	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	s.o1 = geometry.NewPoint()
	s.o2 = geometry.NewPoint()

	s.backgroundMin = geometry.NewPointUsing(x, y)
	s.backgroundMax = geometry.NewPointUsing(x+vw, y+vh)
	s.backgroundColor = rendering.NewPaletteInt64(rendering.DarkGray)
}

// -----------------------------------------------------
// Visuals
// -----------------------------------------------------

func (s *sceneSplash) Draw(context api.IRenderContext) {
	// Transform vertices if anything has changed.
	if s.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformPoint(s.backgroundMin, s.o1)
		context.TransformPoint(s.backgroundMax, s.o2)
		s.SetDirty(false) // Node is no longer dirty
	}

	// Draw background first.
	context.SetDrawColor(s.backgroundColor)
	context.RenderAARectangle(s.o1, s.o2, api.FILLED)
}

// --------------------------------------------------------
// Transitioning
// --------------------------------------------------------

func (s *sceneSplash) TransitionAction() int {
	// Basically this scene never transitions to any node.
	return api.SceneNoAction
}
