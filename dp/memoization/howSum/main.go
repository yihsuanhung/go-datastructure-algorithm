package main

import "fmt"

func main() {
	// fmt.Println(howSum(7, []int{2, 3}))
	// fmt.Println(howSum(7, []int{5, 3, 4, 7}))
	// fmt.Println(howSum(7, []int{2, 4}))
	// fmt.Println(howSum(8, []int{2, 3, 5}))
	fmt.Println(howSum(300, []int{7, 14}))
}

func howSum(target int, nums []int) []int {
	// To initialize a map, use the built in make function
	memo := make(map[int][]int)
	return mHowSum(target, nums, memo)
}

func mHowSum(target int, nums []int, memo map[int][]int) []int {
	// hit base cases
	if m, ok := memo[target]; ok {
		return m
	}
	if target == 0 {
		return []int{}
	}
	if target < 0 {
		return nil
	}

	// recursive
	for _, num := range nums {
		remaining := target - num
		remainingResult := mHowSum(remaining, nums, memo)
		if remainingResult != nil {
			m := append(remainingResult, num)
			memo[target] = m
			return m
		}
	}
	memo[target] = nil
	return nil

}
