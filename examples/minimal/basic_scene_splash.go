package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
)

type sceneSplash struct {
	nodes.Node
	nodes.Scene
	count int
}

func NewBasicSplashScene(name string, replacement api.INode) api.INode {
	o := new(sceneSplash)
	o.Initialize(name)
	o.SetReplacement(replacement)
	return o
}

// --------------------------------------------------------
// Transitioning
// --------------------------------------------------------

func (s *sceneSplash) Transition() int {
	s.count++
	if s.count < 3 {
		return api.SceneNoAction
	} else {
		return api.SceneReplaceTake
	}
}
