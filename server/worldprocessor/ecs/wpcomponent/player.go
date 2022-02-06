package wpcomponent

import (
	"github.com/justclimber/fda/common/ecs/component"
	"github.com/justclimber/fda/server/command"
)

const CPlayer component.Key = "pl"

type Player struct {
	Delay int
	CmdCh chan command.Command
}

func NewPlayer(delay int, cmdCh chan command.Command) *Player {
	return &Player{Delay: delay, CmdCh: cmdCh}
}

func (p *Player) Key() component.Key {
	return CPlayer
}
