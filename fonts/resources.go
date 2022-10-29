package fonts

import (
	"embed"
	"io"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

var (
	JetBrainsMonoRegular         FontName = "JetBrainsMono-Regular.ttf"
	JetBrainsMonoMedium          FontName = "JetBrainsMono-Medium.ttf"
	JetBrainsMonoMediumItalic    FontName = "JetBrainsMono-Medium-Italic.ttf"
	JetBrainsMonoItalic          FontName = "JetBrainsMono-Italic.ttf"
	JetBrainsMonoExtraBold       FontName = "JetBrainsMono-ExtraBold.ttf"
	JetBrainsMonoExtraBoldItalic FontName = "JetBrainsMono-ExtraBold-Italic.ttf"
	JetBrainsMonoBold            FontName = "JetBrainsMono-Bold.ttf"
	JetBrainsMonoBoldItalic      FontName = "JetBrainsMono-Bold-Italic.ttf"
)

type FontName string

//go:embed *
var Fonts embed.FS

func FontFace(name FontName, points float64) font.Face {
	f, err := Fonts.Open(string(name))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fontBytes, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	parsedFont, err := truetype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(parsedFont, &truetype.Options{
		Size: points,
	})
	return face
}
