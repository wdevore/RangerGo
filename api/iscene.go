package api

const (
	// SceneNoAction means no action is taken when transitioning
	SceneNoAction = 0
	// SceneReplace ...
	SceneReplace = 1
	// SceneReplaceTake ...
	SceneReplaceTake = 2
	// SceneReplaceTakeUnregister ...
	SceneReplaceTakeUnregister = 3
)

// IScene scene management
type IScene interface {
	Transition() int

	GetReplacement() INode
}
