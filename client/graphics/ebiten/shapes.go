package ebiten

import (
	"image"
	"image/color"

	"github.com/golang/geo/r2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	emptyImage    = ebiten.NewImage(3, 3)
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
)

func init() {
	emptyImage.Fill(color.White)
}

func DrawRect(r r2.Rect, image *ebiten.Image) {
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
	vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)
	for i := range vs {
		vs[i].SrcX = 1
		vs[i].SrcY = 1
		vs[i].ColorR = 0x56 / float32(0xff)
		vs[i].ColorG = 0x56 / float32(0xff)
		vs[i].ColorB = 0x56 / float32(0xff)
	}
	image.DrawTriangles(vs, is, emptySubImage, op)
}
