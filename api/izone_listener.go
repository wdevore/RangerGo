package api

// IZoneListener is for objects wanting to be notified of Zone events
type IZoneListener interface {
	Notify(state, id int)
}
