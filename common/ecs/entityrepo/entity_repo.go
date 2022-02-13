package entityrepo

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
)

type EAddress struct {
	Mask       component.Mask
	ChunkIndex int
	Index      int
}

type Chunked struct {
	ecgroups    map[component.Mask]ECGroup
	entityIndex map[entity.Id]EAddress
}

func NewChunked(ecgroups map[component.Mask]ECGroup) *Chunked {
	return &Chunked{
		ecgroups:    ecgroups,
		entityIndex: map[entity.Id]EAddress{},
	}
}

func (c *Chunked) Add(e entity.MaskedEntity) {
	ecgroup := c.ecgroups[e.Mask()]

	chunkIndex, index := ecgroup.Add(e)
	c.entityIndex[e.EId()] = EAddress{
		Mask:       e.Mask(),
		ChunkIndex: chunkIndex,
		Index:      index,
	}
}

type ECGroup interface {
	Add(e entity.MaskedEntity) (int, int)
	Get(addr EAddress) entity.MaskedEntity
}

type Chunk interface {
	Size() int
}

func (c *Chunked) GetECGroupsWithMask(mask component.Mask) []ECGroup {
	var cgs []ECGroup
	for m, cgroup := range c.ecgroups {
		if m.Contains(mask) {
			cgs = append(cgs, cgroup)
		}
	}
	return cgs
}

func (c *Chunked) Get(id entity.Id) (entity.MaskedEntity, bool) {
	addr, ok := c.entityIndex[id]
	if !ok {
		return nil, false
	}

	return c.ecgroups[addr.Mask].Get(addr), true
}
