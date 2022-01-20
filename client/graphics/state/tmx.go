package state

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"

	"github.com/justclimber/fda/client/graphics"
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

			op := &ebiten.DrawImageOptions{}
			x, y := layer.GetTilePosition(i)
			op.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(tileImage, op)
		}
	}
}

func (t *TmxExample) Update() (graphics.ScreenState, error) {
	return nil, nil
}
