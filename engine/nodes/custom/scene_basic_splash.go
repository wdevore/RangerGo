package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
)

// Note: this is a very basic splash Node used pretty much for just
// engine development. You should actually supply your own splash node,
// and example can be found in the examples folder.
type sceneBasicSplash struct {
	nodes.Node
	nodes.Scene
}

// NewBasicSplashScene returns an IScene node of base type INode
func NewBasicSplashScene(name string, replacement api.INode) api.INode {
	o := new(sceneBasicSplash)
	o.Initialize(name)
	o.SetReplacement(replacement)
	return o
}

// --------------------------------------------------------
// Transitioning
// --------------------------------------------------------

// Transition indicates what to transition to next
func (s *sceneBasicSplash) Transition() int {
	return api.SceneReplaceTake
}
