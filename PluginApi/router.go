package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Router struct
{}

var server http.Server

func (p Router) Run() (err error) {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Welcome to bonusplay.pl API"))
	})

	log.Info("Starting api on port 3020")
	server = http.Server{
		Addr: ":3020",
		Handler: router,
	}
	return server.ListenAndServe()
}

func (p Router) Cancel() {
	log.Debug("API shutting down")
	_ = server.Shutdown(context.Background())
}

// exported plugin
//noinspection GoUnusedGlobalVariable
var Plugin Router