package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func DownloadFromFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		if url != "" {
			StartDownload(url, filepath.Base(url), false)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
