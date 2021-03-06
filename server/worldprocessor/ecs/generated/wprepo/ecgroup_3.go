// Code generated by entityrepo generator. DO NOT EDIT.
package wprepo

import (
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
)

type ECGroupMask3 struct {
	Chunks []*ChunkMask3
	last   *ChunkMask3
}

func NewECGroupMask3() *ECGroupMask3 {
	chunks := []*ChunkMask3{{}}
	return &ECGroupMask3{
		Chunks: chunks,
		last:   chunks[0],
	}
}

func (eg *ECGroupMask3) Add(e entity.MaskedEntity) (int, int) {
    em := e.(EntityMask3)
	if eg.last.Size == chunkSizeMask3 {
		eg.last = &ChunkMask3{}
		eg.Chunks = append(eg.Chunks, eg.last)
	}
	eg.last.Add(em)
	return len(eg.Chunks) - 1, eg.last.Size - 1
}

func (eg *ECGroupMask3) Get(addr entityrepo.EAddress) entity.MaskedEntity {
	chunk := eg.Chunks[addr.ChunkIndex]

	return EntityMask3{
		Id: chunk.Ids[addr.Index],
        Position: chunk.Position[addr.Index],
        Moving: chunk.Moving[addr.Index],
	}
}