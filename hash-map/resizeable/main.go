package main

import "fmt"

// Define the Entry struct
type Entry struct {
	Key   int
	Value string
	Next  *Entry
}

// Define the HashMap struct with dynamic array
type HashMap struct {
	array []*Entry
	size  int
}

// Define max load factor before resizing
const LoadFactor = 0.75
const InitialCapacity = 4

// Use this function to create a new HashMap
func NewHashMap() *HashMap {
	return &HashMap{
		array: make([]*Entry, InitialCapacity),
		size:  0,
	}
}

func main() {
	hashMap := NewHashMap()
	hashMap.Put(1, "Ilham")
	hashMap.Put(2, "Aka")
	hashMap.Put(3, "Danny")
	hashMap.Put(4, "Haped")
	hashMap.Put(5, "Samsul")

	fmt.Println(hashMap.Get(5))
}

// Hash function
func hashFunc(key int, capacity int) int {
	return key % capacity
}

// Put function - inserts a value into the hash map by key
func (h *HashMap) Put(key int, value string) {
	// Check load factor before insertion
	if float64(h.size)/float64(len(h.array)) >= LoadFactor {
		h.resize()
	}

	index := hashFunc(key, len(h.array))

	// If place is empty, insert new entry
	if h.array[index] == nil {
		h.array[index] = &Entry{Key: key, Value: value}
		h.size++
		return
	}

	// We have a collision, handle it by adding into the linked list
	e := h.array[index]
	for e.Next != nil {
		e = e.Next
	}
	e.Next = &Entry{Key: key, Value: value}
	h.size++
}

func (h *HashMap) Get(key int) (string, bool) {
	index := hashFunc(key, len(h.array))
	if h.array[index] == nil {
		return "", false
	}

	e := h.array[index]
	for e != nil {
		if e.Key == key {
			return e.Value, true
		}
		e = e.Next
	}

	return "", false
}

// Resize the HashMap by creating a new larger array
// and rehashing all the existing entries
func (h *HashMap) resize() {
	newArray := make([]*Entry, len(h.array)*2)
	for _, e := range h.array {
		// Traverse the linked list
		if e != nil {
			p := e
			for p != nil {
				index := hashFunc(p.Key, len(newArray))
				if newArray[index] == nil {
					newArray[index] = &Entry{
						Key:   p.Key,
						Value: p.Value,
					}
				} else {
					// Handle collision
					newP := newArray[index]
					for newP.Next != nil {
						newP = newP.Next
					}
					newP.Next = &Entry{Key: p.Key, Value: p.Value}
				}
				p = p.Next
			}
		}
	}
	h.array = newArray
}
