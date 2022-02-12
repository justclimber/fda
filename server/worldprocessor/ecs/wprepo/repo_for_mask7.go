package wprepo

import (
	"github.com/justclimber/fda/common/ecs"
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/common/ecs/entityrepo"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type RepoForMask6 struct {
	cGroups  []entityrepo.CGroup
	repoLink ecs.EntityRepo
}

func NewRepoForMask6(repoLink ecs.EntityRepo) *RepoForMask6 {
	return &RepoForMask6{
		repoLink: repoLink,
	}
}

func (ci *RepoForMask6) InitRepoLink(mask component.Mask) {
	ci.cGroups = ci.repoLink.GetCGroupsWithMask(mask)
}

func (ci *RepoForMask6) Iterate(f func(
	moving wpcomponent.Moving,
	position wpcomponent.Player,
) (*wpcomponent.Moving, *wpcomponent.Player)) {
	for _, cGroup := range ci.cGroups {
		switch cg := cGroup.(type) {
		case *CGroup7:
			for _, chunk := range cg.Chunks {
				for k := 0; k < chunk.Size; k++ {
					newMoving, newPlayer := f(chunk.Moving[k], chunk.Player[k])
					if newMoving != nil {
						chunk.Moving[k] = *newMoving
					}
					if newPlayer != nil {
						chunk.Player[k] = *newPlayer
					}
				}
			}
		}
	}
}
