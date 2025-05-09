package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/omkar-nag/socialapp/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type config struct {
	addr string
	db   dbConfig
	env  string
}

func (a *application) run(mux http.Handler) error {

	srv := &http.Server{Addr: a.config.addr, Handler: mux, WriteTimeout: time.Second * 30, ReadTimeout: time.Second * 10, IdleTimeout: time.Minute}
	log.Printf("Server started at %s", a.config.addr)
	return srv.ListenAndServe()
}

func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		//  /v1/posts
		r.Route("/posts", func(r chi.Router) {
			r.Post("/", app.createPostHandler)
		})
		r.Route("/users", func(r chi.Router) {
			r.Post("/", app.createUserHandler)
		})
	})

	return r
}
