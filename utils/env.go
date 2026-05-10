package utils

import "strconv"

func GetEnvInt(stringValue string) (int, error) {
	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func GetEnvFloat(stringValue string) (float64, error) {
	float64Value, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return 0, err
	}
	return float64Value, nil
}
