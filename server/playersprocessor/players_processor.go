package playersprocessor

import (
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/player"
	"github.com/justclimber/fda/server/worldlog"
)

type PlayersProcessor struct {
	players  []*player.Player
	wpLink   *internalapi.PpWpLink
	currTick tick.Tick
}

func NewPlayersProcessor(wpLink *internalapi.PpWpLink) *PlayersProcessor {
	return &PlayersProcessor{
		wpLink:  wpLink,
		players: []*player.Player{},
	}
}

func (p *PlayersProcessor) Run() error {
	for {
		err, stop := p.updateWorld()
		if err != nil {
			return err
		}
		if stop {
			return nil
		}
	}
}

func (p *PlayersProcessor) AddPlayer(pl *player.Player) {
	p.players[pl.Id] = pl
}

func (p *PlayersProcessor) updateWorld() (error, bool) {
	return p.applyLogs(<-p.wpLink.LogsCh)
}

func (p *PlayersProcessor) applyLogs(logs *worldlog.Logs) (error, bool) {
	p.currTick = logs.Entries[len(logs.Entries)-1].Tick
	return nil, false
}
