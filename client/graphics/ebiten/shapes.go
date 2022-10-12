package ebiten

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/justclimber/fda/common/fgeom"
)

var (
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	emptyImage.Fill(color.White)
}

func DrawRect(r fgeom.Rect, image *ebiten.Image, clr color.Color) {
	var path vector.Path
	ps := r.Vertices()
	first := true
	for _, p := range ps {
		if first {
			path.MoveTo(float32(p.X), float32(p.Y))
			first = false
			continue
		}
		path.LineTo(float32(p.X), float32(p.Y))
	}
	op := &ebiten.DrawTrianglesOptions{
		FillRule: ebiten.EvenOdd,
	}
	op.ColorM.ScaleWithColor(clr)
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
	}
	image.DrawTriangles(vs, is, emptySubImage, op)
}
