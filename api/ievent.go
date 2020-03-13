package api

const (
	// IOTypeKeyboard is a keyboard event
	IOTypeKeyboard = 0
	// IOTypeMouseMotion is a mouse event
	IOTypeMouseMotion = 1
	// IOTypeMouseButton is a mouse event
	IOTypeMouseButton = 2
	// IOTypeMouseWheel is a mouse event
	IOTypeMouseWheel = 3
)

// IEvent represents IO event system
type IEvent interface {
	BeenHandled() bool
	Handled(mark bool)

	SetMousePosition(x, y int32)
	GetMousePosition() (x, y int32)
	SetMouseRelMovement(x, y int32)
	GetMouseRelMovement() (x, y int32)

	SetDirection(uint32)
	GetDirection() uint32

	SetType(int)
	GetType() int

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
