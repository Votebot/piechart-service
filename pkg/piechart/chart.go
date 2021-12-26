package piechart

import (
	"bytes"
	"github.com/getsentry/sentry-go"
	"github.com/niggelgame/go-chart/v2"
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

// Icon Padding 10x10px to the bottom right corner
func addVotebotLogoRenderable(r chart.Renderer, canvasBox chart.Box, _ chart.Style) {
	paddingRight := 10
	paddingBottom := 10

	backgroundColor := chart.ColorAlternateGray
	highlightColor := chart.ColorWhite

	chart.Draw.Box(r, chart.Box{Top: canvasBox.Bottom - (13 + paddingBottom), Left: canvasBox.Right - (15 + paddingRight), Bottom: canvasBox.Bottom - paddingBottom, Right: canvasBox.Right - (11 + paddingRight)}, chart.Style{FillColor: highlightColor})

	// 2x2px corner cutout in first bar
	r.SetFillColor(backgroundColor)
	r.MoveTo(canvasBox.Right-(15+paddingRight), canvasBox.Bottom-(2+paddingBottom))
	r.LineTo(canvasBox.Right-(13+paddingRight), canvasBox.Bottom-paddingBottom)
	r.LineTo(canvasBox.Right-(15+paddingRight), canvasBox.Bottom-paddingBottom)
	r.LineTo(canvasBox.Right-(15+paddingRight), canvasBox.Bottom-(2+paddingBottom))
	r.FillStroke()

	// Add one pixel padding at bottom for outlined form
	chart.Draw.Box(r, chart.Box{Top: canvasBox.Bottom - (20 + paddingBottom), Left: canvasBox.Right - (9 + paddingRight), Bottom: canvasBox.Bottom - (paddingBottom + 1), Right: canvasBox.Right - (6 + paddingRight)}, chart.Style{FillColor: backgroundColor, StrokeColor: highlightColor, StrokeWidth: 1})

	chart.Draw.Box(r, chart.Box{Top: canvasBox.Bottom - (16 + paddingBottom), Left: canvasBox.Right - (4 + paddingRight), Bottom: canvasBox.Bottom - paddingBottom, Right: canvasBox.Right - paddingRight}, chart.Style{FillColor: highlightColor})
}

func CreateChart(cfg Config) []byte {
	pie := chart.PieChart{
		Title:           cfg.Title,
		TitleStyle:      chart.Style{Hidden: false},
		TitleInsetChart: true,
		Width:           cfg.Width,
		Height:          cfg.Height,
		Values:          cfg.GetValues(),
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

	return imageBuffer.Bytes()
}
