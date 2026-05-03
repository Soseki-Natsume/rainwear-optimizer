package main

type WeatherResponse struct {
	Hourly struct {
		Time                     []string  `json:"time"`
		WindSpeed10m             []float64 `json:"wind_speed_10m"`
		PrecipitationProbability []float64 `json:"precipitation_probability"`
		Precipitation            []float64 `json:"precipitation"`
		WeatherCode              []int     `json:"weather_code"`
	} `json:"hourly"`
}

type WeatherConfig struct {
	Latitude  float64
	Longitude float64
}
