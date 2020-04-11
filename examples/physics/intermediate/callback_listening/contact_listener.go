package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
)

const (
	entityBoundary  = uint16(0x0001)
	entityCircle    = uint16(0x0002)
	entityTriangle  = uint16(0x0004)
	entityRectangle = uint16(0x0008)
)

// ContactListener
type contactListener struct {
	components []api.IContactListener
}

func newContactListener() box2d.B2ContactListenerInterface {
	o := new(contactListener)
	o.components = []api.IContactListener{}
	return o
}

func (c *contactListener) addListener(component api.IContactListener) {
	c.components = append(c.components, component)
}

func (c *contactListener) BeginContact(contact box2d.B2ContactInterface) {
	// --------------------------------------
	// Note: This callback is called only if the filter-callback passed
	// back a true, meaning we can collide, and hence we now react to
	// the collision. Is this example I simply repeat what the filter
	// was showing and print a simple report.
	// --------------------------------------

	fixA := contact.GetFixtureA()
	fixB := contact.GetFixtureB()

	dataA := fixA.GetUserData()
	dataB := fixB.GetUserData()

	// Here the question is: can fixB collide with fixA
	canCollide := (fixB.GetFilterData().CategoryBits & fixA.GetFilterData().MaskBits) != 0

	if canCollide {
		if dataA != nil && dataB != nil {
			// a := dataA.(api.INode)
			// b := dataB.(api.INode)
			// s := fmt.Sprintf("%s can collide with %s\n", a, b) // uncomment me
			// fmt.Println(s)
		}
	} else {
		if dataA != nil && dataB != nil {
			// a := dataA.(api.INode)
			// b := dataB.(api.INode)
			// s := fmt.Sprintf("%s can NOT collide with %s\n", a, b)// uncomment me
			// fmt.Println(s)
		}
	}

	if dataA != nil && dataB != nil {
		for _, l := range c.components {
			handled := l.HandleBeginContact(dataA.(api.INode), dataB.(api.INode))
			if handled {
				break
			}
		}
	}
}

func (c *contactListener) EndContact(contact box2d.B2ContactInterface) {
	fixA := contact.GetFixtureA()
	fixB := contact.GetFixtureB()

	dataA := fixA.GetUserData()
	dataB := fixB.GetUserData()

	if dataA != nil && dataB != nil {
		for _, l := range c.components {
			handled := l.HandleEndContact(dataA.(api.INode), dataB.(api.INode))
			if handled {
				break
			}
		}
	}
}

func (c *contactListener) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {
	// fmt.Println("PreSolve")

}

func (c *contactListener) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {
	// fmt.Println("PostSolve")

}
