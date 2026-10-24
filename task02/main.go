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

	fmt.Println("Анализ числа")
	fmt.Println("Введите q в любой момент, чтобы выйти.")

	for {
		число, хочетВыйти, err := спроситьЧисло(scanner, "Введите число N: ")
		if хочетВыйти || err != nil {
			завершитьПрограмму(err)
			return
		}

		fmt.Println(полныйОтвет(число))

		ещёРаз, хочетВыйти, err := спроситьПовтор(scanner, "Проверить ещё число? [y/n]: ")
		if хочетВыйти || err != nil || !ещёРаз {
			завершитьПрограмму(err)
			return
		}
		fmt.Println()
	}
}

func завершитьПрограмму(err error) {
	if err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
		os.Exit(1)
	}
	fmt.Println("Выход.")
}

func спроситьЧисло(scanner *bufio.Scanner, подсказка string) (int64, bool, error) {
	for {
		строка, err := прочитатьСтроку(scanner, подсказка)
		if err != nil {
			return 0, false, err
		}
		if командаВыхода(строка) {
			return 0, true, nil
		}
		число, err := строкаВЦелоеЧисло(строка)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}
		return число, false, nil
	}
}

func спроситьПовтор(scanner *bufio.Scanner, подсказка string) (bool, bool, error) {
	for {
		строка, err := прочитатьСтроку(scanner, подсказка)
		if err != nil {
			return false, false, err
		}
		if командаВыхода(строка) {
			return false, true, nil
		}
		switch strings.ToLower(strings.TrimSpace(строка)) {
		case "y", "yes", "д", "да", "":
			return true, false, nil
		case "n", "no", "н", "нет":
			return false, false, nil
		default:
			fmt.Println("Введите y (да) или n (нет).")
		}
	}
}

func прочитатьСтроку(scanner *bufio.Scanner, подсказка string) (string, error) {
	fmt.Print(подсказка)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return scanner.Text(), nil
}
