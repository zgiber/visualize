package visualize

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/zgiber/visualize/fonts"
	"golang.org/x/image/font"
)

type DonutOptions struct {
	Width       int
	Value       float64
	MaxValue    float64
	InnerRadius float64
	OuterRadius float64
	Color       color.RGBA
	BgColor     color.RGBA
	TextColor   color.RGBA
	FontFace    font.Face
}

type donutSectionOptions struct {
	innerRadius float64
	outerRadius float64
	startAngle  float64
	endAngle    float64
	color       color.RGBA
}

func Donut(opts DonutOptions) ([]byte, error) {
	if opts.FontFace == nil {
		fontSize := float64(opts.Width) / 6
		opts.FontFace = fonts.FontFace(fonts.JetBrainsMonoMedium, fontSize)
	}

	if opts.InnerRadius == 0 {
		opts.InnerRadius = float64(opts.Width/2) * 0.75
	}

	if opts.OuterRadius == 0 {
		opts.OuterRadius = float64(opts.Width / 2)
	}

	img := image.NewRGBA(image.Rect(0, 0, opts.Width, opts.Width))
	gc := draw2dimg.NewGraphicContext(img)
	drawDonut(gc, img, opts)

	dc := gg.NewContext(opts.Width, opts.Width)
	dc.DrawImage(img, 0, 0)
	dc.SetFontFace(opts.FontFace)
	dc.SetColor(color.Black)

	text := fmt.Sprintf("%2.2f", opts.Value)
	tw, th := dc.MeasureString(text)
	fmt.Println(float64(opts.Width/2)+th/2, float64(opts.Width)/2-tw/2)

	dc.DrawString(text, float64(opts.Width)/2-(tw/2), float64(opts.Width/2)+(th/2))

	b := &bytes.Buffer{}
	err := dc.EncodePNG(b)
	return b.Bytes(), err
}

func drawDonut(gc *draw2dimg.GraphicContext, dest *image.RGBA, opts DonutOptions) {
	drawDounutSection(gc, dest, bgOptions(opts))
	drawDounutSection(gc, dest, donutSectionOptions{
		innerRadius: opts.InnerRadius,
		outerRadius: opts.OuterRadius,
		startAngle:  0,
		endAngle:    (opts.Value / opts.MaxValue) * 360,
		color:       opts.Color,
	})
}

func drawDounutSection(gc *draw2dimg.GraphicContext, dest *image.RGBA, opts donutSectionOptions) {
	centerX := float64(dest.Rect.Dx()) / 2
	centerY := float64(dest.Rect.Dy()) / 2
	outerRadius := opts.outerRadius
	innerRadius := opts.innerRadius

	angle := math.Abs(opts.endAngle - opts.startAngle)

	startX, startY := pointOnCircle(centerX, centerY, outerRadius, opts.startAngle)
	fmt.Println(startX, startY)
	gc.MoveTo(startX, startY)
	gc.ArcTo(centerX, centerY, outerRadius, outerRadius, d2r(270+opts.startAngle), d2r(angle))
	gc.LineTo(pointOnCircle(centerX, centerY, innerRadius, angle+opts.startAngle))
	gc.ArcTo(centerX, centerY, innerRadius, innerRadius, d2r(270+opts.startAngle+angle), d2r(angle*-1))
	gc.LineTo(startX, startY)
	gc.Close()

	gc.SetFillColor(opts.color)
	gc.Fill()
}

func d2r(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func pointOnCircle(cx, cy, r, angle float64) (float64, float64) {
	x := r * math.Sin(d2r(angle))
	y := r * math.Cos(d2r(angle))

	angle = angle - float64(int(angle/360)*360)

	if d2r(180) < angle && angle < d2r(360) {
		x = -1 * x
	}

	if d2r(270) < angle || angle < d2r(90) {
		y = -1 * y
	}

	return cx + x, cy + y
}

func bgOptions(opts DonutOptions) donutSectionOptions {
	return donutSectionOptions{
		innerRadius: opts.InnerRadius + opts.InnerRadius/1000,
		outerRadius: opts.OuterRadius - opts.OuterRadius/1000,
		startAngle:  0,
		endAngle:    360,
		color:       opts.BgColor,
	}
}
