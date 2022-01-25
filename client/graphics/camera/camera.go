package camera

import (
	"github.com/justclimber/fda/common/fgeom"
)

const lowestScaleFactor = 0.3
const highestScaleFactor = 2

type Camera struct {
	viewRect    fgeom.Rect
	scaleFactor float64
}

func NewCamera(viewRect fgeom.Rect) *Camera {
	return &Camera{
		viewRect:    viewRect,
		scaleFactor: 1,
	}
}

func (c *Camera) ViewRect() fgeom.Rect {
	if c.scaleFactor != 1 {
		return c.viewRect.ScaleFromCenter(c.scaleFactor)
	}
	return c.viewRect
}

func (c *Camera) Move(p fgeom.Point) {
	if p == fgeom.EmptyPoint {
		return
	}
	c.viewRect = c.viewRect.Move(p)
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
