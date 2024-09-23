package main

import (
	"fmt"
)

// longestConsecutive finds the length of the longest consecutive sequence in an array of integers.
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Create a set to store unique elements
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0

	for num := range numSet {
		// Check if this is the start of a sequence
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			// Count the length of the sequence
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			// Update the longest streak found
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}

func main() {
	nums1 := []int{2, 20, 4, 10, 3, 4, 5}
	fmt.Println("Output:", longestConsecutive(nums1))
}


