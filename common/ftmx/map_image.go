package ftmx

import (
	"bytes"
	"fmt"
	"image"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"

	ebitenhelper "github.com/justclimber/fda/client/graphics/ebiten"
	"github.com/justclimber/fda/common/fgeom"
)

type camera interface {
	ViewRect() fgeom.Rect
	ScaleFactor() float64
}

type MapImage struct {
	img *ebiten.Image
}

func NewMapImage(tmxFileName, tileImageFileName string, fileSystem fs.ReadFileFS) (*MapImage, error) {
	m := &MapImage{}
	if err := m.load(tmxFileName, tileImageFileName, fileSystem); err != nil {
		return nil, err
	}
	return m, nil
}

func (m MapImage) ImageUnderCamera(camera camera) (*ebiten.Image, fgeom.Point) {
	viewRect := camera.ViewRect()
	img := m.img.SubImage(viewRect.ToImageRect()).(*ebiten.Image)

	return img, m.leftTopOffsetIfOutOfBound(viewRect.Lo(), camera.ScaleFactor())
}

func (m MapImage) leftTopOffsetIfOutOfBound(leftTop fgeom.Point, scaleFactor float64) fgeom.Point {
	var x, y float64
	if leftTop.X < 0 {
		x = -leftTop.X / scaleFactor
	}
	if leftTop.Y < 0 {
		y = -leftTop.Y / scaleFactor
	}
	return fgeom.Point{
		X: x,
		Y: y,
	}
}

func (m *MapImage) load(tmxFileName, tileImageFileName string, fileSystem fs.ReadFileFS) error {
	tiledMap, err := tiled.LoadFile(tmxFileName, tiled.WithFileSystem(fileSystem))
	if err != nil {
		return fmt.Errorf("can't load tmx file: %v", err)
	}

	tilesImage, err := m.tilesImage(err, fileSystem, tileImageFileName)
	if err != nil {
		return err
	}

	img := ebiten.NewImage(tiledMap.Width*tiledMap.TileWidth, tiledMap.Height*tiledMap.TileHeight)
	for _, layer := range tiledMap.Layers {
		for i, tile := range layer.Tiles {
			if tile.Nil {
				continue
			}
			tileRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := tilesImage.SubImage(tileRect).(*ebiten.Image)

			x, y := layer.GetTilePosition(i)
			img.DrawImage(tileImage, ebitenhelper.WithIntOffset(x, y, nil))
		}
	}
	m.img = img
	return nil
}

func (m *MapImage) tilesImage(err error, fileSystem fs.ReadFileFS, tileImageFileName string) (*ebiten.Image, error) {
	tilesPng, err := fileSystem.ReadFile(tileImageFileName)
	if err != nil {
		return nil, fmt.Errorf("can't read file: %v", err)
	}

	tilesImg, _, err := image.Decode(bytes.NewReader(tilesPng))
	if err != nil {
		return nil, fmt.Errorf("can't decode file to image: %v", err)
	}

	tilesImage := ebiten.NewImageFromImage(tilesImg)
	return tilesImage, nil
}
