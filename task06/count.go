package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type numberCounts struct {
	positive int
	negative int
	zero     int
}

func countBySign(counts numberCounts, number float64) numberCounts {
	if number > 0 {
		counts.positive++
	} else if number < 0 {
		counts.negative++
	} else {
		counts.zero++
	}
	return counts
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
