package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPositive(n int64) bool {
	return n > 0
}

func isNegative(n int64) bool {
	return n < 0
}

func isEven(n int64) bool {
	return n%2 == 0
}

func isDivisibleBy(n, divisor int64) bool {
	if divisor == 0 {
		return false
	}
	return n%divisor == 0
}

func signMessage(n int64) string {
	if isPositive(n) {
		return "Число положительное."
	}
	if isNegative(n) {
		return "Число отрицательное."
	}
	return "Число равно нулю: оно ни положительное, ни отрицательное."
}

func evenOddMessage(n int64) string {
	if isEven(n) {
		return "Число чётное."
	}
	return "Число нечётное."
}

func divisionMessage(n int64) string {
	divisibleBy3 := isDivisibleBy(n, 3)
	divisibleBy5 := isDivisibleBy(n, 5)
	switch {
	case divisibleBy3 && divisibleBy5:
		return "Число делится и на 3, и на 5."
	case divisibleBy3:
		return "Число делится на 3, но не делится на 5."
	case divisibleBy5:
		return "Число делится на 5, но не делится на 3."
	default:
		return "Число не делится ни на 3, ни на 5."
	}
}

func fullReport(n int64) string {
	return signMessage(n) + "\n" + evenOddMessage(n) + "\n" + divisionMessage(n)
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
