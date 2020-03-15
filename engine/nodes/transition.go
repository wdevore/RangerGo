package nodes

import "fmt"

// Transition holds properties for scene transitions.
type Transition struct {
	pauseTime     float64
	pauseCnt      float64
	canTransition bool // true = ready to transition
}

// Reset resets the internal timing properties and readies for another cycle.
func (t *Transition) Reset() {
	t.pauseCnt = 0.0
	t.canTransition = false
}

// SetPauseTime sets how long the pause lasts, in milliseconds
func (t *Transition) SetPauseTime(milliseconds float64) {
	t.pauseTime = milliseconds
}

// Inc used rarely to manually increment the internal counter
func (t *Transition) Inc(dt float64) {
	t.pauseCnt += dt
}

// UpdateTransition the internal timer
func (t *Transition) UpdateTransition(dt float64) {
	fmt.Println(t.pauseTime)
	t.pauseCnt += dt
	if t.pauseCnt >= t.pauseTime {
		t.canTransition = true
	}
}

// ReadyToTransition indicates if the node can transition to another scene
func (t *Transition) ReadyToTransition() bool {
	return t.canTransition
}
