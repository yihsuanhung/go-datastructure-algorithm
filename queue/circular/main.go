package main

import "fmt"

type CircularQueue struct {
	list []int
	cap  int
	head int
	tail int
}

func New(n int) *CircularQueue {
	cq := &CircularQueue{
		list: make([]int, n),
		head: -1,
		tail: -1,
		cap:  n,
	}
	return cq
}

// enqueue要從尾巴加進去
func (cq *CircularQueue) EnQueue(value int) error {
	// 滿了不能 enqueue
	if cq.IsFull() {
		return fmt.Errorf("queue is full")
	}
	if cq.IsEmpty() {
		cq.head = 0
		cq.tail = 0
		cq.list[0] = value
		return nil
	} else {
		tailInx := (cq.tail + 1) % cq.cap
		cq.list[tailInx] = value
		cq.tail = tailInx
		return nil
	}
}

// dequeue要從頭去除
func (cq *CircularQueue) DeQueue() (int, error) {
	// 空的不能dequeue
	if cq.IsEmpty() {
		return -1, fmt.Errorf("queue is empty")
	}

	pop := cq.list[cq.head]
	cq.list[cq.head] = 0

	if cq.head == cq.tail {
		// 頭尾指向同一個（代表queue內只有一個元素）
		cq.head = -1
		cq.tail = -1
	} else {
		// 頭尾不同，帶表queue內有不只一個元素
		cq.head = (cq.head + 1) % cq.cap
	}
	return pop, nil
}

func (cq *CircularQueue) Front() int {
	front := cq.list[cq.head]
	return front
}

func (cq *CircularQueue) Rear() int {
	rear := cq.list[cq.tail]
	return rear

}

func (cq *CircularQueue) IsEmpty() bool {
	// 只有兩個都是0的情況才會是空
	return cq.head == -1 && cq.tail == -1

}

func (cq *CircularQueue) IsFull() bool {
	// 因為是從尾巴加入，所以尾巴的下一個如果是頭的話，就是滿了
	// 尾巴的下一個一定要%cap，因為如果超過len要從頭開始算
	return cq.head == (cq.tail+1)%cq.cap
}

func main() {
	cq := New(4)

	cq.EnQueue(1)
	cq.EnQueue(2)
	cq.DeQueue()
	cq.EnQueue(3)
	cq.EnQueue(4)
	cq.DeQueue()
	cq.DeQueue()
	cq.DeQueue()

	list := cq.list

	fmt.Println(list, "head", cq.head, "tail", cq.tail)
}
