package main

import (
	"fmt"

	"github.com/yihsuanhung/go-datastructure/linkedlist"
)

func print(l *linkedlist.LinkedList) {
	for current := l.Head(); current != nil; current = current.Next {
		fmt.Printf("%v -> ", current.Value)
	}
	fmt.Println("nil")
}

func main() {
	ll := linkedlist.New()
	node1 := linkedlist.NewNode("node1")
	node2 := linkedlist.NewNode("node2")
	node4 := linkedlist.NewNode("node4")
	ll.Prepend(node1)
	ll.Prepend(node2)
	ll.Append(node4)

	ll.Remove(node1)

	print(ll)

}
