package utils

import (
	"time"
)

func (r *RateLimiter) Read(p []byte) (int, error) {
	n, err := r.reader.Read(p)
	if err != nil {
		return 0, err
	}

	r.bytesRead += int64(n)

	now := time.Now()
	elapsed := now.Sub(r.lastTime)

	expectedTime := time.Duration(r.bytesRead) * time.Second / time.Duration(r.limitBps)
	if elapsed < expectedTime {
		time.Sleep(expectedTime - elapsed)
	}

	r.lastTime = time.Now()
	return n, err
}
