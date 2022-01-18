package graphics

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type MainGameWindow struct {
	title         string
	width, height int
	screenState   ScreenState
}

func NewMainGameWindow(title string, width, height int, state ScreenState) *MainGameWindow {
	return &MainGameWindow{
		title:       title,
		width:       width,
		height:      height,
		screenState: state,
	}
}

func (w *MainGameWindow) Draw(screen *ebiten.Image) {
	w.screenState.Draw(screen)
}

func (w *MainGameWindow) Update() error {
	newState, err := w.screenState.Update()
	if err != nil {
		return err
	}

	if newState != nil {
		w.screenState = newState
	}
	return nil
}

func (w *MainGameWindow) GetSize() (int, int) {
	return w.width, w.height
}

func (w *MainGameWindow) Title() string {
	return w.title
}

func (w *MainGameWindow) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func Run(w *MainGameWindow) {
	ebiten.SetWindowSize(w.GetSize())
	ebiten.SetWindowTitle(w.Title())
	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}
}
