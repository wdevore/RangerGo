package api

// INode is an abstract object that represents SceneGraph nodes
type INode interface {
	ID() int

	// Initialize configures default properties.
	Initialize(name string)

	// InitializeWithID configures default properties.
	InitializeWithID(id int, name string)

	// Visit traverses "down" the heirarchy while space-mappings traverses upward.
	Visit(context IRenderContext, interpolation float64)

	SetParent(INode)
	Parent() INode

	CalcTransform() IAffineTransform

	Interpolate(interpolation float64)

	Draw(context IRenderContext)

	EnterNode(INodeManager)
	ExitNode(INodeManager)

	IsVisible() bool

	IsDirty() bool
	SetDirty(dirty bool)
	// RippleDirty passes the dirty flag downward to children.
	RippleDirty(dirty bool)

	Handle(IEvent) bool

	// IScene
	ITransform
	IGroup

	GetBucket() []IPoint

	Update(dt float64)
}
