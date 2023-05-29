package main

import (
	"fmt"
)

// Queue is our main data structure which is a slice
type Queue struct {
	items []interface{}
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	return &Queue{items: make([]interface{}, 0)}
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(value interface{}) {
	q.items = append(q.items, value)
}

// Dequeue removes an item from the front of the queue and returns its value
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	value := q.items[0]
	q.items = q.items[1:]

	return value, true
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	val, ok := queue.Dequeue()

	for ok {
		fmt.Println(val)
		val, ok = queue.Dequeue()
	}
}
