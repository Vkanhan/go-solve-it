package main

import (
	"fmt"
	"sort"
)

// Bucket sort approach
// topKFrequent finds the top K most frequent elements in the nums slice.
func topKFrequent(nums []int, k int) []int {
	result := []int{}

	frequencyMap := make(map[int]int)

	// Create a slice of slices where the index represents the frequency
	// and the value is a slice of elements with that frequency
	buckets := make([][]int, len(nums)+1)

	for _, num := range nums {
		frequencyMap[num]++
	}

	for element, frequency := range frequencyMap {
		buckets[frequency] = append(buckets[frequency], element)
	}
	for i := len(nums); i >= 0; i-- {
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

	store := map[int]int{}

	for _, elem := range nums {
		store[elem]++
	}

	for i := 0; i < k; i++ {
		var maxResult int
		var maxKey int

		for key, value := range store {
			if value > maxResult {
				maxResult = value
				maxKey = key
			}
		}
		result = append(result, maxKey)
		delete(store, maxKey)
	}
	return result
}

func topKSort(nums []int, k int) []int {
	set := make(map[int]int)

	for _, num := range nums {
		set[num]++
	}

	result := make([]int, 0, len(set))
	for key := range set {
		result = append(result, key)
	}

	sort.Slice(result, func(i, j int) bool {
		return set[result[i]] > set[result[j]]
	})
	return result[:k]
}

func main() {
	nums := []int{1, 2, 2, 2, 3, 3}
	k := 2

	fmt.Println(topKFrequent(nums, k))

	fmt.Println(topKFrequentElement(nums, k))

	fmt.Println(topKSort(nums, k))
}
