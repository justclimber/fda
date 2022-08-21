package widget

import (
	img "image"

	"github.com/justclimber/fda/client/graphics/ui/image"
	"github.com/justclimber/fda/client/graphics/ui/input"

	"github.com/hajimehoshi/ebiten/v2"
)

type Container struct {
	DebugLabel          string
	BackgroundImage     *image.NineSlice
	AutoDisableChildren bool

	widgetOpts  []Opt
	layout      Layouter
	layoutDirty bool

	init     *MultiOnce
	widget   *Widget
	children []PreferredSizeLocateableWidget
}

type ContainerOpt func(c *Container)

type RemoveChildFunc func()

type ContainerOptions struct {
}

var ContainerOpts ContainerOptions

type PreferredSizeLocateableWidget interface {
	HasWidget
	PreferredSizer
	Locateable
}

func NewContainer(debugLabel string, opts ...ContainerOpt) *Container {
	c := &Container{
		DebugLabel: debugLabel,
		init:       &MultiOnce{},
	}

	c.init.Append(c.createWidget)

	for _, o := range opts {
		o(c)
	}

	return c
}

func (o ContainerOptions) WidgetOpts(opts ...Opt) ContainerOpt {
	return func(c *Container) {
		c.widgetOpts = append(c.widgetOpts, opts...)
	}
}

func (o ContainerOptions) BackgroundImage(i *image.NineSlice) ContainerOpt {
	return func(c *Container) {
		c.BackgroundImage = i
	}
}

func (o ContainerOptions) AutoDisableChildren() ContainerOpt {
	return func(c *Container) {
		c.AutoDisableChildren = true
	}
}

func (o ContainerOptions) Layout(layout Layouter) ContainerOpt {
	return func(c *Container) {
		c.layout = layout
	}
}

func (c *Container) AddChild(child PreferredSizeLocateableWidget) RemoveChildFunc {
	c.init.Do()

	if child == nil {
		panic("cannot add nil child")
	}

	c.children = append(c.children, child)

	child.GetWidget().parent = c

	c.RequestRelayout()

	return func() {
		c.removeChild(child)
	}
}

func (c *Container) removeChild(child PreferredSizeLocateableWidget) {
	index := -1
	for i, ch := range c.children {
		if ch == child {
			index = i
			break
		}
	}

	if index < 0 {
		return
	}

	c.children = append(c.children[:index], c.children[index+1:]...)

	child.GetWidget().parent = nil

	c.RequestRelayout()
}

func (c *Container) RequestRelayout() {
	c.init.Do()

	c.layoutDirty = true

	for _, ch := range c.children {
		if r, ok := ch.(Relayoutable); ok {
			r.RequestRelayout()
		}
	}
}

func (c *Container) GetWidget() *Widget {
	c.init.Do()
	return c.widget
}

func (c *Container) PreferredSize() (int, int) {
	c.init.Do()

	if c.layout == nil {
		return 50, 50
	}

	return c.layout.PreferredSize(c.children)
}

func (c *Container) SetLocation(rect img.Rectangle) {
	c.init.Do()
	c.widget.Rect = rect
}

func (c *Container) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	c.init.Do()

	if c.AutoDisableChildren {
		for _, ch := range c.children {
			ch.GetWidget().Disabled = c.widget.Disabled
		}
	}

	c.widget.Render(screen, def, debugMode)

	c.doLayout()

	c.draw(screen)

	for _, ch := range c.children {
		if cr, ok := ch.(Renderer); ok {
			cr.Render(screen, def, debugMode)
		}
	}
}

func (c *Container) RenderInputLayerDebug(screen *ebiten.Image) {
	c.widget.RenderInputLayerDebug(screen)
	for _, ch := range c.children {
		if w, ok := ch.(HasWidget); ok {
			w.GetWidget().RenderInputLayerDebug(screen)
		}
	}
}

func (c *Container) RenderWidgetSizeDebug(screen *ebiten.Image) {
	c.widget.RenderWidgetRectDebug(screen)
	for _, ch := range c.children {
		if cd, ok := ch.(DebugDrawer); ok {
			cd.RenderWidgetSizeDebug(screen)
		}
		if hw, ok := ch.(HasWidget); ok {
			hw.GetWidget().RenderWidgetRectDebug(screen)
		}
	}
}

func (c *Container) doLayout() {
	if c.layout != nil && c.layoutDirty {
		c.layout.Layout(c.children, c.widget.Rect)
		c.layoutDirty = false
	}
}

func (c *Container) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	c.init.Do()

	for _, ch := range c.children {
		if il, ok := ch.(input.Layerer); ok {
			il.SetupInputLayer(def)
		}
	}
}

func (c *Container) draw(screen *ebiten.Image) {
	if c.BackgroundImage != nil {
		c.BackgroundImage.Draw(screen, c.widget.Rect.Dx(), c.widget.Rect.Dy(), c.widget.drawImageOptions)
	}
}

func (c *Container) createWidget() {
	c.widget = NewWidget(c.widgetOpts...)
	c.widgetOpts = nil
}

// WidgetAt implements WidgetLocator.
func (c *Container) WidgetAt(x int, y int) HasWidget {
	c.init.Do()

	p := img.Point{X: x, Y: y}

	if !p.In(c.GetWidget().Rect) {
		return nil
	}

	for _, ch := range c.children {
		if wl, ok := ch.(Locater); ok {
			w := wl.WidgetAt(x, y)
			if w != nil {
				return w
			}

			continue
		}

		if p.In(ch.GetWidget().Rect) {
			return ch
		}
	}

	return c
}
