package main

import (
	"container/heap"
	"fmt"
)

// IntegerHeap is a type
type IntegerHeap []int

// IntegerHeap method - gets the length of integerHeap
func (iheap IntegerHeap) Len() int { return len(iheap) }

// IntegerHeap method - checks if element of i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool { return iheap[i] > iheap[j] } // We are using max heap here

// IntegerHeap method - swaps the element of i to j index
func (iheap IntegerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

// IntegerHeap method - pushes the item
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
	fmt.Println(iheap)
}

// IntegerHeap method - pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var previous IntegerHeap = *iheap
	var n int = len(previous)
	var x1 interface{} = previous[n-1]
	fmt.Println("previous", previous, n, x1)
	*iheap = previous[0 : n-1]
	fmt.Println("iheap", iheap)
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5, 8, 9}
	heap.Init(intHeap)
	heap.Push(intHeap, 3)
	//fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Println(heap.Pop(intHeap))
	}
}
