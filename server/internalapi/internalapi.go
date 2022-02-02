package internalapi

import (
	"github.com/justclimber/fda/server/worldlog"
)

type PpWpLink struct {
	LogsCh chan *worldlog.Logs
}

func NewPpWpLink() *PpWpLink {
	return &PpWpLink{
		LogsCh: make(chan *worldlog.Logs, 1),
	}
}
