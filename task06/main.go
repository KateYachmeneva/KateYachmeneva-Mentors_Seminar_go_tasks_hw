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

	fmt.Println("Количество положительных, отрицательных чисел и нулей")
	fmt.Println("Введите q в любой момент, чтобы выйти.")

	for {
		count, exit, err := readCount(scanner, "Введите количество чисел N: ")
		if exit || err != nil {
			printExit(err)
			return
		}

		counts, exit, err := readCounts(scanner, count)
		if exit || err != nil {
			printExit(err)
			return
		}

		fmt.Printf("Положительных: %d\n", counts.positive)
		fmt.Printf("Отрицательных: %d\n", counts.negative)
		fmt.Printf("Нулей: %d\n", counts.zero)

		again, exit, err := readYesOrNo(scanner, "Проверить другую последовательность? [y/n]: ")
		if exit || err != nil || !again {
			printExit(err)
			return
		}
		fmt.Println()
	}
}

func readCounts(scanner *bufio.Scanner, count int) (numberCounts, bool, error) {
	var counts numberCounts
	for i := 1; i <= count; i++ {
		value, done, err := readNumber(scanner, fmt.Sprintf("Число %d: ", i))
		if done || err != nil {
			return numberCounts{}, done, err
		}
		counts = countBySign(counts, value)
	}
	return counts, false, nil
}

func printExit(err error) {
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
		os.Exit(1)
	}
	fmt.Println("Выход.")
}

func readCount(scanner *bufio.Scanner, prompt string) (int, bool, error) {
	for {
		line, err := readLine(scanner, prompt)
		if err != nil {
			return 0, false, err
		}
		if isExit(line) {
			return 0, true, nil
		}
		value, err := parseCount(line)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		return value, false, nil
	}
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
		value, err := parseNumber(line)
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
