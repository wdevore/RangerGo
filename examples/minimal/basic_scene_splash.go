package main

import (
	"github.com/wdevore/RangerGo/api"
	"github.com/wdevore/RangerGo/engine/geometry"
	"github.com/wdevore/RangerGo/engine/nodes"
	"github.com/wdevore/RangerGo/engine/rendering"
)

type sceneSplash struct {
	nodes.Node
	nodes.Scene
	nodes.Transition

	api.IRender

	// mesh api.IMesh
	mesh *geometry.Mesh

	oddColor  api.IPalette
	evenColor api.IPalette
}

func newBasicSplashScene(name string, replacement api.INode) api.INode {
	o := new(sceneSplash)
	o.Initialize(name)
	o.SetReplacement(replacement)

	o.mesh = geometry.NewMesh()
	return o
}

func (s *sceneSplash) Build(world api.IWorld) {
	vw, vh := world.ViewSize().Components()
	y := -vh / 2.0
	w := 320.0

	// Construct grid of rectangles
	for y <= vh {
		x := -vw / 2.0
		for x <= vw {
			s.mesh.AddVertex(x, y) // top-left
			s.mesh.AddVertex(x+w, y+w)
			x += w
		}
		y += w
	}

	s.mesh.Build()

	s.SetPauseTime(30000.0)

	s.oddColor = rendering.NewPaletteInt64(rendering.DarkGray)
	s.evenColor = rendering.NewPaletteInt64(rendering.LightGray)
}

// --------------------------------------------------------
// Timing
// --------------------------------------------------------

func (s *sceneSplash) Update(dt float64) {
	s.Transition.UpdateTransition(dt)
}

// --------------------------------------------------------
// Transitioning
// --------------------------------------------------------

func (s *sceneSplash) TransitionAction() int {
	if s.Transition.ReadyToTransition() {
		return api.SceneReplaceTake
	}

	return api.SceneNoAction
}

// -----------------------------------------------------
// Scene lifecycles
// -----------------------------------------------------

// EnterNode called when a node is entering the stage
func (s *sceneSplash) EnterNode(man api.INodeManager) {
	// fmt.Println("sceneSplash enter")
	s.Transition.Reset()
	man.RegisterTarget(s)
}

// ExitNode called when a node is exiting stage
func (s *sceneSplash) ExitNode(man api.INodeManager) {
	man.UnRegisterTarget(s)
}

// -----------------------------------------------------
// Visuals
// -----------------------------------------------------

func (s *sceneSplash) Draw(context api.IRenderContext) {
	s.Node.Draw(context)

	// Transform vertices if anything has changed.
	if s.IsDirty() {
		// Transform this node's vertices using the context
		context.TransformMesh(s.mesh)
		s.SetDirty(false) // Node's bucket is no longer dirty
	}

	// Draw background first. The background is a grid of squares.
	context.RenderCheckerBoard(s.mesh, s.oddColor, s.evenColor)

	// Manually building and rendering a line. This isn't the idiomatic Ranger way!
	// This is just a simple demo scene.
	black := rendering.NewPaletteInt64(rendering.Black)
	context.SetDrawColor(black)

	p1 := geometry.NewPoint()
	p2 := geometry.NewPointUsing(100.0, 100.0)
	o1 := geometry.NewPoint()
	o2 := geometry.NewPoint()
	context.TransformPoints(p1, p2, o1, o2)
	context.RenderLine(o1.X(), o1.Y(), o2.X(), o2.Y())
}
