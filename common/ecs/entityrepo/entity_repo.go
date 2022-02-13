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
	cgroups     map[component.Mask]CGroup
	entityIndex map[entity.Id]EAddress
}

func NewChunked(cgroups map[component.Mask]CGroup) *Chunked {
	return &Chunked{
		cgroups:     cgroups,
		entityIndex: map[entity.Id]EAddress{},
	}
}

func (c *Chunked) Add(e entity.MaskedEntity) {
	cgroup := c.cgroups[e.Mask()]

	chunkIndex, index := cgroup.Add(e)
	c.entityIndex[e.EId()] = EAddress{
		Mask:       e.Mask(),
		ChunkIndex: chunkIndex,
		Index:      index,
	}
}

type CGroup interface {
	Add(e entity.MaskedEntity) (int, int)
	Get(addr EAddress) entity.MaskedEntity
}

type Chunk interface {
	Size() int
}

func (c *Chunked) GetCGroupsWithMask(mask component.Mask) []CGroup {
	var cgs []CGroup
	for m, cgroup := range c.cgroups {
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

	return c.cgroups[addr.Mask].Get(addr), true
}
