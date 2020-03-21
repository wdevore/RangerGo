package nodes

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/maths"
)

var tViewPoint = geometry.NewPoint()
var comp = maths.NewTransform()
var out = maths.NewTransform()

// MapDeviceToView maps mouse-space device coordinates to view-space
func MapDeviceToView(world api.IWorld, dvx, dvy int32, viewPoint api.IPoint) {
	world.InvViewSpace().TransformCompToPoint(float64(dvx), float64(dvy), viewPoint)
}

// MapDeviceToNode maps mouse-space device coordinates to local node-space
func MapDeviceToNode(world api.IWorld, dvx, dvy int32, node api.INode, localPoint api.IPoint) {
	// Mapping from device to node requires to transforms from to "directions"
	// 1st is upwards transform and the 2nd is downwards transform.

	// downwards from device-space to view-space
	MapDeviceToView(world, dvx, dvy, tViewPoint)

	// Upwards from node to world-space (aka view-space)
	wtn := WorldToNodeTransform(node, nil)

	// Now map view-space point to local-space of node
	wtn.TransformCompToPoint(tViewPoint.X(), tViewPoint.Y(), localPoint)

	// Optional scaling
	// localPoint.SetByComp(localPoint.X() * node.Scale(), localPoint.Y() * node.Scale())
}

// WorldToNodeTransform maps a world-space coordinate to local-space of node
func WorldToNodeTransform(node api.INode, psuedoRoot api.INode) api.IAffineTransform {
	wtn := NodeToWorldTransform(node, psuedoRoot)
	wtn.Invert()
	return wtn
}

// NodeToWorldTransform maps a local-space coordinate to world-space
func NodeToWorldTransform(node api.INode, psuedoRoot api.INode) api.IAffineTransform {
	aft := node.CalcTransform()

	// A transform to accumulate the parent transforms.
	comp.SetByTransform(aft)

	// Iterate "upwards" starting with the child towards the parents
	// starting with this child's parent.
	p := node.Parent()

	for p != nil {
		parentT := p.CalcTransform()

		// Because we are iterating upwards we need to pre-multiply each child.
		// Ex: [child] x [parent]
		// ----------------------------------------------------------
		//           [comp] x [parentT]
		//               |
		//               | out
		//               v
		//             [comp] x [parentT]
		//                 |
		//                 | out
		//                 v
		//               [comp] x [parentT...]
		//
		// This is a pre-multiply order
		// [child] x [parent of child] x [parent of parent of child]...
		//
		// In other words the child is mutiplied "into" the parent.
		maths.Multiply(comp, parentT, out)
		comp.SetByTransform(out)

		if p == psuedoRoot {
			fmt.Println("SpaceMappings: hit psuedoRoot")
			break
		}

		// next parent upwards
		p = p.Parent()
	}

	return comp
}
