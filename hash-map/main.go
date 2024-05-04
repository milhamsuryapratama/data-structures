package main

import (
	"fmt"
)

type Entry struct {
	Key   int
	Value string
	Next  *Entry
}

const ArraySize = 2 // choose any prime number

type HashMap struct {
	array [ArraySize]*Entry
	size  int
}

func main() {
	hashMap := &HashMap{}

	hashMap.Put(1, "Ilham")
	hashMap.Put(2, "Aka")
	hashMap.Put(5, "Danny")
	hashMap.Put(7, "Haped")

	fmt.Println(hashMap.Get(5))
}

func hashFunc(i int) int {
	return i % 2
}

func (h *HashMap) Put(key int, value string) {
	index := hashFunc(key)
	if h.array[index] == nil {
		h.array[index] = &Entry{Key: key, Value: value}
		return
	}

	// We have a collision, handle it by adding to the end of linked list
	e := h.array[index]
	for e.Next != nil {
		e = e.Next
	}
	e.Next = &Entry{Key: key, Value: value}
}

func (h *HashMap) Get(key int) (string, bool) {
	index := hashFunc(key)
	if h.array[index] == nil {
		return "", false
	}

	e := h.array[index]

	// Find the entry for the key in the linked list
	for e != nil {
		if e.Key == key {
			return e.Value, true
		}
		e = e.Next
	}

	return "", false
}
