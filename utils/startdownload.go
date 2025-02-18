package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func StartDownload(url string, limitBps interface{}, outputFilename string, savePath string) {
	startTime := time.Now()
	fmt.Printf("start at %s\n", startTime.Format("2006-01-02 15:04:05"))
	fmt.Printf("sending request, awaiting response... ")

	// Get full save path
	fullPath, err := HandlePath(savePath, outputFilename, url)
	if err != nil {
		fmt.Printf("Error setting up save path: %v\n", err)
		return
	}

	// Send request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("status %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))
		return
	}
	fmt.Printf("status %d %s\n", resp.StatusCode, http.StatusText(resp.StatusCode))

	// Get content length
	contentLength := float64(resp.ContentLength)
	if contentLength < 0 {
		fmt.Printf("content size: -- [~-.--MB]\n")
	} else {
		fmt.Printf("content size: %.0f [~%.2fMB]\n", contentLength, float64(contentLength)/1_048_576)
	}

	fmt.Printf("saving file to: %s\n", fullPath)

	// Create output file
	out, err := os.Create(fullPath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer out.Close()

	if limitBps == nil {
		limitBps = 1 * 1024 * 1024 // 1 MBps default
	}

	limiter := &RateLimiter{
		reader:   resp.Body,
		limitBps: limitBps.(float64),
		lastTime: time.Now(),
	}

	// Download with progress tracking
	progress := &ProgressWriter{total: contentLength, current: 0, startTime: startTime}
	_, err = io.Copy(out, io.TeeReader(limiter, progress))
	if err != nil {
		fmt.Printf("Error downloading file: %v\n", err)
		return
	}

	finishTime := time.Now()
	fmt.Printf("\nDownloaded [%s]\nfinished at %s\n", url, finishTime.Format("2006-01-02 15:04:05"))
}

func FormatSize(size int64) string {
	if size >= 1_048_576 { // 1 MiB
		return fmt.Sprintf("%.2fMiB", float64(size)/1_048_576)
	} else if size >= 1_024 { // 1 KiB
		return fmt.Sprintf("%.2fKiB", float64(size)/1_024)
	}
	return fmt.Sprintf("%dB", size)
}
