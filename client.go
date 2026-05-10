package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchWeather(lat, lon float64) (*WeatherResponse, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=wind_speed_10m,precipitation_probability,precipitation,weather_code&timezone=Asia%%2FTokyo&past_days=0&forecast_days=1", lat, lon)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weather_response WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather_response); err != nil {
		return nil, err
	}

	return &weather_response, nil
}
