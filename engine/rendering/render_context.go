package rendering

import (
	"image/color"
	"math"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
)

type renderState struct {
	clearColor color.RGBA
	drawColor  color.RGBA

	current api.IAffineTransform
}

func newRS() *renderState {
	o := new(renderState)
	o.clearColor = NewPaletteInt64(Black).Color()
	o.drawColor = NewPaletteInt64(White).Color()
	o.current = maths.NewTransform()
	return o
}

type renderContext struct {
	world api.IWorld

	stack    []*renderState
	stackTop int

	clearColor color.RGBA
	drawColor  color.RGBA

	windowSize api.IPoint

	current api.IAffineTransform
	post    api.IAffineTransform // Pre allocated cache
}

const stackDepth = 100

// Scratch working variables
var v1 = geometry.NewPoint()
var v2 = geometry.NewPoint()

// NewRenderContext constructs an IRenderContext object
func NewRenderContext(world api.IWorld) api.IRenderContext {
	o := new(renderContext)
	o.world = world
	o.clearColor = NewPaletteInt64(Orange).Color()
	o.drawColor = NewPaletteInt64(White).Color()
	o.current = maths.NewTransform()
	o.post = maths.NewTransform()
	o.windowSize = world.WindowSize()

	return o
}

func (rc *renderContext) Initialize(world api.IWorld) {
	rc.stack = make([]*renderState, stackDepth)

	for i := 0; i < stackDepth; i++ {
		rc.stack[i] = newRS()
	}

	// Apply centered view-space matrix
	rc.Apply(world.ViewSpace())
}

func (rc *renderContext) Apply(aft api.IAffineTransform) {
	// Concat this transform onto the current transform but don't push it.
	// Use post multiply
	maths.Multiply(aft, rc.current, rc.post)
	rc.current.SetByTransform(rc.post)
}

func (rc *renderContext) Pre() {
	c := rc.clearColor
	renderer := rc.world.Renderer()
	renderer.SetDrawColor(c.R, c.G, c.B, c.A)
	renderer.Clear()

	//Draw checkered board as an clear indicator for debugging
	//NOTE: disable this code for release builds
	//draw_checkerboard(context);
}

func (rc *renderContext) Save() {
	top := rc.stack[rc.stackTop]
	top.clearColor = rc.clearColor
	top.drawColor = rc.drawColor
	top.current.SetByTransform(rc.current)

	rc.stackTop++
}

func (rc *renderContext) Restore() {
	rc.stackTop--

	top := rc.stack[rc.stackTop]
	rc.clearColor = top.clearColor
	rc.drawColor = top.drawColor
	rc.current.SetByTransform(top.current)

	c := rc.clearColor
	renderer := rc.world.Renderer()
	renderer.SetDrawColor(c.R, c.G, c.B, c.A)
}

func (rc *renderContext) Post() {
	renderer := rc.world.Renderer()
	renderer.Present()
}

// =_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.
// Transforms
// =_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.=_.

func (rc *renderContext) TransformPoint(p, out api.IPoint) {
	rc.current.TransformToPoint(p, out)
}

func (rc *renderContext) TransformLine(p1, p2, out1, out2 api.IPoint) {
	rc.current.TransformToPoint(p1, out1)
	rc.current.TransformToPoint(p2, out2)
}

func (rc *renderContext) TransformArray(vertices, bucket []api.IPoint) {
	for i := 0; i < len(vertices); i++ {
		rc.current.TransformToPoint(vertices[i], bucket[i])
	}
}

func (rc *renderContext) TransformMesh(mesh api.IMesh) {
	vertices := mesh.Vertices()
	bucket := mesh.Bucket()
	for i := 0; i < len(vertices); i++ {
		rc.current.TransformToPoint(vertices[i], bucket[i])
	}
}

func (rc *renderContext) TransformPolygon(poly api.IPolygon) {
	vertices := poly.Mesh().Vertices()
	bucket := poly.Mesh().Bucket()
	for i := 0; i < len(vertices); i++ {
		rc.current.TransformToPoint(vertices[i], bucket[i])
	}
}

func (rc *renderContext) DrawPoint(x, y int32) {
	renderer := rc.world.Renderer()
	renderer.DrawPoint(x, y)
}

func (rc *renderContext) DrawLine(x1, y1, x2, y2 int32) {
	renderer := rc.world.Renderer()
	renderer.DrawLine(x1, y1, x2, y2)
}

var sdlRect = &sdl.Rect{}

func (rc *renderContext) DrawRectangle(rect api.IRectangle) {
	renderer := rc.world.Renderer()
	sdlRect.X, sdlRect.Y = rect.Min().ComponentsAsInt32()
	sdlRect.W, sdlRect.H = rect.DimesionsAsInt32()

	renderer.DrawRect(sdlRect)
}

func (rc *renderContext) DrawFilledRectangle(rect api.IRectangle) {
	renderer := rc.world.Renderer()
	sdlRect.X, sdlRect.Y = rect.Min().ComponentsAsInt32()
	sdlRect.W, sdlRect.H = rect.DimesionsAsInt32()

	renderer.FillRect(sdlRect)
}

func (rc *renderContext) DrawCheckerBoard(size int) {
	renderer := rc.world.Renderer()
	flip := false
	col := int32(0)
	row := int32(0)
	w, h := rc.windowSize.ComponentsAsInt32()
	s := int32(size)

	for row < h {
		for col < w {
			if flip {
				renderer.SetDrawColor(100, 100, 100, 255)
			} else {
				renderer.SetDrawColor(80, 80, 80, 255)
			}

			sdlRect.X = col
			sdlRect.Y = row
			sdlRect.W = col + s
			sdlRect.H = row + s
			renderer.FillRect(sdlRect)

			flip = !flip
			col += s
		}

		flip = !flip
		col = 0
		row += s
	}
}

func (rc *renderContext) RenderLine(x1, y1, x2, y2 float64) {
	rc.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2))
}

func (rc *renderContext) RenderMesh(mesh api.IMesh) {
	first := true
	bucket := mesh.Bucket()

	for i := 0; i < len(bucket); i++ {
		if first {
			v1.SetByPoint(bucket[i])
			first = false
			continue
		} else {
			v2.SetByPoint(bucket[i])
			first = true
		}

		rc.DrawLine(
			int32(v1.X()), int32(v1.Y()),
			int32(v2.X()), int32(v2.Y()))
	}
}

var irect = geometry.NewRectangle()

func (rc *renderContext) RenderAARectangle(min, max api.IPoint, fillStyle int) {
	irect.Set(math.Round(min.X()), math.Round(min.Y()), math.Round(max.X()), math.Round(max.Y()))

	if fillStyle == api.FILLED {
		rc.DrawFilledRectangle(irect)
	} else if fillStyle == api.OUTLINED {
		rc.DrawRectangle(irect)
	} else {
		rc.DrawFilledRectangle(irect)
		rc.DrawRectangle(irect)
	}
}
