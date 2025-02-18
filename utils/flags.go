package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseRareLimit(rateStr string) (float64, error) {
	if rateStr == "" {
		return 0, fmt.Errorf("usage: go run cmd/main.go [--rate-limit=<rate-limit>] <URL>")
	}
	var unit string

	if len(rateStr) > 0 {
		lastChar := rateStr[len(rateStr)-1:]
		if strings.ContainsAny(lastChar, "mMkK") {
			unit = lastChar
			rateStr = rateStr[:len(rateStr)-1]
		}
	}

	limitBps, err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		return 0, err
	}

	switch strings.ToLower(unit) {
	case "m":
		limitBps *= 1_024 * 1_024
	case "k":
		limitBps *= 1_024
	default:
		return 0, fmt.Errorf("unit value must be either 'k' or 'M'") 
	}

	return limitBps, nil
}
