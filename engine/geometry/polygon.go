package geometry

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
)

// Polygon is a mesh with additional methods
type Polygon struct {
	mesh api.IMesh
}

// NewPolygon constructs a new IPolygon
func NewPolygon() api.IPolygon {
	o := new(Polygon)
	o.mesh = NewMesh()
	return o
}

// AddVertex adds a point to the mesh
func (p *Polygon) AddVertex(x, y float64) {
	p.mesh.AddVertex(x, y)
}

// Mesh provides access to the underlying mesh
func (p *Polygon) Mesh() api.IMesh {
	return p.mesh
}

// Build builds mesh after it has been define.
func (p *Polygon) Build() {
	p.mesh.Build()
}

// PointInside will return false if the point is on the right/bottom edge and/or outward
// if point is on the right and/or bottom edge it is considered outside and
// the left/top edge is considered inside.
// This is consistant with polygon filling.
func (p *Polygon) PointInside(po api.IPoint) bool {
	i := 0
	c := false
	vertices := p.mesh.Vertices()
	nvert := len(vertices)
	j := 1

	for j < nvert {
		if ((vertices[i].Y() > po.Y()) != (vertices[j].Y() > po.Y())) &&
			(po.X() < (vertices[j].X()-vertices[i].X())*(po.Y()-vertices[i].Y())/
				(vertices[j].Y()-vertices[i].Y())+vertices[i].X()) {
			c = !c
		}
		i++
		j++
	}

	// Last edge to close loop
	i = j - 1
	j = 0
	if ((vertices[i].Y() > po.Y()) != (vertices[j].Y() > po.Y())) &&
		(po.X() < (vertices[j].X()-vertices[i].X())*(po.Y()-vertices[i].Y())/
			(vertices[j].Y()-vertices[i].Y())+vertices[i].X()) {
		c = !c
	}

	return c
}

func (p Polygon) String() string {
	return fmt.Sprintf("%s", p.mesh)
}
