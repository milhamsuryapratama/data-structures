package main

import (
	"encoding/json"
	"log"
)

// QueueItem represents each item in the Queue
type QueueItem struct {
	value interface{}
	next  *QueueItem
}

// Queue is our main data structure for the implementation
type Queue struct {
	head *QueueItem
	tail *QueueItem
	size int
}

// NewQueue creates a new empty queue
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue adds an item to the end of the queue
func (q *Queue) Enqueue(value interface{}) {
	newItem := &QueueItem{value: value}

	log.Println("newItem ", newItem)

	if q.IsEmpty() {
		q.head = newItem
		q.tail = newItem
	} else {
		// {{1, nil},{1, nil}}
		log.Println("q ", q)

		log.Println("q.head before", q.head)
		log.Println("q.tail before", q.tail)
		q.tail.next = newItem
		// {{1, nil},{1, {2, nil}}}
		log.Println("q.tail ", q.tail)
		log.Println("q.head ", q.head)

		q.tail = newItem
		log.Println("q.tail ", q.tail)
	}

	q.size++
	ko, _ := json.MarshalIndent(q, "  ", "  ")
	log.Println(string(ko))
	log.Println("q ", q)
	//log.Println("tail ", q.tail)
	//log.Println("size ", q.size)

	log.Println("===============")
}

// Dequeue removes an item from the front of the queue and returns its value
func (q *Queue) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	value := q.head.value
	q.head = q.head.next
	q.size--

	return value, true
}

// IsEmpty returns true if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func main() {
	queue := NewQueue()

	queue.Enqueue(1)
	queue.Enqueue(2)
	//queue.Enqueue(3)

	//val, ok := queue.Dequeue()
	//
	//for ok {
	//	fmt.Println(val)
	//	val, ok = queue.Dequeue()
	//}
}
