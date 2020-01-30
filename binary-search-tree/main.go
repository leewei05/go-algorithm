package main

import "fmt"

type bst struct {
	root *node
}

type node struct {
	data      int
	leftNode  *node
	rightNode *node
}

func main() {
	b := bst{}
	b.insert(10)
	b.insert(6)
	b.insert(7)
	b.insert(12)
	b.insert(13)

	b.traverseBST()
}

func (b *bst) insert(d int) {
	if b.root == nil {
		newNode := &node{}
		b.root = newNode
		b.root.data = d
		return
	}
	b.insertNode(b.root, d)
}

func (b *bst) insertNode(n *node, d int) {
	newNode := &node{}
	if d > n.data {
		if n.rightNode != nil {
			b.insertNode(n.rightNode, d)
		} else {
			newNode.data = d
			n.rightNode = newNode
		}
	} else if d < n.data {
		if n.leftNode != nil {
			b.insertNode(n.leftNode, d)
		} else {
			newNode.data = d
			n.leftNode = newNode
		}
	}
}

func (b *bst) traverseBST() {
	if b.root == nil {
		fmt.Printf("No node exists.")
		return
	}
	b.printNode(b.root)
}

func (b *bst) printNode(n *node) {
	if n.leftNode != nil {
		b.printNode(n.leftNode)
	}

	fmt.Printf("Node data: %d \n", n.data)

	if n.rightNode != nil {
		b.printNode(n.rightNode)
	}
}

func (b *bst) delete(d int) {

}
