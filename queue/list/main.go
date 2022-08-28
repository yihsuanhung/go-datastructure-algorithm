package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	list *list.List
}

func New() *Queue {
	q := new(Queue)
	q.list = list.New()
	return q
}

func (q *Queue) Enqueue(item int) {
	q.list.PushBack(item)
}

func (q *Queue) Dequeue() error {
	if q.list.Len() > 0 {
		q.list.Remove(q.list.Front())
		return nil
	}
	return fmt.Errorf("Queue is empty")

}

func main() {

	q := New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	for q.list.Len() > 0 {
		fmt.Println(q.list.Front().Value)
		q.Dequeue()
	}

}
