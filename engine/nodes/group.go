package nodes

import "github.com/wdevore/RangerGo/api"

// Group holds the children properties and methods.
type Group struct {
	children []api.INode
}

func (g *Group) initializeGroup() {
	g.children = []api.INode{}
}

// Children returns the children of current node.
// Nodes should override this method for providing any child they contain.
func (g *Group) Children() []api.INode {
	return g.children
}

// AddChild adds a node to this node
func (g *Group) AddChild(child api.INode) {
	if child != nil {
		g.children = append(g.children, child)
	}
}
