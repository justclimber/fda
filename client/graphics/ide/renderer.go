package ide

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/justclimber/fda/client/ide/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

const (
	LineDistanceNormal = 1.5
)

type RenderOptions struct {
	IndentWidth        int
	Face               font.Face
	LineDistanceFactor float64
	TypeColorMap       map[ast.TextType]color.Color
	DefaultColor       color.Color
	Text               PredefinedText
}

type PredefinedText struct {
	ArgDelimiter string
	Assignment   string
	Package      string
	Function     string
}

type textMeasurements struct {
	lineHeight float64
	width      float64
}

func NewRenderer(opts RenderOptions, initialCursorX float64, initialCursorY float64) *Renderer {
	return &Renderer{
		opts:             opts,
		textMeasurements: measureFont(opts.Face, opts.LineDistanceFactor),
		imageOptions:     &ebiten.DrawImageOptions{},
		initialCursorX:   initialCursorX,
		initialCursorY:   initialCursorY,
	}
}

type Renderer struct {
	opts             RenderOptions
	image            *ebiten.Image
	imageOptions     *ebiten.DrawImageOptions
	textMeasurements textMeasurements
	currIndent       int
	offset           int
	lineNumber       int
	cursorX          float64
	cursorY          float64
	initialCursorX   float64
	initialCursorY   float64
}

func (r *Renderer) Draw(image *ebiten.Image) {
	if r.image == nil {
		r.image = image
	}
	r.cursorX = r.initialCursorX
	r.cursorY = r.initialCursorY
}

func (r *Renderer) DrawTab(name string) {
	r.DrawText(name, ast.TypeIdentifier)
	r.NewLine() // fixme: this is temp
	r.NewLine()
}

func (r *Renderer) DrawAssignment() {
	r.DrawText(r.opts.Text.Assignment, ast.TypeSystemSymbols)
}

func (r *Renderer) DrawPackageHeader(name string) {
	r.DrawText(r.opts.Text.Package, ast.TypeKeywords)
	r.DrawText(name, ast.TypeIdentifier)
	r.NewLine()
	r.NewLine()
}

func (r *Renderer) DrawFuncHeader(definition *object.FunctionDefinition) {
	r.DrawText(definition.Name, ast.TypeIdentifier)
	r.DrawAssignment()
	r.DrawText(r.opts.Text.Function, ast.TypeKeywords)
	r.DrawText("()", ast.TypeSystemSymbols)
	r.DrawText(" {", ast.TypeSystemSymbols)
	// todo: input args and returns
	r.IndentIncrease()
	r.NewLine()
}

func (r *Renderer) DrawFuncBottom() {
	r.IndentDecrease()
	r.NewLine()
	r.DrawText("}", ast.TypeSystemSymbols)
}

func (r *Renderer) DrawArgDelimiter() {
	r.DrawText(r.opts.Text.ArgDelimiter, ast.TypeSystemSymbols)
}

func (r *Renderer) NewLine() {
	r.cursorX = r.initialCursorX + float64(r.currIndent*r.opts.IndentWidth)*r.textMeasurements.width
	r.cursorY = r.cursorY + r.textMeasurements.lineHeight
}

func (r *Renderer) IndentIncrease() {
	r.currIndent += 1
}

func (r *Renderer) IndentDecrease() {
	r.currIndent -= 1
}

func (r *Renderer) DrawText(str string, t ast.TextType) {
	r.imageOptions.GeoM.Reset()
	r.imageOptions.GeoM.Translate(r.cursorX, r.cursorY)
	r.imageOptions.ColorM.Reset()
	r.imageOptions.ColorM.ScaleWithColor(r.getColorForType(t))
	text.DrawWithOptions(r.image, str, r.opts.Face, r.imageOptions)
	r.AdvanceCursor(len(str))
}

func (r *Renderer) AdvanceCursor(num int) {
	r.cursorX = r.cursorX + float64(num)*r.textMeasurements.width
}

func measureFont(f font.Face, lineDistanceFactor float64) textMeasurements {
	m := f.Metrics()
	a, _ := f.GlyphAdvance('A')
	return textMeasurements{
		lineHeight: fixedIntToFloat64(m.Height) * lineDistanceFactor,
		width:      fixedIntToFloat64(a + f.Kern('A', 'A')),
	}
}

func fixedIntToFloat64(i fixed.Int26_6) float64 {
	return float64(i) / (1 << 6)
}

func (r *Renderer) getColorForType(t ast.TextType) color.Color {
	c, ok := r.opts.TypeColorMap[t]
	if ok {
		return c
	}
	return r.opts.DefaultColor
}
