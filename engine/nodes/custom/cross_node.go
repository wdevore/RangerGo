package custom

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes"
)

// CrossNode is a basic node shaped like "+"
type CrossNode struct {
	nodes.Node
}

// NewCrossNode constructs a cross shaped node
func NewCrossNode() *CrossNode {
	o := new(CrossNode)
	return o
}

// Draw renders shape
func (n *CrossNode) Draw(context api.IRenderContext) {
	fmt.Println("cross draw")
	n.Node.Draw(context)
}
