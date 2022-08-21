package widget

import (
	"image"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInsetsSimple(t *testing.T) {
	inset := 15
	i := NewInsetsSimple(inset)

	assert.Equal(t, inset, i.Left)
	assert.Equal(t, inset, i.Right)
	assert.Equal(t, inset, i.Top)
	assert.Equal(t, inset, i.Bottom)
}

func TestInsets_Apply(t *testing.T) {
	i := Insets{
		Left:   10,
		Right:  20,
		Top:    30,
		Bottom: 40,
	}
	r := image.Rect(25, 35, 145, 155)

	assert.Equal(t, image.Rect(35, 65, 125, 115), i.Apply(r))
}

func TestInsets_Dx(t *testing.T) {
	i := Insets{
		Left:  10,
		Right: 20,
	}

	assert.Equal(t, 30, i.Dx())
}

func TestInsets_Dy(t *testing.T) {
	i := Insets{
		Top:    30,
		Bottom: 40,
	}

	assert.Equal(t, 70, i.Dy())
}
