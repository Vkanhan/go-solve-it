package main

import (
	"fmt"
	"sort"
)

// In Go, you can't directly "sort" a map because Go's maps are inherently unordered collections.
// However, you can extract the keys or values from the map into a slice and sort that slice according to the desired criteria.
func sortKeyMap(hashMap map[int]int) {
	result := make([]int, 0, len(hashMap))

	for key := range hashMap {
		result = append(result, key)
	}

	sort.Ints(result)

	for _, key := range result {
		fmt.Println(key)
	}
}

func sortValueMap(hashMap map[string]string) {
	result := make([]string, 0, len(hashMap))

	for _, value := range hashMap {
		result = append(result, value)
	}

	sort.Strings(result)

	for _, value := range result {
		fmt.Println(value)
	}
}

func main() {
	hashMap := map[int]int{6: 2, 9: 1, 8: 3, 7: 4}
	sortKeyMap(hashMap)

	hashStore := map[string]string{"red": "apple", "yellow": "banana", "blue": "sky"}
	sortValueMap(hashStore)

}
