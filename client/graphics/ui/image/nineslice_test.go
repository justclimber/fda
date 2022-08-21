package image

import (
	"image/color"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/stretchr/testify/assert"
)

func TestNineSlice_MinSize(t *testing.T) {
	n := NewNineSlice(newImageEmptySize(20, 20, t), [3]int{3, 10, 7}, [3]int{2, 16, 2})
	w, h := n.MinSize()
	assert.Equal(t, 10, w)
	assert.Equal(t, 4, h)

	n = NewNineSliceColor(color.White)
	w, h = n.MinSize()
	assert.Equal(t, 0, w)
	assert.Equal(t, 0, h)

	n = NewNineSliceColor(color.Transparent)
	w, h = n.MinSize()
	assert.Equal(t, 0, w)
	assert.Equal(t, 0, h)
}

func newImageEmptySize(width int, height int, t *testing.T) *ebiten.Image {
	t.Helper()
	return ebiten.NewImage(width, height)
}
