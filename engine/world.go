package engine

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
	"github.com/wdevore/RangerGo/engine/rendering"
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
	context  api.IRenderContext

	vectorFont api.IVectorFont
	rasterFont api.IRasterFont

	workingPath string
}

// NewWorld constructs an IWorld object
func NewWorld(title string, viewScale float64, relativePath string) api.IWorld {
	o := new(world)
	o.title = title

	o.windowPosition = geometry.NewPointUsing(api.WindowPositionX, api.WindowPositionY)
	o.windowSize = geometry.NewPointUsing(api.Width, api.Width/api.DisplayRatio)
	o.windowCentered = true

	o.viewSize = geometry.NewPointUsing(o.windowSize.X()*viewScale, o.windowSize.Y()*viewScale)
	o.viewCentered = true

	o.viewSpace = maths.NewTransform()
	o.invViewSpace = maths.NewTransform()

	o.SetViewSpace()

	o.context = rendering.NewRenderContext(o)
	o.context.Initialize()

	fmt.Println("Display dimensions: ", o.windowSize)
	fmt.Println("View Dimensions: ", o.viewSize)

	path, err := filepath.Abs(relativePath)

	if err != nil {
		log.Fatal(err)
	}

	o.workingPath = path
	fmt.Println("Working path: ", path)

	fmt.Println("Loading Vector font...")
	o.vectorFont = rendering.NewVectorFont()
	o.vectorFont.Initialize("vector_font.data", relativePath)

	fmt.Println("Loading Raster font...")
	o.rasterFont = rendering.NewRasterFont()
	o.rasterFont.Initialize("raster_font.data", relativePath)

	return o
}

func (w *world) WorkingPath() string {
	return w.workingPath
}

func (w *world) SetRenderer(rend *sdl.Renderer) {
	w.renderer = rend
}

func (w *world) Renderer() *sdl.Renderer {
	return w.renderer
}

func (w *world) Context() api.IRenderContext {
	return w.context
}

func (w *world) WindowSize() api.IPoint {
	return w.windowSize
}

func (w *world) ViewSize() api.IPoint {
	return w.viewSize
}

func (w *world) Title() string {
	return w.title
}

func (w *world) VectorFont() api.IVectorFont {
	return w.vectorFont
}

func (w *world) RasterFont() api.IRasterFont {
	return w.rasterFont
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

	w.invViewSpace.SetByTransform(center)
	w.InvViewSpace().Invert()
}

func (w *world) ViewSpace() api.IAffineTransform {
	return w.viewSpace
}

func (w *world) InvViewSpace() api.IAffineTransform {
	return w.invViewSpace
}
