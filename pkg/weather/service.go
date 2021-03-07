package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gitterchris/weather-app-server/pkg/core"
)

const metaweather = "https://www.metaweather.com"

// Service implements core.WeatherService interface
type Service struct {
}

// NewService instantiates a new Service object
func NewService() *Service {
	return &Service{}
}

// Locations queries /api/location/search/?query=(query)
func (s *Service) Locations(searchString string) ([]*core.Location, error) {
	url := fmt.Sprintf("%s/api/location/search/?query=%s", metaweather, searchString)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var locations []*core.Location
	if err := json.NewDecoder(res.Body).Decode(&locations); err != nil {
		return nil, err
	}

	return locations, nil
}

// WeatherForecast queries /api/location/(woeid)/(date)/
// It returns the first consolidated weather from the weather API,
// Assuming that the latest is the most accurate forecast
func (s *Service) WeatherForecast(woeID int, date time.Time) (*core.Forecast, error) {
	year, month, day := date.Date()
	url := fmt.Sprintf(
		"%s/api/location/%d/%d/%d/%d",
		metaweather,
		woeID,
		year,
		int(month),
		day,
	)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var forecasts []*core.Forecast
	if err := json.NewDecoder(res.Body).Decode(&forecasts); err != nil {
		return nil, err
	}

	if len(forecasts) == 0 {
		return nil, errors.New("no forecast for the selected date and ID")
	}

	return forecasts[0], nil
}
