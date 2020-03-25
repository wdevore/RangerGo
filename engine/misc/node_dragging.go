package misc

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
)

type dragState struct {
	dragging bool
	active   bool // Track button down state

	positionDown api.IPoint
	positionUp   api.IPoint
	position     api.IPoint
	delta        api.IPoint
	mapPoint     api.IPoint
}

// NewDragState returns a dragging state object
func NewDragState() api.IDragging {
	o := new(dragState)
	o.positionDown = geometry.NewPoint()
	o.positionUp = geometry.NewPoint()
	o.position = geometry.NewPoint()
	o.delta = geometry.NewPoint()
	o.mapPoint = geometry.NewPoint()
	return o
}

func (d *dragState) IsDragging() bool {
	return d.dragging
}

func (d *dragState) Delta() api.IPoint {
	return d.delta
}

func (d *dragState) SetMotionState(x, y int32, state uint32) {
	if d.active {
		if d.dragging {
			d.delta.SetByComp(float64(x)-d.position.X(), float64(y)-d.position.Y())
			d.position.SetByComp(float64(x), float64(y))
		}
	}
}

func (d *dragState) SetButtonState(x, y int32, button uint8, state uint32) {
	d.active = button == 1 && state == 1

	if d.active {
		d.positionDown.SetByComp(float64(x), float64(y))
		d.position.SetByPoint(d.positionDown)
	} else {
		d.positionUp.SetByComp(float64(x), float64(y))
	}
}

func (d *dragState) SetMotionStateUsing(x, y int32, state uint32, node api.INode) {
	if d.active {
		if d.dragging {
			// We need to map to parent space for dragging because the parent may contain
			// a scaling factor. Note: Using view-space will result in drifting from scale difference.
			nodes.MapDeviceToNode(x, y, node.Parent(), d.mapPoint)

			d.delta.SetByComp(d.mapPoint.X()-d.position.X(), d.mapPoint.Y()-d.position.Y())
			d.position.SetByPoint(d.mapPoint)
		}
	}
}

func (d *dragState) SetButtonStateUsing(x, y int32, button uint8, state uint32, node api.INode) {
	d.active = button == 1
	d.dragging = state == 1
	nodes.MapDeviceToNode(x, y, node.Parent(), d.mapPoint)

	if d.active {
		d.positionDown.SetByPoint(d.mapPoint)
		d.position.SetByPoint(d.mapPoint)
	} else {
		d.positionUp.SetByPoint(d.mapPoint)
	}
}
