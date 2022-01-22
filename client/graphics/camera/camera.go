package camera

import (
	"github.com/golang/geo/r2"

	"github.com/justclimber/fda/common/fgeom"
)

var emptyPoint = r2.Point{}

type Camera struct {
	viewRect r2.Rect
}

func NewCamera(rect r2.Rect, tileSize int) *Camera {
	return &Camera{viewRect: rect.AddPoint(r2.Point{
		X: rect.X.Lo - float64(tileSize),
		Y: rect.Y.Lo - float64(tileSize),
	})}
}

func (c *Camera) Offset(p r2.Point) (r2.Point, bool) {
	if !c.viewRect.ContainsPoint(p) {
		return emptyPoint, true
	}
	return p.Sub(c.viewRect.Lo()), false
}

func (c *Camera) Move(p r2.Point) {
	c.viewRect = fgeom.MoveRect(c.viewRect, p)
}

func (c *Camera) Size() r2.Point {
	return c.viewRect.Size()
}
