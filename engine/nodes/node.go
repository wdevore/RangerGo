package nodes

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
)

var ids = 0

// Node is an embedded type used by all nodes.
type Node struct {
	id      int
	name    string
	visible bool
	dirty   bool

	parent api.INode

	Transform
	Group
}

// ID returns the internally generated Id.
func (n *Node) ID() int {
	return n.id
}

// Initialize called by base objects from their Initialize
func (n *Node) Initialize(name string) {
	n.id = ids
	ids++
	n.name = name
	n.visible = true
	n.dirty = true

	n.initializeTransform()
	n.initializeGroup()
}

// InitializeWithID called by base objects from their Initialize
func (n *Node) InitializeWithID(id int, name string) {
	n.id = id
	n.name = name
	n.visible = true
	n.dirty = true
}

// Visit traverses "down" the heirarchy while space-mappings traverses upward.
func (n *Node) Visit(context api.IRenderContext, interpolation float64) {
	if !n.IsVisible() {
		return
	}

	context.Save()

	// Because position and angles are dependent
	// on lerping we perform interpolation first.
	n.Interpolate(interpolation)

	aft := n.CalcTransform()

	context.Apply(aft)

	children := n.Children()

	if children != nil {
		n.Draw(context)

		for _, child := range children {
			child.Visit(context, interpolation)
		}
	} else {
		n.Draw(context)
	}

	context.Restore()
}

// SetParent binds upward parent.
func (n *Node) SetParent(parent api.INode) {
	n.parent = parent
}

// Parent returns any defined parent
func (n *Node) Parent() api.INode {
	return n.parent
}

// IsVisible indicates visibility, default is "true"
func (n *Node) IsVisible() bool {
	return n.visible
}

// Interpolate is used for blending time based properties.
func (n *Node) Interpolate(interpolation float64) {
	fmt.Println("Node Interpolate")
}

// IsDirty indicates if the node has been modified.
func (n *Node) IsDirty() bool {
	return n.dirty
}

// SetDirty marks a node dirty state.
func (n *Node) SetDirty(dirty bool) {
	n.dirty = dirty
}

// RippleDirty propagates a dirty state to children.
func (n *Node) RippleDirty(dirty bool) {
	for _, child := range n.Children() {
		child.RippleDirty(dirty)
	}

	n.SetDirty(dirty)
}

// Update updates the time properties of a node.
func (n *Node) Update(dt float64) {
}

// Draw renders a node
func (n *Node) Draw(context api.IRenderContext) {
}

// GetBucket returns a buffer for capturing transformed vertices
func (n *Node) GetBucket() []api.IPoint {
	return nil
}

// -----------------------------------------------------
// Scene lifecycles
// -----------------------------------------------------

// EnterNode called when a node is entering the stage
func (n *Node) EnterNode(api.INodeManager) {
}

// ExitNode called when a node is exiting stage
func (n *Node) ExitNode(api.INodeManager) {
}

// -----------------------------------------------------
// ITransform defaults
// -----------------------------------------------------

// CalcTransform calculates a matrix based on the
// current transform properties
func (n *Node) CalcTransform() api.IAffineTransform {
	// aft := n.transform.aft

	if n.IsDirty() {
		pos := n.position
		n.aft.MakeTranslate(pos.X(), pos.Y())

		rot := n.rotation
		if rot != 0.0 {
			n.aft.Rotate(rot)
		}

		s := n.Scale()
		if s != 1.0 {
			n.aft.Scale(s, s)
		}

		// Invert...
		n.aft.InvertTo(n.inverse)
	}

	return n.aft
}

var p = geometry.NewPoint()

// SetPosition overrides transform's method
func (n *Node) SetPosition(x, y float64) {
	n.SetPosition(x, y)
	n.RippleDirty(true)
}

// SetRotation overrides transform's method
func (n *Node) SetRotation(radians float64) {
	n.SetRotation(radians)
	n.RippleDirty(true)
}

// SetScale overrides transform's method
func (n *Node) SetScale(scale float64) {
	n.SetScale(scale)
	n.RippleDirty(true)
}

// -------------------------------------------------------------------
// INodeGroup implementations
// -------------------------------------------------------------------

// -------------------------------------------------------------------
// Misc
// -------------------------------------------------------------------

// PrintTree prints the tree relative to this node.
func PrintTree(node api.INode) {
	fmt.Println("---------- Tree ---------------")
	printBranch(0, node)

	children := node.Children()
	if children != nil {
		printSubTree(children, 1)
	}

	fmt.Println("-------------------------------")
}

func printSubTree(children []api.INode, level int) {
	for _, child := range children {
		subChildren := child.Children()
		printBranch(level, child)
		if subChildren != nil {
			printSubTree(subChildren, level+1)
		}
	}
}

const indent = "   "

func printBranch(level int, node api.INode) {
	for i := 0; i < level; i++ {
		fmt.Print(indent)
	}
	fmt.Println(node)
}

func (n Node) String() string {
	return fmt.Sprintf("|'%s' (%d)", n.name, n.id)
}
