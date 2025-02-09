package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func StartDownload(url, fileName string, background bool) {
	startTime := time.Now()
	fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))

	// Send request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Status: %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		return
	}

	// Get content length
	contentLength := resp.ContentLength
	fmt.Printf("content size: %d [~%.2fMB]\n", contentLength, float64(contentLength)/1024/1024)

	// Create file
	out, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer out.Close()

	// Download file
	progress := &ProgressWriter{total: contentLength, current: 0, startTime: startTime}
	_, err = io.Copy(out, io.TeeReader(resp.Body, progress))
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}

	fmt.Printf("Downloaded [%s]\n", url)
	fmt.Printf("finished at %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

func (pw *ProgressWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	pw.current += int64(n)

	// Calculate progress
	percentage := float64(pw.current) / float64(pw.total) * 100
	elapsed := time.Since(pw.startTime)
	remaining := time.Duration(float64(elapsed) / (float64(pw.current) / float64(n)) * float64(pw.total-pw.current))

	fmt.Printf("\r %d / %d [%.2f%%] %s remaining", pw.current, pw.total, percentage, remaining)
	return n, nil
}
