package geometry

import (
	"fmt"

	"github.com/wdevore/RangerGo/api"
)

type line struct {
	p1 api.IPoint
	p2 api.IPoint
}

// NewLine constructs a new ILine
func NewLine() api.ILine {
	o := new(line)
	o.p1 = NewPoint()
	o.p2 = NewPoint()
	return o
}

// NewLineUsing constructs a new point using components
func NewLineUsing(x1, y1, x2, y2 float64) api.ILine {
	o := new(line)
	o.p1 = NewPointUsing(x1, y1)
	o.p2 = NewPointUsing(y1, y2)
	return o
}

func (l *line) Components() (p1, p2 api.IPoint) {
	return l.p1, l.p2
}

func (l *line) SetP1(x, y float64) {
	l.p1.SetByComp(x, y)
}

func (l *line) SetP2(x, y float64) {
	l.p2.SetByComp(x, y)
}

func (l *line) SetByComp(x1, y1, x2, y2 float64) {
	l.p1.SetByComp(x1, y1)
	l.p2.SetByComp(x2, y2)
}

func (l *line) SetByLine(il api.ILine) {
	p1, p2 := il.Components()
	l.p1.SetByComp(p1.X(), p1.Y())
	l.p2.SetByComp(p2.X(), p2.Y())
}

func (l line) String() string {
	return fmt.Sprintf("%v -> %v", l.p1, l.p2)
}
