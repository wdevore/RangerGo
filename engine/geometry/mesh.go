package geometry

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
)

type mesh struct {
	// Original vertices
	vertices []api.IPoint
	// Transformed vertices
	bucket []api.IPoint
}

// NewMesh constructs a new IMesh
func NewMesh() api.IMesh {
	o := new(mesh)
	o.vertices = []api.IPoint{}
	return o
}

func (m *mesh) Vertices() []api.IPoint {
	return m.vertices
}

func (m *mesh) AddVertex(x, y float64) {
	m.vertices = append(m.vertices, NewPointUsing(x, y))
}

func (m *mesh) Build() {
	// bucket needs to be the same size as vertices
	m.bucket = make([]api.IPoint, len(m.vertices))
}

func (m mesh) String() string {
	s := ""
	for _, v := range m.vertices {
		s += fmt.Sprintf("%v\n", v)
	}

	return s
}
