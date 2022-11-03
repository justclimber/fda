package ebiteninput

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/justclimber/fda/client/graphics/input"
	"github.com/justclimber/fda/common/fgeom"
)

const moveK = 5
const scaleK = 0.01

type EbitenInput struct{}

func NewEbitenInput() *EbitenInput {
	return &EbitenInput{}
}

func (e *EbitenInput) ScrollChange() fgeom.Point {
	if inpututil.KeyPressDuration(ebiten.KeyMetaLeft) > 1 {
		return fgeom.EmptyPoint
	}
	dx, dy := ebiten.Wheel()
	return fgeom.Point{X: dx * -moveK, Y: dy * -moveK}
}

func (e *EbitenInput) ScaleChange() float64 {
	if inpututil.KeyPressDuration(ebiten.KeyMetaLeft) < 1 {
		return 0
	}
	_, dy := ebiten.Wheel()
	return dy * scaleK
}

func (e *EbitenInput) WhichControlArrowsPressed() input.ControlArrow {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		return input.ControlArrowUp
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		return input.ControlArrowLeft
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		return input.ControlArrowDown
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		return input.ControlArrowRight
	}
	return input.ControlArrowNone
}
