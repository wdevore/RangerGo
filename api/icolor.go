package api

// IColor represents Nodes that allow color manipulation.
type IColor interface {
	// SetRed
	SetRed(r int)
	// SetGreen
	SetGreen(r int)
	// SetBlue
	SetBlue(r int)
	// SetAlpha
	SetAlpha(r int)
}
