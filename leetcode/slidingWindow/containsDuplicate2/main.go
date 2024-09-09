package main

import (
	"fmt"
)

// Function to check for duplicates within distance k
func containsDuplicate2BruteForce(nums []int, K int) bool {
	if len(nums) == 0 {
		return false
	}

	for L := 0; L < len(nums); L++ {
		for R := L + 1; R < len(nums) && R < L+K; R++ {
			if nums[L] == nums[R] {
				return true
			}
		}
	}
	return false
}

//Find it through sliding window method
func containsDuplicate2Optimized(nums []int, K int) bool {
	//keys are the elements from nums and values are most recent index at which num is found
	seen := make(map[int]int)

	for i, num := range nums {
		// Check if the current number was seen before and if it's within distance k
		if idx, found := seen[num]; found && i-idx <= K {
			return true
		}
		// Update the map with the current number and its position
		seen[num] = i 
	}
	return false
}

func main() {
	nums := []int{2, 4, 3, 4, 5, 6, 7}
	K := 3
	fmt.Println(containsDuplicate2BruteForce(nums, K))
	fmt.Println(containsDuplicate2Optimized(nums, K))
}
