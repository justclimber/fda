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

func (m *sysMock) DoTick(_ tick.Tick) (error, bool) {
	for _, cc := range m.components {
		cc.c1.num1 = cc.c1.num1 + 20
		cc.c2.str = "changed"
	}
	return nil, false
}

func (m *sysMock) RequiredComponentKeys() []ComponentKey {
	return []ComponentKey{c2key, c1key}
}

type objectiveMock struct {
	curC1           *c1
	objectiveC1Num1 int
}

func NewObjectiveMock(o int) *objectiveMock {
	return &objectiveMock{objectiveC1Num1: o}
}

func (o *objectiveMock) AddEntity(e *Entity, in []interface{}) error {
	if e.Id != 10 {
		return nil
	}

	o.curC1, _ = in[0].(*c1)
	return nil
}

func (o *objectiveMock) RemoveEntity(_ *Entity) {}

func (o *objectiveMock) DoTick(_ tick.Tick) (error, bool) {
	if o.curC1 == nil {
		return errors.New("oops"), false
	}
	return nil, o.curC1.num1 == o.objectiveC1Num1
}

func (o *objectiveMock) RequiredComponentKeys() []ComponentKey {
	return []ComponentKey{c1key}
}
