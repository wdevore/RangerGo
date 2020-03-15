package main

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
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
	o.Transition.SetPauseTime(3000.0)
	return o
}

// --------------------------------------------------------
// Timing
// --------------------------------------------------------

func (s *sceneSplash) Update(dt float64) {
	s.Transition.UpdateTransition(dt)
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
	fmt.Println("basic scene splash exit")
	man.UnRegisterTarget(s)
}
