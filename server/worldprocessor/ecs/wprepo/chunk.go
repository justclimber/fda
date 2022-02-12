package wprepo

import (
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

const chunkSize = 10

type Chunk3 struct {
	Size     int
	Ids      [chunkSize]entity.Id
	Moving   [chunkSize]wpcomponent.Moving
	Position [chunkSize]wpcomponent.Position
}

func (ch *Chunk3) Add(e entity.Entity) {
	ch.Ids[ch.Size] = e.Id
	ch.Moving[ch.Size] = *e.Components[wpcomponent.KeyMoving].(*wpcomponent.Moving)
	ch.Position[ch.Size] = *e.Components[wpcomponent.KeyPosition].(*wpcomponent.Position)
	ch.Size++
}

type Chunk7 struct {
	Size     int
	Ids      [chunkSize]entity.Id
	Moving   [chunkSize]wpcomponent.Moving
	Position [chunkSize]wpcomponent.Position
	Player   [chunkSize]wpcomponent.Player
}

func (ch *Chunk7) Add(e entity.Entity) {
	ch.Ids[ch.Size] = e.Id
	ch.Moving[ch.Size] = *e.Components[wpcomponent.KeyMoving].(*wpcomponent.Moving)
	ch.Position[ch.Size] = *e.Components[wpcomponent.KeyPosition].(*wpcomponent.Position)
	ch.Player[ch.Size] = *e.Components[wpcomponent.KeyPlayer].(*wpcomponent.Player)
	ch.Size++
}
