package api

// IVectorFont is the polygon font defined in assets/vector_font.data
type IVectorFont interface {
	Initialize(dataFile string)

	// Glyph returns an array of vertices that matches the character
	Glyph(char byte) []float64
}
