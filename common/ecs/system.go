package ecs

import (
	"fmt"

	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/tick"
)

type System interface {
	fmt.Stringer
	Init()
	RequiredComponentKeys() []component.Key
	AddEntity(e *entity.Entity, in []interface{}) error
	RemoveEntity(e *entity.Entity)
	DoTick(tick tick.Tick) (error, bool)
}
