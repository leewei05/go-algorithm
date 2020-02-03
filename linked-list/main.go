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
	l := linkedList{}
	l.length()
	l.insertStart(10)
	l.insertStart(12)
	l.insertEnd(6)
	l.insertEnd(9)
	l.length()
	l.traverseList()
	l.remove(12)
	l.traverseList()
	l.length()
	l.clear()
	l.length()
}

func (l *linkedList) traverseList() {
	if l.len == 0 {
		fmt.Println("Linkedlist is empty.")
		return
	}

	current := l.node
	for current != nil {
		fmt.Printf("Node data: %d \n", current.data)
		current = current.next
	}
}

func (l *linkedList) length() {
	fmt.Printf("Total length of linked list: %d \n", l.len)
}

func (l *linkedList) insertStart(d int) {
	newNode := &node{}
	newNode.data = d
	if l.len == 0 {
		l.node = newNode
		l.len++
		fmt.Printf("Insert %d from start \n", d)
		return
	}
	newNode.next = l.node
	l.node = newNode
	l.len++
	fmt.Printf("Insert %d from start \n", d)
}

func (l *linkedList) insertEnd(d int) {
	newNode := &node{}
	newNode.data = d
	if l.len == 0 {
		l.node = newNode
		l.len++
		fmt.Printf("Insert %d \n", d)
		return
	}
	current := l.node
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	l.len++
	fmt.Printf("Insert %d from end \n", d)
}

func (l *linkedList) remove(d int) {
	current := l.node
	previous := &node{}
	if l.node.data == d {
		l.node = nil
		l.node = current.next
	}
	for current != nil {
		if current.data == d {
			current = current.next
			break
		}
		previous = current
		current = current.next
	}

	if current == nil {
		fmt.Printf("There is no: %d \n", d)
		return
	}
	previous.next = current
	current = nil
	l.len--
	fmt.Printf("Remove data: %d \n", d)
}

func (l *linkedList) clear() {
	current := l.node
	for current != nil {
		l.node = current.next
		current = l.node
		l.len--
	}
	fmt.Printf("Clear data \n")
}

func (l *linkedList) reverse() {
	if l.node == nil || l.len == 1 {
		return
	}
	current := l.node
	nullPter := &node{}
	for current.next != nil {
		current.next = nullPter
	}
	l.node = current
}
