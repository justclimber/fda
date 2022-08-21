package image

import (
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

func TestBufferedImage_Image(t *testing.T) {
	b := &BufferedImage{}
	b.Width, b.Height = 100, 100
	i := b.Image()
	w, h := i.Size()
	assert.Equal(t, 100, w)
	assert.Equal(t, 100, h)

	b.Width, b.Height = 150, 70
	i = b.Image()
	w, h = i.Size()
	assert.Equal(t, 150, w)
	assert.Equal(t, 70, h)
}

func TestMaskedRenderBuffer_Draw(t *testing.T) {
	b := NewMaskedRenderBuffer()
	screen := newImageEmptySize(100, 100, t)

	draw := false
	drawMask := false

	b.Draw(screen, func(buf *ebiten.Image) {
		draw = true
	}, func(buf *ebiten.Image) {
		drawMask = true
	})

	assert.True(t, draw)
	assert.True(t, drawMask)
}
