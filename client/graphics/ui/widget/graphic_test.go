package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestGraphic_PreferredSize(t *testing.T) {
	i := newImageEmptySize(47, 11, t)
	g := newGraphic(t, GraphicOpts.Image(i))
	w, h := g.PreferredSize()
	assert.Equal(t, i.Bounds().Dx(), w)
	assert.Equal(t, i.Bounds().Dy(), h)
}

func newGraphic(t *testing.T, opts ...GraphicOpt) *Graphic {
	t.Helper()

	g := NewGraphic(opts...)
	event.ExecuteDeferred()
	render(g, t)
	return g
}
