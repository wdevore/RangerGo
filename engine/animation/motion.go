package animation

// Motion are properties related to interpolation
type Motion struct {
	rate      float64
	timeScale float64

	// Enabled by default
	autoWrap bool
}

// InitializeMotion sets default values
func (m *Motion) InitializeMotion() {
	m.rate = 0.0
	m.timeScale = 1000.0
	m.autoWrap = true
}

// SetRate sets rate
func (m *Motion) SetRate(rate float64) {
	m.rate = rate
}

// SetAutoWrap enables/disable wrapping
func (m *Motion) SetAutoWrap(wrap bool) {
	m.autoWrap = wrap
}

// SetTimeScale set time scaling
func (m *Motion) SetTimeScale(s float64) {
	m.timeScale = s
}
