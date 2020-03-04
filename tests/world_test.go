package lerps

import (
	"testing"

	"github.com/wdevore/RangerGo/engine"
)

func TestRunner(t *testing.T) {
	worldTest(t)
}

func worldTest(t *testing.T) {
	engine.NewWorld("Ranger world")
}
