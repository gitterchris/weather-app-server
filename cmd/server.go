package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/gitterchris/weather-app-server/pkg/weather"
)

type Server struct {
	router *mux.Router
}

func NewServer() *Server {
	return &Server{
		router: mux.NewRouter(),
	}
}

func (s *Server) registerEndpoints() {
	subRouter := s.router.PathPrefix("/api").Subrouter()

	weatherService := weather.NewService()
	weather.RegisterRoutes(subRouter, weatherService)
}

func (s *Server) ServeHTTP() {
	s.registerEndpoints()

	server := &http.Server{
		Handler:      s.router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Panic("Unable to run server", err)
	}
}
