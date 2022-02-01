package ecs

import (
	"errors"
	"fmt"
	"log"

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
		err := ec.checkComponentsAndAddEntity(e, s)
		if err != nil {
			return err
		}
	}
	ec.entities[e.Id] = e
	return nil
}

func (ec *Ecs) checkComponentsAndAddEntity(e *Entity, s System) error {
	requiredComponentKeys := s.RequiredComponentKeys()
	if requiredComponentKeys == nil {
		return nil
	}

	var components []interface{}
	for _, key := range requiredComponentKeys {
		c, ok := e.Components[key]
		if !ok {
			return nil
		}
		components = append(components, c)
	}
	err := s.AddEntity(e, components)
	if err != nil {
		return fmt.Errorf("%s system: %w", s, err)
	}

	log.Printf("%s added to %s system", e, s)
	return nil
}

func (ec *Ecs) DoTick(currentTick tick.Tick) (error, bool) {
	var err error
	var stop bool
	for _, s := range ec.systems {
		err, stop = s.DoTick(currentTick)
		if err != nil {
			return err, false
		}
		if stop {
			break
		}
	}
	return nil, stop
}
