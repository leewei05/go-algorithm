package main

import "fmt"

type linkedList struct {
	len  int
	node *node
}

type node struct {
	data int
	next *node
}

func main() {

}

func (l *linkedList) traverseList() {
	if l.len == 0 {
		fmt.Println("Linkedlist is empty.")
		return
	}

	current := l.node
	for current != nil {
		fmt.Printf("Node data: %d", current.data)
		current = l.node.next
	}
}
