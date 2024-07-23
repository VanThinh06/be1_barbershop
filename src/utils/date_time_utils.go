package utils

import "time"

func ConvertMinutesToMicroseconds(minutes int32) int64 {
	duration := time.Duration(minutes) * time.Minute
	microseconds := duration.Microseconds()
	return microseconds
}
