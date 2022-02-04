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
	debugger *debugger.Debugger
}

func NewPlayersProcessor(wpLink *internalapi.PpWpLink, debugger *debugger.Debugger) *PlayersProcessor {
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
		p.debugger.Printf("Run", "init [tick: %d]", p.currTick)
		select {
		case <-p.wpLink.DoneCh:
			return nil
		case logs := <-p.wpLink.LogsCh:
			p.debugger.Printf("Run", "get logs")
			if err := p.applyLogs(logs); err != nil {
				return err
			}
			p.processPlayers()
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
		p.debugger.Printf("Run", "Send command [tick: %d]", p.currTick)
		pl.SendCommand(cmd)
	}
}

func (p *PlayersProcessor) applyLogs(logs *worldlog.Logs) error {
	if logs == nil {
		return nil
	}
	p.currTick = logs.Entries[len(logs.Entries)-1].Tick
	return nil
}

func (p *PlayersProcessor) processPlayer(_ *player.Player) (command.Command, error) {
	return command.Command{Move: 0.5}, nil
}
