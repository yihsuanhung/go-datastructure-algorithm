package main

import "fmt"

func main() {
	fmt.Println(bestSum(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSum(8, []int{2, 3, 5}))
	fmt.Println(bestSum(8, []int{1, 4, 5}))
	fmt.Println(bestSum(100, []int{1, 2, 5, 25}))

}

func bestSum(m int, nums []int) []int {
	// make table
	table := make([][]int, m+1)
	table[0] = []int{}

	// iteration
	for i := 0; i < m+1; i++ {
		if table[i] != nil {
			for _, num := range nums {
				if i+num < m+1 {
					values := append(table[i], num)
					if table[i+num] == nil || len(values) < len(table[i+num]) {
						table[i+num] = values
					}
				}
			}
		}
	}

	return table[m]
}
