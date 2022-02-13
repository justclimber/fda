package entity

import (
	"github.com/justclimber/fda/common/ecs/component"
)

type Id int64

type MaskedEntity interface {
	EId() Id
	Mask() component.Mask
}
