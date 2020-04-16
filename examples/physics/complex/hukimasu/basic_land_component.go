package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

// BasicLandCompoent represents both the visual and physic components
type BasicLandCompoent struct {
	land api.INode

	b2Body *box2d.B2Body

	categoryBits uint16 // I am a...
	maskBits     uint16 // I can collide with a...
}

// NewBasicLandCompoent constructs a component
func NewBasicLandCompoent(name string, parent api.INode) *BasicLandCompoent {
	o := new(BasicLandCompoent)
	o.land = custom.NewPolygonNode(name, parent.World(), parent)
	glo := o.land.(*custom.PolygonNode)
	glo.EnableHitDetection(false)

	return o
}

// Configure component
func (b *BasicLandCompoent) Configure(sx, sy, posx, posy float64, categoryBits, maskBits uint16, b2World *box2d.B2World) {
	gl := b.land.(*custom.PolygonNode)
	gl.AddVertex(-30.0*sx+posx, -2.5*sy+posy, false)
	gl.AddVertex(-25.0*sx+posx, -2.5*sy+posy, false)
	gl.AddVertex(-20.0*sx+posx, -7.5*sy+posy, false)
	gl.AddVertex(-10.0*sx+posx, -7.5*sy+posy, false)
	gl.AddVertex(-10.0*sx+posx, -5.5*sy+posy, false)
	gl.AddVertex(-5.0*sx+posx, -2.5*sy+posy, false)
	gl.AddVertex(1.0*sx+posx, -2.5*sy+posy, false)
	gl.AddVertex(1.5*sx+posx, -1.0*sy+posy, false)
	gl.AddVertex(7.0*sx+posx, -1.0*sy+posy, false)
	gl.AddVertex(7.5*sx+posx, -0.0*sy+posy, false)
	gl.AddVertex(10.5*sx+posx, -0.0*sy+posy, false)
	gl.AddVertex(11.0*sx+posx, -2.0*sy+posy, false)
	gl.AddVertex(11.0*sx+posx, -5.0*sy+posy, false)
	gl.AddVertex(15.0*sx+posx, -10.0*sy+posy, false)
	gl.AddVertex(20.0*sx+posx, -10.0*sy+posy, true)
	gl.SetOpen(true)

	b.categoryBits = categoryBits
	b.maskBits = maskBits

	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_staticBody

	// An instance of a body to contain Fixtures
	b.b2Body = b2World.CreateBody(&bDef)

	b2ChainShape := box2d.MakeB2ChainShape()

	vertices := []box2d.B2Vec2{}
	verts := gl.Polygon().Mesh().Vertices()

	for _, v := range verts {
		vertices = append(vertices, box2d.B2Vec2{X: v.X(), Y: v.Y()})
	}

	b2ChainShape.CreateChain(vertices, len(vertices))

	fd := box2d.MakeB2FixtureDef()
	fd.Shape = &b2ChainShape
	fd.UserData = b.land
	fd.Filter.CategoryBits = b.categoryBits
	fd.Filter.MaskBits = b.maskBits

	b.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}

// SetColor sets the visual's color
func (b *BasicLandCompoent) SetColor(color api.IPalette) {
	gr := b.land.(*custom.PolygonNode)
	gr.SetColor(color)
}
