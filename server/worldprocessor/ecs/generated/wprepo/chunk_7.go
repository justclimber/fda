// Code generated by entityrepo generator. DO NOT EDIT.
package wprepo

import (
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

const chunkSizeMask7 = 10

type ChunkMask7 struct {
	Size     int
	Ids      [chunkSizeMask7]entity.Id
	Position   [chunkSizeMask7]wpcomponent.Position
	Moving   [chunkSizeMask7]wpcomponent.Moving
	Player   [chunkSizeMask7]wpcomponent.Player
}

func (ch *ChunkMask7) Add(e EntityMask7) {
	ch.Ids[ch.Size] = e.EId()
	ch.Position[ch.Size] = e.Position
	ch.Moving[ch.Size] = e.Moving
	ch.Player[ch.Size] = e.Player
	ch.Size++
}