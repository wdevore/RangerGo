package main

import (
	"github.com/ByteArena/box2d"
	"github.com/wdevore/RangerGo/api"
)

type filterListener struct {
	components []api.IFilterListener
}

func newFilterListener() box2d.B2ContactFilterInterface {
	o := new(filterListener)
	o.components = []api.IFilterListener{}
	return o
}

func (c *filterListener) addListener(component api.IFilterListener) {
	c.components = append(c.components, component)
}

func (c *filterListener) ShouldCollide(fixtureA *box2d.B2Fixture, fixtureB *box2d.B2Fixture) bool {

	// ------------------------------------------------
	// Custom filtering
	// ------------------------------------------------
	// canCollide := (fixtureB.GetFilterData().CategoryBits & fixtureA.GetFilterData().MaskBits) != 0

	// ------------------------------------------------
	// Standard filtering
	// ------------------------------------------------
	filterA := fixtureA.GetFilterData()
	filterB := fixtureB.GetFilterData()

	// ------------------------------------------------
	// Show report to understand filtering
	// ------------------------------------------------
	// fmt.Println("------------------")
	// dataA := fixtureA.GetUserData()
	// dataB := fixtureB.GetUserData()

	// if dataA != nil && dataB != nil {
	// 	a := dataA.(api.INode)
	// 	b := dataB.(api.INode)
	// 	s := fmt.Sprintf("%s, %s\n", a, b)
	// 	fmt.Println(s)
	// }

	// s := fmt.Sprintf("ABits: %016b, %016b (%d, %d)", filterA.CategoryBits, filterA.MaskBits, filterA.CategoryBits, filterA.MaskBits)
	// fmt.Println(s)
	// s = fmt.Sprintf("BBits: %016b, %016b (%d, %d)", filterB.CategoryBits, filterB.MaskBits, filterB.CategoryBits, filterB.MaskBits)
	// fmt.Println(s)

	// // B (I am) to A (collide with)
	fBCtofAM := filterB.CategoryBits & filterA.MaskBits
	// fmt.Println("fBCtofAM: ", fBCtofAM)

	// // A (I am) to B (collide with)
	fACtofBM := filterA.CategoryBits & filterB.MaskBits
	// fmt.Println("fACtofBM: ", fACtofBM)

	if filterA.GroupIndex == filterB.GroupIndex && filterA.GroupIndex != 0 {
		// fmt.Println("GroupIndex triggered: ", filterA.GroupIndex > 0)
		return filterA.GroupIndex > 0
	}

	canCollide := fBCtofAM != 0 && fACtofBM != 0

	return canCollide
}
