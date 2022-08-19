package main

import (
	"github.com/justclimber/fda/client/assets"
	"github.com/justclimber/fda/client/graphics"
	"github.com/justclimber/fda/client/graphics/state"
)

func main() {
	ideState := state.NewIDEState()
	_ = ideState.Setup(assets.EmbeddedFS)
	w := graphics.NewMainGameWindow("Hello world!!!!", 1000, 500, ideState)
	graphics.Run(w)
}
