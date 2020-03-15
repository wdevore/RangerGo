package nodes

import (
	"github.com/wdevore/RangerGo/api"
)

// Scene is an embedded type for nodes of type IScene.
// Boot scenes and Splash scenes are typical examples.
type Scene struct {
	replacement api.INode
}

// SetReplacement sets this node's replacment during transitions.
func (s *Scene) SetReplacement(replacement api.INode) {
	s.replacement = replacement
}

// GetReplacement returns the replacement node for transitions.
func (s *Scene) GetReplacement() api.INode {
	return s.replacement
}
