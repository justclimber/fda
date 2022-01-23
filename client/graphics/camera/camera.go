package camera

import (
	"github.com/golang/geo/r2"

	"github.com/justclimber/fda/common/fgeom"
)

const lowestScaleFactor = 0.3
const highestScaleFactor = 2

type Camera struct {
	viewRect    r2.Rect
	scaleFactor float64
}

func NewCamera(viewRect r2.Rect) *Camera {
	return &Camera{
		viewRect:    viewRect,
		scaleFactor: 1,
	}
}

func (c *Camera) ViewRect() r2.Rect {
	if c.scaleFactor != 1 {
		return fgeom.RectScaleFromCenter(c.viewRect, c.scaleFactor)
	}
	return c.viewRect
}

func (c *Camera) Move(p r2.Point) {
	// todo: do nothing if p is empty
	c.viewRect = fgeom.MoveRect(c.viewRect, p)
}

func (c *Camera) Scale(scaleFactorChange float64) {
	if scaleFactorChange == 0 {
		return
	}
	if c.scaleFactor+scaleFactorChange > highestScaleFactor || c.scaleFactor+scaleFactorChange < lowestScaleFactor {
		return
	}
	c.scaleFactor = c.scaleFactor + scaleFactorChange
}

func (c *Camera) ScaleFactor() float64 {
	return c.scaleFactor
}
