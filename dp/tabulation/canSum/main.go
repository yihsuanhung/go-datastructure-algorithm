package main

import "fmt"

func main() {
	fmt.Println(canSum(7, []int{2, 3}))
	fmt.Println(canSum(7, []int{5, 3, 4, 7}))
	fmt.Println(canSum(7, []int{2, 4}))
	fmt.Println(canSum(8, []int{5, 3, 4}))
	fmt.Println(canSum(300, []int{7, 14}))
}

func canSum(m int, nums []int) bool {
	// make table
	table := make([]bool, m+1)
	table[0] = true

	// iteration
	for i := 0; i < m+1; i++ {
		if table[i] {
			for _, num := range nums {
				if i+num < m+1 {
					table[i+num] = true
				}
			}
		}
	}

	return table[m]
}
