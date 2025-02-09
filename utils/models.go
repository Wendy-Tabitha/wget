package utils

import "time"

type ProgressWriter struct {
	total     int64
	current   int64
	startTime time.Time
}
