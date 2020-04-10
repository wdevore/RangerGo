package api

// IContactListener represents a Box2D contact listener
type IContactListener interface {
	HandleBeginContact(nodeA, nodeB INode) bool
	HandleEndContact(nodeA, nodeB INode) bool
}
