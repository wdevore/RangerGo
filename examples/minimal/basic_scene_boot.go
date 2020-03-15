package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
)

// Note: this is a very basic boot Node used pretty much for just
// engine development. You should actually supply your own boot node,
// and example can be found in the examples folder.
type sceneBoot struct {
	nodes.Node
	nodes.Scene
	nodes.Transition
}

func newBasicBootScene(name string, replacement api.INode) api.INode {
	o := new(sceneBoot)
	o.Initialize(name)
	o.SetReplacement(replacement)
	return o
}

// --------------------------------------------------------
// Transitioning
// --------------------------------------------------------

func (s *sceneBoot) TransitionAction() int {
	return api.SceneReplaceTake
}
