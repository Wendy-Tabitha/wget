package utils

import (
	"fmt"
	"strings"
	"time"
)

func (pw *ProgressWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	pw.current += float64(n)
	arrowLength := 50

	// Check if total size is greater than zero to avoid division by zero
	if pw.total > 0 {
		// Calculate progress
		percentage := float64(pw.current) / float64(pw.total) * 100
		elapsed := time.Since(pw.startTime) // Total elapsed time since the start
		var remaining time.Duration

		// Calculate remaining time only if current bytes read is greater than zero
		if pw.current > 0 {
			remaining = time.Duration(float64(elapsed) / (float64(pw.current) / float64(n)) * (pw.total - pw.current))
		}

		// Calculate download speed
		speed := float64(n) / elapsed.Seconds() // bytes per second
		speedMBps := speed / (1024 * 1024)       // Convert to MBps

		// Create a moving arrow based on progress
		progressLength := int(percentage / 100 * float64(arrowLength))

		// Clamp progressLength to be at least 0
		if progressLength < 0 {
			progressLength = 0
		} else if progressLength > arrowLength {
			progressLength = arrowLength
		}

		arrow := fmt.Sprintf("[%s>%s]",
			strings.Repeat("=", progressLength),
			strings.Repeat(" ", arrowLength-progressLength))

		// Print the progress in the new order
		fmt.Printf("\u001b[2K \r %s / %s %s %.2f%% %.2f MiB/s %s",
			FormatSize(int64(pw.current)),
			FormatSize(int64(pw.total)),
			arrow,
			percentage,
			speedMBps,
			FormatDuration(remaining))

	} else {
		arrow := fmt.Sprintf("[<=>%s]",
			strings.Repeat(" ", arrowLength-3))
		percentage := "--.-%"

		fmt.Printf("\u001b[2K \r %s / --.--KiB %s %s 0.01 MiB/s 0s",
			FormatSize(int64(pw.current)), arrow, percentage)
	}
	return n, nil
}

// FormatDuration formats the duration into a human-readable string
func FormatDuration(d time.Duration) string {
	if d < 0 {
		return "0s"
	}
	return fmt.Sprintf("%ds", int(d.Seconds())) // Display only seconds
}
