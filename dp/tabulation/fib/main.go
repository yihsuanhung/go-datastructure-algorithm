package main

import "fmt"

func main() {
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(50))
}

func fib(n int) int {
	// make a table
	table := make([]int, n+1)

	// initial the table
	table[1] = 1

	// iteration
	for i := 0; i <= n; i++ {
		current := table[i]
		if i+1 <= n {
			table[i+1] += current
		}
		if i+2 <= n {
			table[i+2] += current
		}
	}

	return table[n]
}
