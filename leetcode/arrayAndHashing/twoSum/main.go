package main

import (
	"fmt"
)

// twoSum finds two indices i and j such that nums[i] + nums[j] == target and i != j.
// It returns the indices i and j, ensuring i < j.
func twoSum(nums []int, target int) (int, int) {
	// Create a map to store the number and its index
	numsMap := make(map[int]int)

	// Iterate over the array
	for i, num := range nums {
		// Calculate the complement of the current number with respect to the target
		complement := target - num

		// Check if the complement exists in the map
		if j, found := numsMap[complement]; found {
			// If found, return the indices in the correct order
			if i < j {
				return i, j
			}
			return j, i
		}

		// Store the current number and its index in the map
		numsMap[num] = i
	}

	// Return -1, -1 if no solution is found 
	return -1, -1
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	// Find the indices
	i, j := twoSum(nums, target)

	// Print the indices
	fmt.Printf("Indices: %d, %d\n", i, j)
}
