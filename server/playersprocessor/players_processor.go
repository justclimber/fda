package playersprocessor

import (
	"github.com/justclimber/fda/server/player"
)

type PlayersProcessor struct {
	players []*player.Player
}

func NewPlayersProcessor() *PlayersProcessor {
	return &PlayersProcessor{
		players: []*player.Player{},
	}
}
