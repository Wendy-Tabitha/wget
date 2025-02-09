package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"wget/utils"
)

func main() {
	// Define flags
	background := flag.Bool("B", false, "Run in background")
	outputFile := flag.String("O", "", "Output file name")
	outputPath := flag.String("P", "", "Output directory")
	// rateLimit := flag.String("rate-limit", "", "Limit download speed")
	inputFile := flag.String("i", "", "Input file with URLs")
	// mirror := flag.Bool("mirror", false, "Mirror a website")
	// reject := flag.String("R", "", "Reject file types")
	// exclude := flag.String("X", "", "Exclude directories")
	// convertLinks := flag.Bool("convert-links", false, "Convert links for offline viewing")

	flag.Parse()

	// Handle input file for multiple downloads
	if *inputFile != "" {
		utils.DownloadFromFile(*inputFile)
		return
	}

	// Get the URL from the command line arguments
	if len(flag.Args()) < 1 {
		fmt.Println("No URL provided")
		return
	}
	url := flag.Args()[0]

	// Prepare output file name and path
	var fileName string
	if *outputFile != "" {
		fileName = *outputFile
	} else {
		fileName = filepath.Base(url)
	}

	if *outputPath != "" {
		fileName = filepath.Join(*outputPath, fileName)
	}

	// Start downloading
	utils.StartDownload(url, fileName, *background)
}
