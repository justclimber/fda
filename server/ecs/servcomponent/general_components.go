package servcomponent

import (
	"github.com/justclimber/fda/common/fgeom"
)

type Movable interface {
	Move(p *fgeom.Point)
}

type Engine struct {
	power float64
}

func NewEngine(power float64) *Engine {
	return &Engine{power: power}
}

func (e *Engine) Move(p *fgeom.Point) {
	p.X = p.X + e.power
}

type Position struct {
	Pos *fgeom.Point
}
