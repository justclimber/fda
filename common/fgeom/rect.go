package fgeom

import (
	"image"

	"github.com/golang/geo/r1"
	"github.com/golang/geo/r2"
)

func MoveRect(r r2.Rect, p r2.Point) r2.Rect {
	return r2.Rect{
		X: r1.Interval{
			Lo: r.X.Lo + p.X,
			Hi: r.X.Hi + p.X,
		},
		Y: r1.Interval{
			Lo: r.Y.Lo + p.Y,
			Hi: r.Y.Hi + p.Y,
		},
	}
}

func RectFromPointAndSize(p r2.Point, size r2.Point) r2.Rect {
	return r2.Rect{
		X: r1.Interval{
			Lo: p.X,
			Hi: p.X + size.X,
		},
		Y: r1.Interval{
			Lo: p.Y,
			Hi: p.Y + size.Y,
		},
	}
}

func R2RectToImageRect(r2rect r2.Rect) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: int(r2rect.Lo().X),
			Y: int(r2rect.Lo().Y),
		},
		Max: image.Point{
			X: int(r2rect.Hi().X),
			Y: int(r2rect.Hi().Y),
		},
	}
}
