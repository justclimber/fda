package widget

import (
	"image"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/justclimber/fda/client/graphics/ui/input"
)

type controlMock struct {
	mock.Mock
}

func TestContainer_Render(t *testing.T) {
	w := NewWidget()
	m := controlMock{}
	m.On("GetWidget").Maybe().Return(w)
	m.On("PreferredSize").Maybe().Return(50, 50)
	m.On("SetLocation", mock.Anything).Maybe()
	m.On("Render", mock.Anything, mock.Anything, mock.Anything)

	c := newContainer(t,
		ContainerOpts.Layout(newRowLayout(t)))
	c.AddChild(&m)

	render(c, t)

	m.AssertExpectations(t)
}

func TestContainer_Render_AutoDisableChildren(t *testing.T) {
	w := NewWidget()
	m := controlMock{}
	m.On("GetWidget").Maybe().Return(w)
	m.On("PreferredSize").Maybe().Return(50, 50)
	m.On("SetLocation", mock.Anything).Maybe()
	m.On("Render", mock.Anything, mock.Anything, mock.Anything).Maybe()

	c := newContainer(t,
		ContainerOpts.AutoDisableChildren(),
		ContainerOpts.Layout(newRowLayout(t)))
	c.AddChild(&m)

	c.widget.Disabled = true
	render(c, t)

	assert.True(t, w.Disabled)
}

func TestContainer_SetupInputLayer(t *testing.T) {
	def := func(s input.SetupInputLayerFunc) {
		// nothing to do
	}

	w := NewWidget()
	m := controlMock{}
	m.On("GetWidget").Maybe().Return(w)
	m.On("SetupInputLayer", mock.AnythingOfType("input.DeferredSetupInputLayerFunc"))

	c := newContainer(t,
		ContainerOpts.Layout(newRowLayout(t)))
	c.AddChild(&m)

	c.SetupInputLayer(def)

	m.AssertExpectations(t)
}

func (c *controlMock) GetWidget() *Widget {
	args := c.Called()
	return args.Get(0).(*Widget)
}

func (c *controlMock) PreferredSize() (int, int) {
	c.Called()
	return -1, -1
}

func (c *controlMock) SetLocation(rect image.Rectangle) {
	c.Called(rect)
}

func (c *controlMock) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	c.Called(screen, def, debugMode)
}

func (c *controlMock) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	c.Called(def)
}

func newContainer(t *testing.T, opts ...ContainerOpt) *Container {
	t.Helper()
	return NewContainer("test", opts...)
}
