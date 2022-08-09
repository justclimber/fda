package player

import (
	"github.com/justclimber/fda/common/computer"
	"github.com/justclimber/fda/server/command"
	"github.com/justclimber/fda/server/worldprocessor/ecs/wpcomponent"
)

type Player struct {
	Id       int64
	CmdCh    chan command.Command
	Computer *computer.Computer
}

func NewPlayer(cmdCh chan command.Command, computer *computer.Computer) *Player {
	return &Player{CmdCh: cmdCh, Computer: computer}
}

func NewPlayerWithComponent(delay int, computer *computer.Computer) (*Player, wpcomponent.Player) {
	cmdCh := make(chan command.Command, 1)
	return NewPlayer(cmdCh, computer), wpcomponent.NewPlayer(delay, cmdCh)
}

func (p *Player) SendCommand(command command.Command) {
	p.CmdCh <- command
}
