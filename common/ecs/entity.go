package ecs

import (
	"fmt"
)

type EntityId int64

type Entity struct {
	Id         EntityId
	Components map[ComponentKey]Component
}

func NewEntity(Id EntityId) *Entity {
	return &Entity{
		Id:         Id,
		Components: map[ComponentKey]Component{},
	}
}

func (e *Entity) AddComponent(c Component) {
	e.Components[c.Key()] = c
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity[ID: %d]", e.Id)
}
