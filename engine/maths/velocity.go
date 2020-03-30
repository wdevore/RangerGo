package maths

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
)

type velocity struct {
	magnitude float64

	minMag, maxMag float64

	direction api.IVector

	limitMag bool
}

// NewVelocity constructs a new IVelocity
func NewVelocity() api.IVelocity {
	o := new(velocity)
	o.direction = NewVectorUsing(1.0, 0.0)
	o.limitMag = true
	return o
}

// NewVelocityUsing constructs a new IVelocity using another velocity
func NewVelocityUsing(vel api.IVelocity) api.IVelocity {
	o := NewVelocity()
	o.SetMagnitude(vel.Magnitude())
	o.SetMinMax(vel.Range())
	o.SetDirectionByVector(vel.Direction())
	return o
}

func (v *velocity) SetMin(min float64) {
	v.minMag = min
}

func (v *velocity) SetMax(max float64) {
	v.maxMag = max
}

func (v *velocity) SetMinMax(min, max float64) {
	v.minMag = min
	v.maxMag = max
}

func (v *velocity) SetMagnitude(mag float64) {
	v.magnitude = mag
}

func (v *velocity) Magnitude() float64 {
	return v.magnitude
}

func (v *velocity) Range() (float64, float64) {
	return v.minMag, v.maxMag
}

func (v *velocity) SetDirectionByAngle(radians float64) {
	v.direction.SetByAngle(radians)
}

func (v *velocity) SetDirectionByVector(dir api.IVector) {
	v.direction.SetByVector(dir)
}

func (v *velocity) Direction() api.IVector {
	return v.direction
}

func (v *velocity) ConstrainMagnitude(con bool) {
	v.limitMag = con
}

func (v *velocity) ApplyToPoint(point api.IPoint) {
	// Get actual velocity
	v1.SetByComp(v.direction.X()*v.magnitude, v.direction.Y()*v.magnitude)
	v2.SetByPoint(point)
	Add(v1, v2, v3)
	point.SetByComp(v3.X(), v3.Y())
}

func (v velocity) String() string {
	return fmt.Sprintf("|%3.3f| %v", v.magnitude, v.direction)
}
