package player

import (
	"github.com/justclimber/fda/server/command"
	"github.com/justclimber/fda/server/ecs/servcomponent"
)

type Player struct {
	Id    int64
	CmdCh chan command.Command
}

func NewPlayer(cmdCh chan command.Command) *Player {
	return &Player{CmdCh: cmdCh}
}

func NewPlayerWithComponent(delay int) (*Player, *servcomponent.Player) {
	cmdCh := make(chan command.Command, 1)
	return NewPlayer(cmdCh), servcomponent.NewPlayer(delay, cmdCh)
}

func (p *Player) SendCommand(command command.Command) {
	p.CmdCh <- command
}
