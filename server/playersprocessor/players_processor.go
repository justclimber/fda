package playersprocessor

import (
	"log"

	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/tick"
	"github.com/justclimber/fda/server/command"
	"github.com/justclimber/fda/server/internalapi"
	"github.com/justclimber/fda/server/player"
)

type PlayersProcessor struct {
	players       map[int64]*player.Player
	wpLink        *internalapi.PpWpLink
	currTick      tick.Tick
	entitiesState *EntitiesState
	debugger      *debugger.Nested
}

func NewPlayersProcessor(wpLink *internalapi.PpWpLink, debugger *debugger.Nested) *PlayersProcessor {
	return &PlayersProcessor{
		wpLink:        wpLink,
		players:       make(map[int64]*player.Player),
		entitiesState: NewEntitiesState(),
		debugger:      debugger,
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
		case logBatch := <-p.wpLink.LogsCh:
			p.debugger.LogF("Run", "get logs")
			p.entitiesState.ApplyLogBatch(logBatch)
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

func (p *PlayersProcessor) processPlayer(pl *player.Player) (command.Command, error) {
	return pl.Computer.Run()
}
