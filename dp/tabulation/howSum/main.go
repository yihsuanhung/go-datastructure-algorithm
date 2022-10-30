package main

import "fmt"

func main() {
	fmt.Println(howSum(7, []int{2, 3}))
	fmt.Println(howSum(7, []int{5, 3, 4, 7}))
	fmt.Println(howSum(7, []int{2, 4}), howSum(7, []int{2, 4}) == nil)
	fmt.Println(howSum(8, []int{2, 3, 5}))
	fmt.Println(howSum(300, []int{7, 14}), howSum(300, []int{7, 14}) == nil)
}

func howSum(m int, nums []int) []int {
	// make table
	table := make([][]int, m+1)
	table[0] = []int{}
	// fmt.Println(table[0] == nil, table[1] == nil)

	// iteration
	for i := 0; i < m+1; i++ {
		if table[i] != nil {
			for _, num := range nums {
				if i+num < m+1 {
					values := table[i]
					table[i+num] = append(values, num)
				}
			}
		}
	}

	return table[m]
}
