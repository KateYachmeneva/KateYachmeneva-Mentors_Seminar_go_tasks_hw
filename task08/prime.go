package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	_, found := firstDivisor(n)
	return !found
}

func firstDivisor(n int) (int, bool) {
	if n <= 1 {
		return 0, false
	}
	for divisor := 2; divisor <= n/divisor; divisor++ {
		if n%divisor == 0 {
			return divisor, true
		}
	}
	return 0, false
}

func primeMessage(n int) string {
	if n <= 1 {
		return fmt.Sprintf("%d — не простое число: простыми бывают только числа больше 1.", n)
	}
	if divisor, found := firstDivisor(n); found {
		return fmt.Sprintf("%d — не простое число: делится на %d.", n, divisor)
	}
	return fmt.Sprintf("%d — простое число.", n)
}
