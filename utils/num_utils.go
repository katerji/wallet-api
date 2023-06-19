package utils

import "time"

func IntToUnixTime(num int) int64 {
	now := time.Now()
	duration := time.Duration(num) * time.Second
	result := now.Add(duration)
	return result.Unix()
}
