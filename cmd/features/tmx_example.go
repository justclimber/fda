package main

import (
	_ "embed"
	_ "image/png"
	"log"

	"github.com/golang/geo/r1"
	"github.com/golang/geo/r2"

	"github.com/justclimber/fda/client/assets"
	"github.com/justclimber/fda/client/graphics"
	"github.com/justclimber/fda/client/graphics/camera"
	"github.com/justclimber/fda/client/graphics/input"
	"github.com/justclimber/fda/client/graphics/state"
	"github.com/justclimber/fda/common/ftmx"
)

const screenWidth = 1500
const screenHeight = 1500
const cameraMargin = 100

const tmxFileName = "testmap.tmx"
const tileImageFileName = "tmw_desert_spacing.png"

func main() {
	mapImage, err := ftmx.NewMapImage(tmxFileName, tileImageFileName, assets.EmbeddedFS)
	if err != nil {
		log.Fatalf("can't load map image: %v", err)
	}

	cam := camera.NewCamera(r2.Rect{
		X: r1.Interval{Lo: cameraMargin, Hi: screenWidth - cameraMargin},
		Y: r1.Interval{Lo: cameraMargin, Hi: screenHeight - cameraMargin},
	})
	in := input.NewEbitenInput()
	tmxExampleState := state.NewTmxExample(mapImage, in, cam, r2.Point{
		X: cameraMargin,
		Y: cameraMargin,
	})

	w := graphics.NewMainGameWindow("Tmx Example", screenWidth, screenHeight, tmxExampleState)
	graphics.Run(w)
}
