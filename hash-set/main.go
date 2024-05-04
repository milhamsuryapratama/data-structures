package main

import (
	"fmt"
)

type HashSet map[interface{}]bool

func NewHashSet() HashSet {
	return make(HashSet)
}

func (set HashSet) Add(element interface{}) {
	set[element] = true
}

func (set HashSet) Contains(element interface{}) bool {
	return set[element]
}

func (set HashSet) Remove(element interface{}) {
	delete(set, element)
}

func main() {
	set := NewHashSet()

	set.Add("Ilham")
	set.Add("Aka")

	fmt.Println(set.Contains("Ilham")) // Prints: true

	set.Remove("Ilham")

	fmt.Println(set.Contains("Ilham")) // Prints: false

	set.Remove("Aka")

	fmt.Println(set.Contains("Aka")) // Prints: false
}
