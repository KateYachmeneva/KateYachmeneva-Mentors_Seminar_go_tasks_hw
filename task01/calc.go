package main

import (
	"fmt"
	"strconv"
	"strings"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}

func calculate(a, b float64, operation string) (float64, error) {
	switch strings.TrimSpace(operation) {
	case "+":
		return add(a, b), nil
	case "-":
		return subtract(a, b), nil
	case "*":
		return multiply(a, b), nil
	case "/":
		return divide(a, b)
	default:
		return 0, fmt.Errorf("неизвестная операция %q: допустимы +, -, *, /", operation)
	}
}


func parseFloat(raw string) (float64, error) {
	text := strings.TrimSpace(raw)
	text = strings.ReplaceAll(text, ",", ".")
	if text == "" {
		return 0, fmt.Errorf("пустая строка")
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("%q не является числом", raw)
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
