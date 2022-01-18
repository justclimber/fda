package main

import (
	"github.com/justclimber/fda/client/graphics"
	"github.com/justclimber/fda/client/graphics/state"
)

func main() {
	basicState := state.NewBasicState()
	w := graphics.NewMainGameWindow("Hello world!!!!", 1000, 500, basicState)
	graphics.Run(w)
}
