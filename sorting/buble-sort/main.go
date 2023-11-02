package main

import "fmt"

func main() {
	unsortedList := []int{8, 3, 4, 2, 9, 3, 7, 10, 6}

	fmt.Println("===== [start] bubble sort =====")

	//bubbleSort1(unsortedList)
	//bubbleSort2(unsortedList)
	bubbleSort3(unsortedList)

	fmt.Println("===== [end] bubble sort =====")
}

func bubbleSort1(unsortedList []int) {
	fmt.Print(unsortedList)

	for i := 0; i < len(unsortedList)-1; i++ {
		for j := 0; j < len(unsortedList)-i-1; j++ {
			if unsortedList[j] > unsortedList[j+1] {
				unsortedList[j], unsortedList[j+1] = unsortedList[j+1], unsortedList[j]
			}
		}
	}

	fmt.Println(" -> ", unsortedList)
}

func bubbleSort2(unsortedList []int) {
	firstPointer := 0
	secondPointer := firstPointer + 1

	swapped := true

	for swapped {

		if firstPointer > len(unsortedList)-1 || secondPointer > len(unsortedList)-1 {
			firstPointer = 0
			secondPointer = firstPointer + 1
		}

		firstPointerValue := unsortedList[firstPointer]
		secondPointerValue := unsortedList[secondPointer]

		if firstPointerValue > secondPointerValue {
			unsortedList[firstPointer] = secondPointerValue
			unsortedList[secondPointer] = firstPointerValue

			swapped = true

			firstPointer++
			secondPointer = firstPointer + 1

			continue
		}

		swapped = false
	}
}

func bubbleSort3(unsortedList []int) {
	fmt.Print(unsortedList)

	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(unsortedList)-1; i++ {
			if unsortedList[i] > unsortedList[i+1] {
				unsortedList[i], unsortedList[i+1] = unsortedList[i+1], unsortedList[i]
				swapped = true
			}
		}
	}

	fmt.Println(" -> ", unsortedList)
}
