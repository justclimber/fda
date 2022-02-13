package ecs

import (
	"fmt"

	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
	"github.com/justclimber/fda/common/tick"
)

type nestedDebugger interface {
	LogF(method string, str string, args ...interface{})
}

type emptyDebugger struct{}

func (e *emptyDebugger) LogF(_ string, _ string, _ ...interface{}) {}

type EntityRepo interface {
	Add(e entity.MaskedEntity)
	Get(id entity.Id) (entity.MaskedEntity, bool)
	GetECGroupsWithMask(mask component.Mask) []entityrepo.ECGroup
}

type System interface {
	fmt.Stringer
	DoTick(tick tick.Tick) bool
	Init(tick tick.Tick)
}
