package main

import "fmt"

func main() {
	target := 300
	nums := []int{7, 14}
	fmt.Println(canSum(target, nums))
}

func canSum(target int, nums []int) bool {
	memo := map[int]bool{}
	return mCanSum(target, nums, memo)
}

func mCanSum(target int, nums []int, memo map[int]bool) bool {

	if m, ok := memo[target]; ok {
		return m
	}
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}

	for _, num := range nums {
		remaining := target - num
		if mCanSum(remaining, nums, memo) {
			memo[target] = true
			return true
		}
	}

	memo[target] = false
	return false

}
