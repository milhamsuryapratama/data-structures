package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Creates a new Node
func NewNode(data int) *Node {
	return &Node{Data: data}
}

// Inserts a node into the tree
func (n *Node) Insert(data int) {
	if n == nil {
		return
	} else if data <= n.Data {
		if n.Left == nil {
			n.Left = NewNode(data)
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(data)
		} else {
			n.Right.Insert(data)
		}
	}
}

// Inverts a binary tree
func (n *Node) InvertTree() {
	if n == nil {
		return
	}

	// Swap the left and right subtree.
	n.Left, n.Right = n.Right, n.Left

	// Recursively invert the left and right subtree.
	n.Left.InvertTree()
	n.Right.InvertTree()
}

// Preorder traversal (Root, Left, Right)
func (n *Node) Preorder() {
	if n == nil {
		return
	}

	fmt.Print(n.Data, " ")
	n.Left.Preorder()
	n.Right.Preorder()
}

func main() {
	root := NewNode(4)
	root.Insert(2)
	root.Insert(7)
	root.Insert(1)
	root.Insert(3)
	root.Insert(6)
	root.Insert(9)

	root.Preorder() // prints: 4 2 1 3 7 6 9
	fmt.Println()

	ko, _ := json.MarshalIndent(root, "  ", "  ")
	fmt.Println("root", string(ko))

	root.InvertTree()

	root.Preorder() // prints: 4 7 9 6 2 3 1
	fmt.Println()

	po, _ := json.MarshalIndent(root, "  ", "  ")
	fmt.Println("root", string(po))
}
