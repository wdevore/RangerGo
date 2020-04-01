package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type gameLayer struct {
	nodes.Node

	mesh api.IMesh

	textColor api.IPalette

	oddColor  api.IPalette
	evenColor api.IPalette
}

func newBasicGameLayer(name string) api.INode {
	o := new(gameLayer)
	o.Initialize(name)
	return o
}

func (g *gameLayer) Build(world api.IWorld) {
	vw, vh := world.ViewSize().Components()
	y := -vh / 2.0
	w := 320.0

	g.mesh = geometry.NewMesh()

	// Construct grid of rectangles
	for y <= vh {
		x := -vw / 2.0
		for x <= vw {
			g.mesh.AddVertex(x, y) // top-left
			g.mesh.AddVertex(x+w, y+w)
			x += w
		}
		y += w
	}

	g.mesh.Build()

	g.textColor = rendering.NewPaletteInt64(rendering.LightNavyBlue)
	g.oddColor = rendering.NewPaletteInt64(rendering.DarkGray)
	g.evenColor = rendering.NewPaletteInt64(rendering.LightGray)
}

// --------------------------------------------------------
// Timing
// --------------------------------------------------------

func (g *gameLayer) Update(dt float64) {
}

// -----------------------------------------------------
// Node lifecycles
// -----------------------------------------------------

// EnterNode called when a node is entering the stage
func (g *gameLayer) EnterNode(man api.INodeManager) {
	man.RegisterTarget(g)
}

// ExitNode called when a node is exiting stage
func (g *gameLayer) ExitNode(man api.INodeManager) {
	man.UnRegisterTarget(g)
}

// -----------------------------------------------------
// Visuals
// -----------------------------------------------------

func (g *gameLayer) Draw(context api.IRenderContext) {
	// Transform vertices if anything has changed.
	if g.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformMesh(g.mesh)
		g.SetDirty(false) // Node is no longer dirty
	}

	// Draw background first. The background is a grid of squares.
	context.RenderCheckerBoard(g.mesh, g.oddColor, g.evenColor)

}
