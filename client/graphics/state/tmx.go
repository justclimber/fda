package state

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/justclimber/fda/client/graphics"
	ebitenhelper "github.com/justclimber/fda/client/graphics/ebiten"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/ftmx"
)

type camera interface {
	ViewRect() fgeom.Rect
	Move(p fgeom.Point)
	Scale(scaleFactorChange float64)
	ScaleFactor() float64
}

type input interface {
	ScrollChange() fgeom.Point
	ScaleChange() float64
}

type TmxExample struct {
	mapImage  *ftmx.MapImage
	camera    camera
	input     input
	cameraPos fgeom.Point
}

func NewTmxExample(mapImage *ftmx.MapImage, input input, camera camera, cameraPos fgeom.Point) *TmxExample {
	return &TmxExample{
		mapImage:  mapImage,
		input:     input,
		camera:    camera,
		cameraPos: cameraPos,
	}
}

func (t *TmxExample) Draw(screen *ebiten.Image) {
	imageUnderCamera, offset := t.mapImage.ImageUnderCamera(t.camera)
	op := &ebiten.DrawImageOptions{}
	op = ebitenhelper.WithScale(1/t.camera.ScaleFactor(), op)
	op = ebitenhelper.WithOffset(t.cameraPos, op)
	op = ebitenhelper.WithOffset(offset, op)
	screen.DrawImage(imageUnderCamera, op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f\nScale factor: %0.2f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		t.camera.ScaleFactor(),
	))
}

func (t *TmxExample) Update() (graphics.ScreenState, error) {
	t.camera.Move(t.input.ScrollChange())
	t.camera.Scale(t.input.ScaleChange())
	return nil, nil
}
