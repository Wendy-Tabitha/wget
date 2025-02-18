package main

import (
	"flag"
	"fmt"
	"os"

	"wget/utils"
)

func printUsage() {
	fmt.Println("wget - a simple wget clone")
	fmt.Println("\nUsage:")
	fmt.Println("  go run cmd/main.go [options] <URL>")
	fmt.Println("\nOptions:")
	fmt.Println("  -O=filename     Save file with a different name")
	fmt.Println("  -P=directory    Save file in specified directory (e.g., 'downloads' or '~/Downloads')")
	fmt.Println("  -rate-limit=N   Limit download speed (e.g., '500k' or '1M')")
	fmt.Println("\nExamples:")
	fmt.Println("  go run cmd/main.go https://example.com/file.jpg")
	fmt.Println("  go run cmd/main.go -O=custom.jpg https://example.com/file.jpg")
	fmt.Println("  go run cmd/main.go -P=downloads -rate-limit=500k https://example.com/file.jpg")
}

func main() {
	// Define flags
	rateLimiterFlag := flag.String("rate-limit", "1M", "rate limit in bytes per second")
	outputFlag := flag.String("O", "", "output filename")
	pathFlag := flag.String("P", "", "directory to save the downloaded file")

	// Parse flags
	flag.Parse()

	// If no arguments provided, print usage and exit normally
	if len(os.Args) == 1 {
		printUsage()
		return
	}

	// Check for URL argument
	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		os.Exit(1)
	}

	url := args[0]

	// Parse rate limit
	limitBps, err := utils.ParseRareLimit(*rateLimiterFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Start download
	utils.StartDownload(url, limitBps, *outputFlag, *pathFlag)
}
