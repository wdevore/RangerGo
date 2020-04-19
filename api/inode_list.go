package api

// INodeList is a simple list collection
type INodeList interface {
	Items() []INode
	DeleteAt(i int, slice []INode)
	FindFirstElement(node INode, slice []INode) int
	Add(node INode)
	Remove(node INode)
}
