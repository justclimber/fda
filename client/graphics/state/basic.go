package state

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/justclimber/fda/client/graphics"
)

type Basic struct{}

func NewBasicState() *Basic {
	return &Basic{}
}

func (b *Basic) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World from Basic State!")
}

func (b *Basic) Update() (graphics.ScreenState, error) {
	return nil, nil
}
