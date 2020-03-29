package custom

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
)

// AnchorNode is generally used as a focal point about which transforms
// take place.
// This node is as basic as it gets, it has no visual and only collects
// children.
// Some anchor nodes have a Filter as a parent.
type AnchorNode struct {
	nodes.Node
}

// NewAnchorNode constructs a axis aligned bounding box node
func NewAnchorNode(name string, parent api.INode) *AnchorNode {
	o := new(AnchorNode)
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	return o
}

// Build configures the node
func (a *AnchorNode) Build(world api.IWorld) {
	a.Node.Build(world)
}
