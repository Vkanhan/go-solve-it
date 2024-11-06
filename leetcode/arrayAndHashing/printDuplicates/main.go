package main

import (
	"fmt"
)

// printAllDuplicates finds and returns all duplicate numbers 
func printAllDuplicates(nums []int) []int {
	duplicates := []int{}

	hashSet := make(map[int]bool)

	seen := make(map[int]bool)

	for _, num := range nums {
		if hashSet[num] && !seen[num] {
			// Add the number to the duplicates slice
			duplicates = append(duplicates, num)
			seen[num] = true
		}
		
		hashSet[num] = true
	}
	return duplicates
}

func main() {
	nums := []int{1, 3, 5, 1, 3, 8, 7, 7, 8, 1}
	fmt.Println(printAllDuplicates(nums))
}
