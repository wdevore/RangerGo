package api

import "image/color"

// IEngine is the main engine API
type IEngine interface {
	// Configure constructs a display using SDL and display it.
	Configure( /*TODO Game object*/ )

	// Start launches the game loop
	Start()

	// Ends shuts down the engine
	End()

	// DisplaySize returns the application's window dimensions.
	DisplaySize() (int, int)

	// SetClearColor sets the background clear color
	SetClearColor(color color.RGBA)

	// PushStart pushes the given node onto the stack as the
	// first scene to start once the engine's configuration in complete.
	PushStart(INode)
}
