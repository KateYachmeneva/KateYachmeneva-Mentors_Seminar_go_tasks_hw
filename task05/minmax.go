package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func smaller(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func larger(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func validCount(n int) error {
	if n < 1 {
		return fmt.Errorf("количество чисел должно быть больше нуля")
	}
	return nil
}

func parseCount(raw string) (int, error) {
	text := strings.TrimSpace(raw)
	if text == "" {
		return 0, fmt.Errorf("пустая строка")
	}
	if strings.Contains(text, ".") || strings.Contains(text, ",") {
		return 0, fmt.Errorf("%q не целое число", strings.TrimSpace(raw))
	}
	value, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("%q не является целым числом", strings.TrimSpace(raw))
	}
	if err := validCount(value); err != nil {
		return 0, err
	}
	return value, nil
}

func parseNumber(raw string) (float64, error) {
	text := strings.TrimSpace(raw)
	text = strings.ReplaceAll(text, ",", ".")
	if text == "" {
		return 0, fmt.Errorf("пустая строка")
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil || math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, fmt.Errorf("%q не является числом", strings.TrimSpace(raw))
	}
	return value, nil
}

func isExit(text string) bool {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "q", "quit", "exit", "й", "выход":
		return true
	default:
		return false
	}
}

func formatNumber(value float64) string {
	text := strconv.FormatFloat(value, 'f', 10, 64)
	text = strings.TrimRight(text, "0")
	text = strings.TrimRight(text, ".")
	if text == "-0" {
		return "0"
	}
	return text
}
