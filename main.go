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

	config.Latitude = getEnvFloat(os.Getenv("APP_LAT"))
	config.Longitude = getEnvFloat(os.Getenv("APP_LON"))

	weather, err := fetchWeather(config.Latitude, config.Longitude)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(weather)
}

func getEnvFloat(stringValue string) float64 {
	float64Value, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return 0
	}
	return float64Value
}
