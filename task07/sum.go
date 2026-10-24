package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isEven(n int64) bool {
	return n%2 == 0
}

func isDivisibleBy(n, divisor int64) bool {
	if divisor == 0 {
		return false
	}
	return n%divisor == 0
}

// sumSuitable складывает числа от 1 до n, которые чётные и делятся на 3.
func sumSuitable(n int64) (int64, error) {
	if n < 1 {
		return 0, fmt.Errorf("N должно быть целым числом больше нуля")
	}
	var sum int64
	for number := int64(1); number <= n; number++ {
		if isEven(number) && isDivisibleBy(number, 3) {
			sum += number
		}
	}
	return sum, nil
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
