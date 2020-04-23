package main

import (
	"math"

	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/misc"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

// ZoneCircleNode is a basic vector circle shape.
type ZoneCircleNode struct {
	nodes.Node

	segments int

	innerRadius float64
	outerRadius float64

	innerCircle api.IPolygon
	outerCircle api.IPolygon

	innerColor   api.IPalette
	outerColor   api.IPalette
	enteredColor api.IPalette

	zone      api.IZone // CircleZone
	zoneState int

	tweenEnabled      bool
	tweenZoomIn       *gween.Tween
	tweenZoomOut      *gween.Tween
	isFinished        bool
	tweenCurrentValue float32

	zoomTo   float64
	zoomFrom float64
	duration float64

	// Typically a zone manager would be the subscriber
	subscribers []api.IZoneListener

	zoneMan *zoneManager
}

// NewZoneCircleNode constructs a circle shaped node
func NewZoneCircleNode(name string, world api.IWorld, parent api.INode, zoneMan *zoneManager) api.INode {
	o := new(ZoneCircleNode)
	o.zoneMan = zoneMan
	o.Initialize(name)
	o.SetParent(parent)
	parent.AddChild(o)
	o.Build(world)
	return o
}

// Build configures the node
func (z *ZoneCircleNode) Build(world api.IWorld) {
	z.Node.Build(world)

	z.subscribers = []api.IZoneListener{}

	z.innerColor = rendering.NewPaletteInt64(rendering.LightGray)
	z.outerColor = rendering.NewPaletteInt64(rendering.Silver)
	z.enteredColor = rendering.NewPaletteInt64(rendering.LightPurple)

	z.isFinished = true
}

// Configure circles, if radius is 1 then diameter is 2
func (z *ZoneCircleNode) Configure(segments int, innerRadius, outerRadius float64) {
	z.segments = segments // typically 12

	z.zone = misc.NewCircleZone()
	z.SetRadi(innerRadius, outerRadius)

	step := math.Pi / float64(z.segments)

	z.innerCircle = geometry.NewPolygon()
	z.outerCircle = geometry.NewPolygon()
	for i := 0.0; i < 2.0*math.Pi; i += step {
		ix := math.Cos(i) * z.innerRadius
		iy := math.Sin(i) * z.innerRadius
		z.innerCircle.AddVertex(ix, iy)
		ox := math.Cos(i) * z.outerRadius
		oy := math.Sin(i) * z.outerRadius
		z.outerCircle.AddVertex(ox, oy)
	}
	z.innerCircle.Build()
	z.outerCircle.Build()
}

// RequestNotification asks for notification when an event on the zone
// happens. ZoneManager is a subscriber.
func (z *ZoneCircleNode) RequestNotification(listener api.IZoneListener) {
	z.subscribers = append(z.subscribers, listener)
}

// SetTweenRange sets the from and to values
func (z *ZoneCircleNode) SetTweenRange(from, to float64) {
	z.zoomTo = to
	z.zoomFrom = from
}

// SetTweenDuration sets the animation time duration
func (z *ZoneCircleNode) SetTweenDuration(duration float64) {
	z.duration = duration
}

// SetPosition sets position of zone
func (z *ZoneCircleNode) SetPosition(x, y float64) {
	z.Node.SetPosition(x, y)
	cr := z.zone.(*misc.CircleZone)
	cr.SetPosition(x, y)
}

// SetRadi sets circle's inner and outer radi
func (z *ZoneCircleNode) SetRadi(innerRadius, outerRadius float64) {
	z.innerRadius = innerRadius
	z.outerRadius = outerRadius
	cr := z.zone.(*misc.CircleZone)
	cr.SetRadi(innerRadius, outerRadius)
}

// SetSegments sets how many segments on the circle (default = 12)
func (z *ZoneCircleNode) SetSegments(segments int) {
	z.segments = segments
}

// SetInnerColor sets circle's inner color (default = LightGray)
func (z *ZoneCircleNode) SetInnerColor(color api.IPalette) {
	z.innerColor = color
}

// SetOuterColor sets circle's outer color (default = Silver)
func (z *ZoneCircleNode) SetOuterColor(color api.IPalette) {
	z.outerColor = color
}

// UpdateCheck forces the zone to update based on a given point
func (z *ZoneCircleNode) UpdateCheck(point api.IPoint) (state, id int) {
	newState, stateChanged := z.zone.Update(point)
	id = z.ID()

	if stateChanged {
		z.zoneState = newState

		// Send message to listeners. The "id" is a self identifier.
		// Most likely the ZoneManager
		for _, listener := range z.subscribers {
			listener.Notify(z.zoneState, id)
		}

		z.createTween(z.zoneState, id)
	}

	return newState, id
}

// TweenUpdate updates any tweens if enabled
func (z *ZoneCircleNode) TweenUpdate(msPerUpdate float64) (float64, bool) {
	if z.tweenEnabled {
		switch z.zoneState {
		case api.CrossStateEntered:
			z.tweenCurrentValue, z.isFinished = z.tweenZoomIn.Update(float32(msPerUpdate))
			if z.isFinished {
				z.tweenEnabled = false
				z.zoneMan.SetAnimationActive(false)
			}
		case api.CrossStateExited:
			z.tweenCurrentValue, z.isFinished = z.tweenZoomOut.Update(float32(msPerUpdate))
			if z.isFinished {
				z.tweenEnabled = false
				z.zoneMan.SetAnimationActive(false)
			}
		}
	}

	return float64(z.tweenCurrentValue), z.isFinished
}

// ----------------------------------------------------------
// IZoneListener implementation
// ----------------------------------------------------------

func (z *ZoneCircleNode) createTween(state, id int) {
	// Whenever an tween needs to be created we are animating from a
	// "begin" value to an "end" value--regardless of zooming In or Out.

	// The "to" should also we the final zoom value.
	// The "from" should be whatever the initial "to" value is OR
	//   the current value of the ZoomNode if an animation was in progress
	//   during the Enter event (aka tweenEnabled).

	// fmt.Println("Scale: ", z.zoneMan.ZoomScale(), ", current: ", z.tweenCurrentValue, ", from: ", z.zoomFrom, ", to: ", z.zoomTo)
	switch state {
	case api.CrossStateEntered:
		if !z.isFinished {
			z.tweenZoomIn = gween.New(float32(z.zoneMan.ZoomScale()), float32(z.zoomTo), float32(z.duration), ease.InOutQuad)
		} else {
			if z.zoneMan.AnimationActive() {
				z.tweenZoomIn = gween.New(float32(z.zoneMan.ZoomScale()), float32(z.zoomTo), float32(z.duration), ease.InOutQuad)
			} else {
				z.tweenZoomIn = gween.New(float32(z.zoomFrom), float32(z.zoomTo), float32(z.duration), ease.InOutQuad)
			}
		}
		z.innerColor = rendering.NewPaletteInt64(rendering.Lime)
		z.tweenEnabled = true
		z.zoneMan.SetAnimationActive(true)
	case api.CrossStateExited:
		if !z.isFinished {
			z.tweenZoomOut = gween.New(float32(z.zoneMan.ZoomScale()), float32(z.zoomFrom), float32(z.duration), ease.InOutQuad)
		} else {
			if z.zoneMan.AnimationActive() {
				z.tweenZoomOut = gween.New(float32(z.zoneMan.ZoomScale()), float32(z.zoomFrom), float32(z.duration), ease.InOutQuad)
			} else {
				z.tweenZoomOut = gween.New(float32(z.zoomTo), float32(z.zoomFrom), float32(z.duration), ease.InOutQuad)
			}
		}
		z.innerColor = rendering.NewPaletteInt64(rendering.LightGray)
		z.tweenEnabled = true
		z.zoneMan.SetAnimationActive(true)
	}
}

// Draw renders shape
func (z *ZoneCircleNode) Draw(context api.IRenderContext) {
	if z.IsDirty() {
		context.TransformPolygon(z.innerCircle)
		context.TransformPolygon(z.outerCircle)
		z.SetDirty(false)
	}

	context.SetDrawColor(z.innerColor)
	context.RenderPolygon(z.innerCircle, api.CLOSED)

	context.SetDrawColor(z.outerColor)
	context.RenderPolygon(z.outerCircle, api.CLOSED)
}
