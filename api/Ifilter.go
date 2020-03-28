package api

// IFilter represents Transform Filter nodes
type IFilter interface {
	Build(IWorld)
	VisitFilter(context IRenderContext, interpolation float64)

	InheritOnlyRotation()
	InheritOnlyScale()
	InheritOnlyTranslation()
	InheritRotationAndTranslation()
	InheritAll()
}
