package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(item interface{}) { *h = append(*h, item.(int)) }
func (h *MaxHeap) Pop() interface{} {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

func main() {
	h := &MaxHeap{1, 5, 3, 7, 2}
	heap.Init(h)
	heap.Push(h, 8)
	heap.Push(h, 11)
	// heap.Push(h, 0)

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

}
