package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(item interface{}) { *h = append(*h, item.(int)) }
func (h *MinHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

func main() {
	minHeap := &MinHeap{1, 5, 3, 7, 2}
	heap.Init(minHeap)
	heap.Push(minHeap, 8)
	// heap.Push(minHeap, 0)

	fmt.Println(heap.Pop(minHeap))
}
