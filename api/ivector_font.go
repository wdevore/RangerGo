package api

// IVectorFont is the polygon font defined in assets/vector_font.data
type IVectorFont interface {
	Initialize(dataFile string, relativePath string)

	HorizontalOffset() float64
	VerticalOffset() float64
	Scale() float64

	// Glyph returns an array of vertices that matches the character
	Glyph(char byte) []float64
}
