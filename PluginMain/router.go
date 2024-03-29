package main

import (
	"context"
	"github.com/BonusPlay/VueHoster/util"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
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
		util.ServeFile("main_files/index.html", w)
	})

	router.Get("/cv_en", func(w http.ResponseWriter, r *http.Request) {
		util.ServeFile("main_files/cv_en.pdf", w)
	})

	router.Get("/cv_pl", func(w http.ResponseWriter, r *http.Request) {
		util.ServeFile("main_files/cv_pl.pdf", w)
	})

	router.Get("/github", http.RedirectHandler("https://github.com/BonusPlay", 301).ServeHTTP)
	router.Get("/facebook", http.RedirectHandler("https://facebook.com/BonusPlay3", 301).ServeHTTP)
	router.Get("/discord", http.RedirectHandler("https://discordapp.com/invite/tYk4PW5", 301).ServeHTTP)
	router.Get("/youtube", http.RedirectHandler("https://www.youtube.com/user/adamklis1975", 301).ServeHTTP)
	router.Get("/asktoask", http.RedirectHandler("https://www.youtube.com/watch?v=53zkBvL4ZB4", 301).ServeHTTP)
	router.Get("/why", http.RedirectHandler("https://www.youtube.com/watch?v=VPpIjhtgGj0", 301).ServeHTTP)
	router.Get("/linkedin", http.RedirectHandler("https://www.linkedin.com/in/adam-kliś", 301).ServeHTTP)

	// static files
	workDir, _ := os.Getwd()
	staticDir := filepath.Join(workDir, "main_files")

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {

		log.Info("main_files" + r.URL.Path)

		if _, err := os.Stat("main_files" + r.URL.Path); os.IsNotExist(err) {
			http.ServeFile(w, r, "main_files/index.html")
		} else {
			http.FileServer(http.Dir(staticDir)).ServeHTTP(w, r)
		}
	})

	log.Info("Starting main on port 3010")
	server = http.Server{
		Addr: ":3010",
		Handler: router,
	}
	return server.ListenAndServe()
}

func (p Router) Cancel() {
	log.Debug("Main shutting down")
	_ = server.Shutdown(context.Background())
}

// exported plugin
//noinspection GoUnusedGlobalVariable
var Plugin Router