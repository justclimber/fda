package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"

	"github.com/justclimber/fda/client/assets"
	"github.com/justclimber/fda/client/graphics"
	"github.com/justclimber/fda/client/graphics/state"
)

func main() {
	const screenWidth = 640
	const screenHeight = 640

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
	tmxExampleState := state.NewTmxExample(tilesEbitenImage, tiledMap)

	w := graphics.NewMainGameWindow("Tmx Example", screenWidth, screenHeight, tmxExampleState)
	graphics.Run(w)
}
