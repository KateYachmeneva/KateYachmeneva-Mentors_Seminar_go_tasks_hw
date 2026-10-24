package main

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for divisor := 2; divisor <= n/divisor; divisor++ {
		if n%divisor == 0 {
			return false
		}
	}
	return true
}
