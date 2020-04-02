package api

// INode is an abstract object that represents SceneGraph nodes
type INode interface {
	ID() int

	// Initialize configures default properties.
	Initialize(name string)

	// InitializeWithID configures default properties.
	InitializeWithID(id int, name string)

	Build(IWorld)
	World() IWorld

	HasParent() bool
	SetParent(INode)
	Parent() INode

	CalcTransform() IAffineTransform

	Interpolate(interpolation float64)

	EnterNode(INodeManager)
	ExitNode(INodeManager)

	IsVisible() bool
	SetVisible(bool)

	IsDirty() bool
	SetDirty(dirty bool)
	// RippleDirty passes the dirty flag downward to children.
	RippleDirty(dirty bool)

	Handle(IEvent) bool

	// IScene
	ITransform

	// Children
	IGroup

	GetBucket() []IPoint

	Update(msPerUpdate, secPerUpdate float64)
}
