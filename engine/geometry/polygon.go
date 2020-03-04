package geometry

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
)

type polygon struct {
	mesh api.IMesh
}

// NewPolygon constructs a new IPolygon
func NewPolygon() api.IPolygon {
	o := new(polygon)
	o.mesh = NewMesh()
	return o
}

func (p *polygon) AddVertex(x, y float64) {
	p.mesh.AddVertex(x, y)
}

func (p *polygon) Mesh() api.IMesh {
	return p.mesh
}

func (p *polygon) Build() {
	p.mesh.Build()
}

// PointInside will return false if the point is on the right/bottom edge and/or outward
// if point is on the right and/or bottom edge it is considered outside and
// the left/top edge is considered inside.
// This is consistant with polygon filling.
func (p *polygon) PointInside(po api.IPoint) bool {
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

func (p polygon) String() string {
	return fmt.Sprintf("%s", p.mesh)
}
