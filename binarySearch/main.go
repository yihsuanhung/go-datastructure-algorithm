package main

import "fmt"

func BinarySearch(array []int, target int) int {

	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			right = mid + 1
		} else {
			left = mid - 1
		}
	}
	return -1
}

func main() {
	array := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	index := BinarySearch(array, 7)
	fmt.Println(index)
}
