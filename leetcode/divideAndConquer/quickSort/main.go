package main

import (
	"fmt"
)

// QuickSort sorts an array using the Quick Sort algorithm.
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr // Base case: arrays with 0 or 1 element are already sorted
	}

	// Choose a pivot (here, we choose the last element)
	pivot := arr[len(arr)-1]
	left := []int{}
	right := []int{}

	// Partitioning the array into left and right sub-arrays
	for _, v := range arr[:len(arr)-1] {
		if v < pivot {
			left = append(left, v) // Elements less than pivot go to left
		} else {
			right = append(right, v) // Elements greater or equal go to right
		}
	}

	// Recursively apply QuickSort to left and right, then combine
	return append(append(QuickSort(left), pivot), QuickSort(right)...)
}

func main() {
	arr := []int{7, 2, 1, 6, 8, 5, 3, 4}
	fmt.Println("Unsorted array:", arr)

	sortedArr := QuickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
