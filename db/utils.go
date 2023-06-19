package db

import "strconv"

func BytesToString(item any) string {
	return string(item.([]byte))
}

func StringToFloat64(item any) float64 {
	float, _ := strconv.ParseFloat(string(item.([]uint8)), 64)
	return float
}
