package main

import (
	"errors"
	"fmt"
)

// Node represents an element in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents a singly linked list
type LinkedList struct {
	head *Node
	size int
}

// NewLinkedList creates a new empty singly linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, size: 0}
}

// Append adds a new element to the end of the list
func (ll *LinkedList) Append(value int) *LinkedList {
	newNode := &Node{data: value, next: nil}
	if ll.head == nil {
		ll.head = newNode
	} else {
		current := ll.head

		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.size++

	return ll
}

// Display outputs the contents of the linked list
func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d", current.data)
		if current.next != nil {
			fmt.Print(" -> ")
		}
		current = current.next
	}
	fmt.Println()
}

// Delete a node with the given data
func (ll *LinkedList) Delete(data int) error {
	if ll.head == nil {
		return errors.New("list is empty")
	}

	if ll.head.data == data {
		ll.head = ll.head.next
		return nil
	}

	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return nil
		}
		current = current.next
	}

	return errors.New("element not found")
}

func (ll *LinkedList) Find(data int) error {
	if ll.head == nil {
		return errors.New("list is empty")
	}

	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			fmt.Println("data found ", current.next)
			return nil
		}
		current = current.next
	}

	return errors.New("element not found")
}

func main() {
	ll := NewLinkedList()
	ll.Append(10).
		Append(20).
		Append(30).
		Append(40).
		Append(50)

	ll.Display()

	ll.Delete(30)

	ll.Display()

	ll.Find(30)
}
