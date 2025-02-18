package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func HandlePath(savePath, outputFile, url string) (string, error) {
	// Get filename from either -O flag or URL
	filename := outputFile
	if filename == "" {
		filename = filepath.Base(url)
	}

	// If no path specified, use current directory
	if savePath == "" {
		return filename, nil
	}

	// Remove leading '/' to make it relative to current directory
	savePath = strings.TrimPrefix(savePath, "/")

	// Handle home directory expansion
	if strings.HasPrefix(savePath, "~/") || savePath == "~" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		if savePath == "~" {
			savePath = homeDir
		} else {
			savePath = filepath.Join(homeDir, savePath[2:])
		}
	} else {
		// If not starting with ~, make it relative to current directory
		currentDir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		savePath = filepath.Join(currentDir, savePath)
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(savePath, 0755); err != nil {
		return "", err
	}

	return filepath.Join(savePath, filename), nil
}