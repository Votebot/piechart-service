package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/votebot/piechart-service/pkg/config"
	"github.com/votebot/piechart-service/pkg/server"
	"go.uber.org/zap"
	"log"
)

func main() {
	cfg := config.LoadConfig()
	initializeLogger(cfg.Dev)

	err := sentry.Init(sentry.ClientOptions{Dsn: cfg.SentryDsn})
	if err != nil {
		zap.L().Warn("sentry could not be initialized: ", zap.Error(err))
	}

	srv := server.CreateServer(cfg.BindAddress)

	srv.Start()
}

func initializeLogger(dev bool) {
	var (
		l   *zap.Logger
		err error
	)
	if dev {
		l, err = zap.NewDevelopment()
	} else {
		l, err = zap.NewProduction()
	}
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal("failed creating logger: ", err)
		return
	}
	zap.ReplaceGlobals(l)
}