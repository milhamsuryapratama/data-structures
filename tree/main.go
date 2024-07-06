package main

import (
	"encoding/json"
	"fmt"
)

// Define a tree node
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Initialize a new Node
func NewNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

// Insert a Node in the Tree
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

// To check if a value is present in the Tree
func (n *Node) Contains(data int) bool {
	if n == nil {
		return false
	} else if data == n.Data {
		return true
	} else if data < n.Data {
		return n.Left.Contains(data)
	} else {
		return n.Right.Contains(data)
	}
}

func main() {
	node := NewNode(5)
	node.Insert(4)
	node.Insert(7)
	node.Insert(2)
	node.Insert(9)
	node.Insert(3)

	ko, _ := json.MarshalIndent(node, "  ", "  ")
	fmt.Println("node", string(ko))

	fmt.Println(node.Contains(7)) // expected to print 'true'
	fmt.Println(node.Contains(1)) // expected to print 'false'
}
