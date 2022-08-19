package font

import (
	"io/fs"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func LoadFont(fileName string, size float64, fileSystem fs.ReadFileFS) (font.Face, error) {
	fontData, err := fileSystem.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}
