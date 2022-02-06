package wpcomponent

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/fgeom"
)

const CMovable component.Key = "mov"

type Movable interface {
	Move(p *fgeom.Point)
}

type PowerSettable interface {
	SetPower(p float64)
}

type Engine struct {
	power float64
}

func (e *Engine) Key() component.Key {
	return CMovable
}

func NewEngine(power float64) *Engine {
	return &Engine{power: power}
}

func (e *Engine) Move(p *fgeom.Point) {
	p.X = p.X + e.power
}

func (e *Engine) SetPower(p float64) {
	e.power = p
}
