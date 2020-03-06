package custom

import (
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
}

// // AffineTransform see ITransform
// func (n *CrossNode) AffineTransform() api.IAffineTransform {
// 	return nil
// }

// // InverseTransform see ITransform
// func (n *CrossNode) InverseTransform() api.IAffineTransform {
// 	return nil
// }

// // CalcFilteredTransform see ITransform
// func (n *CrossNode) CalcFilteredTransform() {
// }
