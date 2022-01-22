package input

import (
	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
)

const k = 5

type EbitenInput struct{}

func NewEbitenInput() *EbitenInput {
	return &EbitenInput{}
}

func (e *EbitenInput) ScrollChange() r2.Point {
	dx, dy := ebiten.Wheel()
	return r2.Point{X: dx * -k, Y: dy * -k}
}
