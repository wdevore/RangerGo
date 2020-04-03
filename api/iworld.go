package api

import "github.com/veandco/go-sdl2/sdl"

const (
	// ----------------------------------------------
	// Physics
	// ----------------------------------------------

	// PTM is Pixels-to-Meters which isn't used in Ranger. It is
	// here as an example from pixel based engines. I wouldn't
	// use it, but instead use STM below.
	// Box2D uses the MKS(meters/kilograms/seconds) unit system.
	PTM = 1.0 / 30.0 // 1 MKS = 30 GUs

	// RangerScale is a value you change according to your desires.
	// The default is 30.0. For example
	RangerScale = 30.0

	// STM is the Scale-to-MKS ratio.
	// Because Ranger uses transforms we don't think in terms of
	// pixels but rather in terms of spaces. Ranger's View-space
	// --without any scaling--is equal to physic-space (aka Box2D-space)
	// Thus if we want, for example, everything is ranger scaled up
	// then we need to scale it back down to physic-space and that
	// is what STM is for.
	STM = 1.0 / RangerScale // 1 MKS = 30 GUs

	// VelocityIterations is a resolution adjustment
	VelocityIterations = 8

	// PositionIterations is a resolution adjustment
	PositionIterations = 3
	// ----------------------------------------------

	// ----------------------------------------------
	// Display and View
	// ----------------------------------------------

	// DisplayRatio is the aspect ratio.
	DisplayRatio = 16.0 / 9.0
	// Width is the horizontal size
	Width = 1024.0 + 512.0

	// ViewScale is a global scale
	ViewScale = 0.25

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
	InvViewSpace() IAffineTransform

	VectorFont() IVectorFont
	RasterFont() IRasterFont

	WorkingPath() string
}
