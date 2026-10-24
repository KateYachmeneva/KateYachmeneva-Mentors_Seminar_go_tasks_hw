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

	fmt.Println("Проверка простого числа")
	fmt.Println("Введите q в любой момент, чтобы выйти.")

	for {
		n, exit, err := readNumber(scanner, "Введите число: ")
		if exit || err != nil {
			printExit(err)
			return
		}

		fmt.Println(primeMessage(n))

		again, exit, err := readYesOrNo(scanner, "Проверить другое число? [y/n]: ")
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

func readNumber(scanner *bufio.Scanner, prompt string) (int, bool, error) {
	for {
		line, err := readLine(scanner, prompt)
		if err != nil {
			return 0, false, err
		}
		if isExit(line) {
			return 0, true, nil
		}
		value, err := parseInteger(line)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		return value, false, nil
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

func parseInteger(raw string) (int, error) {
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
