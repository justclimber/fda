package ecs

import (
	"errors"

	"github.com/justclimber/fda/common/tick"
)

const (
	c1key ComponentKey = "c1"
	c2key ComponentKey = "c2"
	c3key ComponentKey = "c3"
)

type c1 struct{ num1 int }
type c2 struct{ str string }
type c3 struct{ num2 float64 }

type components struct {
	c1 *c1
	c2 *c2
}

type sysMock struct {
	components map[EntityId]components
}

func (m *sysMock) AddEntity(e *Entity, in []interface{}) error {
	if len(in) != 2 {
		return errors.New("incorrect components count on input")
	}
	c2, ok := in[0].(*c2)
	if !ok {
		return errors.New("incorrect components on input")
	}
	c1, ok := in[1].(*c1)
	if !ok {
		return errors.New("incorrect components on input")
	}
	m.components[e.Id] = components{
		c1: c1,
		c2: c2,
	}
	return nil
}

func (m *sysMock) RemoveEntity(e *Entity) {
}

func (m *sysMock) DoTick(_ tick.Tick) error {
	for _, cc := range m.components {
		cc.c1.num1 = 20
		cc.c2.str = "changed"
	}
	return nil
}

func (m *sysMock) RequiredComponentKeys() []ComponentKey {
	return []ComponentKey{c2key, c1key}
}
