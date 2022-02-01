package ecs

import (
	"fmt"

	"github.com/justclimber/fda/common/tick"
)

type System interface {
	fmt.Stringer
	AddEntity(e *Entity, in []interface{}) error
	RemoveEntity(e *Entity)
	DoTick(tick tick.Tick) (error, bool)
	RequiredComponentKeys() []ComponentKey
}
