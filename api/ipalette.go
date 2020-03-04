package api

import "image/color"

// IPalette represents colors
type IPalette interface {
	// Components return each component
	Components() (r, g, b, a uint8)

	// AsUInt64 as 64nit unsigned
	AsUInt64() uint64

	// Color as image.RGBA
	Color() color.RGBA

	// R is red component
	R() uint8
	// G is green component
	G() uint8
	// B is blus component
	B() uint8
	// A is alpha component
	A() uint8

	// SetRed
	SetRed(r int)
	// SetGreen
	SetGreen(r int)
	// SetBlue
	SetBlue(r int)
	// SetAlpha
	SetAlpha(r int)
}
