package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to group anagrams without sorting
func groupAnagrams(strs []string) [][]string {
	// Initialize a map to store the grouped anagrams
	anagrams := make(map[[26]int][]string)

	for _, str := range strs {
		// Create a character frequency array (size 26 for each lowercase letter)
		var charCount [26]int

		// Count the frequency of each character in the string
		for _, char := range str {
			//The expression char - 'a' takes the ASCII value of the character char
			//and subtracts the ASCII value of 'a' (which is 97).
			charCount[char-'a']++
		}

		// Use the character frequency array as the key in the map
		anagrams[charCount] = append(anagrams[charCount], str)
	}

	// Initialize a result slice to store the grouped anagrams
	result := make([][]string, 0, len(anagrams))

	// Collect all the lists of anagrams from the map and add them to the result slice
	for _, group := range anagrams {
		result = append(result, group)
	}

	return result
}

// Function to group anagrams by sorting
func GroupAnagrams(strs []string) [][]string {
	anagram := make(map[string][]string)

	for _, str := range strs {
		sorted := strings.Split(str, "")
		sort.Strings(sorted)
		sortedStr := strings.Join(sorted, "")

		anagram[sortedStr] = append(anagram[sortedStr], str)
	}

	result := make([][]string, 0, len(anagram))

	for _, group := range anagram {
		result = append(result, group)
	}

	return result
}

func groupAnagramsBySorting(strs []string) [][]string {
	anagram := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		anagram[sortedStr] = append(anagram[sortedStr], str)
	}

	result := make([][]string, 0, len(anagram))
	for _, group := range anagram {
		result = append(result, group)
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
	fmt.Println(GroupAnagrams(strs))
	fmt.Println(groupAnagramsBySorting(strs))
}
