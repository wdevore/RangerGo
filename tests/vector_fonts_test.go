package vectorfonts

import (
	// "fmt"

	"testing"

	"github.com/wdevore/RangerGo/engine/rendering"
)

func TestRunner(t *testing.T) {
	vf := rendering.NewVectorFont()
	vf.Initialize("vector_font.data")
}
