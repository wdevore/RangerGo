package api

// ZoomValue returns the current zoom value
type ZoomValue func() float64

const (
	// CrossStateNone means has remained in the side since the last check.
	CrossStateNone = 0
	// CrossStateEntered Object has entered zone.
	CrossStateEntered = 1
	// CrossStateExited Object has exit zone.
	CrossStateExited = 2
	// CrossStateInside Object is currently inside
	CrossStateInside = 3
	// CrossStateOutside Object hasn't entered yet
	CrossStateOutside = 4

	// ZoneStateObjectIsOutside object is currently outside of a zone
	ZoneStateObjectIsOutside = 0
	// ZoneStateEnteredOuter Object has entered the outer region
	ZoneStateEnteredOuter = 1
	// ZoneStateEnteredInner Object has entered the inner region
	ZoneStateEnteredInner = 2
	// ZoneStateExitedInner Object has exited the inner region
	ZoneStateExitedInner = 3
	// ZoneStateExitedOuter Object has exited the outer region
	ZoneStateExitedOuter = 4
	// ZoneStateObjectIsInside object is currently inside of a zone
	ZoneStateObjectIsInside = 5

	// ZoneActionNone Object is inside both regions
	ZoneActionNone = 0
	// ZoneActionInward Object is heading inward
	ZoneActionInward = 1
	// ZoneActionOutward Object is heading outward
	ZoneActionOutward = 2
)

// IZone an area with two regions: inner and outer
type IZone interface {
	Update(position IPoint) (int, bool)

	State() int
}
