package main

import "fmt"

type Queue struct {
	list []int
}

func (q *Queue) Enqueue(item int) {
	q.list = append(q.list, item)
}

func (q *Queue) Dequeue() (int, error) {

	if len(q.list) == 0 {
		return -1, fmt.Errorf("Queue is empty")
	}

	removeItem := q.list[0]
	q.list = q.list[1:]

	return removeItem, nil
}

func main() {
	q := Queue{}
	fmt.Println(q)
	q.Enqueue(100)
	fmt.Println(q)
	q.Enqueue(200)
	fmt.Println(q)
	q.Enqueue(300)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)

}
