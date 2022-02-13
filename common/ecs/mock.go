package ecs

/*

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

const chunkSize = 10

type Moving2 struct {
	repoForMask3 *RepoForMask3
}

func NewMoving2(compIterator *RepoForMask3) *Moving2 {
	m := &Moving2{repoForMask3: compIterator}
	m.repoForMask3.InitRepoLink(m.mask())
	return m
}

func (m *Moving2) String() string   { return "Moving" }
func (m *Moving2) Init(_ tick.Tick) {}

func (m *Moving2) mask() component.Mask {
	return component.NewMask([]component.Key{wpcomponent.KeyMoving, wpcomponent.KeyPosition})
}

func (m *Moving2) DoTick(_ tick.Tick) bool {
	m.repoForMask3.Iterate(func(mov wpcomponent.Moving, p wpcomponent.Position) (*wpcomponent.Moving, *wpcomponent.Position) {
		if mov.D.Empty() {
			return nil, nil
		}
		return nil, &wpcomponent.Position{Pos: p.Pos.Add(mov.D)}
	})
	return false
}

type RepoForMask3 struct {
	cGroups  []entityrepo.CGroup
	repoLink EntityRepo
}

func NewRepoForMask3(repoLink EntityRepo) *RepoForMask3 {
	return &RepoForMask3{
		repoLink: repoLink,
	}
}

func (ci *RepoForMask3) InitRepoLink(mask component.Mask) {
	ci.cGroups = ci.repoLink.GetCGroupsWithMask(mask)
}

func (ci *RepoForMask3) Iterate(f func(
	moving wpcomponent.Moving,
	position wpcomponent.Position,
) (*wpcomponent.Moving, *wpcomponent.Position)) {
	for _, cGroup := range ci.cGroups {
		switch cg := cGroup.(type) {
		case *CGroup3:
			for _, chunk := range cg.Chunks {
				for k := 0; k < chunk.Size; k++ {
					newMoving, newPosition := f(chunk.Moving[k], chunk.Position[k])
					if newMoving != nil {
						chunk.Moving[k] = *newMoving
					}
					if newPosition != nil {
						chunk.Position[k] = *newPosition
					}
				}
			}
		case *CGroup7:
			for _, chunk := range cg.Chunks {
				for k := 0; k < chunk.Size; k++ {
					newMoving, newPosition := f(chunk.Moving[k], chunk.Position[k])
					if newMoving != nil {
						chunk.Moving[k] = *newMoving
					}
					if newPosition != nil {
						chunk.Position[k] = *newPosition
					}
				}
			}
		}
	}
}

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

type Chunk3 struct {
	Size     int
	Ids      [chunkSize]entity.Id
	Moving   [chunkSize]wpcomponent.Moving
	Position [chunkSize]wpcomponent.Position
}

type Chunk7 struct {
	Size     int
	Ids      [chunkSize]entity.Id
	Moving   [chunkSize]wpcomponent.Moving
	Position [chunkSize]wpcomponent.Position
	Player   [chunkSize]wpcomponent.Player
}

func (ch *Chunk3) Add(e entity.Entity) {
	ch.Ids[ch.Size] = e.Id
	ch.Moving[ch.Size] = *e.Components[wpcomponent.KeyMoving].(*wpcomponent.Moving)
	ch.Position[ch.Size] = *e.Components[wpcomponent.KeyPosition].(*wpcomponent.Position)
	ch.Size++
}

func (ch *Chunk7) Add(e entity.Entity) {
	ch.Ids[ch.Size] = e.Id
	ch.Moving[ch.Size] = *e.Components[wpcomponent.KeyMoving].(*wpcomponent.Moving)
	ch.Position[ch.Size] = *e.Components[wpcomponent.KeyPosition].(*wpcomponent.Position)
	ch.Player[ch.Size] = *e.Components[wpcomponent.KeyPlayer].(*wpcomponent.Player)
	ch.Size++
}
*/
