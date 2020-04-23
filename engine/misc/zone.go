package misc

import (
	"github.com/wdevore/RangerGo/api"
)

// Zone is a trigger region or area with an inner and out boundary
// Zones collect subscribers such that the subscribers can be
// notified of enter/exit states.
type Zone struct {
	prevState int
	state     int // CrossState

	// inner region has been visited since being outside or exiting
	innerAccessed bool
}

// NewZone constructs a new Zone
func NewZone() *Zone {
	o := new(Zone)
	return o
}

// InitializeZone initializes the base zone
func (z *Zone) InitializeZone() {
	z.state = api.CrossStateNone
	z.prevState = api.CrossStateNone
}

// UpdateState checks if an object crossed into/out-of the zone, and updates
// and returns the internal state. In addition, it will send messages to
// listeners if state changed.
func (z *Zone) UpdateState(currentState int) (state int, stateChanged bool) {
	stateChanged = false

	if z.prevState == api.ZoneStateEnteredOuter {
		if currentState == api.ZoneStateEnteredInner && !z.innerAccessed {
			// fmt.Println("zone: entered")
			z.state = api.CrossStateEntered
			z.innerAccessed = true
			stateChanged = true
		}
	} else if currentState == api.ZoneStateObjectIsOutside {
		if z.innerAccessed {
			// fmt.Println("zone: exited")
			z.state = api.CrossStateExited
			z.innerAccessed = false
			stateChanged = true
		}
	}

	z.prevState = currentState

	return z.state, stateChanged
}

// State returns the current state
func (z *Zone) State() int {
	return z.state
}
