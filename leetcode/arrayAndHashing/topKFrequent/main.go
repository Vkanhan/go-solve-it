package main

import (
	"fmt"
)

//Bucket sort approach
// topKFrequent finds the top K most frequent elements in the nums slice.
func topKFrequent(nums []int, k int) []int {
	result := []int{}

	// Create a map to count the frequency of each element in nums
	frequencyMap := make(map[int]int)

	// Create a slice of slices where the index represents the frequency
	// and the value is a slice of elements with that frequency
	buckets := make([][]int, len(nums)+1)

	// Count the frequency of each element in nums
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Group elements by their frequency using the buckets slice
	for element, frequency := range frequencyMap {
		buckets[frequency] = append(buckets[frequency], element)
	}

	// Iterate over the buckets from the highest frequency to the lowest
	for i := len(nums); i >= 0; i-- {
		// Append elements to the result slice until we have K elements
		for _, element := range buckets[i] {
			result = append(result, element)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}

func topKFrequentElement(nums []int, k int) []int {
    var result []int

	// Create a map to count the frequency of each element in the nums slice.
	store := map[int]int{}

	 // Loop through the nums slice and update the frequency count in the map.
	for _, elem := range nums {
		store[elem]++
	}

	 // Loop k times to find and store the k most frequent elements
	for i := 0; i < k; i++ {
		// Variables to keep track of the element with the maximum frequency in the current iteration.
		var maxResult int
		var maxKey int

		  // Loop through the map to find the element with the maximum frequency.
		for key, value := range store {
			if value > maxResult {
				maxResult = value
				maxKey = key
			}
		}
		// Append the element with the highest frequency to the result slice
		result = append(result, maxKey)

		 // Remove the element with the highest frequency from the map to find the next most frequent element in the next iteration.
		delete(store, maxKey)
	}
	return result
}

func main() {
	nums := []int{1, 2, 2, 2, 3, 3}
	k := 2

	fmt.Println(topKFrequent(nums, k))

	fmt.Println(topKFrequentElement(nums, k))
}
