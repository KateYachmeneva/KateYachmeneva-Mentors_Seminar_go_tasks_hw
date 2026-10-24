package main

import (
	"fmt"
	"strconv"
	"strings"
)

type sequenceStats struct {
	sum       int64
	average   float64
	maximum   int64
	minimum   int64
	evenCount int
	oddCount  int
}

func isEven(n int64) bool {
	return n%2 == 0
}

func isOdd(n int64) bool {
	return n%2 != 0
}

func calculateStats(numbers []int64) (sequenceStats, error) {
	if len(numbers) == 0 {
		return sequenceStats{}, fmt.Errorf("нет чисел")
	}

	stats := sequenceStats{
		minimum: numbers[0],
		maximum: numbers[0],
	}
	for _, number := range numbers {
		stats.sum += number
		if number < stats.minimum {
			stats.minimum = number
		}
		if number > stats.maximum {
			stats.maximum = number
		}
		if isEven(number) {
			stats.evenCount++
		} else if isOdd(number) {
			stats.oddCount++
		}
	}
	stats.average = float64(stats.sum) / float64(len(numbers))
	return stats, nil
}

func formatAverage(value float64) string {
	text := strconv.FormatFloat(value, 'f', 10, 64)
	text = strings.TrimRight(text, "0")
	text = strings.TrimRight(text, ".")
	if text == "-0" {
		return "0"
	}
	return text
}

func countIsValid(count int) error {
	if count < 1 {
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
	count, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("%q не является целым числом", strings.TrimSpace(raw))
	}
	if err := countIsValid(count); err != nil {
		return 0, err
	}
	return count, nil
}

func parseInteger(raw string) (int64, error) {
	text := strings.TrimSpace(raw)
	if text == "" {
		return 0, fmt.Errorf("пустая строка")
	}
	if strings.Contains(text, ".") || strings.Contains(text, ",") {
		return 0, fmt.Errorf("%q не целое число", strings.TrimSpace(raw))
	}
	value, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%q не является целым числом", strings.TrimSpace(raw))
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
