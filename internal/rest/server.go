package rest

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/MarySmirnova/create_pdf/internal/config"
	"github.com/MarySmirnova/create_pdf/internal/form"
	"github.com/MarySmirnova/create_pdf/internal/reports"
)

type Worker struct {
	f8949Generator *reports.Report8949Generator

	httpServer *http.Server
}

func NewWorker(cfg config.REST, f8949Generator *reports.Report8949Generator) *Worker {
	w := &Worker{
		f8949Generator: f8949Generator,
	}

	handler := mux.NewRouter()
	handler.Name("generate_report").Methods(http.MethodPost).Path("/generate").HandlerFunc(w.generateReportHandler)

	w.httpServer = &http.Server{
		Addr:         cfg.Listen,
		Handler:      handler,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	return w
}

func (worker *Worker) GetHTTPServer() *http.Server {
	return worker.httpServer
}

func (worker *Worker) generateReportHandler(w http.ResponseWriter, r *http.Request) {
	var request form.Report

	// parsing request data
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.WithError(err).Warn("unable to parse the request") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: validate request

	// execute business logic
	report, err := worker.f8949Generator.Generate(request)
	if err != nil {
		log.WithError(err).Warn("unable to generate the report") // TODO: prepare public error description
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// preparing and sending the response
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename=f6949.pdf")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(report)
}
