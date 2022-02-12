package entityrepo

import (
	"log"

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

func (c *Chunked) Add(e entity.Entity) {
	cgroup, ok := c.cgroups[e.CMask]
	if !ok {
		log.Fatalf("CGroup for mask %s doesn't exist", e.CMask)
	}
	chunkIndex, index := cgroup.Add(e)
	c.entityIndex[e.Id] = EAddress{
		Mask:       e.CMask,
		ChunkIndex: chunkIndex,
		Index:      index,
	}
}

type CGroup interface {
	Add(e entity.Entity) (int, int)
	Get(addr EAddress) (entity.Entity, bool)
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

func (c *Chunked) Get(id entity.Id) (entity.Entity, bool) {
	addr, ok := c.entityIndex[id]
	if !ok {
		return entity.Entity{}, false
	}

	cgroup, ok := c.cgroups[addr.Mask]
	if !ok {
		log.Fatalf("CGroup for mask %s doesn't exist", addr.Mask)
	}

	return cgroup.Get(addr)
}
