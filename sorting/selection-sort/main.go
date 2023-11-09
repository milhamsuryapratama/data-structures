package main

import "fmt"

func main() {
	data := []int{9, 4, 7, 2, 8, 1, 0, 6, 5, 3, 90, 52, 14, 87}

	fmt.Println(data)

	// init minimum value
	var minValue = data[0]
	var minValueIdx int

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[minValueIdx] > data[j] {
				minValue = data[j]
				minValueIdx = j
			}
		}

		data[minValueIdx] = data[i]
		data[i] = minValue

		minValue = data[i+1]
		minValueIdx = i + 1
	}

	fmt.Println(data)
}
