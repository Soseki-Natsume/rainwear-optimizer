package main

import (
	"fmt"
	"os"

	"github.com/Soseki-Natsume/rainwear-optimizer/utils"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	var config WeatherConfig
	var err error

	config.Latitude, err = utils.GetEnvFloat(os.Getenv("APP_LAT"))
	if err != nil {
		fmt.Println(err)
		return
	}
	config.Longitude, err = utils.GetEnvFloat(os.Getenv("APP_LON"))
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
