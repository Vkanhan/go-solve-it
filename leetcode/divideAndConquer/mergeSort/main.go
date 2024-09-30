package main

import (
	"fmt"
)

//mergeSort sorts the slice by divide and conquer 
func mergeSort(items []int) []int {
    // Base case: if the slice has 1 or 0 elements, it is already sorted
    if len(items) < 2 {
        return items
    }

    // Recursively split the slice into two halves and sort each half
    first := mergeSort(items[:len(items)/2])     
    second := mergeSort(items[len(items)/2:])    

    // Merge the two sorted halves into one sorted slice
    return merge(first, second)
}

//merge merges 2 slices to form a single slice
func merge(a []int, b []int) []int {

    final := make([]int, 0, len(a)+len(b))

    i := 0  // Index for slice 'a'
    j := 0  // Index for slice 'b'

    // Iterate through both slices until one is fully traversed
    for i < len(a) && j < len(b) {
        // Append the smaller element from 'a' or 'b' to 'final'
        if a[i] < b[j] {
            final = append(final, a[i])
            i++  
        } else {
            final = append(final, b[j])
            j++  
        }
    }

	//Append the remaining elements of a and b if any
   final = append(final, a[i:]...)
   final = append(final, b[j:]...)

    return final
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
