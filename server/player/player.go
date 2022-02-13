package player

import (
	"github.com/justclimber/fda/server/command"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type Player struct {
	Id    int64
	CmdCh chan command.Command
}

func NewPlayer(cmdCh chan command.Command) *Player {
	return &Player{CmdCh: cmdCh}
}

func NewPlayerWithComponent(delay int) (*Player, wpcomponent.Player) {
	cmdCh := make(chan command.Command, 1)
	return NewPlayer(cmdCh), wpcomponent.NewPlayer(delay, cmdCh)
}

func (p *Player) SendCommand(command command.Command) {
	p.CmdCh <- command
}
