package visualize

import (
	"bytes"
	"image"
	"image/color"
	"time"

	"github.com/fogleman/gg"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/zgiber/visualize/fonts"
	"golang.org/x/image/font"
)

type TimeChartData struct {
	Name   string
	Values []TimeChartDatapoint
	Color  color.RGBA
}

type TimeChartDatapoint struct {
	Time  time.Time
	Value float64
}

type TimeChartOptions struct {
	Width     int
	Height    int
	Values    []TimeChartData
	Title     string
	TextColor color.RGBA
	FontFace  font.Face
}

func TimeChart(opts TimeChartOptions) ([]byte, error) {
	if opts.FontFace == nil {
		fontSize := float64(opts.Width) / 6
		opts.FontFace = fonts.FontFace(fonts.JetBrainsMonoMedium, fontSize)
	}

	dc := gg.NewContext(opts.Width, opts.Width)
	drawTimeChart(dc, opts)

	img := image.NewRGBA(image.Rect(0, 0, opts.Width, opts.Width))
	dc.DrawImage(img, 0, 0)

	// draw the textual elements

	b := &bytes.Buffer{}
	err := dc.EncodePNG(b)
	return b.Bytes(), err
}

func drawTimeChart(dc *gg.Context, opts TimeChartOptions) {}

func drawTimeChartGrid(gc *draw2dimg.GraphicContext, dest *image.RGBA, opts TimeChartOptions) {}

func drawTimeChartLegend(gc *draw2dimg.GraphicContext, dest *image.RGBA, opts TimeChartOptions) {}
