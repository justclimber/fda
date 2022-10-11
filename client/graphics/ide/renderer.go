package ide

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	"github.com/justclimber/fda/client/ide/ast"
)

type RenderOptions struct {
	ArgDelimiterStr string
	AssignmentStr   string
	IndentWidth     int
	Face            font.Face
	TypeColorMap    map[ast.TextType]color.Color
}

type textMeasurements struct {
	lineHeight float64
	width      float64
}

func NewRenderer(opts RenderOptions) *Renderer {
	return &Renderer{
		opts:             opts,
		textMeasurements: measureFont(opts.Face),
		op:               &ebiten.DrawImageOptions{},
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
}

func (r *Renderer) SetImage(image *ebiten.Image) {
	if r.image == nil {
		r.image = image
	}
	r.cursorX = 50
	r.cursorY = 50
}

func (r *Renderer) DrawAssignment() {
	r.DrawText(r.opts.AssignmentStr, ast.TypeSystemSymbols)
}

func (r *Renderer) DrawArgDelimiter() {
	r.DrawText(r.opts.ArgDelimiterStr, ast.TypeSystemSymbols)
}

func (r *Renderer) NewLine() {}

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

func measureFont(f font.Face) textMeasurements {
	m := f.Metrics()
	a, _ := f.GlyphAdvance('A')
	return textMeasurements{
		lineHeight: fixedIntToFloat64(m.Height),
		width:      fixedIntToFloat64(a + f.Kern('A', 'A')),
	}
}

func fixedIntToFloat64(i fixed.Int26_6) float64 {
	return float64(i) / (1 << 6)
}
