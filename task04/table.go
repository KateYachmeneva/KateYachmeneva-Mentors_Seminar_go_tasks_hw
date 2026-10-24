package main

import "fmt"


func printTable(n int) {
	for multiplier := 1; multiplier <= 10; multiplier++ {
		fmt.Printf("%d x %d = %d\n", n, multiplier, n*multiplier)
	}
}
