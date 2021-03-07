package core

import "time"

// Location represents the response from the location weather API
type Location struct {
	ID          int    `json:"woeid"`
	Title       string `json:"title"`
	Type        string `json:"location_type"`
	Coordinates string `json:"latt_long"`
	Distance    int    `json:"distance,omitempty"`
}

// Forecast represents the response from the location day weather API
type Forecast struct {
	ID                   int     `json:"id"`
	ApplicableDate       string  `json:"applicable_date"`
	State                string  `json:"weather_state_name"`
	StateAbbreviation    string  `json:"weather_state_abbr"`
	WindSpeed            float64 `json:"wind_speed"`
	WindDirection        float64 `json:"wind_direction"`
	WindDirectionCompass string  `json:"wind_direction_compass"`
	MinTemp              float64 `json:"min_temp"`
	MaxTemp              float64 `json:"max_temp"`
	TheTemp              float64 `json:"the_temp"`
	AirPressure          float64 `json:"air_pressure"`
	Humidity             float64 `json:"humidity"`
	Visibility           float64 `json:"visibility"`
	Predictability       int     `json:"predictability"`
}

// WeatherService defines the contract for the weather services
type WeatherService interface {
	Locations(searchString string) ([]*Location, error)
	WeatherForecast(woeID int, date time.Time) (*Forecast, error)
}
