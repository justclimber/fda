package state

import (
	"embed"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	"github.com/justclimber/fda/client/graphics"
	ui "github.com/justclimber/fda/client/graphics/ui"
	"github.com/justclimber/fda/client/graphics/ui/font"
	"github.com/justclimber/fda/client/graphics/ui/widget"
)

type IDEState struct {
	ui *ui.UI
}

func NewIDEState() *IDEState {
	return &IDEState{}
}

func (is *IDEState) Draw(screen *ebiten.Image) {
	is.ui.Draw(screen)
}

func (is *IDEState) Update() (graphics.ScreenState, error) {
	return nil, nil
}

func (is *IDEState) Setup(assets embed.FS) error {
	f, err := font.LoadFont("NotoSans-Regular.ttf", 14, assets)
	if err != nil {
		return err
	}
	rootContainer := widget.NewContainer(
		"root",
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Stretch([]bool{true}, []bool{false, true, false}),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(20)),
			widget.GridLayoutOpts.Spacing(0, 20))),
	)
	is.ui = &ui.UI{Container: rootContainer}

	rootContainer.AddChild(widget.NewText(
		widget.TextOpts.Text(
			"Header",
			f,
			colornames.White,
		),
	))

	mainContainer := widget.NewContainer(
		"main",
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
		)),
	)
	rootContainer.AddChild(mainContainer)

	is.ui = &ui.UI{Container: rootContainer}

	return nil
}
