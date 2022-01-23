package state

import (
	"fmt"

	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/justclimber/fda/client/graphics"
	ebitenhelper "github.com/justclimber/fda/client/graphics/ebiten"
	"github.com/justclimber/fda/common/ftmx"
)

type camera interface {
	ViewRect() r2.Rect
	Move(p r2.Point)
}

type input interface {
	ScrollChange() r2.Point
}

type TmxExample struct {
	mapImage  *ftmx.MapImage
	camera    camera
	input     input
	cameraPos r2.Point
}

func NewTmxExample(mapImage *ftmx.MapImage, input input, camera camera, cameraPos r2.Point) *TmxExample {
	return &TmxExample{
		mapImage:  mapImage,
		input:     input,
		camera:    camera,
		cameraPos: cameraPos,
	}
}

func (t *TmxExample) Draw(screen *ebiten.Image) {
	imageUnderCamera := t.mapImage.ImageUnderCamera(t.camera)
	screen.DrawImage(imageUnderCamera, ebitenhelper.WithOffset(t.cameraPos, nil))

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
	))
}

func (t *TmxExample) Update() (graphics.ScreenState, error) {
	t.camera.Move(t.input.ScrollChange())
	return nil, nil
}
