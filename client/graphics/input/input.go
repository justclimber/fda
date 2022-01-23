package input

import (
	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const moveK = 5
const scaleK = 0.01

type EbitenInput struct{}

func NewEbitenInput() *EbitenInput {
	return &EbitenInput{}
}

func (e *EbitenInput) ScrollChange() r2.Point {
	if inpututil.KeyPressDuration(ebiten.KeyMetaLeft) > 1 {
		return r2.Point{}
	}
	dx, dy := ebiten.Wheel()
	return r2.Point{X: dx * -moveK, Y: dy * -moveK}
}

func (e *EbitenInput) ScaleChange() float64 {
	if inpututil.KeyPressDuration(ebiten.KeyMetaLeft) < 1 {
		return 0
	}
	_, dy := ebiten.Wheel()
	return dy * scaleK
}
