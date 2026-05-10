package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	var config WeatherConfig
	var err error

	config.Latitude, err = getEnvFloat(os.Getenv("APP_LAT"))
	if err != nil {
		fmt.Println(err)
		return
	}
	config.Longitude, err = getEnvFloat(os.Getenv("APP_LON"))
	if err != nil {
		fmt.Println(err)
		return
	}

	weather, err := fetchWeather(config.Latitude, config.Longitude)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(weather.Hourly.WeatherCode)
	fmt.Println(weather.Hourly.WindSpeed10m)
	fmt.Println(judgeWeather(weather))
}

func getEnvFloat(stringValue string) (float64, error) {
	float64Value, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return 0, err
	}
	return float64Value, nil
}
