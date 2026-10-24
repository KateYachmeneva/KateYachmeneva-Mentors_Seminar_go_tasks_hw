package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func topUpBalance(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, fmt.Errorf("сумма пополнения должна быть больше нуля")
	}
	return roundToKopecks(balance + amount), nil
}

func takeCash(balance, amount float64) (float64, error) {
	if amount <= 0 {
		return balance, fmt.Errorf("сумма снятия должна быть больше нуля")
	}
	if amount > balance {
		return balance, fmt.Errorf("недостаточно средств")
	}
	return roundToKopecks(balance - amount), nil
}

func roundToKopecks(value float64) float64 {
	return math.Round(value*100) / 100
}

func balanceText(value float64) string {
	return fmt.Sprintf("%.2f руб.", value)
}

func amountFromText(raw string) (float64, error) {
	text := strings.TrimSpace(raw)
	text = strings.ReplaceAll(text, ",", ".")
	if text == "" {
		return 0, fmt.Errorf("пустая строка")
	}
	value, err := strconv.ParseFloat(text, 64)
	if err != nil || math.IsNaN(value) || math.IsInf(value, 0) {
		return 0, fmt.Errorf("%q не является числом", strings.TrimSpace(raw))
	}
	return roundToKopecks(value), nil
}

func balanceAllowed(balance float64) error {
	if balance < 0 {
		return fmt.Errorf("начальный баланс не может быть отрицательным")
	}
	return nil
}

func userQuits(text string) bool {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "q", "quit", "exit", "й", "выход":
		return true
	default:
		return false
	}
}
