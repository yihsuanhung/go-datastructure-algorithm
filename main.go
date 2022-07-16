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
	node := linkedlist.NewNode("node1")
	ll.Prepend(node)

	print(ll)

}
