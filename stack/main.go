package main

import (
	"fmt"
	"sort"
)

type Stack struct {
	list []int
}

func New() *Stack {
	s := new(Stack)
	s.list = []int{}
	return s
}

func (s *Stack) IsEmpty() bool {
	return len(s.list) == 0
}

func (s *Stack) Push(item int) {
	s.list = append(s.list, item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}
	item := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return item, nil
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("stack is empyt")
	}
	return s.list[len(s.list)-1], nil
}

func (s *Stack) GetMin() (int, error) {
	if s.IsEmpty() {
		return -1, fmt.Errorf("stack is empty")
	}
	sorted := make([]int, len(s.list))
	copy(sorted, s.list)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i] < sorted[j]
	})
	return sorted[0], nil
}

func main() {
	s := New()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	min, _ := s.GetMin()
	// top, _ := s.Top()

	fmt.Println(s, min)
	s.Pop()
	fmt.Println(s, min)
	// fmt.Println(s, "min", min, top)
}
