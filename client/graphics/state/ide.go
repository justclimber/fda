package state

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/justclimber/fda/client/graphics"
)

type IDE struct{}

func NewIDEState() *IDE {
	return &IDE{}
}

func (i *IDE) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello World from IDE State!")
}

func (i *IDE) Update() (graphics.ScreenState, error) {
	return nil, nil
}
