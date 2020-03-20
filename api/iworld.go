package api

import "github.com/veandco/go-sdl2/sdl"

const (
	// DisplayRatio is the aspect ratio.
	DisplayRatio = 16.0 / 9.0
	// Width is the horizontal size
	Width = 1024.0 + 512.0

	// ViewScale is a global scale
	ViewScale = 1.5

	// WindowPositionX is a default position
	WindowPositionX = 1000.0
	// WindowPositionY is a default position
	WindowPositionY = 100.0
)

// IWorld represents app window properties
type IWorld interface {
	SetRenderer(*sdl.Renderer)
	Renderer() *sdl.Renderer
	Context() IRenderContext

	// WindowSize is the device window dimensions.
	WindowSize() IPoint
	ViewSize() IPoint

	// Title is the window title
	Title() string

	// SetViewSpace configures the view-space matrix
	SetViewSpace()

	// ViewSpace returns the view-space matrix
	ViewSpace() IAffineTransform

	VectorFont() IVectorFont
	RasterFont() IRasterFont
}
