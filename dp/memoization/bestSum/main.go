package main

import "fmt"

func main() {
	fmt.Println(bestSum(7, []int{5, 3, 4, 7}))
	fmt.Println(bestSum(8, []int{2, 3, 5}))
	fmt.Println(bestSum(8, []int{1, 4, 5}))
	fmt.Println(bestSum(100, []int{1, 2, 5, 25}))
}

func bestSum(target int, nums []int) []int {
	memo := make(map[int][]int)
	return mBestSum(target, nums, memo)
}

func mBestSum(target int, nums []int, memo map[int][]int) []int {
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

	var best []int // nil

	// recursive
	for _, num := range nums {
		remaining := target - num
		remainingResult := mBestSum(remaining, nums, memo)

		if remainingResult != nil {
			combination := append(remainingResult, num)
			if best == nil || len(combination) < len(best) {
				best = combination
			}
		}
	}

	memo[target] = best
	return best

}
