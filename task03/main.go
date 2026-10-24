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

	fmt.Println("Сумма и среднее чисел от 1 до N")
	fmt.Println("Введите q в любой момент, чтобы выйти.")

	for {
		n, exit, err := readNumber(scanner, "Введите число N: ")
		if exit || err != nil {
			printExit(err)
			return
		}

		sum, average, err := sumAndAverage(n)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		fmt.Printf("Сумма чисел от 1 до %d: %s\n", n, formatNumber(sum))
		fmt.Printf("Среднее значение: %s\n", formatNumber(average))

		again, exit, err := readYesOrNo(scanner, "Посчитать для другого N? [y/n]: ")
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

func readNumber(scanner *bufio.Scanner, prompt string) (int64, bool, error) {
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
