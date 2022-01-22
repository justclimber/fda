package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/golang/geo/r1"
	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"

	"github.com/justclimber/fda/client/assets"
	"github.com/justclimber/fda/client/graphics"
	"github.com/justclimber/fda/client/graphics/camera"
	"github.com/justclimber/fda/client/graphics/input"
	"github.com/justclimber/fda/client/graphics/state"
)

const screenWidth = 1600
const screenHeight = 1500
const cameraMargin = 100

func main() {

	tilesPng, err := assets.EmbeddedFS.ReadFile("tmw_desert_spacing.png")
	if err != nil {
		log.Fatalf("can't read file: %v", err)
	}

	tilesImg, _, err := image.Decode(bytes.NewReader(tilesPng))
	if err != nil {
		log.Fatalf("can't decode file to image: %v", err)
	}

	tilesEbitenImage := ebiten.NewImageFromImage(tilesImg)

	tiledMap, err := tiled.LoadFile("testmap.tmx", tiled.WithFileSystem(assets.EmbeddedFS))
	if err != nil {
		log.Fatalf("can't load tmx file: %v", err)
	}

	tileSize := tiledMap.TileWidth
	cam := camera.NewCamera(r2.Rect{
		X: r1.Interval{Lo: cameraMargin, Hi: screenWidth - cameraMargin},
		Y: r1.Interval{Lo: cameraMargin, Hi: screenHeight - cameraMargin},
	}, tileSize)
	in := input.NewEbitenInput()
	tmxExampleState := state.NewTmxExample(tilesEbitenImage, tiledMap, in, cam, r2.Point{
		X: cameraMargin,
		Y: cameraMargin,
	})

	w := graphics.NewMainGameWindow("Tmx Example", screenWidth, screenHeight, tmxExampleState)
	graphics.Run(w)
}
