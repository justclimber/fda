package wprepo

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/ecs/entityrepo"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type RepoForMask3 struct {
	cGroups  []entityrepo.CGroup
	repoLink ecs.EntityRepo
}

func NewRepoForMask3(repoLink ecs.EntityRepo) *RepoForMask3 {
	return &RepoForMask3{
		repoLink: repoLink,
	}
}

func (ci *RepoForMask3) InitRepoLink(mask component.Mask) {
	ci.cGroups = ci.repoLink.GetCGroupsWithMask(mask)
}

func (ci *RepoForMask3) Iterate(f func(
	id entity.Id,
	moving wpcomponent.Moving,
	position wpcomponent.Position,
) (*wpcomponent.Moving, *wpcomponent.Position)) {
	for _, cGroup := range ci.cGroups {
		switch cg := cGroup.(type) {
		case *CGroup3:
			for _, chunk := range cg.Chunks {
				for k := 0; k < chunk.Size; k++ {
					newMoving, newPosition := f(chunk.Ids[k], chunk.Moving[k], chunk.Position[k])
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
					newMoving, newPosition := f(chunk.Ids[k], chunk.Moving[k], chunk.Position[k])
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
