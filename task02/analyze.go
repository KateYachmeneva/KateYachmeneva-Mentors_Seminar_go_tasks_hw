package main

import (
	"fmt"
	"strconv"
	"strings"
)

func числоПоложительное(число int64) bool {
	return число > 0
}

func числоОтрицательное(число int64) bool {
	return число < 0
}

func числоЧётное(число int64) bool {
	return число%2 == 0
}

func числоДелитсяНа(число, делитель int64) bool {
	if делитель == 0 {
		return false
	}
	return число%делитель == 0
}

func текстПроЗнак(число int64) string {
	if числоПоложительное(число) {
		return "Число положительное."
	}
	if числоОтрицательное(число) {
		return "Число отрицательное."
	}
	return "Число равно нулю: оно ни положительное, ни отрицательное."
}

func текстПроЧётность(число int64) string {
	if числоЧётное(число) {
		return "Число чётное."
	}
	return "Число нечётное."
}

func текстПроДелимостьНа3И5(число int64) string {
	делитсяНа3 := числоДелитсяНа(число, 3)
	делитсяНа5 := числоДелитсяНа(число, 5)
	switch {
	case делитсяНа3 && делитсяНа5:
		return "Число делится и на 3, и на 5."
	case делитсяНа3:
		return "Число делится на 3, но не делится на 5."
	case делитсяНа5:
		return "Число делится на 5, но не делится на 3."
	default:
		return "Число не делится ни на 3, ни на 5."
	}
}

func полныйОтвет(число int64) string {
	return текстПроЗнак(число) + "\n" + текстПроЧётность(число) + "\n" + текстПроДелимостьНа3И5(число)
}

func строкаВЦелоеЧисло(строка string) (int64, error) {
	текст := strings.TrimSpace(строка)
	if текст == "" {
		return 0, fmt.Errorf("пустая строка")
	}
	if strings.Contains(текст, ".") || strings.Contains(текст, ",") {
		return 0, fmt.Errorf("%q не целое число", strings.TrimSpace(строка))
	}
	число, err := strconv.ParseInt(текст, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%q не является целым числом", strings.TrimSpace(строка))
	}
	return число, nil
}

func командаВыхода(текст string) bool {
	switch strings.ToLower(strings.TrimSpace(текст)) {
	case "q", "quit", "exit", "й", "выход":
		return true
	default:
		return false
	}
}
