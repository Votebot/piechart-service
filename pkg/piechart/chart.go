package piechart

import (
	"bytes"
	"github.com/getsentry/sentry-go"
	"github.com/wcharczuk/go-chart"
	"go.uber.org/zap"
)

type Config struct {
	Votes  []Part `json:"votes"`
	Title  string `json:"title"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Part struct {
	VoteCount int    `json:"vote_count"`
	Title     string `json:"title"`
}

func (c Config) GetValues() []chart.Value {
	var arr []chart.Value

	for _, vote := range c.Votes {
		arr = append(arr, chart.Value{Value: float64(vote.VoteCount), Label: vote.Title})
	}

	return arr
}

func addVotebotLogoRenderable(r chart.Renderer, canvasBox chart.Box, defaults chart.Style) {
	chart.Draw.Box(r, chart.Box{Top: canvasBox.Height() - 23, Left: canvasBox.Width() - 20, Bottom: canvasBox.Height() - 10, Right: canvasBox.Width() - 16}, chart.Style{FillColor: chart.ColorWhite})

	// Move to function
	r.SetFillColor(chart.ColorAlternateGray)
	r.MoveTo(canvasBox.Width()-20, canvasBox.Height()-12)
	r.LineTo(canvasBox.Width()-18, canvasBox.Height()-10)
	r.LineTo(canvasBox.Width()-20, canvasBox.Height()-10)
	r.LineTo(canvasBox.Width()-20, canvasBox.Height()-12)
	r.FillStroke()

	chart.Draw.Box(r, chart.Box{Top: canvasBox.Height() - 30, Left: canvasBox.Width() - 14, Bottom: canvasBox.Height() - 11, Right: canvasBox.Width() - 10}, chart.Style{FillColor: chart.ColorAlternateGray, StrokeColor: chart.ColorWhite, StrokeWidth: 1})
	chart.Draw.Box(r, chart.Box{Top: canvasBox.Height() - 26, Left: canvasBox.Width() - 8, Bottom: canvasBox.Height() - 10, Right: canvasBox.Width() - 4}, chart.Style{FillColor: chart.ColorWhite})
}

/*drawTriangle(
	r,
	Point{
		x: canvasBox.Width() - 20,
		y: canvasBox.Height() - 19,
	}, Point{
		x: canvasBox.Width() - 16,
		y: canvasBox.Height() - 10,
	}, Point{
		x: canvasBox.Width() - 20,
		y: canvasBox.Height() - 10,
	},
	chart.ColorGreen,
	canvasBox,
)*/

/*type Point struct {
	x int
	y int
}

func drawTriangle(r chart.Renderer, p1, p2, p3 Point, color drawing.Color, canvasBox chart.Box) {
	r.SetFillColor(color)

	defer r.ResetStyle()

	r.MoveTo(p1.x, p1.y)
	r.LineTo(p2.x, p2.y)
	r.LineTo(p3.x, p3.x)
	r.LineTo(p1.x, p1.y)
	r.FillStroke()


}*/

func CreateChart(cfg Config) []byte {
	pie := chart.PieChart{
		Title:  cfg.Title,
		TitleStyle: chart.Style{Show: true},
		Width:  cfg.Width,
		Height: cfg.Height,
		Values: cfg.GetValues(),
		Elements: []chart.Renderable{
			addVotebotLogoRenderable,
		},
		Background: chart.Style{FillColor: chart.ColorAlternateGray, StrokeColor: chart.ColorAlternateGray},
		Canvas:     chart.Style{FillColor: chart.ColorAlternateGray},
	}

	imageBuffer := bytes.NewBuffer([]byte{})
	err := pie.Render(chart.PNG, imageBuffer)
	if err != nil {
		zap.L().Error("could not create pie chart", zap.Error(err))
		sentry.CaptureException(err)
		return nil
	}

	/*	var imageBuffer []byte
		_, err = collector.Write(imageBuffer)

		// image, err := collector.Image()
		if err != nil {
			zap.L().Error("could not create pie chart image", zap.Error(err))
			sentry.CaptureException(err)
			return nil
		}*/

	return imageBuffer.Bytes()
}
