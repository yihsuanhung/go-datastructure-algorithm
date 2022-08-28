package main

import "fmt"

// 遞迴把array拆成兩半，並且使用merge function來排序
func mergeSort(array []int) []int {
	// 若array只剩下一個，就不用再拆了
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2
	l := mergeSort(array[:mid])
	r := mergeSort(array[mid:])

	// 全部拆開後，逐個merge
	return merge(l, r)
}

// 把兩個sorted array由小至大合併成一個
func merge(a []int, b []int) []int {

	merged := []int{}

	// i控制a, j控制b
	i, j := 0, 0

	// ab都還沒合併完成之前執行，此回圈會持續執行到ab其中一個被合併完
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	// 若是b先被合併完，把a所有元素加到merged內
	for ; i < len(a); i++ {
		merged = append(merged, a[i])
	}
	// 同上，把b所有元素加到merged內
	for ; j < len(b); j++ {
		merged = append(merged, b[j])
	}

	return merged
}

func main() {
	array := []int{4, 6, 7, 1, 7, 3, 10, 6, 3, 2, 8, 9}
	sorted := mergeSort(array)
	fmt.Println("合併前", array)
	fmt.Println("合併後", sorted)
}
