package utils

import (
	"io"
	"time"
)

type ProgressWriter struct {
	total     float64
	current   float64
	startTime time.Time
}

type RateLimiter struct {
	reader    io.Reader
	limitBps  float64
	bytesRead int64
	lastTime  time.Time
}

// PathConfig holds path-related configuration
type PathConfig struct {
	SavePath      string
	OutputFile    string
	OriginalURL   string
}

func (pc *PathConfig) IsEmpty() bool {
	return pc.SavePath == "" && pc.OutputFile == ""
}

