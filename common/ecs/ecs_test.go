package ecs

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

func TestEcs(t *testing.T) {
	entityId1 := entity.Id(12)
	entityId2 := entity.Id(89)
	entityId3 := entity.Id(5)
	curTick := tick.Tick(55)

	e1 := entity.NewEntity(entityId1)
	d1 := fgeom.Point{X: 1}
	pos1 := fgeom.Point{X: 2}
	e1.AddComponent(&wpcomponent.Moving{D: d1})
	e1.AddComponent(&wpcomponent.Position{Pos: pos1})
	mask3 := e1.CMask

	e2 := entity.NewEntity(entityId2)
	d2 := fgeom.Point{X: -3}
	pos2 := fgeom.Point{X: 10}
	e2.AddComponent(&wpcomponent.Moving{D: d2})
	e2.AddComponent(&wpcomponent.Position{Pos: pos2})
	e2.AddComponent(&wpcomponent.Player{Delay: 10})

	e3 := entity.NewEntity(entityId3)
	d3 := fgeom.Point{X: -30}
	pos3 := fgeom.Point{X: 100}
	e3.AddComponent(&wpcomponent.Moving{D: d3})
	e3.AddComponent(&wpcomponent.Position{Pos: pos3})
	e3.AddComponent(&wpcomponent.Player{Delay: 10})

	mask7 := e2.CMask

	cg3 := NewCGroup3()
	cg7 := NewCGroup7()
	repo := entityrepo.NewChunked(map[component.Mask]entityrepo.CGroup{
		mask3: cg3,
		mask7: cg7,
	})

	repoForMask3 := NewRepoForMask3(repo)
	movSys := NewMoving2(repoForMask3)

	ec, err := NewEcs([]System{movSys}, repo, &emptyDebugger{})
	require.NoError(t, err)

	ec.AddEntity(e1)
	ec.AddEntity(e2)
	ec.AddEntity(e3)

	stop := ec.DoTick(curTick)
	require.False(t, stop)
	require.Equal(t, pos1.Add(d1), cg3.Chunks[0].Position[0].Pos, "check pos for entity 1")
	require.Equal(t, pos2.Add(d2), cg7.Chunks[0].Position[0].Pos, "check pos for entity 2")
	require.Equal(t, pos3.Add(d3), cg7.Chunks[0].Position[1].Pos, "check pos for entity 3")

	e1n, ok := repo.Get(entityId1)
	require.True(t, ok)
	require.Equal(t, pos1.Add(d1), e1n.Components[wpcomponent.KeyPosition].(*wpcomponent.Position).Pos)

	e2n, ok := repo.Get(entityId2)
	require.True(t, ok)
	require.Equal(t, pos2.Add(d2), e2n.Components[wpcomponent.KeyPosition].(*wpcomponent.Position).Pos)
}
