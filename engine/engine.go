package engine

import "github.com/wdevore/RangerGo/api"

type engine struct {
	// App window size
	width  int
	height int
}

// New constructs a Engine object.
// The Engine runs the main loop.
func New(width, height int) api.IEngine {
	o := new(engine)
	o.width = width
	o.height = height
	return o
}

func (e *engine) Configure() {

}

func (e *engine) Start() {

}

// DisplaySize see api.go for docs
func (e *engine) DisplaySize() (w, h int) {
	return e.width, e.height
}
