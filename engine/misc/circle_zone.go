package misc

import (
	"math"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
)

// The inner region is considered entering. The outer region is considered exiting.
// First an object must "enter" the inner region first before an exit can
// be considered to occur. Until an object enters the inner region exiting
// the outer region does nothing.

// CircleZone is a trigger region or area with an inner and out circular boundary
type CircleZone struct {
	Zone

	innerRadius float64
	outerRadius float64

	position api.IPoint
}

// NewCircleZone constructs a new CircleZone
func NewCircleZone() api.IZone {
	o := new(CircleZone)
	o.InitializeZone()
	o.position = geometry.NewPoint()
	return o
}

// SetRadi sets zone's inner and outer radi
func (z *CircleZone) SetRadi(innerRadius, outerRadius float64) {
	z.innerRadius = innerRadius
	z.outerRadius = outerRadius
}

// SetPosition sets the center location of zone.
func (z *CircleZone) SetPosition(x, y float64) {
	z.position.SetByComp(x, y)
}

// Update changes zone's state based on point and returns new state
func (z *CircleZone) Update(position api.IPoint) (state int, stateChanged bool) {
	// First check if inside/outside using given position
	state = z.PointInside(position)
	return z.UpdateState(state)
}

// PointInside checks if point is inside either inner or outer radius
func (z *CircleZone) PointInside(point api.IPoint) int {
	distance := z.DistanceFromCenter(point)

	if distance < z.outerRadius && distance < z.innerRadius {
		return api.ZoneStateEnteredInner
	}

	if distance < z.outerRadius && distance > z.innerRadius {
		return api.ZoneStateEnteredOuter
	}

	return api.ZoneStateObjectIsOutside
}

// DistanceFromCenter returns distance from point to circle center
func (z *CircleZone) DistanceFromCenter(point api.IPoint) float64 {
	dx := z.position.X() - point.X()
	dy := z.position.Y() - point.Y()
	return math.Sqrt(dx*dx + dy*dy)
}

// DistanceFromEdge returns distance from point to outer radius
func (z *CircleZone) DistanceFromEdge(point api.IPoint) float64 {
	dx := z.position.X() - point.X()
	dy := z.position.Y() - point.Y()
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance - z.outerRadius
}
