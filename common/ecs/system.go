package ecs

import (
	"github.com/justclimber/fda/common/tick"
)

type System interface {
	AddEntity(e *Entity, in []interface{}) error
	RemoveEntity(e *Entity)
	DoTick(tick tick.Tick) (error, bool)
	RequiredComponentKeys() []ComponentKey
}
