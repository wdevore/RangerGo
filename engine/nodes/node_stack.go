package nodes

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
)

// A stack of nodes
type nodeStack struct {
	nodes []api.INode

	nextNode    api.INode
	runningNode api.INode
}

func newNodeStack() *nodeStack {
	o := new(nodeStack)
	return o
}

func (n *nodeStack) isEmpty() bool {
	return len(n.nodes) == 0
}

func (n *nodeStack) hasNextNode() bool {
	return n.nextNode != nil
}

func (n *nodeStack) hasRunningNode() bool {
	return n.runningNode != nil
}

func (n *nodeStack) clearNextNode() {
	n.nextNode = nil
}

func (n *nodeStack) clearRunningNode() {
	n.runningNode = nil
}

func (n *nodeStack) push(node api.INode) {
	n.nextNode = node
	fmt.Println("NodeStack: pushing ", n.nextNode)
	n.nodes = append(n.nodes, node)
}

func (n *nodeStack) pop() {
	if !n.isEmpty() {
		topI := len(n.nodes) - 1 // Top element index
		n.nextNode = n.nodes[topI]
		n.nodes = n.nodes[:topI] // Pop
		fmt.Println("NodeStack: popped ", n.nextNode)
	} else {
		fmt.Println("NodeStack -- no nodes to pop")
	}
}

func (n *nodeStack) replace(replacement api.INode) {
	n.nextNode = replacement

	// Replacement is the act of popping and pushing. i.e. replacing
	// the stack top with the new node.
	if !n.isEmpty() {
		topI := len(n.nodes) - 1
		top := n.nodes[topI]     // Top element
		n.nodes = n.nodes[:topI] // Pop
		fmt.Println("NodeStack: popped ", top, " for replacement with ", replacement)
	} else {
		fmt.Println("NodeStack: replace WARNING, nothing popped")
	}

	n.nodes = append(n.nodes, replacement)
}
