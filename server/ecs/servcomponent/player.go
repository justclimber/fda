package servcomponent

import (
	"github.com/justclimber/fda/server/command"
)

type Player struct {
	Delay int
	CmdCh chan command.Command
}

func NewPlayer(delay int, cmdCh chan command.Command) *Player {
	return &Player{Delay: delay, CmdCh: cmdCh}
}
