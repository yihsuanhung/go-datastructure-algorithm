package linkedlist

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

func Append() {}

func Pop() {}

func PopFirst() {}

func (ll LinkedList) Head() *Node {
	return ll.head
}

func (ll LinkedList) Tail() *Node {
	return ll.tail
}

func Remove() {}
