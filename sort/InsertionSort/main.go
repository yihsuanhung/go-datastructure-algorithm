package main

import "fmt"

func InsertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		current := array[i]
		j := i - 1

		for j >= 0 && array[j] > current {
			array[j+1] = array[j]
			j--
		}

		array[j+1] = current
	}
}

func main() {
	array := []int{1, 5, 6, 3, 2, 5, 4, 7}
	fmt.Println(array)
	InsertionSort(array)
	fmt.Println(array)
}
