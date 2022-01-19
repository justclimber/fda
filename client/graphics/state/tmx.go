package state

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"

	"github.com/justclimber/fda/client/graphics"
)

const tileSize = 32

type TmxExample struct {
	screenWidth  int
	screenHeight int
	tilesImage   *ebiten.Image
	tiledMap     *tiled.Map
}

func NewTmxExample(screenWidth, screenHeight int, tilesImage *ebiten.Image, tiledMap *tiled.Map) *TmxExample {
	return &TmxExample{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		tilesImage:   tilesImage,
		tiledMap:     tiledMap,
	}
}

func (t *TmxExample) Draw(screen *ebiten.Image) {
	xNum := 20
	for _, layer := range t.tiledMap.Layers {
		for i, tile := range layer.Tiles {
			if tile.Nil {
				continue
			}
			tileRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := t.tilesImage.SubImage(tileRect).(*ebiten.Image)

			op := &ebiten.DrawImageOptions{}
			tx := float64((i % xNum) * tileSize)
			ty := float64((i / xNum) * tileSize)
			op.GeoM.Translate(tx, ty)

			screen.DrawImage(tileImage, op)
		}
	}
}

func (t *TmxExample) Update() (graphics.ScreenState, error) {
	return nil, nil
}
