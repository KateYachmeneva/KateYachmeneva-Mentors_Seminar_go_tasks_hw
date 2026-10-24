package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Калькулятор: +, -, *, /")
	fmt.Println("Введите q в любой момент, чтобы выйти.")

	for {
		a, exit, err := readNumber(scanner, "Первое число: ")
		if exit || err != nil {
			printExit(err)
			return
		}

		operation, exit, err := readOperation(scanner, "Операция (+, -, *, /): ")
		if exit || err != nil {
			printExit(err)
			return
		}

		b, exit, err := readNumber(scanner, "Второе число: ")
		if exit || err != nil {
			printExit(err)
			return
		}

		result, err := calculate(a, b, operation)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Printf("Результат: %s\n", formatNumber(result))
		}

		again, exit, err := readYesOrNo(scanner, "Посчитать ещё раз? [y/n]: ")
		if exit || err != nil || !again {
			printExit(err)
			return
		}
		fmt.Println()
	}
}

func printExit(err error) {
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
		os.Exit(1)
	}
	fmt.Println("Выход.")
}

func readNumber(scanner *bufio.Scanner, prompt string) (float64, bool, error) {
	for {
		line, err := readLine(scanner, prompt)
		if err != nil {
			return 0, false, err
		}
		if isExit(line) {
			return 0, true, nil
		}
		value, err := parseFloat(line)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		return value, false, nil
	}
}

func readOperation(scanner *bufio.Scanner, prompt string) (string, bool, error) {
	for {
		line, err := readLine(scanner, prompt)
		if err != nil {
			return "", false, err
		}
		if isExit(line) {
			return "", true, nil
		}
		operation := strings.TrimSpace(line)
		switch operation {
		case "+", "-", "*", "/":
			return operation, false, nil
		default:
			fmt.Printf("Ошибка: неизвестная операция %q. Допустимы +, -, *, /.\n", line)
		}
	}
}

func readYesOrNo(scanner *bufio.Scanner, prompt string) (bool, bool, error) {
	for {
		line, err := readLine(scanner, prompt)
		if err != nil {
			return false, false, err
		}
		if isExit(line) {
			return false, true, nil
		}
		switch strings.ToLower(strings.TrimSpace(line)) {
		case "y", "yes", "д", "да", "":
			return true, false, nil
		case "n", "no", "н", "нет":
			return false, false, nil
		default:
			fmt.Println("Введите y (да) или n (нет).")
		}
	}
}

func readLine(scanner *bufio.Scanner, prompt string) (string, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return scanner.Text(), nil
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
