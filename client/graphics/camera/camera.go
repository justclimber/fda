package camera

import (
	"github.com/golang/geo/r2"

	"github.com/justclimber/fda/common/fgeom"
)

type Camera struct {
	viewRect r2.Rect
}

func NewCamera(viewRect r2.Rect) *Camera {
	return &Camera{viewRect: viewRect}
}

func (c *Camera) ViewRect() r2.Rect {
	return c.viewRect
}

func (c *Camera) Move(p r2.Point) {
	c.viewRect = fgeom.MoveRect(c.viewRect, p)
}
