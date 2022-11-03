package state

import (
	"embed"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"

	"github.com/justclimber/fda/client/graphics"
	"github.com/justclimber/fda/client/graphics/ebiteninput"
	iderenderer "github.com/justclimber/fda/client/graphics/ide"
	"github.com/justclimber/fda/client/graphics/ui/font"
	"github.com/justclimber/fda/client/ide"
	"github.com/justclimber/fda/client/ide/ast"
	"github.com/justclimber/fda/client/ide/program"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/lang/executor/object"
)

type IDEState struct {
	ideObj      *ide.IDE
	ideRenderer *iderenderer.Renderer
}

func NewIDEState() *IDEState {
	return &IDEState{}
}

func (is *IDEState) Draw(screen *ebiten.Image) {
	is.ideRenderer.Draw(screen)
	is.ideRenderer.HighlightActiveNode(screen)
	is.ideRenderer.DebugPrintIndex(screen)
}

func (is *IDEState) Update() (graphics.ScreenState, error) {
	is.ideRenderer.HandleUserInput()
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
	identifier3 := ast.NewIdentifier(0, "testVar3")
	expr1 := ast.NewNumInt(0, 124)
	expr2 := ast.NewNumInt(0, 100000)
	expr3 := ast.NewNumInt(0, 100)
	assignment1 := ast.NewAssignment(0, []*ast.Identifier{identifier1, identifier2}, expr1)
	assignment2 := ast.NewAssignment(0, []*ast.Identifier{identifier2}, expr2)
	ifStatement := ast.NewIfStatement(
		123,
		ast.NewNumInt(0, 333),
		ast.NewStatementsBlock(0, []ast.Stmt{assignment1}),
		ast.NewStatementsBlock(0, []ast.Stmt{assignment2}),
	)
	assignment3 := ast.NewAssignment(0, []*ast.Identifier{identifier3}, expr3)
	funcBody := ast.NewStatementsBlock(0, []ast.Stmt{ifStatement, assignment3})
	function := ast.NewFunction(0, funcDefinition, funcBody)
	pkg := ast.NewPackage(0, packageName)
	_ = pkg.RegisterFunction(function)
	pkgist := program.NewPackagist()
	_ = pkgist.RegisterPackage(pkg)
	prog := program.NewProgram(pkgist)
	tab1 := ide.NewTab("tab1", pkg)
	tab2 := ide.NewTab("some_tab1", nil)
	tab3 := ide.NewTab("some_tab2", nil)
	is.ideObj = ide.NewIDE(prog, []*ide.Tab{tab1, tab2, tab3}, 0)

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
			IfStart:      "if ",
			IfElse:       "else",
		},
		TabOptions: iderenderer.TabOptions{
			HeaderSpacing:         4,
			HeaderPadding:         5,
			BodyPadding:           5,
			Size:                  fgeom.Point{X: 500, Y: 300},
			HeaderBackgroundColor: color.RGBA{R: 20, G: 20, B: 20, A: 0xff},
			BodyBackgroundColor:   color.RGBA{R: 35, G: 35, B: 35, A: 0xff},
			TabColor:              color.RGBA{R: 35, G: 35, B: 35, A: 0xff},
		},
	}, fgeom.Point{X: 5, Y: 5}, ebiteninput.NewEbitenInput())

	is.ideObj.Draw(is.ideRenderer)

	return nil
}
