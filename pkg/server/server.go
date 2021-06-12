package server

import (
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/votebot/piechart-service/pkg/piechart"
	"go.uber.org/zap"
)

type Server struct {
	address string
}

func (s *Server) createImage(c *fiber.Ctx) (err error) {
	var chartConfig piechart.Config
	err = json.Unmarshal(c.Body(), &chartConfig)

	if err != nil {
		c.Status(400)
		_ = c.SendString("invalid body")
		zap.L().Warn("cannot parse JSON", zap.Error(err))
		return err
	}

	imgBytes := piechart.CreateChart(chartConfig)
	c.Set("Content-type", "image/png")
	err = c.Send(imgBytes)
	if err != nil {
		c.Status(500)
		_ = c.SendString("send error")
		zap.L().Warn("cannot send image", zap.Error(err))
		return err
	}

	return err
}

func (s *Server) Start() {
	app := fiber.New()

	app.Post("/create", s.createImage)

	err := app.Listen(s.address)

	if err != nil {
		sentry.CaptureException(err)
		zap.L().Fatal("failed to serve", zap.Error(err))
	}
}

func CreateServer(address string) *Server {
	return &Server{
		address: address,
	}
}
