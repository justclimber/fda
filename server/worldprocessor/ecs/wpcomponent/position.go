package wpcomponent

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/fgeom"
)

type Position struct {
	Pos fgeom.Point
}

func NewPosition(p fgeom.Point) *Position {
	return &Position{Pos: p}
}

func (p *Position) Key() component.Key {
	return KeyPosition
}
