package main

import (
	"fmt"
)

// printAllDuplicates finds and returns all duplicate numbers 
func printAllDuplicates(nums []int) []int {
	// Slice to store the duplicate numbers
	duplicates := []int{}

	// Map to track seen numbers
	hashSet := make(map[int]bool)

	// Map to track already recorded duplicates to avoid adding duplicates multiple times
	seen := make(map[int]bool)

	// Iterate through each number in the input slice
	for _, num := range nums {
		
		//if the number is already in the hashSet indicates its a duplicate and hasnt been recorded in duplicate slice
		if hashSet[num] && !seen[num] {
			// Add the number to the duplicates slice
			duplicates = append(duplicates, num)

			// Mark the number as seen in the duplicates slice to avoid adding it again
			seen[num] = true
		}
		
		// Mark the number as seen in the hashSet
		hashSet[num] = true
	}

	// Return the slice containing all duplicate numbers
	return duplicates
}

func main() {
	// Example input slice
	nums := []int{1, 3, 5, 1, 3, 8, 7, 7, 8, 1}

	// Print all duplicates found in the input slice
	fmt.Println(printAllDuplicates(nums))
}
