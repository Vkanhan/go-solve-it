package main

import "fmt"

// SelectionSort sorts an array using the selection sort algorithm
func SelectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        // Find the minimum element in the unsorted part
        minIdx := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }
        // Swap the found minimum element with the first element of the unsorted part
        arr[i], arr[minIdx] = arr[minIdx], arr[i]
    }
}

func main() {
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("Original array:", arr)
    
    SelectionSort(arr)
    
    fmt.Println("Sorted array:", arr)
}
