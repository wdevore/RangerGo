package api

const (
	// IOTypeKeyboard is a keyboard event
	IOTypeKeyboard = 0
	// IOTypeMouseMotion is a mouse event
	IOTypeMouseMotion = 1024
	// IOTypeMouseButtonDown is a mouse event
	IOTypeMouseButtonDown = 1025
	// IOTypeMouseButtonUp is a mouse event
	IOTypeMouseButtonUp = 1026
	// IOTypeMouseWheel is a mouse event
	IOTypeMouseWheel = 1027
)

// IEvent represents IO event system
type IEvent interface {
	Reset()

	BeenHandled() bool
	Handled(mark bool)

	SetMousePosition(x, y int32)
	GetMousePosition() (x, y int32)
	SetMouseRelMovement(x, y int32)
	GetMouseRelMovement() (x, y int32)

	SetDirection(uint32)
	GetDirection() uint32

	SetType(uint32)
	GetType() uint32

	SetClicks(uint8)
	GetClicks() uint8

	SetButton(uint8)
	GetButton() uint8

	SetWhich(uint32)
	GetWhich() uint32

	SetState(uint32)
	GetState() uint32
	SetRepeat(uint8)
	GetRepeat() uint8

	SetKeyScan(uint32)
	GetKeyScan() uint32
	SetKeyCode(uint32)
	GetKeyCode() uint32
	SetKeyMotif(uint32)
	GetKeyMotif() uint32
}
