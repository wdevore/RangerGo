package rendering

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/wdevore/RangerGo/api"
)

// --------------------------------------------------------------
// Internal glyph
// --------------------------------------------------------------

type vectorGlyph struct {
	vertices []float64
}

func newVectorGlyph() *vectorGlyph {
	o := new(vectorGlyph)
	o.vertices = []float64{}
	return o
}

func (vg *vectorGlyph) addVector(x1, y1, x2, y2 float64) {
	vg.vertices = append(vg.vertices, x1, y1, x2, y2)
}

// --------------------------------------------------------------
// Font built from glyphs
// --------------------------------------------------------------

type vectorFont struct {
	vectors []*vectorGlyph
	glyphs  map[byte]int // Maps into glyph array

	horizontalOffset float64
	verticalOffset   float64
	scale            float64
}

// NewVectorFont constructs an IVectorFont object
func NewVectorFont() api.IVectorFont {
	o := new(vectorFont)
	o.vectors = []*vectorGlyph{}
	o.glyphs = make(map[byte]int)
	o.horizontalOffset = 1.2
	o.verticalOffset = 1.2
	o.scale = 3.0
	return o
}

func (v *vectorFont) Initialize(dataFile string, relativePath string) {
	dataPath, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(dataPath + "/assets/" + dataFile)
	defer file.Close()

	if err != nil {
		log.Fatalf("VectorFont: failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	fmt.Println("Opened vector font file")

	var glyph *vectorGlyph
	idx := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 1 {
			// Add character to glyph dictionary
			gIdx := byte(line[0])
			v.glyphs[gIdx] = idx
			// Start new glyph for character
			glyph = newVectorGlyph()
			idx++
			continue
		}

		if line == "||" {
			// Finished glyph vector
			v.vectors = append(v.vectors, glyph)
			continue
		}

		// readlines until end of pixel marker: "||"
		if line != "||" {
			ele := strings.Split(line, " ")
			v1, _ := strconv.ParseFloat(ele[0], 64)
			v2, _ := strconv.ParseFloat(ele[1], 64)
			v3, _ := strconv.ParseFloat(ele[2], 64)
			v4, _ := strconv.ParseFloat(ele[3], 64)
			glyph.addVector(v1, v2, v3, v4)
		}
	}
}

func (v *vectorFont) HorizontalOffset() float64 {
	return v.horizontalOffset
}

func (v *vectorFont) VerticalOffset() float64 {
	return v.verticalOffset
}

func (v *vectorFont) Scale() float64 {
	return v.scale
}

func (v *vectorFont) Glyph(char byte) []float64 {
	glyph := v.glyphs[char]
	return v.vectors[glyph].vertices
}
