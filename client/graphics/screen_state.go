package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ScreenState interface {
	Draw(screen *ebiten.Image)
	Update() (ScreenState, error)
}
