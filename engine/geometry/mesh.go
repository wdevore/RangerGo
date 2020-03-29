package geometry

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
)

// Mesh is a polygon object
type Mesh struct {
	// Original vertices
	vertices []api.IPoint
	// Transformed vertices
	bucket []api.IPoint
}

// NewMesh constructs a new IMesh
func NewMesh() *Mesh {
	o := new(Mesh)
	o.vertices = []api.IPoint{}
	return o
}

// Vertices are the untransformed vertices
func (m *Mesh) Vertices() []api.IPoint {
	return m.vertices
}

// Bucket is the transformed vertices
func (m *Mesh) Bucket() []api.IPoint {
	return m.bucket
}

// AddVertex adds a vertex to the mesh
func (m *Mesh) AddVertex(x, y float64) {
	m.vertices = append(m.vertices, NewPointUsing(x, y))
}

// SetVertex updates a vertex on the mesh
func (m *Mesh) SetVertex(x, y float64, index int) {
	m.vertices[index].SetByComp(x, y)
}

// Build sizes the transform bucket
func (m *Mesh) Build() {
	// bucket needs to be the same size as vertices
	m.bucket = make([]api.IPoint, len(m.vertices))
	for i := range m.vertices {
		m.bucket[i] = NewPoint()
	}
}

func (m Mesh) String() string {
	s := ""
	for _, v := range m.vertices {
		s += fmt.Sprintf("%v\n", v)
	}

	return s
}
