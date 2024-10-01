package main 

import(
	"fmt"
)

// insertionSort sorts an array of integers using the insertion sort algorithm.
func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func main() {
	unsorted := []int{6, 2, 5, 33, 1, 0}
    fmt.Println(insertionSort(unsorted))
}
