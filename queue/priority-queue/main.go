package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Value    string
	Priority int
}

type PriorityQueue []Item

// Enqueue inserts the value into the queue based on priority
func (pq *PriorityQueue) Enqueue(value string, priority int) {
	newPQItem := Item{
		Value:    value,
		Priority: priority,
	}
	*pq = append(*pq, newPQItem)

	// Sorting Slice to hold priority queue property
	sort.Slice(*pq, func(i, j int) bool {
		return (*pq)[i].Priority > (*pq)[j].Priority
	})
}

// Dequeue extracts the element with the highest priority and returns it
func (pq *PriorityQueue) Dequeue() string {
	if pq.IsEmpty() {
		return ""
	}
	item := (*pq)[0].Value
	*pq = (*pq)[1:]
	return item
}

// IsEmpty checks whether the queue is empty
func (pq *PriorityQueue) IsEmpty() bool {
	return len(*pq) == 0
}

func main() {
	pq := PriorityQueue{}

	pq.Enqueue("apple", 1)
	pq.Enqueue("banana", 2)
	pq.Enqueue("cherry", 2)

	for !pq.IsEmpty() {
		fmt.Println(pq.Dequeue())
	}
}
