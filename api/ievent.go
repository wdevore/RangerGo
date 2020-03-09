package api

// IEvent represents IO event system
type IEvent interface {
	Handle() bool
}
