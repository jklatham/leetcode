package main

import "fmt"

type Tree struct {
	root *Node
}

type Node struct {
	key   byte
	left  *Node
	right *Node
}

// Tree

func (t *Tree) insert(data byte) {
	if t.root == nil {
		t.root = &Node{key: data}
	} else {
		t.root.insert(data)
	}
}

// Node
func (n *Node) insert(data byte) {
	if data <= n.key {
		if n.left == nil {
			n.left = &Node{key: data}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{key: data}
		} else {
			n.right.insert(data)
		}
	}
}

//Pre Order
func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%c ", n.key)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

//Post Order
func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%c ", n.key)
	}
}

func printInOrder(n *Node) {
	if n == nil {
		return
	} else {
		printInOrder(n.left)
		fmt.Printf("%c ", n.key)
		printInOrder(n.right)
	}
}

func main() {
	fmt.Println("Binary Search Tree")
	var t Tree
	t.insert('F')
	t.insert('B')
	t.insert('A')
	t.insert('D')
	t.insert('C')
	t.insert('E')
	t.insert('G')
	t.insert('H')
	t.insert('I')

	fmt.Printf("Pre Order: ")
	printPreOrder(t.root)
	fmt.Println()
	fmt.Printf("Post Order: ")
	printPostOrder(t.root)
	fmt.Println()
	fmt.Printf("In Order: ")
	printInOrder(t.root)
	fmt.Println()
}
