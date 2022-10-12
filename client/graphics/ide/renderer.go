package ide

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	ebiten2 "github.com/justclimber/fda/client/graphics/ebiten"
	"github.com/justclimber/fda/client/ide/ast"
	"github.com/justclimber/fda/common/fgeom"
	"github.com/justclimber/fda/common/lang/executor/object"
)

const (
	LineDistanceNormal = 1.5
)

type Options struct {
	IndentWidth        int
	Face               font.Face
	LineDistanceFactor float64
	TypeColorMap       map[ast.TextType]color.Color
	DefaultColor       color.Color
	Text               PredefinedText
	TabOptions         TabOptions
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
	ascent     float64
}

type TabOptions struct {
	HeaderPadding         int
	BodyPadding           int
	Size                  fgeom.Point
	HeaderBackgroundColor color.Color
	TabColor              color.Color
	BodyBackgroundColor   color.Color
}

func NewRenderer(opts Options, topLeft fgeom.Point) *Renderer {
	m := measureFont(opts.Face, opts.LineDistanceFactor)
	tabBodyX := topLeft.X
	tabBodyY := topLeft.Y + 2*float64(opts.TabOptions.HeaderPadding) + m.lineHeight
	return &Renderer{
		opts:             opts,
		textMeasurements: m,
		imageOptions:     &ebiten.DrawImageOptions{},
		tabBodyX:         tabBodyX,
		tabBodyY:         tabBodyY,
		topLeft:          topLeft,
	}
}

type Renderer struct {
	opts             Options
	image            *ebiten.Image
	imageOptions     *ebiten.DrawImageOptions
	textMeasurements textMeasurements
	currIndent       int
	offset           int
	lineNumber       int
	cursorX          float64
	cursorY          float64
	tabBodyX         float64
	tabBodyY         float64
	topLeft          fgeom.Point
}

func (r *Renderer) Draw(image *ebiten.Image) {
	if r.image == nil {
		r.image = image
	}
	r.cursorX = r.tabBodyX + float64(r.opts.TabOptions.BodyPadding)
	r.cursorY = r.tabBodyY + float64(r.opts.TabOptions.BodyPadding) + r.textMeasurements.ascent
}

func (r *Renderer) DrawTab(name string) {
	tabHeight := float64(r.opts.TabOptions.HeaderPadding)*2 + r.textMeasurements.lineHeight
	tabHeaderRect := fgeom.RectFromPointAndSize(r.topLeft, fgeom.Point{
		X: r.opts.TabOptions.Size.X,
		Y: tabHeight,
	})
	tabBodyRect := fgeom.RectFromPointAndSize(r.topLeft.Add(fgeom.Point{Y: tabHeight}), r.opts.TabOptions.Size)
	ebiten2.DrawRect(tabHeaderRect, r.image, r.opts.TabOptions.HeaderBackgroundColor)
	ebiten2.DrawRect(tabBodyRect, r.image, r.opts.TabOptions.BodyBackgroundColor)

	x := r.topLeft.X + float64(r.opts.TabOptions.HeaderPadding)
	y := r.topLeft.Y + r.textMeasurements.ascent*r.opts.LineDistanceFactor + float64(r.opts.TabOptions.HeaderPadding)
	r.imageOptions.GeoM.Reset()
	r.imageOptions.GeoM.Translate(x, y)
	r.imageOptions.ColorM.Reset()
	r.imageOptions.ColorM.ScaleWithColor(r.getColorForType(ast.TypeIdentifier))
	text.DrawWithOptions(r.image, name, r.opts.Face, r.imageOptions)
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
	r.cursorX = r.tabBodyX + float64(r.currIndent*r.opts.IndentWidth)*r.textMeasurements.width + float64(r.opts.TabOptions.BodyPadding)
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
		ascent:     fixedIntToFloat64(m.Ascent),
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
