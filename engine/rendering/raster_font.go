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
// Font built from glyphs
// --------------------------------------------------------------

type rasterFont struct {
	pixels [][]uint8
	glyphs map[byte]int // Maps into glyph array
}

// NewRasterFont constructs an IRasterFont object
func NewRasterFont() api.IRasterFont {
	o := new(rasterFont)
	o.glyphs = make(map[byte]int)
	return o
}

func (r *rasterFont) Initialize(dataFile string, relativePath string) {
	dataPath, err := filepath.Abs(relativePath)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(dataPath + "/assets/" + dataFile)
	defer file.Close()

	if err != nil {
		log.Fatalf("RasterFont: failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	fmt.Println("Opened raster font file")

	idx := 0

	for scanner.Scan() {
		line := scanner.Text()

		// ele[0] is the ascill character itself,
		// the rest of the line contains the pixels
		ele := strings.Split(line, " ")

		// Add character to glyph dictionary
		gIdx := byte(ele[0][0])
		r.glyphs[gIdx] = idx
		idx++

		// Add data to raw data array
		px1, _ := strconv.ParseInt(ele[1], 0, 8)
		px2, _ := strconv.ParseInt(ele[2], 0, 8)
		px3, _ := strconv.ParseInt(ele[3], 0, 8)
		px4, _ := strconv.ParseInt(ele[4], 0, 8)
		px5, _ := strconv.ParseInt(ele[5], 0, 8)
		px6, _ := strconv.ParseInt(ele[6], 0, 8)
		px7, _ := strconv.ParseInt(ele[7], 0, 8)
		px8, _ := strconv.ParseInt(ele[8], 0, 8)

		px := []uint8{uint8(px1), uint8(px2), uint8(px3), uint8(px4), uint8(px5), uint8(px6), uint8(px7), uint8(px8)}

		r.pixels = append(r.pixels, px)
	}
}

func (r *rasterFont) Glyph(char byte) []uint8 {
	glyph := r.glyphs[char]
	return r.pixels[glyph]
}

func (r *rasterFont) GlyphWidth() int {
	return 8
}
