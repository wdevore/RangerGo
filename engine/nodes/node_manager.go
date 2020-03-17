package nodes

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type nodeManager struct {
	clearBackground bool

	context api.IRenderContext

	// Stack of node
	stack *nodeStack

	timingTargets []api.INode
	eventTargets  []api.INode
}

// NewNodeManager constructs a manager for node.
// It manages the lifecycle and events
func NewNodeManager(world api.IWorld) api.INodeManager {
	o := new(nodeManager)
	o.clearBackground = true

	o.context = rendering.NewRenderContext(world)
	o.context.Initialize()

	o.stack = newNodeStack()

	o.timingTargets = []api.INode{}
	o.eventTargets = []api.INode{}

	return o
}

func (m *nodeManager) PreVisit() {
	// Typically Scenes/Layers will clear the background themselves so the default
	// is to NOT perform a clear here.
	if m.clearBackground {
		// If vsync is enabled then this takes nearly 1/fps milliseconds.
		// For example, 60fps -> 1/60 = ~16.666ms
		m.context.Pre()
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

	// This saves view-space matrix
	m.context.Save()

	// DEBUG
	// If mouse coords changed then update view coords.
	// self.global_data.update_view_coords(&mut self.context);

	runningScene := m.stack.runningNode.(api.IScene)

	action := runningScene.TransitionAction()

	if action == api.SceneReplaceTake {
		repl := runningScene.GetReplacement()
		fmt.Println("NodeManager: SceneReplaceTake with ", repl)
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
	Visit(m.stack.runningNode, m.context, interpolation)

	m.context.Restore()

	// fmt.Println("NodeManager: done visiting ", m.stack.runningNode)
	return true // continue to draw.
}

func (m *nodeManager) PostVisit() {
	m.context.Post()
}

func (m *nodeManager) PopNode() {
	m.stack.pop()
}

func (m *nodeManager) PushNode(node api.INode) {
	m.stack.nextNode = node
	// fmt.Println("NodeManager: pushing ", node)
	m.stack.push(node)
}

func (m *nodeManager) ReplaceNode(node api.INode) {
	m.stack.replace(node)
}

// --------------------------------------------------------------------------
// Timing
// --------------------------------------------------------------------------

func (m *nodeManager) Update(dt float64) {
	for _, target := range m.timingTargets {
		target.Update(dt)
	}
}

func (m *nodeManager) RegisterTarget(target api.INode) {
	fmt.Println("NodeManager registering ", target)
	m.timingTargets = append(m.timingTargets, target)
}

func (m *nodeManager) UnRegisterTarget(target api.INode) {
	idx := findFirstElement(target, m.timingTargets)

	if idx >= 0 {
		fmt.Println("UnRegistering idx:(", idx, ") ", m.timingTargets[idx], " target")
		deleteAt(idx, m.timingTargets)
	} else {
		fmt.Println("Unable to UnRegister ", target, " target")
	}
}

// --------------------------------------------------------------------------
// IO events
// --------------------------------------------------------------------------

func (m *nodeManager) RegisterEventTarget(target api.INode) {
	fmt.Println("Register ", target, " event target")
	m.eventTargets = append(m.eventTargets, target)
}

func (m *nodeManager) UnRegisterEventTarget(target api.INode) {
	idx := findFirstElement(target, m.eventTargets)

	if idx >= 0 {
		fmt.Println("UnRegistering event idx:(", idx, ") ", m.eventTargets[idx], " target")
		deleteAt(idx, m.eventTargets)
	} else {
		fmt.Println("Unable to UnRegister event ", target, " target")
	}
}

func findFirstElement(node api.INode, slice []api.INode) int {
	for idx, item := range slice {
		if item.ID() == node.ID() {
			return idx
		}
	}

	return -1
}

func (m *nodeManager) RouteEvents(event api.IEvent) {
	for _, target := range m.eventTargets {
		handled := target.Handle(event)

		if handled {
			break
		}
	}
}

func deleteAt(i int, slice []api.INode) {
	// Remove the element at index i from a.
	copy(slice[i:], slice[i+1:]) // Shift a[i+1:] left one index.
	slice[len(slice)-1] = nil    // Erase last element (write zero value).
	slice = slice[:len(slice)-1] // Truncate slice.
}

func (m *nodeManager) setNextNode() {
	if m.stack.hasRunningNode() {
		m.exitNodes(m.stack.runningNode)
	}

	m.stack.runningNode = m.stack.nextNode
	m.stack.clearNextNode()

	fmt.Println("NodeManager: new running node ", m.stack.runningNode)

	m.enterNodes(m.stack.runningNode)
}

// -----------------------------------------------------
// Scene lifecycles
// -----------------------------------------------------

func (m *nodeManager) enterNodes(node api.INode) {
	fmt.Println("NodeManager: enter-node ", node)
	node.EnterNode(m)

	children := node.Children()
	for _, child := range children {
		m.enterNodes(child)
	}
}

func (m *nodeManager) exitNodes(node api.INode) {
	fmt.Println("NodeManager: exit-node ", node)
	node.ExitNode(m)

	children := node.Children()
	for _, child := range children {
		m.exitNodes(child)
	}
}

func (m *nodeManager) Debug() {
	black := rendering.NewPaletteInt64(rendering.Black)
	m.context.SetDrawColor(black)
	// renderer.SetDrawColor(0, 0, 0, 255)

	m.context.DrawLine(500, 500, 800, 850)
}

func (m nodeManager) String() string {
	return fmt.Sprintf("%s", m.stack)
}
