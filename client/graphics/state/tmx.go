package state

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"

	"github.com/justclimber/fda/client/graphics"
	ebitenhelper "github.com/justclimber/fda/client/graphics/ebiten"
)

type TmxExample struct {
	tilesImage *ebiten.Image
	tiledMap   *tiled.Map
}

func NewTmxExample(tilesImage *ebiten.Image, tiledMap *tiled.Map) *TmxExample {
	return &TmxExample{
		tilesImage: tilesImage,
		tiledMap:   tiledMap,
	}
}

func (t *TmxExample) Draw(screen *ebiten.Image) {
	for _, layer := range t.tiledMap.Layers {
		for i, tile := range layer.Tiles {
			if tile.Nil {
				continue
			}
			tileRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := t.tilesImage.SubImage(tileRect).(*ebiten.Image)

			x, y := layer.GetTilePosition(i)
			screen.DrawImage(tileImage, ebitenhelper.WithIntOffset(x, y))
		}
	}
}

func (t *TmxExample) Update() (graphics.ScreenState, error) {
	return nil, nil
}
