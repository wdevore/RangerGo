package api

// IEngine is the main engine API
type IEngine interface {
	Configure()
	Start()

	// DisplaySize returns the application's window dimensions.
	DisplaySize() (int, int)
}
