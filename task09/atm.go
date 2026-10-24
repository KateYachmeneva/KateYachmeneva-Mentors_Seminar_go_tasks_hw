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

type accountOperation struct {
	title   string
	amount  float64
	balance float64
}

func addOperation(history []accountOperation, title string, amount, balance float64) []accountOperation {
	return append(history, accountOperation{
		title:   title,
		amount:  amount,
		balance: balance,
	})
}

func historyText(history []accountOperation) string {
	if len(history) == 0 {
		return "История операций пуста."
	}
	var builder strings.Builder
	builder.WriteString("История операций:")
	for index, operation := range history {
		fmt.Fprintf(&builder, "\n%d. %s на %s, баланс: %s", index+1, operation.title, balanceText(operation.amount), balanceText(operation.balance))
	}
	return builder.String()
}

func userQuits(text string) bool {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "q", "quit", "exit", "й", "выход":
		return true
	default:
		return false
	}
}
