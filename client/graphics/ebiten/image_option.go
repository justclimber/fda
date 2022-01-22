package ebiten

import (
	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
)

func WithIntOffset(x, y int) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	return op
}

func WithOffset(p r2.Point, op *ebiten.DrawImageOptions) *ebiten.DrawImageOptions {
	if op == nil {
		op = &ebiten.DrawImageOptions{}
	}
	op.GeoM.Translate(p.X, p.Y)
	return op
}
