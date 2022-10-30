package main

import "fmt"

func main() {

	fmt.Println(fib(0))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	memo := map[int]int{}
	return memFib(n, memo)
}

func memFib(n int, memo map[int]int) int {

	if m, ok := memo[n]; ok {
		return m
	}
	if n == 1 || n == 2 {
		return 1
	}

	memo[n] = memFib(n-2, memo) + memFib(n-1, memo)

	return memo[n]
}
