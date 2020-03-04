package api

// INode is an abstract object that represents SceneGraph nodes
type INode interface {
	// Initialize configures default properties.
	Initialize(id int, name string)

	// Visit traverses "down" the heirarchy while space-mappings traverses upward.
	Visit(context IRenderContext, interpolation float64)

	CalcTransform() IAffineTransform

	Interpolate(interpolation float64)

	Draw(context IRenderContext)

	EnterNode()
	ExitNode()

	IOEvent()
	IsVisible() bool
	IsDirty() bool
	SetDirty(dirty bool)
	RippleDirty(dirty bool)

	ITransform

	INodeGroup

	GetBucket()
	Update()
}
