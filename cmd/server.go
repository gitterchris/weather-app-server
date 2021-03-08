package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/gitterchris/weather-app-server/pkg/weather"
)

type server struct {
	router *mux.Router
}

func newServer() *server {
	return &server{
		router: mux.NewRouter(),
	}
}

func (s *server) registerEndpoints() {
	subRouter := s.router.PathPrefix("/api").Subrouter()

	weatherService := weather.NewService()
	weather.RegisterRoutes(subRouter, weatherService)
}

func (s *server) ServeHTTP() {
	s.registerEndpoints()

	s.router.Use(
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
		),
	)

	httpServer := &http.Server{
		Handler:      s.router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Panic("Unable to run server", err)
	}
}
