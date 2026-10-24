package main

import (
	"fmt"
	"strconv"
	"strings"
)

// sumFrom1ToN складывает 1 + 2 + ... + n в цикле.
func sumFrom1ToN(n int64) float64 {
	var sum float64
	for number := int64(1); number <= n; number++ {
		sum += float64(number)
	}
	return sum
}

// averageFrom1ToN возвращает среднее чисел от 1 до n.
func averageFrom1ToN(n int64, sum float64) float64 {
	return sum / float64(n)
}

// sumAndAverage возвращает сумму чисел от 1 до n и их среднее.
func sumAndAverage(n int64) (sum float64, average float64, err error) {
	if n < 1 {
		return 0, 0, fmt.Errorf("N должно быть целым числом больше нуля")
	}
	sum = sumFrom1ToN(n)
	average = averageFrom1ToN(n, sum)
	return sum, average, nil
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

func formatNumber(value float64) string {
	text := strconv.FormatFloat(value, 'f', 10, 64)
	text = strings.TrimRight(text, "0")
	text = strings.TrimRight(text, ".")
	if text == "-0" {
		return "0"
	}
	return text
}
