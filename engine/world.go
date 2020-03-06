package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
)

type world struct {
	title string

	windowPosition api.IPoint
	windowSize     api.IPoint
	windowCentered bool

	viewSize     api.IPoint
	viewCentered bool

	viewSpace    api.IAffineTransform
	invViewSpace api.IAffineTransform

	renderer *sdl.Renderer
}

// NewWorld constructs an IWorld object
func NewWorld(title string) api.IWorld {
	o := new(world)
	o.title = title

	o.windowPosition = geometry.NewPointUsing(api.WindowPositionX, api.WindowPositionY)
	o.windowSize = geometry.NewPointUsing(api.Width, api.Width/api.DisplayRatio)
	o.windowCentered = true

	o.viewSize = geometry.NewPointUsing(o.windowSize.X()*api.ViewScale, o.windowSize.Y()*api.ViewScale)
	o.viewCentered = true

	o.viewSpace = maths.NewTransform()
	o.invViewSpace = maths.NewTransform()

	o.SetViewSpace()

	fmt.Println("Display dimensions: ", o.windowSize)
	fmt.Println("View Dimensions: ", o.viewSize)

	return o
}

func (w *world) SetRenderer(rend *sdl.Renderer) {
	w.renderer = rend
}

func (w *world) Renderer() *sdl.Renderer {
	return w.renderer
}

func (w *world) WindowSize() api.IPoint {
	return w.windowSize
}

func (w *world) Title() string {
	return w.title
}

func (w *world) SetViewSpace() {
	center := maths.NewTransform()

	// What separates world from view is the ratio between the device (aka window)
	// and an optional centering translation.
	widthRatio := w.windowSize.X() / w.viewSize.X()
	heightRatio := w.windowSize.Y() / w.viewSize.Y()

	if w.viewCentered {
		center.MakeTranslate(w.windowSize.X()/2.0, w.windowSize.Y()/2.0)
	}

	center.Scale(widthRatio, heightRatio)

	w.viewSpace.SetByTransform(center)

	w.invViewSpace.SetByTransform(w.invViewSpace)
}

func (w *world) ViewSpace() api.IAffineTransform {
	return w.viewSpace
}
