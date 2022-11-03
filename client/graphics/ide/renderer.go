package ide

import (
	"fmt"
	"image/color"
	"reflect"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	ebiten2 "github.com/justclimber/fda/client/graphics/ebiten"
	"github.com/justclimber/fda/client/graphics/input"
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
	IfStart      string
	IfElse       string
	Assignment   string
	Package      string
	Function     string
}

type textMeasurements struct {
	lineHeight float64 // letter max height * LineDistanceFactor
	width      float64 // width of one letter space
	ascent     float64 // height of Uppercase letter for determining Y offset of letter drawing
}

type TabOptions struct {
	HeaderSpacing         float64
	HeaderPadding         float64
	BodyPadding           float64
	Size                  fgeom.Point
	HeaderBackgroundColor color.Color
	TabColor              color.Color
	BodyBackgroundColor   color.Color
}

func NewRenderer(opts Options, topLeft fgeom.Point, input input.Input) *Renderer {
	m := measureFont(opts.Face, opts.LineDistanceFactor)
	tabBodyX := topLeft.X
	tabBodyY := topLeft.Y + 2*float64(opts.TabOptions.HeaderPadding) + m.lineHeight
	current := &IndexNode{
		slug:      "package",
		yInterval: fgeom.Interval[int]{Lo: 0, Hi: 0},
		xInterval: fgeom.Interval[int]{Lo: 0, Hi: 0},
	}
	r := &Renderer{
		opts:             opts,
		textMeasurements: m,
		image:            ebiten.NewImage(600, 600), // todo get from options
		imageOptions:     &ebiten.DrawImageOptions{},
		tabBodyX:         tabBodyX,
		tabBodyY:         tabBodyY,
		topLeft:          topLeft,
		indexCurrent:     current,
		indexRoot:        current,
		indexActive:      current,
		userInput:        input,
	}
	r.cursorX = r.tabBodyX + r.opts.TabOptions.BodyPadding
	r.cursorY = r.tabBodyY + r.opts.TabOptions.BodyPadding + r.textMeasurements.ascent
	return r
}

type Renderer struct {
	opts             Options
	image            *ebiten.Image
	imageOptions     *ebiten.DrawImageOptions
	userInput        input.Input
	textMeasurements textMeasurements
	currIndent       int
	offset           int
	lineNumber       int
	cursorX          float64
	cursorY          float64
	tabBodyX         float64
	tabBodyY         float64
	topLeft          fgeom.Point

	indexRoot        *IndexNode
	indexCurrent     *IndexNode
	indexActive      *IndexNode
	indexHasSiblings bool
}

func (r *Renderer) Draw(image *ebiten.Image) {
	image.DrawImage(r.image, &ebiten.DrawImageOptions{})
}

func (r *Renderer) DrawHeaderTab() {
	tabHeight := float64(r.opts.TabOptions.HeaderPadding)*2 + r.textMeasurements.lineHeight
	tabHeaderRect := fgeom.RectFromPointAndSize(r.topLeft, fgeom.Point{
		X: r.opts.TabOptions.Size.X,
		Y: tabHeight,
	})
	ebiten2.DrawRect(tabHeaderRect, r.image, r.opts.TabOptions.HeaderBackgroundColor)
}

func (r *Renderer) DrawInactiveTab(name string, offset float64) float64 {
	return r.drawTab(name, offset, false)
}

func (r *Renderer) DrawActiveTab(name string, offset float64) float64 {
	return r.drawTab(name, offset, true)
}

func (r *Renderer) drawTab(name string, offset float64, isActive bool) float64 {
	tabWidth := float64(len(name))*r.textMeasurements.width + 2*r.opts.TabOptions.HeaderPadding
	tabHeight := 0.
	if isActive {
		tabHeight = r.fullTabHeight()
	} else {
		tabHeight = r.textMeasurements.lineHeight + r.opts.TabOptions.HeaderPadding
	}
	tabSize := fgeom.Point{
		X: tabWidth,
		Y: tabHeight,
	}
	tabTopLeft := r.topLeft.Add(fgeom.Point{X: offset})
	tabHeaderRect := fgeom.RectFromPointAndSize(tabTopLeft, tabSize)
	ebiten2.DrawRect(tabHeaderRect, r.image, r.opts.TabOptions.TabColor)

	x := tabTopLeft.X + r.opts.TabOptions.HeaderPadding
	y := tabTopLeft.Y + r.textMeasurements.ascent*r.opts.LineDistanceFactor + r.opts.TabOptions.HeaderPadding
	r.imageOptions.GeoM.Reset()
	r.imageOptions.GeoM.Translate(x, y)
	r.imageOptions.ColorM.Reset()
	r.imageOptions.ColorM.ScaleWithColor(r.getColorForType(ast.TypeIdentifier))
	text.DrawWithOptions(r.image, name, r.opts.Face, r.imageOptions)
	return offset + tabWidth + r.opts.TabOptions.HeaderSpacing
}

func (r *Renderer) fullTabHeight() float64 {
	return float64(r.opts.TabOptions.HeaderPadding)*2 + r.textMeasurements.lineHeight
}

func (r *Renderer) DrawTabBody() {
	tabBodyRect := fgeom.RectFromPointAndSize(r.topLeft.Add(fgeom.Point{Y: r.fullTabHeight()}), r.opts.TabOptions.Size)
	ebiten2.DrawRect(tabBodyRect, r.image, r.opts.TabOptions.BodyBackgroundColor)
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

func (r *Renderer) DrawIfStart() {
	r.DrawText(r.opts.Text.IfStart, ast.TypeKeywords)
}

func (r *Renderer) DrawIfElse() {
	r.IndentDecrease()
	r.NewLine()
	r.DrawText("} ", ast.TypeSystemSymbols)
	r.DrawText(r.opts.Text.IfElse, ast.TypeKeywords)
	r.DrawText(" {", ast.TypeSystemSymbols)
	r.IndentIncrease()
	r.NewLine()
}

func (r *Renderer) DrawIfMid() {
	r.DrawText(" {", ast.TypeSystemSymbols)
	r.IndentIncrease()
	r.NewLine()
}

func (r *Renderer) DrawIfEnd() {
	r.IndentDecrease()
	r.NewLine()
	r.DrawText("}", ast.TypeSystemSymbols)
}

func (r *Renderer) NewLine() {
	r.cursorX = r.tabBodyX + float64(r.currIndent*r.opts.IndentWidth)*r.textMeasurements.width + r.opts.TabOptions.BodyPadding
	r.cursorY = r.cursorY + r.textMeasurements.lineHeight
	r.lineNumber++
	r.offset = r.currIndent * r.opts.IndentWidth
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
	r.offset = r.offset + num
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

func (r *Renderer) HighlightActiveNode(screen *ebiten.Image) {
	rect := fgeom.Rect{X: fgeom.Interval[float64]{
		Lo: float64(r.indexActive.xInterval.Lo)*r.textMeasurements.width + r.tabBodyX + r.opts.TabOptions.BodyPadding - 3,
		Hi: float64(r.indexActive.xInterval.Hi)*r.textMeasurements.width + r.tabBodyX + r.opts.TabOptions.BodyPadding + 4,
	}, Y: fgeom.Interval[float64]{
		Lo: float64(r.indexActive.yInterval.Lo)*r.textMeasurements.lineHeight + r.tabBodyY + r.opts.TabOptions.BodyPadding - 5,
		Hi: float64(r.indexActive.yInterval.Hi+1)*r.textMeasurements.lineHeight + r.tabBodyY + r.opts.TabOptions.BodyPadding - 2,
	}}
	// todo color from opts
	ebiten2.DrawRect(rect, screen, color.RGBA{R: 0x66, G: 0x99, B: 0xcc, A: 0x25})
}

type IndexNode struct {
	node       ast.DrawableNode
	slug       string
	yInterval  fgeom.Interval[int]
	xInterval  fgeom.Interval[int]
	firstChild *IndexNode
	next       *IndexNode
	prev       *IndexNode
	parent     *IndexNode
}

func (r *Renderer) StartContainerNode() {
	r.indexHasSiblings = false
}

func (r *Renderer) EndContainerNode() {
	r.indexHasSiblings = true
	r.indexCurrent = r.indexCurrent.parent
	r.goNextToEndForCurrent()
}

func (r *Renderer) goNextToEndForCurrent() {
	if r.indexCurrent.next != nil {
		r.indexCurrent = r.indexCurrent.next
		r.goNextToEndForCurrent()
	}
}

func (r *Renderer) StartSiblingNode(n ast.DrawableNode, slug string) func() {
	newNode := &IndexNode{
		node:      n,
		slug:      slug,
		yInterval: fgeom.Interval[int]{Lo: r.lineNumber},
		xInterval: fgeom.Interval[int]{Lo: r.offset},
	}

	if r.indexHasSiblings {
		r.indexCurrent.next = newNode
		newNode.parent = r.indexCurrent.parent
		newNode.prev = r.indexCurrent
	} else {
		r.indexCurrent.firstChild = newNode
		newNode.parent = r.indexCurrent
		r.indexHasSiblings = true
	}

	r.indexCurrent = newNode
	return func() {
		newNode.yInterval.Hi = r.lineNumber
		updateXHiIfGraterRecursive(r.offset, newNode)
	}
}

func updateXHiIfGraterRecursive(x int, n *IndexNode) {
	if x > n.xInterval.Hi {
		n.xInterval.Hi = x
	}
	if n.parent != nil {
		updateXHiIfGraterRecursive(x, n.parent)
	}
}

func (r *Renderer) HandleUserInput() {
	if r.indexActive == nil {
		r.indexActive = r.indexRoot
	}
	r.nodeNavigationByControls(r.userInput.WhichControlArrowsPressed())
}

func (r *Renderer) nodeNavigationByControls(ctrl input.ControlArrow) {
	switch ctrl {
	case input.ControlArrowDown:
		if r.indexActive.firstChild != nil {
			r.indexActive = r.indexActive.firstChild
		}
	case input.ControlArrowRight:
		if r.indexActive.next != nil {
			r.indexActive = r.indexActive.next
		} else if r.indexActive.firstChild != nil {
			r.nodeNavigationByControls(input.ControlArrowDown)
		}
	case input.ControlArrowLeft:
		if r.indexActive.prev != nil {
			r.indexActive = r.indexActive.prev
		} else if r.indexActive.parent != nil {
			r.nodeNavigationByControls(input.ControlArrowUp)
		}
	case input.ControlArrowUp:
		if r.indexActive.parent != nil {
			r.indexActive = r.indexActive.parent
		}
	}
}

func (r *Renderer) DebugPrintIndex(screen *ebiten.Image) {
	str := r.sprintfNode(r.indexRoot, 3)
	ebitenutil.DebugPrintAt(screen, str, 500, 10)
}

func (r *Renderer) sprintfNode(n *IndexNode, indent int) string {
	var children, siblings, indentStr string
	if n.firstChild != nil {
		children = r.sprintfNode(n.firstChild, indent+1)
	}
	if n.next != nil {
		siblings = r.sprintfNode(n.next, indent)
	}
	result := n.slug
	if r.indexActive == n {
		result = "[x] " + result
		indentStr = strings.Repeat(" ", (indent-1)*3)
	} else {
		indentStr = strings.Repeat(" ", indent*3)
	}
	if n.node != nil {
		result = result + ": " + reflect.TypeOf(n.node).String()
	}
	result = indentStr + result
	if len(children) != 0 {
		result = fmt.Sprintf("%s\n%s", result, children)
	}
	if len(siblings) != 0 {
		result = fmt.Sprintf("%s\n%s", result, siblings)
	}
	return result
}
