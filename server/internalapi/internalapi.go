package internalapi

import (
	"github.com/justclimber/fda/server/command"
	"github.com/justclimber/fda/server/worldlog"
)

type PpWpLink struct {
	LogsCh chan *worldlog.Logs
	CmdsCh chan map[int64]command.Command
	DoneCh chan bool
	SyncCh chan bool
}

func NewPpWpLink() *PpWpLink {
	return &PpWpLink{
		DoneCh: make(chan bool),
		SyncCh: make(chan bool),
		LogsCh: make(chan *worldlog.Logs, 1),
		CmdsCh: make(chan map[int64]command.Command, 1),
	}
}
