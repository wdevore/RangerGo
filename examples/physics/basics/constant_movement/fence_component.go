package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/nodes/custom"
)

// FenceComponent represents both the visual and physic components
type FenceComponent struct {
	bottom api.INode
	right  api.INode
	top    api.INode
	left   api.INode

	b2Body  *box2d.B2Body
	b2Shape box2d.B2EdgeShape
}

// NewFenceComponent constructs a component
func NewFenceComponent(name string, parent api.INode) *FenceComponent {
	o := new(FenceComponent)

	o.bottom = custom.NewLineNode(name, parent.World(), parent)
	gln := o.bottom.(*custom.LineNode)
	gln.SetPoints(-1.0, 0.0, 1.0, 0.0) // Set by unit coordinates

	o.right = custom.NewLineNode(name, parent.World(), parent)
	gln = o.right.(*custom.LineNode)
	gln.SetPoints(0.0, 1.0, 0.0, -1.0) // Set by unit coordinates

	o.top = custom.NewLineNode(name, parent.World(), parent)
	gln = o.top.(*custom.LineNode)
	gln.SetPoints(1.0, 0.0, -1.0, 0.0) // Set by unit coordinates

	o.left = custom.NewLineNode(name, parent.World(), parent)
	gln = o.left.(*custom.LineNode)
	gln.SetPoints(0.0, -1.0, 0.0, 1.0) // Set by unit coordinates

	return o
}

// Configure component
func (f *FenceComponent) Configure(b2World *box2d.B2World) {
	// A body def used to create bodies
	bDef := box2d.MakeB2BodyDef()
	bDef.Type = box2d.B2BodyType.B2_staticBody

	// An instance of a body to contain Fixtures
	f.b2Body = b2World.CreateBody(&bDef)
}

// SetColor sets the visual's color
func (f *FenceComponent) SetColor(color api.IPalette) {
	gr := f.bottom.(*custom.LineNode)
	gr.SetColor(color)
	gr = f.right.(*custom.LineNode)
	gr.SetColor(color)
	gr = f.top.(*custom.LineNode)
	gr.SetColor(color)
	gr = f.left.(*custom.LineNode)
	gr.SetColor(color)
}

// SetScale sets the component's length
func (f *FenceComponent) SetScale(scale float64) {
	f.bottom.SetPosition(0.0, scale)
	f.bottom.SetScale(scale)
	f.right.SetPosition(scale, 0.0)
	f.right.SetScale(scale)
	f.top.SetPosition(0.0, -scale)
	f.top.SetScale(scale)
	f.left.SetPosition(-scale, 0.0)
	f.left.SetScale(scale)

	fd := box2d.MakeB2FixtureDef()

	// Bottom fixture
	f.b2Shape = box2d.MakeB2EdgeShape()
	f.b2Shape.Set(box2d.MakeB2Vec2(-scale, scale), box2d.MakeB2Vec2(scale, scale))
	fd.Shape = &f.b2Shape
	f.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body

	// Right fixture
	f.b2Shape = box2d.MakeB2EdgeShape()
	f.b2Shape.Set(box2d.MakeB2Vec2(scale, scale), box2d.MakeB2Vec2(scale, -scale))
	fd.Shape = &f.b2Shape
	f.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body

	// Top fixture
	f.b2Shape = box2d.MakeB2EdgeShape()
	f.b2Shape.Set(box2d.MakeB2Vec2(scale, -scale), box2d.MakeB2Vec2(-scale, -scale))
	fd.Shape = &f.b2Shape
	f.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body

	// Left fixture
	f.b2Shape = box2d.MakeB2EdgeShape()
	f.b2Shape.Set(box2d.MakeB2Vec2(-scale, -scale), box2d.MakeB2Vec2(-scale, scale))
	fd.Shape = &f.b2Shape
	f.b2Body.CreateFixtureFromDef(&fd) // attach Fixture to body
}
