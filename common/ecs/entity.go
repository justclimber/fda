package ecs

import (
	"fmt"
)

type ComponentKey string
type EntityId int64

type Entity struct {
	Id         EntityId
	Components map[ComponentKey]interface{}
}

func (e *Entity) AddComponent(key ComponentKey, component interface{}) {
	e.Components[key] = component
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity[ID: %d]", e.Id)
}
