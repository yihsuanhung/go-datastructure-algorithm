package linkedlist

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func New() *LinkedList {
	ll := new(LinkedList)
	return ll
}

func NewNode(Value interface{}) *Node {
	node := new(Node)
	node.Value = Value
	return node
}

func (ll *LinkedList) isEmpty() bool {
	return ll.head == nil
}

func (ll *LinkedList) Prepend(node *Node) {
	isEmpyt := ll.isEmpty()

	if isEmpyt {
		// assign new node as tail and head
		ll.head = node
		ll.tail = node
	} else {
		// assign new node to head and reassign tail
		node.Next = ll.head
		ll.head = node
	}

}

func (ll *LinkedList) Append(node *Node) {
	isEmpty := ll.isEmpty()

	if isEmpty {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.Next = node
		ll.tail = node
	}
}

func (ll *LinkedList) Pop() (*Node, error) {
	isEmpty := ll.isEmpty()

	if isEmpty {
		return nil, fmt.Errorf("list is empty")
	}

	if ll.head == ll.tail {
		n := ll.head
		ll.head = nil
		ll.tail = nil
		return n, nil
	}

	var previous *Node

	for current := ll.head; current != ll.tail; current = current.Next {

		previous = current

	}

	n := previous.Next
	previous.Next = nil
	ll.tail = previous

	return n, nil

}

func (ll *LinkedList) PopFirst() (*Node, error) {

	isEmpty := ll.isEmpty()

	if isEmpty {
		return nil, fmt.Errorf("list is empty")
	}

	if ll.head == ll.tail {
		n := ll.head
		ll.head = nil
		ll.tail = nil
		return n, nil
	}

	n := ll.head

	ll.head = ll.head.Next
	return n, nil

}

func (ll *LinkedList) Head() *Node {
	return ll.head
}

func (ll *LinkedList) Tail() *Node {
	return ll.tail
}

func (ll *LinkedList) Remove(node *Node) error {

	if node == ll.head {
		_, err := ll.PopFirst()
		return err
	}

	if node == ll.tail {
		ll.Pop()
		return nil
	}

	var previous *Node

	current := ll.head

	for ; current != node && current != nil; current = current.Next {

		previous = current

	}

	if current == nil {
		return fmt.Errorf("%v is not in the linked list", node)
	}

	previous.Next = previous.Next.Next
	return nil

}
