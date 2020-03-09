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
}

// NewNodeManager constructs a manager for node.
// It manages the lifecycle and events
func NewNodeManager(world api.IWorld) api.INodeManager {
	o := new(nodeManager)
	o.clearBackground = true

	o.context = rendering.NewRenderContext(world)
	o.context.Initialize()

	o.stack = newNodeStack()

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

	if m.stack.hasNextNode() {
		m.setNextNode()
	}

	// This will save view-space matrix
	m.context.Save()

	// DEBUG
	// If mouse coords changed then update view coords.
	// self.global_data.update_view_coords(&mut self.context);

	action := m.stack.runningNode.Transition()

	if action == api.SceneReplaceTake {
		repl := m.stack.runningNode.GetReplacement()
		m.stack.replace(repl)
	}

	// Immediately switch to the new runner
	if m.stack.hasNextNode() {
		m.setNextNode()
	}

	// Visit the running node
	m.stack.runningNode.Visit(m.context, interpolation)

	m.context.Restore()

	return true // continue to draw.
}

func (m *nodeManager) PostVisit() {

}

func (m *nodeManager) Update(dt float64) {

}

func (m *nodeManager) PushNode(node api.INode) {

}

func (m *nodeManager) PopNode() {

}

func (m *nodeManager) RegisterTarget(target api.INode) {

}

func (m *nodeManager) UnRegisterTarget(target api.INode) {

}

func (m *nodeManager) RegisterEventTarget(target api.INode) {

}

func (m *nodeManager) UnRegisterEventTarget(target api.INode) {

}

func (m *nodeManager) setNextNode() {

}

func (m *nodeManager) enterNodes() {

}

func (m *nodeManager) exitNodes() {

}
