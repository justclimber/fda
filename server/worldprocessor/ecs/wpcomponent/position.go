package wpcomponent

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/fgeom"
)

const CPosition ecs.ComponentKey = "pos"

type Position struct {
	Pos *fgeom.Point
}

func NewPosition(p *fgeom.Point) *Position {
	return &Position{Pos: p}
}

func (p *Position) Key() ecs.ComponentKey {
	return CPosition
}
