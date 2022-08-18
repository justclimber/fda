package main

import (
	"github.com/justclimber/fda/client/graphics"
	"github.com/justclimber/fda/client/graphics/state"
)

func main() {
	ideState := state.NewIDEState()
	w := graphics.NewMainGameWindow("Hello world!!!!", 1000, 500, ideState)
	graphics.Run(w)
}
