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

	api.IRender

	backgroundColor api.IPalette
	backgroundMin   api.IPoint
	backgroundMax   api.IPoint

	textColor api.IPalette
}

func newBasicSplashScene(name string, replacement api.INode) api.INode {
	o := new(sceneSplash)
	o.Initialize(name)
	o.SetReplacement(replacement)
	return o
}

func (s *sceneSplash) Build(world api.IWorld) {
	vw, vh := world.ViewSize().Components()
	x := -vw / 2.0
	y := -vh / 2.0

	s.backgroundMin = geometry.NewPointUsing(x, y)
	s.backgroundMax = geometry.NewPointUsing(x+vw, y+vh)

	s.SetPauseTime(30000.0)

	s.backgroundColor = rendering.NewPaletteInt64(rendering.LightGray)
	s.textColor = rendering.NewPaletteInt64(rendering.LightNavyBlue)
}

// --------------------------------------------------------
// Timing
// --------------------------------------------------------

func (s *sceneSplash) Update(msPerUpdate, secPerUpdate float64) {
	s.Transition.UpdateTransition(msPerUpdate)
}

// --------------------------------------------------------
// Transitioning
// --------------------------------------------------------

func (s *sceneSplash) TransitionAction() int {
	if s.Transition.ReadyToTransition() {
		return api.SceneReplaceTake
	}

	return api.SceneNoAction
}

// -----------------------------------------------------
// Scene lifecycles
// -----------------------------------------------------

// EnterNode called when a node is entering the stage
func (s *sceneSplash) EnterNode(man api.INodeManager) {
	s.Transition.Reset()
	man.RegisterTarget(s)
}

// ExitNode called when a node is exiting stage
func (s *sceneSplash) ExitNode(man api.INodeManager) {
	man.UnRegisterTarget(s)
}

// -----------------------------------------------------
// Visuals
// -----------------------------------------------------

var o1 = geometry.NewPoint()
var o2 = geometry.NewPoint()

func (s *sceneSplash) Draw(context api.IRenderContext) {
	// Transform vertices if anything has changed.
	if s.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformPoint(s.backgroundMin, o1)
		context.TransformPoint(s.backgroundMax, o2)
		s.SetDirty(false) // Node is no longer dirty
	}

	// Draw background first.
	context.SetDrawColor(s.backgroundColor)
	context.RenderAARectangle(o1, o2, api.FILLED)

	context.SetDrawColor(s.textColor)
	context.DrawText(450.0, 250.0, "Splash Scene", 6, 1, false)

	s.Node.Draw(context)
}
