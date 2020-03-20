package api

// A simple Unicode raster 8x8 font
// The raw font data was ported from a Rust crate:
// https://crates.io/crates/font8x8/0.2.3

// IRasterFont is the bitmap raster font defined in assets/raster_font.data
type IRasterFont interface {
	Initialize(dataFile string, relativePath string)

	// Glyph returns an array of vertices that matches the character
	Glyph(char byte) []uint8

	GlyphWidth() int
}
