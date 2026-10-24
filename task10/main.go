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

	fmt.Println("Статистика последовательности")
	fmt.Println("Введите q в любой момент, чтобы выйти.")

	for {
		count, exit, err := readCount(scanner, "Введите количество чисел N: ")
		if exit || err != nil {
			finishProgram(err)
			return
		}

		numbers, exit, err := readNumbers(scanner, count)
		if exit || err != nil {
			finishProgram(err)
			return
		}

		stats, err := calculateStats(numbers)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		printStats(stats)

		again, exit, err := readAgain(scanner, "Посчитать для другой последовательности? [y/n]: ")
		if exit || err != nil || !again {
			finishProgram(err)
			return
		}
		fmt.Println()
	}
}

func printStats(stats sequenceStats) {
	fmt.Printf("Сумма: %d\n", stats.sum)
	fmt.Printf("Среднее значение: %s\n", formatAverage(stats.average))
	fmt.Printf("Максимальное значение: %d\n", stats.maximum)
	fmt.Printf("Минимальное значение: %d\n", stats.minimum)
	fmt.Printf("Чётных: %d\n", stats.evenCount)
	fmt.Printf("Нечётных: %d\n", stats.oddCount)
}

func readNumbers(scanner *bufio.Scanner, count int) ([]int64, bool, error) {
	numbers := make([]int64, 0, count)
	for index := 1; index <= count; index++ {
		number, exit, err := readInteger(scanner, fmt.Sprintf("Число %d: ", index))
		if exit || err != nil {
			return nil, exit, err
		}
		numbers = append(numbers, number)
	}
	return numbers, false, nil
}

func finishProgram(err error) {
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
		count, err := parseCount(line)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		return count, false, nil
	}
}

func readInteger(scanner *bufio.Scanner, prompt string) (int64, bool, error) {
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

func readAgain(scanner *bufio.Scanner, prompt string) (bool, bool, error) {
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
