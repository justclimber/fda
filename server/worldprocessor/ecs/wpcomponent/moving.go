package wpcomponent

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/fgeom"
)

type Moving struct {
	D fgeom.Point
}

func NewMoving(d fgeom.Point) Moving {
	return Moving{D: d}
}

func (m Moving) Key() component.Key {
	return KeyMoving
}
