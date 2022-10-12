package state

import (
	"embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	"github.com/justclimber/fda/client/graphics"
	iderenderer "github.com/justclimber/fda/client/graphics/ide"
	ui "github.com/justclimber/fda/client/graphics/ui"
	"github.com/justclimber/fda/client/graphics/ui/font"
	"github.com/justclimber/fda/client/graphics/ui/widget"
	"github.com/justclimber/fda/client/ide"
	"github.com/justclimber/fda/client/ide/ast"
	"github.com/justclimber/fda/client/ide/program"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/lang/executor/object"
)

type IDEState struct {
	ui          *ui.UI
	ideObj      *ide.IDE
	ideRenderer *iderenderer.Renderer
}

func NewIDEState() *IDEState {
	return &IDEState{}
}

func (is *IDEState) Draw(screen *ebiten.Image) {
	//is.ui.Draw(screen)
	is.ideRenderer.Draw(screen)
	is.ideObj.Render(is.ideRenderer)
}

func (is *IDEState) Update() (graphics.ScreenState, error) {
	is.ui.HandleInput()
	return nil, nil
}

func (is *IDEState) Setup(assets embed.FS) error {
	f, err := font.LoadFont("FiraCode-Regular.ttf", 14, assets)
	if err != nil {
		return err
	}

	packageName := "main"
	funcDefinition := object.NewFunctionDefinition("testFunc", packageName, nil, nil)
	identifier1 := ast.NewIdentifier(0, "testVar1")
	identifier2 := ast.NewIdentifier(0, "testVar21")
	expr1 := ast.NewNumInt(0, 124)
	expr2 := ast.NewNumInt(0, 100000)
	assignment1 := ast.NewAssignment(0, []*ast.Identifier{identifier1, identifier2}, expr1)
	assignment2 := ast.NewAssignment(0, []*ast.Identifier{identifier2}, expr2)
	funcBody := ast.NewStatementsBlock(0, []ast.Stmt{assignment1, assignment2})
	function := ast.NewFunction(0, funcDefinition, funcBody)
	pkg := ast.NewPackage(0, packageName)
	_ = pkg.RegisterFunction(function)
	pkgist := program.NewPackagist()
	_ = pkgist.RegisterPackage(pkg)
	prog := program.NewProgram(pkgist)
	tab := ide.NewTab("tab1", pkg)
	is.ideObj = ide.NewIDE(prog, []*ide.Tab{tab}, 0)

	is.ideRenderer = iderenderer.NewRenderer(iderenderer.Options{
		IndentWidth:        3,
		Face:               f,
		LineDistanceFactor: iderenderer.LineDistanceNormal,
		DefaultColor:       colornames.White,
		TypeColorMap: map[ast.TextType]color.Color{
			ast.TypeSystemSymbols: colornames.White,
			ast.TypeKeywords:      colornames.Orange,
			ast.TypeIdentifier:    colornames.Lightcoral,
			ast.TypeNumbers:       colornames.Aqua,
		},
		Text: iderenderer.PredefinedText{
			ArgDelimiter: ", ",
			Assignment:   ": ",
			Package:      "package ",
			Function:     "func",
		},
		TabOptions: iderenderer.TabOptions{
			HeaderPadding:         3,
			BodyPadding:           3,
			Size:                  fgeom.Point{X: 500, Y: 300},
			HeaderBackgroundColor: color.RGBA{R: 20, G: 20, B: 20, A: 0xff},
			BodyBackgroundColor:   color.RGBA{R: 35, G: 35, B: 35, A: 0xff},
		},
	}, fgeom.Point{X: 5, Y: 5})

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
		widget.ContainerOpts.WidgetOpts(widget.Opts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
		)),
	)

	textInputOpts := []widget.TextInputOpt{
		widget.TextInputOpts.WidgetOpts(widget.Opts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.TextInputOpts.Color(&widget.TextInputColor{
			Idle:          colornames.White,
			Disabled:      colornames.White,
			Caret:         colornames.White,
			DisabledCaret: colornames.White,
		}),
		widget.TextInputOpts.Padding(widget.Insets{
			Left:   13,
			Right:  13,
			Top:    7,
			Bottom: 7,
		}),
		widget.TextInputOpts.Face(f),
		widget.TextInputOpts.CaretOpts(
			widget.CaretOpts.Size(f, 2),
		),
		widget.TextInputOpts.Placeholder("Enter text here"),
	}

	mainContainer.AddChild(widget.NewTextInput(textInputOpts...))

	rootContainer.AddChild(mainContainer)

	is.ui = &ui.UI{Container: rootContainer}

	return nil
}
