package entity

import (
	"fmt"

	"github.com/justclimber/fda/common/ecs/component"
)

type Id int64

type Entity struct {
	Id         Id
	Components map[component.Key]component.Component
	CMask      component.Mask
}

func NewEntity(Id Id) Entity {
	return Entity{
		Id:         Id,
		Components: map[component.Key]component.Component{},
	}
}

func (e *Entity) AddComponent(c component.Component) {
	e.Components[c.Key()] = c
	e.CMask.Add(c.Key())
}

func (e Entity) String() string {
	return fmt.Sprintf("Entity[ID: %d]", e.Id)
}