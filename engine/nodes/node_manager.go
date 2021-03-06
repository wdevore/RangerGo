package nodes

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
)

type nodeManager struct {
	world api.IWorld

	clearBackground bool

	// Stack of node
	stack *nodeStack

	timingTargets api.INodeList
	eventTargets  api.INodeList
}

// NewNodeManager constructs a manager for node.
// It manages the lifecycle and events
func NewNodeManager(world api.IWorld) api.INodeManager {
	o := new(nodeManager)
	o.world = world

	// It is very rare that the manager would clear the background
	// because almost all nodes will handle clearing/painting their
	// own backgrounds.
	o.clearBackground = false

	o.stack = newNodeStack()

	o.timingTargets = NewNodeList()
	o.eventTargets = NewNodeList()

	return o
}

func (m *nodeManager) ClearEnabled(clear bool) {
	m.clearBackground = clear
}

func (m *nodeManager) PreVisit() {
	// Typically Scenes/Layers will clear the background themselves so the default
	// is to NOT perform a clear here.
	if m.clearBackground {
		// If vsync is enabled then this takes nearly 1/fps milliseconds.
		// For example, 60fps -> 1/60 = ~16.666ms
		m.world.Context().Pre()
	}
}

func (m *nodeManager) Visit(interpolation float64) bool {
	if m.stack.isEmpty() {
		fmt.Println("NodeManager: no more nodes to visit.")
		return false
	}

	// fmt.Println("NodeManager: visiting ", m.stack.runningNode)

	if m.stack.hasNextNode() {
		m.setNextNode()
	}

	context := m.world.Context()

	// This saves view-space matrix
	context.Save()

	runningScene := m.stack.runningNode.(api.IScene)

	action := runningScene.TransitionAction()

	if action == api.SceneReplaceTake {
		repl := runningScene.GetReplacement()
		// fmt.Println("NodeManager: SceneReplaceTake with ", repl)
		if repl != nil {
			m.stack.replace(repl)
			// Immediately switch to the new replacement node
			if m.stack.hasNextNode() {
				m.setNextNode()
			}
		} else {
			m.exitNodes(m.stack.runningNode)
			m.stack.pop()
		}
	}

	// Visit the running node
	Visit(m.stack.runningNode, context, interpolation)

	context.Restore()

	return true // continue to draw.
}

func (m *nodeManager) PostVisit() {
	m.world.Context().Post()
}

func (m *nodeManager) PopNode() api.INode {
	return m.stack.pop()
}

func (m *nodeManager) PushNode(node api.INode) {
	m.stack.nextNode = node
	m.stack.push(node)
}

func (m *nodeManager) ReplaceNode(node api.INode) {
	m.stack.replace(node)
}

// --------------------------------------------------------------------------
// Timing
// --------------------------------------------------------------------------

func (m *nodeManager) Update(msPerUpdate, secPerUpdate float64) {
	for _, target := range m.timingTargets.Items() {
		target.Update(msPerUpdate, secPerUpdate)
	}
}

func (m *nodeManager) RegisterTarget(target api.INode) {
	m.timingTargets.Add(target)
}

func (m *nodeManager) UnRegisterTarget(target api.INode) {
	m.timingTargets.Remove(target)
}

// --------------------------------------------------------------------------
// IO events
// --------------------------------------------------------------------------

func (m *nodeManager) RegisterEventTarget(target api.INode) {
	m.eventTargets.Add(target)
}

func (m *nodeManager) UnRegisterEventTarget(target api.INode) {
	m.eventTargets.Remove(target)
}

func (m *nodeManager) RouteEvents(event api.IEvent) {
	if m.eventTargets == nil {
		return
	}

	for _, target := range m.eventTargets.Items() {
		handled := target.Handle(event)

		if handled {
			break
		}
	}
}

func (m *nodeManager) setNextNode() {
	if m.stack.hasRunningNode() {
		m.exitNodes(m.stack.runningNode)
	}

	m.stack.runningNode = m.stack.nextNode
	m.stack.clearNextNode()

	// fmt.Println("NodeManager: new running node ", m.stack.runningNode)

	m.enterNodes(m.stack.runningNode)
}

// End cleans up NodeManager by clearing the stack and calling all Exits
func (m *nodeManager) End() {
	// Dump the stack

	n := m.PopNode()

	for n != nil {
		m.exitNodes(n)
		n = m.PopNode()
	}

	m.eventTargets = nil
}

// -----------------------------------------------------
// Scene lifecycles
// -----------------------------------------------------

func (m *nodeManager) enterNodes(node api.INode) {
	// fmt.Println("NodeManager: enter-node ", node)
	node.EnterNode(m)

	children := node.Children()
	for _, child := range children {
		m.enterNodes(child)
	}
}

func (m *nodeManager) exitNodes(node api.INode) {
	// fmt.Println("NodeManager: exit-node ", node)
	node.ExitNode(m)

	children := node.Children()
	for _, child := range children {
		m.exitNodes(child)
	}
}

func (m *nodeManager) Debug() {
}

func (m nodeManager) String() string {
	return fmt.Sprintf("%s", m.stack)
}

// DeleteAt removes an item from the slice
func DeleteAt(i int, slice []api.INode) {
	// Remove the element at index i from slice.
	copy(slice[i:], slice[i+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = nil    // Erase last element (write zero value).
	slice = slice[:len(slice)-1] // Truncate slice.
}

// FindFirstElement finds the first item in the slice
func FindFirstElement(node api.INode, slice []api.INode) int {
	for idx, item := range slice {
		if item.ID() == node.ID() {
			return idx
		}
	}

	return -1
}
