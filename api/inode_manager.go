package api

// INodeManager manages node on a stack and forms a SceneGraph
type INodeManager interface {
	ClearEnabled(bool)

	PreVisit()
	Visit(interpolation float64) bool
	PostVisit()

	Update(msPerUpdate, secPerUpdate float64)

	PushNode(INode)
	PopNode() INode
	ReplaceNode(INode)

	RouteEvents(IEvent)

	RegisterTarget(target INode)
	UnRegisterTarget(target INode)

	RegisterEventTarget(target INode)
	UnRegisterEventTarget(target INode)

	End()

	Debug()
}
