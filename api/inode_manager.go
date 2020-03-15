package api

// INodeManager manages node on a stack and forms a SceneGraph
type INodeManager interface {
	PreVisit()
	Visit(interpolation float64) bool
	PostVisit()

	Update(dt float64)

	PushNode(INode)
	PopNode()
	ReplaceNode(INode)

	RouteEvents(IEvent)

	RegisterTarget(target INode)
	UnRegisterTarget(target INode)

	RegisterEventTarget(target INode)
	UnRegisterEventTarget(target INode)
}
