package api

// ITransition scene timing and transitions
type ITransition interface {
	Reset()

	SetPauseTime(milliseconds float64)
	Inc(dt float64)
	UpdateTransition(dt float64)

	ReadyToTransition() bool
}
