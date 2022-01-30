package ecs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/tick"
)

func TestEcs_NewShouldBeWithAtLeastOneSystem(t *testing.T) {
	_, err := NewEcs([]System{})
	assert.Equal(t, ErrNewEcsShouldBeWithAtLeastOneSystem, err)
}

func TestEcs_AddEntityWithMock(t *testing.T) {
	sysMock := &sysMock{components: make(map[EntityId]components)}
	ec, err := NewEcs([]System{sysMock})
	require.NoError(t, err)

	e := &Entity{
		Id: 10,
		Components: map[ComponentKey]interface{}{
			c1key: &c1{num1: 54},
			c2key: &c2{str: "foo"},
			c3key: &c3{num2: 5.4},
		},
	}
	err = ec.AddEntity(e)
	require.NoError(t, err, "fail to add entity to ecs")

	c, ok := sysMock.components[e.Id]
	require.True(t, ok, "entity must be in system")
	require.Equal(t, 54, c.c1.num1, "check component data")
	require.Equal(t, "foo", c.c2.str, "check component data")
}

func TestEcs_DoTickWithMock(t *testing.T) {
	sysMock := &sysMock{components: make(map[EntityId]components)}
	ec, err := NewEcs([]System{sysMock})
	require.NoError(t, err)

	c1c := &c1{num1: 54}
	c2c := &c2{str: "foo"}
	c3c := &c3{num2: 5.4}

	e := &Entity{
		Id: 10,
		Components: map[ComponentKey]interface{}{
			c1key: c1c,
			c2key: c2c,
			c3key: c3c,
		},
	}
	err = ec.AddEntity(e)
	require.NoError(t, err, "fail to add entity to ecs")

	err = ec.DoTick(tick.Tick(10))
	require.NoError(t, err, "ecs do tick error")

	assert.Equal(t, 20, c1c.num1, "expect component data to be changed during ecs->system tick")
	assert.Equal(t, "changed", c2c.str, "expect component data to be changed during ecs->system tick")
}
