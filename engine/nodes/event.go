package nodes

import "fmt"

// Event handles IO events
type Event struct {
	eTtimeStamp  uint32
	eType        uint32
	eWhich       uint32
	eState       uint32
	eClicks      uint8
	eButton      uint8
	eDir         uint32
	eRepeat      uint8
	eKeyScancode uint32
	eKeycode     uint32
	eKeyModif    uint32
	mx, my       int32
	mxRel, myRel int32
	handled      bool
}

// NewEvent returns a new Event object
func NewEvent() *Event {
	o := new(Event)
	return o
}

// Reset clears all properties
func (e *Event) Reset() {
	e.eTtimeStamp = 0
	e.eType = 0
	e.eWhich = 0
	e.eState = 0
	e.eClicks = 0
	e.eButton = 0
	e.eDir = 0
	e.eRepeat = 0
	e.eKeyScancode = 0
	e.eKeycode = 0
	e.eKeyModif = 0
	e.mx = 0
	e.my = 0
	e.handled = false
}

// Handled marks event as handled to stop event bubbling
func (e *Event) Handled(mark bool) {
	e.handled = mark
}

// BeenHandled indicates if the event has already been taken care of.
func (e *Event) BeenHandled() bool {
	return e.handled
}

// SetMousePosition set mx, my
func (e *Event) SetMousePosition(x, y int32) {
	e.mx = x
	e.my = y
}

// GetMousePosition gets mx, my
func (e *Event) GetMousePosition() (x, y int32) {
	return e.mx, e.my
}

// SetMouseRelMovement set mxRel, myRel
func (e *Event) SetMouseRelMovement(x, y int32) {
	e.mxRel = x
	e.myRel = y
}

// GetMouseRelMovement gets
func (e *Event) GetMouseRelMovement() (x, y int32) {
	return e.mxRel, e.myRel
}

// SetType sets
func (e *Event) SetType(eType uint32) {
	e.eType = eType
}

// GetType gets
func (e *Event) GetType() uint32 {
	return e.eType
}

// SetClicks sets
func (e *Event) SetClicks(eClicks uint8) {
	e.eClicks = eClicks
}

// GetClicks gets
func (e *Event) GetClicks() uint8 {
	return e.eClicks
}

// SetButton sets
func (e *Event) SetButton(eButton uint8) {
	e.eButton = eButton
}

// GetButton gets
func (e *Event) GetButton() uint8 {
	return e.eButton
}

// SetDirection sets
func (e *Event) SetDirection(eDir uint32) {
	e.eDir = eDir
}

// GetDirection gets
func (e *Event) GetDirection() uint32 {
	return e.eDir
}

// SetWhich sets
func (e *Event) SetWhich(eWhich uint32) {
	e.eWhich = eWhich
}

// GetWhich gets
func (e *Event) GetWhich() uint32 {
	return e.eWhich
}

// SetState sets
func (e *Event) SetState(state uint32) {
	e.eState = state
}

// GetState gets
func (e *Event) GetState() uint32 {
	return e.eState
}

// SetRepeat sets
func (e *Event) SetRepeat(rep uint8) {
	e.eRepeat = rep
}

// GetRepeat gets
func (e *Event) GetRepeat() uint8 {
	return e.eRepeat
}

// SetKeyScan sets
func (e *Event) SetKeyScan(code uint32) {
	e.eKeyScancode = code
}

// GetKeyScan gets
func (e *Event) GetKeyScan() uint32 {
	return e.eKeyScancode
}

// SetKeyCode sets
func (e *Event) SetKeyCode(code uint32) {
	e.eKeycode = code
}

// GetKeyCode gets
func (e *Event) GetKeyCode() uint32 {
	return e.eKeycode
}

// SetKeyMotif sets
func (e *Event) SetKeyMotif(code uint32) {
	e.eKeyModif = code
}

// GetKeyMotif gets
func (e *Event) GetKeyMotif() uint32 {
	return e.eKeyModif
}

func (e Event) String() string {
	s := "----------Event---------\n"
	s += fmt.Sprintf("mx: %d, my: %d\n", e.mx, e.my)
	s += fmt.Sprintf("mxRel: %d, myRel: %d\n", e.mxRel, e.myRel)
	s += fmt.Sprintf("Button: %d\n", e.eButton)
	s += fmt.Sprintf("Dir: %d\n", e.eDir)
	s += fmt.Sprintf("Which: %d 0x%0x\n", e.eWhich, e.eWhich)
	s += fmt.Sprintf("State: %d 0x%0x\n", e.eState, e.eState)
	s += fmt.Sprintf("KeyScan: %d\n", e.eKeyScancode)
	s += fmt.Sprintf("KeyCode: %d\n", e.eKeycode)
	s += fmt.Sprintf("Type: (%d) 0x%0x", e.eType, e.eType)
	return s
}
