package wprepo

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type CGroup3 struct {
	Chunks []*Chunk3
	last   *Chunk3
}

func NewCGroup3() *CGroup3 {
	chunks := []*Chunk3{{}}
	return &CGroup3{
		Chunks: chunks,
		last:   chunks[0],
	}
}

func (cg *CGroup3) Add(e entity.Entity) (int, int) {
	if cg.last.Size == chunkSize {
		cg.last = &Chunk3{}
		cg.Chunks = append(cg.Chunks, cg.last)
	}
	cg.last.Add(e)
	return len(cg.Chunks) - 1, cg.last.Size - 1
}

func (cg *CGroup3) Get(addr entityrepo.EAddress) (entity.Entity, bool) {
	if addr.ChunkIndex >= len(cg.Chunks) {
		return entity.Entity{}, false
	}

	chunk := cg.Chunks[addr.ChunkIndex]

	if addr.Index >= chunk.Size {
		return entity.Entity{}, false
	}

	return entity.Entity{
		Id: chunk.Ids[addr.Index],
		Components: map[component.Key]component.Component{
			wpcomponent.KeyPosition: &chunk.Position[addr.Index],
			wpcomponent.KeyMoving:   &chunk.Moving[addr.Index],
		},
		CMask: component.NewMask([]component.Key{wpcomponent.KeyPosition, wpcomponent.KeyMoving}),
	}, true
}

type CGroup7 struct {
	Chunks []*Chunk7
	last   *Chunk7
}

func NewCGroup7() *CGroup7 {
	chunks := []*Chunk7{{}}
	return &CGroup7{
		Chunks: chunks,
		last:   chunks[0],
	}
}

func (cg *CGroup7) Add(e entity.Entity) (int, int) {
	if cg.last.Size == chunkSize {
		cg.last = &Chunk7{}
		cg.Chunks = append(cg.Chunks, cg.last)
	}
	cg.last.Add(e)
	return len(cg.Chunks) - 1, cg.last.Size - 1
}

func (cg *CGroup7) Get(addr entityrepo.EAddress) (entity.Entity, bool) {
	if addr.ChunkIndex >= len(cg.Chunks) {
		return entity.Entity{}, false
	}

	chunk := cg.Chunks[addr.ChunkIndex]

	if addr.Index >= chunk.Size {
		return entity.Entity{}, false
	}

	return entity.Entity{
		Id: chunk.Ids[addr.Index],
		Components: map[component.Key]component.Component{
			wpcomponent.KeyPosition: &chunk.Position[addr.Index],
			wpcomponent.KeyMoving:   &chunk.Moving[addr.Index],
		},
		CMask: component.NewMask([]component.Key{wpcomponent.KeyPosition, wpcomponent.KeyMoving}),
	}, true
}
