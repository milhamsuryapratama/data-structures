package main

import (
	"fmt"
)

// Node structure for doubly linked list
type Node struct {
	data int
	prev *Node
	next *Node
}

// DoublyLinkedList structure
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// newNode initializes a new node
func newNode(data int) *Node {
	return &Node{data: data, prev: nil, next: nil}
}

// insert appends a new node at the end of the list
func (dll *DoublyLinkedList) insert(data int) {
	newNode := newNode(data)

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}

	//log.Println("head ", dll.head)
	//log.Println("tail ", dll.tail)
	//log.Println("newNode ", newNode)

	newNode.prev = dll.tail
	//log.Println("newNode.prev ", newNode)
	dll.tail.next = newNode
	//log.Println("dll.tail.next ", dll.tail)
	dll.tail = newNode
	//log.Println("dll.tail ", dll.tail)
}

// delete removes a node with the given data
func (dll *DoublyLinkedList) delete(data int) bool {
	current := dll.head

	for current != nil {
		if current.data == data {
			//log.Println("current ", current)
			if current.prev != nil {
				//log.Println("current.prev.next ", current.prev.next)
				//log.Println("current.next ", current.next)
				current.prev.next = current.next
			} else {
				dll.head = current.next
			}

			if current.next != nil {
				//log.Println("current.next.prev ", current.next.prev)
				//log.Println("current.prev ", current.prev)
				current.next.prev = current.prev
			} else {
				dll.tail = current.prev
			}

			return true
		}
		current = current.next
	}
	return false
}

// display prints the elements of the doubly linked list
func (dll *DoublyLinkedList) display() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d <-> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	dll := DoublyLinkedList{}

	dll.insert(1)
	dll.insert(2)
	dll.insert(3)
	dll.display() // 1 <-> 2 <-> 3 <-> nil

	if dll.delete(2) {
		fmt.Println("Deleted 2")
	} else {
		fmt.Println("2 not found")
	}
	dll.display() // 1 <-> 3 <-> nil
}
