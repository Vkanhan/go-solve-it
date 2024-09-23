package main

import (
	"fmt"
)

// findMin returns the minimum element in a rotated sorted array using binary search.
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// If mid element is greater than the right element,
		// then the smallest element is in the right half
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// Else, the smallest element is in the left half (including mid)
			right = mid
		}
	}
	
	// When left == right, we have found the smallest element
	return nums[left]
}

func main() {
	nums1 := []int{3, 4, 5, 6, 1, 2}
	fmt.Println("Minimum element:", findMin(nums1))

	nums2 := []int{4, 5, 0, 1, 2, 3}
	fmt.Println("Minimum element:", findMin(nums2))
}
