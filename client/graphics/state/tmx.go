package state

import (
	"fmt"

	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"

	"github.com/justclimber/fda/client/graphics"
	ebitenhelper "github.com/justclimber/fda/client/graphics/ebiten"
	"github.com/justclimber/fda/common/fgeom"
)

type camera interface {
	Offset(p r2.Point) (r2.Point, bool)
	Size() r2.Point
	Move(p r2.Point)
}

type input interface {
	ScrollChange() r2.Point
}

type TmxExample struct {
	tilesImage    *ebiten.Image
	tiledMap      *tiled.Map
	camera        camera
	input         input
	rectForCamera r2.Rect
	cameraPos     r2.Point
}

func NewTmxExample(tilesImage *ebiten.Image, tiledMap *tiled.Map, input input, camera camera, cameraPos r2.Point) *TmxExample {
	return &TmxExample{
		tilesImage:    tilesImage,
		tiledMap:      tiledMap,
		input:         input,
		camera:        camera,
		rectForCamera: fgeom.RectFromPointAndSize(cameraPos, camera.Size()),
		cameraPos:     cameraPos,
	}
}

func (t *TmxExample) Draw(screen *ebiten.Image) {
	tilesRendered := 0
	ebitenhelper.DrawRect(t.rectForCamera, screen)
	for _, layer := range t.tiledMap.Layers {
		for i, tile := range layer.Tiles {
			if tile.Nil {
				continue
			}
			x, y := layer.GetTilePosition(i)
			p, isOutOfBounds := t.camera.Offset(r2.Point{
				X: float64(x),
				Y: float64(y),
			})
			if isOutOfBounds {
				continue
			}
			tileRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := t.tilesImage.SubImage(tileRect).(*ebiten.Image)
			p = p.Add(t.cameraPos).Sub(r2.Point{
				X: float64(t.tiledMap.TileWidth),
				Y: float64(t.tiledMap.TileWidth),
			})
			op := &ebiten.DrawImageOptions{}
			op.CompositeMode = ebiten.CompositeModeSourceIn
			//op.CompositeMode = ebiten.CompositeModeLighter
			screen.DrawImage(tileImage, ebitenhelper.WithOffset(p, op))
			tilesRendered++
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"TPS: %0.2f\nFPS: %0.2f\nTiles rendered: %d",
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		tilesRendered,
	))
}

func (t *TmxExample) Update() (graphics.ScreenState, error) {
	t.camera.Move(t.input.ScrollChange())
	return nil, nil
}
