package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/chatex-com/process-manager"
	log "github.com/sirupsen/logrus"

	"github.com/MarySmirnova/create_pdf/internal"
	"github.com/MarySmirnova/create_pdf/internal/config"
)

var cfg config.Application

func init() {
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	log.SetLevel(level)

	process.SetLogger(&PMLogger{Logger: log.StandardLogger()})
}

func main() {
	app, err := internal.NewApplication(cfg)
	if err != nil {
		panic(err)
	}

	app.Run()
}
