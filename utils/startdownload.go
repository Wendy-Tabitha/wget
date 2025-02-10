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
	fmt.Printf("--%s--  %s\n", startTime.Format("2006-01-02 15:04:05"), url)

	// Send request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP request sent, awaiting response... %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		return
	}

	// Get content length
	contentLength := resp.ContentLength
	if contentLength < 0 {
		contentLength = 0 // Handle unknown content length
	}
	fmt.Printf("Length: %d (%s) [%s]\n", contentLength, formatSize(contentLength), resp.Header.Get("Content-Type"))

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

	fmt.Printf("\n%s 100%%[========================>]  %s   in %s\n", fileName, formatSize(contentLength), time.Since(startTime))
	fmt.Printf("'%s' saved [%d/%d]\n", fileName, contentLength, contentLength)
}

func (pw *ProgressWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	pw.current += int64(n)

	// Calculate progress
	percentage := float64(pw.current) / float64(pw.total) * 100
	elapsed := time.Since(pw.startTime)
	remaining := time.Duration(float64(elapsed) / (float64(pw.current) / float64(n)) * float64(pw.total-pw.current))

	fmt.Printf("\r %s %d / %d [%.2f%%] %s remaining", formatSize(pw.current), pw.current, pw.total, percentage, remaining)
	return n, nil
}

func formatSize(size int64) string {
	if size >= 1_048_576 { // 1 MiB
		return fmt.Sprintf("%.2fMiB", float64(size)/1_048_576)
	} else if size >= 1_024 { // 1 KiB
		return fmt.Sprintf("%.2fKiB", float64(size)/1_024)
	}
	return fmt.Sprintf("%dB", size)
}