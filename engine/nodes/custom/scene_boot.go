package custom

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
}

// NewBasicBootScene returns an IScene node of base type INode
func NewBasicBootScene(name string, replacement api.INode) api.INode {
	o := new(sceneBoot)
	o.Initialize(name)
	o.SetReplacement(replacement)
	return o
}

// --------------------------------------------------------
// Transitioning
// --------------------------------------------------------

// Transition indicates what to transition to next
func (s *sceneBoot) Transition() int {
	return api.SceneReplaceTake
}
