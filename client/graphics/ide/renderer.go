package ide

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/justclimber/fda/client/ide/ast"
)

const (
	LineDistanceNormal = 1.5
)

type RenderOptions struct {
	ArgDelimiterStr    string
	AssignmentStr      string
	IndentWidth        int
	Face               font.Face
	LineDistanceFactor float64
	TypeColorMap       map[ast.TextType]color.Color
}

type textMeasurements struct {
	lineHeight float64
	width      float64
}

func NewRenderer(opts RenderOptions, initialCursorX float64, initialCursorY float64) *Renderer {
	return &Renderer{
		opts:             opts,
		textMeasurements: measureFont(opts.Face, opts.LineDistanceFactor),
		op:               &ebiten.DrawImageOptions{},
		initialCursorX:   initialCursorX,
		initialCursorY:   initialCursorY,
	}
}

type Renderer struct {
	opts             RenderOptions
	image            *ebiten.Image
	op               *ebiten.DrawImageOptions
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

func (r *Renderer) DrawAssignment() {
	r.DrawText(r.opts.AssignmentStr, ast.TypeSystemSymbols)
}

func (r *Renderer) DrawArgDelimiter() {
	r.DrawText(r.opts.ArgDelimiterStr, ast.TypeSystemSymbols)
}

func (r *Renderer) NewLine() {
	r.cursorX = r.initialCursorX
	r.cursorY = r.cursorY + r.textMeasurements.lineHeight
}

func (r *Renderer) IndentIncrease() {}

func (r *Renderer) IndentDecrease() {}

func (r *Renderer) DrawText(str string, t ast.TextType) {
	r.op.GeoM.Reset()
	r.op.GeoM.Translate(r.cursorX, r.cursorY)
	clr := r.opts.TypeColorMap[t]
	r.op.ColorM.Reset()
	r.op.ColorM.ScaleWithColor(clr)
	text.DrawWithOptions(r.image, str, r.opts.Face, r.op)
	r.Advance(len(str))
}

func (r *Renderer) Advance(num int) {
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
