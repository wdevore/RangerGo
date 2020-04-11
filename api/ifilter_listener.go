package api

// IFilterListener represents a Box2D filter listener
type IFilterListener interface {
	ShouldCollide(nodeA, nodeB INode) bool
}
