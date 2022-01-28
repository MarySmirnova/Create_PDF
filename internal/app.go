package internal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chatex-com/process-manager"

	"github.com/MarySmirnova/create_pdf/internal/config"
	"github.com/MarySmirnova/create_pdf/internal/reports"
	"github.com/MarySmirnova/create_pdf/internal/rest"
)

type Application struct {
	sigChan <-chan os.Signal
	cfg     config.Application
	manager *process.Manager
}

func NewApplication(cfg config.Application) (*Application, error) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	app := &Application{
		sigChan: sigChan,
		cfg:     cfg,
		manager: process.NewManager(),
	}

	if err := app.bootstrap(); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *Application) bootstrap() error {
	// init dependencies
	// init workers

	if err := a.bootstrapRestWorker(); err != nil {
		return err
	}

	return nil
}

func (a *Application) bootstrapRestWorker() error {
	worker := rest.NewWorker(a.cfg.REST, reports.NewReport8949Generator(a.cfg.ReportTemplatesPath, a.cfg.Report8949))
	a.manager.AddWorker(process.NewServerWorker("REST Server", worker.GetHTTPServer()))

	return nil
}

func (a *Application) Run() {
	a.manager.StartAll()
	a.registerShutdown()
}

func (a *Application) registerShutdown() {
	go func(manager *process.Manager) {
		<-a.sigChan

		manager.StopAll()
	}(a.manager)

	a.manager.AwaitAll()
}
