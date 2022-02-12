package wpsystem

import (
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/ecs/entity"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/worldprocessor/ecs/generated/wprepo"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type PlayerCommands struct {
	entityRepo *wprepo.RepoForMask6
	delay      int
	debugger   *debugger.Nested
}

func NewPlayerCommands(
	repoForMask6 *wprepo.RepoForMask6,
	delay int,
	debugger *debugger.Nested,
) *PlayerCommands {
	return &PlayerCommands{
		entityRepo: repoForMask6,
		delay:      delay,
		debugger:   debugger,
	}
}

func (p *PlayerCommands) String() string   { return "PlayerCommands" }
func (p *PlayerCommands) Init(_ tick.Tick) {}

func (p *PlayerCommands) DoTick(_ tick.Tick) bool {
	p.entityRepo.Iterate(func(_ entity.Id, mov wpcomponent.Moving, pl wpcomponent.Player) (*wpcomponent.Moving, *wpcomponent.Player) {
		pl.Delay--
		if pl.Delay > 0 {
			return nil, &pl
		}
		select {
		case cmd := <-pl.CmdCh:
			p.debugger.LogF("DoTick", "get commands")
			newMov := wpcomponent.Moving{D: mov.D.Add(fgeom.Point{X: cmd.Move})}
			return &newMov, &pl
		default:
		}
		return nil, nil
	})
	return false
}
