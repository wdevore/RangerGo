package rendering

import (
	"fmt"
	"image/color"

	"github.com/wdevore/RangerGo/api"
)

type palette struct {
	c color.RGBA
}

// NewPalette constructs an IPalette color object
func NewPalette() api.IPalette {
	o := new(palette)
	o.c.R = 0
	o.c.G = 0
	o.c.B = 0
	o.c.A = 255
	return o
}

// NewPaletteRGB constructs an RGB color object
func NewPaletteRGB(r, g, b int) api.IPalette {
	o := new(palette)
	o.c.R = uint8(r)
	o.c.G = uint8(g)
	o.c.B = uint8(b)
	o.c.A = 255
	return o
}

// NewPaletteRGBA constructs an RGB color object with alpha
func NewPaletteRGBA(r, g, b, a int) api.IPalette {
	o := NewPaletteRGB(r, g, b)
	o.SetAlpha(a)
	return o
}

// NewPaletteInt64 constructs an RGB color object a single 64bit int
func NewPaletteInt64(c uint64) api.IPalette {
	o := new(palette)
	o.c.R = uint8((c & 0xff000000) >> 24)
	o.c.G = uint8((c & 0x00ff0000) >> 16)
	o.c.B = uint8((c & 0x0000ff00) >> 8)
	o.c.A = uint8(c & 0x000000ff)
	return o
}

func (p *palette) AsUInt64() uint64 {
	return uint64((uint64(p.c.R) << 24) | (uint64(p.c.G) << 16) | (uint64(p.c.B) << 8) | uint64(p.c.A))
}

func (p *palette) Color() color.RGBA {
	return p.c
}

func (p *palette) Components() (r, g, b, a uint8) {
	return r, g, b, a
}

func (p *palette) R() uint8 {
	return p.c.R
}

func (p *palette) G() uint8 {
	return p.c.G
}

func (p *palette) B() uint8 {
	return p.c.B
}

func (p *palette) A() uint8 {
	return p.c.A
}

func (p *palette) SetRed(c int) {
	p.c.R = uint8(c)
}

func (p *palette) SetGreen(c int) {
	p.c.G = uint8(c)
}

func (p *palette) SetBlue(c int) {
	p.c.B = uint8(c)
}

func (p *palette) SetAlpha(c int) {
	p.c.A = uint8(c)
}

func (p palette) String() string {
	return fmt.Sprintf("{%d,%d,%d,%d} : %08x", p.c.R, p.c.G, p.c.B, p.c.A, p.AsUInt64())
}
