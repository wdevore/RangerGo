package lerps

import (
	// "fmt"

	"fmt"
	"testing"

	"github.com/wdevore/RangerGo/engine/rendering"
)

func TestRunner(t *testing.T) {
	c := rendering.NewPaletteRGB(255, 0, 0)

	fmt.Println(c)

	c = rendering.NewPaletteInt64(rendering.Aqua)
	fmt.Println(c)

}
