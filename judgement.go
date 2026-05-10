package main

import (
	"os"
	"strconv"
)

type Rainwear string

const (
	unnecessary     Rainwear = "不要"
	foldingUmbrella Rainwear = "折り畳み傘"
	umbrella        Rainwear = "傘"
	raincoat        Rainwear = "レインコート"
)

const (
	msToKmh = 3.6 // m/s*3.6 = km/h
)

func judgeWeather(weather *WeatherResponse) Rainwear {
	departureTime := getEnvInt(os.Getenv("DEPARTURE_TIME"))
	// 12時間先までの降水量と風速を確認する
	checkHours := departureTime + 12
	// 同日まで確認可能とする
	if checkHours > 24 {
		checkHours = 24
	}

	weatherCode := weather.Hourly.WeatherCode[departureTime]
	// Exception for rain(source: https://www.jodc.go.jp/data_format/weather-code_j.html)
	// 0-19: no precipitation here
	// 30-35: sandstorm or duststorm
	// 40-49: fog or ice fog
	// 上記以外は大体雨
	if weatherCode < 20 || weatherCode >= 30 && weatherCode < 36 || weatherCode >= 40 && weatherCode < 50 {
		return unnecessary
	}
	// 70-79: solid precipitation not in showers
	if weatherCode >= 70 && weatherCode < 80 {
		return umbrella
	}

	maxWindSpeed10m := weather.Hourly.WindSpeed10m[departureTime]

	// 計算量が小さいため、slices.Maxを使用せずループで最大値を求める
	for i := departureTime; i <= checkHours; i++ {
		if maxWindSpeed10m < weather.Hourly.WindSpeed10m[i] {
			maxWindSpeed10m = weather.Hourly.WindSpeed10m[i]
		}
	}

	// APIで取得できるのはkm/hなので、m/sをkm/hに変換して比較
	if maxWindSpeed10m <= 3.0*msToKmh {
		return foldingUmbrella
	} else if maxWindSpeed10m <= 7.0*msToKmh {
		return umbrella
	} else {
		return raincoat
	}
}

func getEnvInt(stringValue string) int {
	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0
	}
	return intValue
}
