package ecs

import (
	"errors"

	"github.com/justclimber/fda/common/tick"
)

var ErrNewEcsShouldBeWithAtLeastOneSystem = errors.New("should be at least one system")

type Ecs struct {
	systems  []System
	entities map[EntityId]*Entity
}

func NewEcs(systems []System) (*Ecs, error) {
	if len(systems) == 0 {
		return nil, ErrNewEcsShouldBeWithAtLeastOneSystem
	}
	return &Ecs{
		systems:  systems,
		entities: make(map[EntityId]*Entity),
	}, nil
}

func (ec *Ecs) AddEntity(e *Entity) error {
	for _, s := range ec.systems {
		err := ec.addEntityComponents(e, s)
		if err != nil {
			return err
		}
	}
	ec.entities[e.Id] = e
	return nil
}

func (ec *Ecs) addEntityComponents(e *Entity, s System) error {
	var components []interface{}
	for _, key := range s.RequiredComponentKeys() {
		c, ok := e.Components[key]
		if !ok {
			return nil
		}
		components = append(components, c)
	}
	return s.AddEntity(e, components)
}

func (ec *Ecs) DoTick(currentTick tick.Tick) error {
	for _, s := range ec.systems {
		if err := s.DoTick(currentTick); err != nil {
			return err
		}
	}
	return nil
}
