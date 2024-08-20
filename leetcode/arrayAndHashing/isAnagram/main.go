package main

import (
	"fmt"
)

// isAnagram checks if two strings are anagrams of each other
func isAnagram(s string, t string) bool {
	// If the lengths of the two strings are different, they cannot be anagrams
	if len(s) != len(t) {
		return false
	}

	// Create a map to count the occurrences of each character in the first string
	charCount := make(map[rune]int)

	// Iterate over each character in the first string
	for _, char := range s {
		// Increment the count for this character in the map
		charCount[char]++
	}

	// Iterate over each character in the second string
	for _, char := range t {
		// Decrement the count for this character in the map
		charCount[char]--
		// If the count goes below zero, it means there are more occurrences of this character in the second string
		if charCount[char] < 0 {
			return false
		}
	}

	// Check the counts for all characters
	for _, count := range charCount {
		// If any character count is not zero, the strings are not anagrams
		if count != 0 {
			return false
		}
	}

	// If all counts are zero, the strings are anagrams
	return true
}

func main() {
	
	s1, t1 := "anagram", "nagaram"
	fmt.Println(isAnagram(s1, t1)) 

	s2, t2 := "rat", "car"
	fmt.Println(isAnagram(s2, t2)) 
}
