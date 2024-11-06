package main

import (
	"fmt"
	"sort"
)

func sortByLength(strs []string) {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
}

func main() {
	strs := []string{"apple", "banana", "kiwi", "grape"}
	sortByLength(strs)
	fmt.Println(strs)
}
