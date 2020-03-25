package api

// IDragging represents the node dragging behaviour
type IDragging interface {
	IsDragging() bool

	Delta() IPoint
	SetMotionState(x, y int32, state uint32)
	SetButtonState(x, y int32, button uint8, state uint32)

	SetMotionStateUsing(x, y int32, state uint32, node INode)
	SetButtonStateUsing(x, y int32, button uint8, state uint32, node INode)
}
