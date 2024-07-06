package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Data     int
	Children []*Node
}

// Creates a new Node
func NewNode(data int) *Node {
	return &Node{Data: data}
}

// Adds a child to a Node
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

// Check if a value is present in the Tree
func (n *Node) Contains(data int) bool {
	if n.Data == data {
		return true
	}

	for _, child := range n.Children {
		if child.Contains(data) {
			return true
		}
	}

	return false
}

func main() {
	root := NewNode(10)

	child1 := NewNode(20)
	child2 := NewNode(30)
	child3 := NewNode(40)

	root.AddChild(child1)
	root.AddChild(child2)
	root.AddChild(child3)

	child2Parent := NewNode(50)
	child1.AddChild(child2Parent)
	child2Parent.AddChild(NewNode(80))

	child2ParentRight := NewNode(60)
	child1.AddChild(child2ParentRight)
	child2ParentRight.AddChild(NewNode(90))

	child3.AddChild(NewNode(70))

	data, _ := json.MarshalIndent(root, "", "  ")
	fmt.Println(string(data))

	fmt.Println(root.Contains(90)) // prints: true
}
