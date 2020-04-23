package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

// zoneManager handles zones
// The ZM coordinates between zones and any animations created by them.
// When a zone is entered all other zones' animations must stop
type zoneManager struct {
	parent api.INode

	zones         []api.INode
	enteredZoneID int

	// Zooming
	zoom api.INode

	zoomScale float64

	animationActive bool
}

// newZoneManager creates a zone manager
func newZoneManager(parent api.INode) *zoneManager {
	o := new(zoneManager)
	o.parent = parent
	return o
}

func (z *zoneManager) Build(world api.IWorld) {
	z.zoom = custom.NewZoomNode("ZoomNode", world, z.parent)
	gz := z.zoom.(*custom.ZoomNode)
	gz.SetStepSize(0.05)

	zone := NewZoneCircleNode("RightCircleZone", z.parent.World(), z.zoom, z)
	z.zones = append(z.zones, zone)
	zone.SetID(objectRightZone)
	gr := zone.(*ZoneCircleNode)
	gr.SetTweenRange(1.0, 3.0)
	gr.SetTweenDuration(1000.0)
	gr.Configure(12, 13.0, 15.0)
	gr.RequestNotification(z)
	gr.SetPosition(30.0, 20.0)

	zone = NewZoneCircleNode("LeftCircleZone", z.parent.World(), z.zoom, z)
	z.zones = append(z.zones, zone)
	zone.SetID(objectLeftZone)
	gr = zone.(*ZoneCircleNode)
	gr.SetTweenRange(1.0, 2.0)
	gr.SetTweenDuration(1000.0)
	gr.Configure(12, 7.0, 10.0)
	gr.RequestNotification(z)
	gr.SetPosition(-30.0, 20.0)
	// gr.SetPosition(0.0, 15.0)
}

// GetZoom returns zoom INode
func (z *zoneManager) GetZoom() api.INode {
	return z.zoom
}

// UpdateCheck updates zone tweens
func (z *zoneManager) UpdateCheck(point api.IPoint, msPerUpdate float64) {
	isFinished := true

	for _, zone := range z.zones {
		grz := zone.(*ZoneCircleNode)
		grz.UpdateCheck(point)

		// Animate only the zone that was enter. The other zone's
		// animation is frozen/stopped.
		if z.enteredZoneID == zone.ID() {
			z.zoomScale, isFinished = grz.TweenUpdate(msPerUpdate)
			if !isFinished {
				gz := z.zoom.(*custom.ZoomNode)
				gz.ScaleTo(float64(z.zoomScale))
			}
		}
	}
}

func (z *zoneManager) AnimationActive() bool {
	return z.animationActive
}

func (z *zoneManager) SetAnimationActive(active bool) {
	z.animationActive = active
}

func (z *zoneManager) ZoomScale() float64 {
	gz := z.zoom.(*custom.ZoomNode)

	return gz.ZoomScale()
}

// ----------------------------------------------------------
// IZoneListener implementation
// ----------------------------------------------------------

// Notify receives messages from IZone objects
func (z *zoneManager) Notify(state, id int) {
	if state != api.CrossStateEntered {
		return
	}

	z.enteredZoneID = id

	// fmt.Println("ZM notified: ", z.enteredZoneID)

	// Find zone that matches "id"
	for _, zone := range z.zones {
		if z.enteredZoneID == zone.ID() {
			gz := z.zoom.(*custom.ZoomNode)
			gz.SetFocalPoint(zone.Position().X(), zone.Position().Y())
			break
		}
	}
}
