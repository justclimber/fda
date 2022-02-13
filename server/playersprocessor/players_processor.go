package playersprocessor

import (
	"log"

	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/command"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/player"
	"github.com/justclimber/fda/server/worldlog"
)

type PlayersProcessor struct {
	players  map[int64]*player.Player
	wpLink   *internalapi.PpWpLink
	currTick tick.Tick
	debugger *debugger.Nested
}

func NewPlayersProcessor(wpLink *internalapi.PpWpLink, debugger *debugger.Nested) *PlayersProcessor {
	return &PlayersProcessor{
		wpLink:   wpLink,
		players:  make(map[int64]*player.Player),
		debugger: debugger,
	}
}

func (p *PlayersProcessor) AddPlayer(pl *player.Player) {
	p.players[pl.Id] = pl
}

func (p *PlayersProcessor) Run() error {
	// todo: init world?
	for {
		p.debugger.LogF("Run", "[tick: %d]", p.currTick)
		select {
		case <-p.wpLink.DoneCh:
			return nil
		case logs := <-p.wpLink.LogsCh:
			p.debugger.LogF("Run", "get logs")
			if err := p.applyLogs(logs); err != nil {
				return err
			}
			p.processPlayers()
			p.debugger.LogF("Run", "send sync")
			p.wpLink.SyncCh <- true
		}
	}
}

func (p *PlayersProcessor) processPlayers() {
	for _, pl := range p.players {
		cmd, err := p.processPlayer(pl)
		if err != nil {
			log.Println(err)
		}
		p.debugger.LogF("Run", "Send command [tick: %d]", p.currTick)
		pl.SendCommand(cmd)
	}
}

func (p *PlayersProcessor) applyLogs(logs worldlog.LogBatch) error {
	p.currTick = logs.StartTick
	return nil
}

func (p *PlayersProcessor) processPlayer(_ *player.Player) (command.Command, error) {
	return command.Command{Move: 0.5}, nil
}
