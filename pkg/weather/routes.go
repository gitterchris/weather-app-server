package weather

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gitterchris/weather-app-server/pkg/core"
	"github.com/gorilla/mux"
)

type routes struct {
	service core.WeatherService
}

// RegisterRoutes registers the weather routes to the server
func RegisterRoutes(router *mux.Router, service core.WeatherService) {
	r := routes{
		service: service,
	}
	s := router.PathPrefix("/weather").Subrouter()

	s.Handle("/locations", core.AppHandler(r.locations)).Methods(http.MethodGet)
	s.Handle("/locations/{id}", core.AppHandler(r.location)).Methods(http.MethodGet)
}

func (ro *routes) locations(w http.ResponseWriter, r *http.Request) *core.ServerError {
	q := r.URL.Query().Get("q")
	if q == "" {
		return core.NewServerError(
			"query string cannot be empty",
			http.StatusBadRequest,
			errors.New("invalid parameter"),
		)
	}

	locations, err := ro.service.Locations(q)
	if err != nil {
		return core.NewServerError(
			"cannot retrieve location information from the weather API",
			http.StatusInternalServerError,
			err,
		)
	}

	json.NewEncoder(w).Encode(locations)
	return nil
}

func (ro *routes) location(w http.ResponseWriter, r *http.Request) *core.ServerError {
	date, err := time.Parse(time.RFC3339, r.URL.Query().Get("date"))
	if err != nil {
		return core.NewServerError(
			"invalid date query string parameter",
			http.StatusBadRequest,
			err,
		)
	}

	id := mux.Vars(r)["id"]
	woeID, err := strconv.Atoi(id)
	if err != nil {
		return core.NewServerError(
			"invalid woeID. Cannot parse into int",
			http.StatusBadRequest,
			err,
		)
	}

	location, err := ro.service.WeatherForecast(woeID, date)
	if err != nil {
		return core.NewServerError(
			"cannot retrieve weather information from the weather API for the specified date and location",
			http.StatusInternalServerError,
			err,
		)
	}

	json.NewEncoder(w).Encode(location)
	return nil
}
