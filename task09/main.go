package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Мини-банкомат")
	fmt.Println("Введите q в любой момент, чтобы выйти.")

	balance, exit, err := askStartingBalance(scanner, "Введите начальный баланс: ")
	if exit || err != nil {
		finishSession(err)
		return
	}

	for {
		showActions()
		choice, readErr := askText(scanner, "Выберите действие: ")
		if readErr != nil {
			finishSession(readErr)
			return
		}
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Printf("Баланс: %s\n", balanceText(balance))
		case "2":
			amount, done, readErr := askAmount(scanner, "Сумма пополнения: ")
			if done || readErr != nil {
				finishSession(readErr)
				return
			}
			updated, opErr := topUpBalance(balance, amount)
			if opErr != nil {
				fmt.Println("Ошибка:", opErr)
				continue
			}
			balance = updated
			fmt.Printf("Баланс: %s\n", balanceText(balance))
		case "3":
			amount, done, readErr := askAmount(scanner, "Сумма снятия: ")
			if done || readErr != nil {
				finishSession(readErr)
				return
			}
			updated, opErr := takeCash(balance, amount)
			if opErr != nil {
				fmt.Println("Ошибка:", opErr)
				continue
			}
			balance = updated
			fmt.Printf("Баланс: %s\n", balanceText(balance))
		case "q", "Q", "quit", "exit", "й", "выход":
			fmt.Printf("Итоговый баланс: %s\n", balanceText(balance))
			finishSession(nil)
			return
		default:
			fmt.Println("Неизвестное действие. Введите 1, 2, 3 или q.")
		}
	}
}

func showActions() {
	fmt.Println()
	fmt.Println("1. Посмотреть баланс")
	fmt.Println("2. Пополнить баланс")
	fmt.Println("3. Снять средства")
	fmt.Println("q. Выход")
}

func finishSession(err error) {
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
		os.Exit(1)
	}
	fmt.Println("Выход.")
}

func askStartingBalance(scanner *bufio.Scanner, prompt string) (float64, bool, error) {
	for {
		line, err := askText(scanner, prompt)
		if err != nil {
			return 0, false, err
		}
		if userQuits(line) {
			return 0, true, nil
		}
		value, err := amountFromText(line)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		if err := balanceAllowed(value); err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		return value, false, nil
	}
}

func askAmount(scanner *bufio.Scanner, prompt string) (float64, bool, error) {
	for {
		line, err := askText(scanner, prompt)
		if err != nil {
			return 0, false, err
		}
		if userQuits(line) {
			return 0, true, nil
		}
		value, err := amountFromText(line)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		return value, false, nil
	}
}

func askText(scanner *bufio.Scanner, prompt string) (string, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return scanner.Text(), nil
}
